package requests

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/vashish1/OnlineClassPortal/pkg/hub"
	"github.com/vashish1/OnlineClassPortal/pkg/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type messageBytes []byte

// handleCreateNewRoom creates a new room for user.
func (msg messageBytes) handleCreateNewRoom() {
	var newRoom models.NewRoomRequest
	if err := json.Unmarshal(msg, &newRoom); err != nil {
		log.Errorln("could not convert to required New Room Request struct")
		return
	}

	roomID, err := createNewRoom(newRoom)
	if err != nil {
		log.Errorln("unable to create a new room for user:", newRoom.Email, "err:", err.Error())
		return
	}

	// Broadcast a joined message.
	userJoinedMessage := models.Joined{
		RoomID:      roomID,
		Email:       newRoom.Email,
		RoomName:    newRoom.RoomName,
		MessageType: models.JoinRoomMsgType,
	}

	jsonContent, err := json.Marshal(userJoinedMessage)
	if err != nil {
		log.Errorln("could not marshal to jsonByte while creating room", err.Error())
		return
	}

	hub.HubConstruct.SendMessage(jsonContent, newRoom.Email)
}

func (msg messageBytes) handleRequestUserToJoinRoom() {
	var request joinRequest
	if err := json.Unmarshal(msg, &request); err != nil {
		log.Errorln("could not convert to required Joined Request struct, err:", err)
		return
	}

	for _, user := range request.Users {
		roomRegisteredUser, err := request.requestUserToJoinRoom(user)
		if err != nil {
			log.Errorln("error while requesting to room, err:", err)
			continue
		}

		data := struct {
			MsgType       string `json:"msgType"`
			RequesterID   string `json:"requesterID"`
			RequesterName string `json:"requesterName"`
			UserRequested string `json:"userRequested"`
			RoomID        string `json:"roomID"`
			RoomName      string `json:"roomName"`
		}{
			values.SentRoomRequestMsgType,
			request.RequestingUserID, request.RequestingUserName,
			user, request.RoomID, request.RoomName,
		}

		jsonContent, err := json.Marshal(data)
		if err != nil {
			log.Errorln("could not marshal to RequestUsersToJoinRoom, err:", err)
			continue
		}

		// Send back RequestUsersToJoinRoom signal to everyone registered in room.
		for _, roomRegisteredUser := range roomRegisteredUser {
			HubConstruct.sendMessage(jsonContent, roomRegisteredUser)
		}

		HubConstruct.sendMessage(jsonContent, user)
	}
}

// handleUserAcceptRoomRequest accepts room join request.
func (msg messageBytes) handleUserAcceptRoomRequest() {
	var roomRequest joined
	if err := json.Unmarshal(msg, &roomRequest); err != nil {
		log.Errorln("could not convert to required Join Room Request struct, err:", err)
		return
	}

	values.MapEmailToName.Mutex.RLock()
	roomRequest.Name = values.MapEmailToName.Mapper[roomRequest.Email]
	values.MapEmailToName.Mutex.RUnlock()

	users, err := roomRequest.acceptRoomRequest()
	if err != nil {
		log.Errorln("could not accept room request", err)
		return
	}

	for _, user := range users {
		HubConstruct.sendMessage(msg, user)
	}
}

// handleRequestAllMessages fetches messages given the specified message room ID.
func (msg messageBytes) handleRequestMessages(user string) {
	roomContent := room{}
	if err := json.Unmarshal(msg, &roomContent); err != nil {
		log.Errorln("could not unmarshall file on handle request partitioned message, err:", err)
		return
	}

	if err := roomContent.getPartitionedMessageInRoom(); err != nil {
		log.Errorln("could not get all messages in room, err:", err)
		return
	}

	// Strip off unnecessary message information sent to serverDB.
	for index := range roomContent.Messages {
		roomContent.Messages[index].RoomID = ""
	}
	roomContent.RoomIcon = ""

	data := struct {
		MsgType     string `json:"msgType"`
		RoomContent room   `json:"roomPageDetails"`
	}{
		values.RequestMessages,
		roomContent,
	}

	jsonContent, err := json.Marshal(&data)
	if err != nil {
		log.Errorln("could not marshal images, err:", err)
		return
	}

	HubConstruct.sendMessage(jsonContent, user)
}

// handleNewMessage broadcasts users message to all online users and also saves to database.
func (msg messageBytes) handleNewMessage() {
	var newMessage Message
	if err := json.Unmarshal(msg, &newMessage); err != nil {
		log.Errorln("could not convert to required New Message struct", err)
		return
	}

	newMessage.Time = time.Now().Format(values.TimeLayout)
	// Save message to database ensuring user is registered to room.
	registeredUsers, err := newMessage.saveMessageContent()
	if err != nil {
		log.Errorln("error saving msg to db, err:", err, "user:", newMessage.UserID)
		return
	}

	jsonContent, err := json.Marshal(newMessage)
	if err != nil {
		log.Errorln("error converting message to json, err:", err)
		return
	}

	// Message is sent back to all online users including sender.
	for _, registeredUser := range registeredUsers {
		HubConstruct.sendMessage(jsonContent, registeredUser)
	}
}

// handleExitRoom exits requesters joined room and also notifies all room users.
func (msg messageBytes) handleExitRoom(author string) {
	data := struct {
		Email  string `json:"userID"`
		RoomID string `json:"roomID"`
	}{}

	if err := json.Unmarshal(msg, &data); err != nil {
		log.Errorln("could not retrieve json on exit room request", err)
		return
	}

	user := User{Email: data.Email}
	registeredUsers, err := user.exitRoom(data.RoomID)
	if err != nil {
		log.Errorln("error exiting room", err)
		return
	}

	// Broadcast to all online users of a room exit.
	for _, registeredUser := range registeredUsers {
		HubConstruct.sendMessage(msg, registeredUser)
	}

	HubConstruct.sendMessage(msg, author)
}

// handleNewFileUpload creates a new file content in database.
// If file create is a success, a file upload success is sent to client to send next chunk.
// next chunk could be the next preceding file chunk if another user has uploaded file content.
// If file upload error, send back error message to user
func (msg messageBytes) handleNewFileUpload() {
	file := file{}
	if err := json.Unmarshal(msg, &file); err != nil {
		log.Errorln("error unmarshalling byte in handle new file upload, err:", err)
		return
	}

	data := struct {
		MsgType      string `json:"msgType"`
		ErrorMessage string `json:"errorMsg,omitempty"`
		RecentHash   string `json:"recentHash"`
		FileName     string `json:"fileName,omitempty"`
		FileHash     string `json:"fileHash"`
		Chunk        int    `json:"nextChunk"`
	}{}

	data.FileName = file.FileName
	data.FileHash = file.UniqueFileHash
	user := file.User

	if err := file.uploadNewFile(); err == mongo.ErrNoDocuments || err == nil {
		// Send next file chunk and current hash which is a "".
		data.MsgType = values.UploadFileChunkMsgType

		// Resume file chunk upload if Current chunk is greater than 0.
		if file.Chunks > 0 {
			data.Chunk = file.Chunks + 1
		} else {
			data.Chunk = file.Chunks
		}

	} else {
		log.Println("error on handle new file upload calling UploadNewFile, error:", err)
		data.ErrorMessage = values.ErrFileUpload.Error()
		data.MsgType = values.UploadFileErrorMsgType
	}

	jsonContent, err := json.Marshal(&data)
	if err != nil {
		log.Println("Error sending marshalled ")
		return
	}

	HubConstruct.sendMessage(jsonContent, user)
}

func (msg messageBytes) handleUploadFileChunk() {
	data := struct {
		MsgType         string `json:"msgType"`
		User            string `json:"userID"`
		FileName        string `json:"fileName"`
		File            string `json:"file,omitempty"`
		NewChunkHash    string `json:"newChunkHash,omitempty"`
		RecentChunkHash string `json:"recentChunkHash,omitempty"`
		ChunkIndex      int    `json:"chunkIndex,omitempty"`
		NextChunk       int    `json:"nextChunk"`
		FileHash        string `json:"fileHash"`
	}{}

	if err := json.Unmarshal(msg, &data); err != nil {
		log.Errorln("error unmarshalling file in handle upload chunk, err:", err)
		return
	}

	file := fileChunks{
		UniqueFileHash:     data.NewChunkHash,
		FileBinary:         data.File,
		ChunkIndex:         data.ChunkIndex,
		CompressedFileHash: data.FileHash,
	}

	userID := data.User
	var recentFileExist bool
	// If file upload is a new file, set recent file exist as true.
	if data.RecentChunkHash == "" {
		recentFileExist = true
	} else {
		recentFileExist = fileChunks{UniqueFileHash: data.RecentChunkHash}.fileChunkExists()
	}

	data.RecentChunkHash, data.File, data.NewChunkHash = "", "", ""
	data.NextChunk, data.ChunkIndex = 0, 0

	fileHash := sha256.Sum256([]byte(file.FileBinary))
	// Check if client sent file hash is same as server generated Hash.
	if hex.EncodeToString(fileHash[:]) != file.UniqueFileHash || !recentFileExist {
		fmt.Println("Invalid unique hash", hex.EncodeToString(fileHash[:]), recentFileExist)
		data.MsgType = "UploadError"

		// Re-request for current chunk index.
		jsonContent, err := json.Marshal(&data)
		if err != nil {
			log.Println("Could not generate jsonContent to re-request file chunk")
			return
		}

		HubConstruct.sendMessage(jsonContent, data.User)

		return
	}

	if err := file.addFileChunk(); err != nil {
		// What could be cases where err is not nil.
		// File could have already been added to database?.
		// We still request for next file chunk, if when we receive a new fille chunk,
		// so that when we notice file corruption, we re-request from corrupted stage.
		log.Println(err)
	}

	data.NextChunk = file.ChunkIndex + 1

	jsonContent, err := json.Marshal(&data)
	if err != nil {
		log.Println("Error sending marshalled ")
		return
	}

	HubConstruct.sendMessage(jsonContent, userID)
}

// handleUploadFileUploadComplete is called when file chunk uploads is complete.
// File accessibility is broadcasted to other users in the room so as to download
// file.
func (msg messageBytes) handleUploadFileUploadComplete() {
	data := struct {
		MsgType  string `json:"msgType"`
		UserID   string `json:"userID"`
		UserName string `json:"name"`
		FileName string `json:"fileName"`
		FileSize string `json:"fileSize"`
		FileHash string `json:"fileHash"`
		RoomID   string `json:"roomID"`
	}{}

	if err := json.Unmarshal(msg, &data); err != nil {
		log.Errorln("error unmarshalling file in handle upload file complete, err:", err)
		return
	}

	data.MsgType = values.UploadFileSuccessMsgType
	values.MapEmailToName.Mutex.RLock()
	data.UserName = values.MapEmailToName.Mapper[data.UserID]
	values.MapEmailToName.Mutex.RUnlock()

	roomUsers, err := Message{
		RoomID:  data.RoomID,
		UserID:  data.UserID,
		Name:    data.UserName,
		Message: data.FileName,
		Time:    time.Now().Format(values.TimeLayout),
		Type:    values.MessageTypeFile,
		Size:    data.FileSize,
		Hash:    data.FileHash,
	}.saveMessageContent()

	if err != nil {
		log.Errorln("error saving message content in handle upload file, err:", err)
		return
	}

	jsonContent, err := json.Marshal(&data)
	if err != nil {
		log.Errorln(err)
		return
	}

	for _, roomUser := range roomUsers {
		if roomUser == data.UserID {
			continue
		}

		HubConstruct.sendMessage(jsonContent, roomUser)
	}
}

func (msg messageBytes) handleRequestDownload(author string) {
	file := file{}
	if err := json.Unmarshal(msg, &file); err != nil {
		log.Println(err)
		return
	}

	fileName := file.FileName

	if err := file.retrieveFileInformation(); err != nil {
		file.MsgType = values.DownloadFileErrorMsgType
	}

	file.FileName = fileName

	jsonContent, err := json.Marshal(&file)
	if err != nil {
		log.Println(err)
	}

	HubConstruct.sendMessage(jsonContent, author)
}

func (msg messageBytes) handleFileDownload(author string) {
	file := fileChunks{}
	if err := json.Unmarshal(msg, &file); err != nil {
		log.Errorln("error unmarshalling on handle file download, err:", err)
		return
	}

	fileName := file.FileName

	if err := file.retrieveFileChunk(); err != nil {
		log.Println("error retrieving file", err)
		// Send download file error message to client so as to stop download.
		file = fileChunks{}
		file.MsgType = values.DownloadFileErrorMsgType
	} else {
		file.MsgType = values.DownloadFileChunkMsgType
	}

	file.FileName = fileName

	jsonContent, err := json.Marshal(&file)
	if err != nil {
		log.Println(err)
	}

	HubConstruct.sendMessage(jsonContent, author)
}

// handleSearchUser returns registered users that match searchText.
func handleSearchUser(searchText, user string) {
	data := struct {
		UsersFound interface{} `json:"fetchedUsers"`
		MsgType    string      `json:"msgType"`
	}{
		getUser(searchText, user),
		values.SearchUserMsgType,
	}

	jsonContent, err := json.Marshal(&data)
	if err != nil {
		log.Errorln("error while converting search user result to json", err)
		return
	}

	HubConstruct.sendMessage(jsonContent, user)
}

// handleLoadUserContent loads all users contents on page load.
// All rooms joined and users requests are loaded through WS.
func handleLoadUserContent(email string) {
	userInfo := User{Email: email}
	if err := userInfo.getUser(); err != nil {
		log.Errorln("Could not fetch users room", email)
		return
	}

	usersAssociate, err := userInfo.getAllUsersAssociates()
	if err != nil {
		log.Errorln("error getting users associates in handleLoadUserContent, err:", err)
		return
	}

	HubConstruct.users.mutex.RLock()
	values.MapEmailToName.Mutex.RLock()
	isUserOnline := make(map[string]associateStatus)
	for _, associate := range usersAssociate {
		_, isOnline := HubConstruct.users.users[associate]
		isUserOnline[associate] = associateStatus{Name: values.MapEmailToName.Mapper[associate], IsOnline: isOnline}
	}
	values.MapEmailToName.Mutex.RUnlock()
	HubConstruct.users.mutex.RUnlock()

	request := map[string]interface{}{
		"msgType":      values.WebsocketOpenMsgType,
		"joinedRooms":  userInfo.RoomsJoined,
		"joinRequests": userInfo.JoinRequest,
		"associates":   isUserOnline,
	}

	if data, err := json.Marshal(request); err == nil {
		HubConstruct.sendMessage(data, email)
	}
}

// broadcastOnlineStatusToAllUserRoom broadcasts users availability status to all users joined rooms.
func broadcastOnlineStatusToAllUserRoom(userEmail string, online bool) {
	user := User{Email: userEmail}
	associates, err := user.getAllUsersAssociates()
	if err != nil {
		log.Errorln("could not get users associate", err)
		return
	}

	values.MapEmailToName.Mutex.RLock()
	for _, assassociateEmail := range associates {
		msg := map[string]interface{}{
			"msgType": values.OnlineStatusMsgType,
			"userID":  userEmail,
			"status":  online,
		}

		if data, err := json.Marshal(msg); err == nil {
			HubConstruct.sendMessage(data, assassociateEmail)
		}
	}

	values.MapEmailToName.Mutex.RUnlock()
}
