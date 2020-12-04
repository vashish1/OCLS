package requests

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/vashish1/OnlineClassPortal/pkg/database"
	"github.com/vashish1/OnlineClassPortal/pkg/models"
	"github.com/vashish1/OnlineClassPortal/vendor/go.mongodb.org/mongo-driver/bson"
	"github.com/vashish1/OnlineClassPortal/vendor/go.mongodb.org/mongo-driver/mongo/options"
	"github.com/vashish1/OnlineClassPortal/vendor/golang.org/x/crypto/bcrypt"
	"go.mongodb.org/mongo-driver/mongo"
)

func updateRoomsJoinedByUsers(email string,roomID, roomName string) error {
	if err := getUser(email); err != nil {
		return err
	}

	var roomJoined = models.RoomsJoined{RoomID: roomID, RoomName: roomName}
	b.RoomsJoined = append(b.RoomsJoined, roomJoined)

	_, err := db.Collection(values.UsersCollectionName).UpdateOne(ctx, bson.M{"_id": b.Email},
		bson.M{"$set": bson.M{"roomsJoined": b.RoomsJoined}})

	return err
}

func (b *User) getAllUsersAssociates() ([]string, error) {
	if err := b.getUser(); err != nil {
		return nil, err
	}

	usersChannel := make(chan []string)
	done := make(chan struct{})
	users := make([]string, 0)
	registeredUser := make(map[string]bool)

	go func() {
		for {
			data, ok := <-usersChannel
			if ok {
				for _, user := range data {
					if _, exist := registeredUser[user]; !exist && user != b.Email {
						users = append(users, user)
						registeredUser[user] = true
					}
				}

				continue
			}

			close(done)
			break
		}
	}()

	for _, roomJoined := range b.RoomsJoined {
		var room room
		result := db.Collection(values.RoomsCollectionName).FindOne(ctx, bson.M{
			"_id": roomJoined.RoomID,
		})

		if err := result.Decode(&room); err != nil {
			close(usersChannel)
			<-done
			return nil, err
		}

		usersChannel <- room.RegisteredUsers
	}

	close(usersChannel)
	<-done

	return users, nil
}

func (b User) exitRoom(roomID string) ([]string, error) {
	if err := b.getUser(); err != nil {
		return nil, err
	}

	// Confirm if indeed user is registered to room.
	var roomExist bool
	for i, roomJoined := range b.RoomsJoined {
		if roomJoined.RoomID == roomID {
			roomExist = true

			if len(b.RoomsJoined)-1 > i {
				b.RoomsJoined = append(b.RoomsJoined[:i], b.RoomsJoined[i+1:]...)
			} else {
				b.RoomsJoined = b.RoomsJoined[:i]
			}
			break
		}
	}
	if !roomExist {
		return nil, values.ErrUserNotRegisteredToRoom
	}

	// Update room joined by user in DB.
	_, err := db.Collection(values.UsersCollectionName).UpdateOne(ctx, bson.M{"_id": b.Email},
		bson.M{"$set": bson.M{"roomsJoined": b.RoomsJoined}})
	if err != nil {
		return nil, err
	}

	room := room{RoomID: roomID}
	result := db.Collection(values.RoomsCollectionName).FindOne(ctx, bson.M{"_id": room.RoomID})

	if err := result.Decode(&room); err != nil {
		return nil, err
	}

	_, err = Message{
		UserID:  b.Email,
		RoomID:  room.RoomID,
		Type:    values.MessageTypeInfo,
		Message: b.Email + " Left the room",
	}.saveMessageContent()

	if err != nil {
		return nil, err
	}

	for i, user := range room.RegisteredUsers {
		if user == b.Email {
			if len(room.RegisteredUsers)-1 > i {
				room.RegisteredUsers = append(room.RegisteredUsers[:i], room.RegisteredUsers[i+1:]...)
			} else {
				room.RegisteredUsers = room.RegisteredUsers[:i]
			}

			break
		}
	}

	_, err = db.Collection(values.RoomsCollectionName).UpdateOne(ctx, bson.M{
		"_id": room.RoomID,
	}, bson.M{"$set": bson.M{"registeredUsers": room.RegisteredUsers}})

	return room.RegisteredUsers, err
}

func (b User) CreateUserLogin(password string, w http.ResponseWriter) (string, error) {
	if err := b.getUser(); err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword(b.Password, []byte(password)); err != nil {
		return "", err
	}

	token, err := CookieDetail{
		Email:      b.Email,
		Collection: values.UsersCollectionName,
		CookieName: values.UserCookieName,
		Path:       "/",
		Data: CookieData{
			Email: b.Email,
		},
	}.GenerateCookie(w)

	return token, err
}

func (b User) validateUser(uniqueID string) error {
	if err := b.getUser(); err != nil {
		return err
	}

	if b.UUID != uniqueID {
		return values.ErrIncorrectUUID
	}
	return nil
}

// saveMessageContent saves users messages to database.
// Messages are stored using individual Insertion so that all message in room can be retrieved in partition.
func (b Message) saveMessageContent() ([]string, error) {
	var roomDetails room

	opts := options.FindOneAndUpdate().SetProjection(bson.M{"roomName": 0}).
		SetReturnDocument(options.After)

	result := db.Collection(values.RoomsCollectionName).
		FindOneAndUpdate(ctx, bson.M{"_id": b.RoomID}, bson.M{"$inc": bson.M{"messageCount": 1}}, opts)

	if err := result.Decode(&roomDetails); err != nil {
		return nil, err
	}

	b.Index = roomDetails.MessageCount

	var userExists bool
	for _, user := range roomDetails.RegisteredUsers {
		if b.UserID == user {
			userExists = true
			break
		}
	}

	if !userExists {
		return nil, values.ErrInvalidUser
	}

	if _, err := db.Collection(values.MessageCollectionName).InsertOne(ctx, b); err != nil {
		return nil, err
	}

	return roomDetails.RegisteredUsers, nil
}

// getPartitionedMessageInRoom retrieves messages for a particular room in the DB.
// Retrieved messages are collected in partitions of 20s per room messages. First message index received by client
// is to be sent by client so that the next recurring messages are fetched from database. If total message count in DB is say 30
// It is assumed that if the client has retrieved messages from index 30-50 where total message from room is 50
// and want to fetch next count messages, mmessages from  index 10 to 30 is retrieved.
// If FirstLoad is indicated, last 20 message count is retrieved.
func (b *room) getPartitionedMessageInRoom() error {
	if b.FirstLoad {
		messsageCountResult := db.Collection(values.RoomsCollectionName).
			FindOne(ctx, bson.M{"_id": b.RoomID})

		if err := messsageCountResult.Decode(&b); err != nil {
			return err
		}
	}

	var messages []Message

	result, err := db.Collection(values.MessageCollectionName).
		Find(ctx, bson.M{"roomID": b.RoomID, "index": bson.M{"$gt": b.MessageCount - 19, "$lt": b.MessageCount + 1}}, options.Find().SetProjection(bson.M{"roomID": 0}))

	if err != nil {
		return err
	}

	if err := result.All(ctx, &messages); err != nil {
		return err
	}

	b.Messages = messages
	return nil
}

func createNewRoom(b models.NewRoomRequest) (string, error) {
	var chats models.Room

	chats.RoomID = uuid.New().String()
	chats.RoomName = b.RoomName
	chats.RegisteredUsers = append(chats.RegisteredUsers, b.Email)

	message := models.Message{
		RoomID:  chats.RoomID,
		Message: b.Email + " Joined",
		Type:    models.MessageTypeInfo,
	}

	if ok := database.InsertIntoDb(database.RoomDb(), chats); !ok {
		return "", errors.New("chats were not inserted in DB")
	}

	if ok := database.InsertIntoDb(database.MessageDb(), message); !ok {
		return "", errors.New("Messages were not inserted in DB")
	}

	
	if err := updateRoomsJoinedByUsers(b.Email,chats.RoomID, chats.RoomName); err != nil {
		return "", err
	}

	return chats.RoomID, nil
}

// acceptRoomRequest accept room join request from a requesting user.
func (b joined) acceptRoomRequest() ([]string, error) {
	result := db.Collection(values.UsersCollectionName).FindOne(ctx, bson.M{
		"_id": b.Email,
	})

	var user User
	err := result.Decode(&user)
	if err != nil {
		return nil, err
	}

	var joinRequestLegit bool
	// Confirm if join request was sent to user and saved to DB.
	for i, request := range user.JoinRequest {
		if request.RoomID == b.RoomID {
			joinRequestLegit = true
			user.JoinRequest = append(user.JoinRequest[:i], user.JoinRequest[i+1:]...)
			break
		}
	}

	if !joinRequestLegit {
		return nil, values.ErrIllicitJoinRequest
	}

	var messages room
	result = db.Collection(values.RoomsCollectionName).FindOne(ctx, bson.M{
		"_id": b.RoomID,
	})

	if err := result.Decode(&messages); err != nil {
		return nil, err
	}

	// Room request is declined.
	if !b.Joined {
		message := Message{
			UserID:  b.RequesterID,
			RoomID:  b.RoomID,
			Message: b.Email + "Refused join request.",
			Type:    values.MessageTypeInfo,
		}

		if _, err := message.saveMessageContent(); err != nil {
			return nil, err
		}

		return messages.RegisteredUsers, nil
	}

	message := Message{
		UserID:  b.RequesterID,
		RoomID:  b.RoomID,
		Message: b.Email + " Accepted join request.",
		Type:    values.MessageTypeInfo,
	}

	if _, err := message.saveMessageContent(); err != nil {
		return nil, err
	}

	messages.RegisteredUsers = append(messages.RegisteredUsers, b.Email)

	_, err = db.Collection(values.RoomsCollectionName).UpdateOne(ctx, bson.M{
		"_id": b.RoomID,
	}, bson.M{"$set": bson.M{"registeredUsers": messages.RegisteredUsers}})

	// Save rooms joined by user to user collection.
	if b.Joined {
		user.RoomsJoined = append(user.RoomsJoined, roomsJoined{RoomID: b.RoomID, RoomName: b.RoomName})

		_, err = db.Collection(values.UsersCollectionName).UpdateOne(ctx, bson.M{"_id": b.Email},
			bson.M{"$set": bson.M{"joinRequest": user.JoinRequest, "roomsJoined": user.RoomsJoined}})

		if err != nil {
			return nil, err
		}
	}

	return messages.RegisteredUsers, err
}

// requestUserToJoinRoom confirms sends join request to user.
func (b joinRequest) requestUserToJoinRoom(userToJoinEmail string) ([]string, error) {
	var room room
	result := db.Collection(values.RoomsCollectionName).FindOne(ctx, bson.M{"_id": b.RoomID})

	if err := result.Decode(&room); err != nil {
		return nil, err
	}

	// Confirm if person making the request is part of the room.
	var requesterLegit bool
	for _, registeredUser := range room.RegisteredUsers {
		if registeredUser == b.RequestingUserID {
			requesterLegit = true
			break

		} else if registeredUser == userToJoinEmail {
			return nil, values.ErrUserExistInRoom
		}
	}

	if !requesterLegit {
		return nil, errors.New("invalid user made a RequestUsersToJoinRoom request Name: " + b.RequestingUserID)
	}

	result = db.Collection(values.UsersCollectionName).FindOne(ctx, bson.M{"_id": userToJoinEmail})
	var user User

	if err := result.Decode(&user); err != nil {
		return nil, err
	}

	// Return error if user is already requested.
	for _, request := range user.JoinRequest {
		if b.RoomID == request.RoomID {
			return nil, values.ErrUserAlreadyRequested
		}
	}

	user.JoinRequest = append(user.JoinRequest, b)

	_, err := db.Collection(values.UsersCollectionName).UpdateOne(ctx, bson.M{"_id": userToJoinEmail},
		bson.M{"$set": bson.M{"joinRequest": user.JoinRequest}})

	if err != nil {
		return nil, err
	}

	_, err = Message{
		RoomID:  b.RoomID,
		UserID:  b.RequestingUserID,
		Message: fmt.Sprintf("%s was requested to join the room by %s", userToJoinEmail, b.RequestingUserID),
		Type:    values.MessageTypeInfo,
	}.saveMessageContent()

	return room.RegisteredUsers, err
}

// UploadNewFile create a NewFile content to database and returns file content if one
// has already been created.
// Chunks is set to zero so that if user wants to retrieve
func (b *file) uploadNewFile() error {
	result := db.Collection(values.FilesCollectionName).FindOne(ctx, bson.M{"_id": b.UniqueFileHash}) //, b, options.FindOneAndReplace().SetUpsert(true))

	if result.Err() == mongo.ErrNoDocuments {
		_, err := db.Collection(values.FilesCollectionName).InsertOne(ctx, b)
		return err
	}

	if err := result.Decode(&b); err != nil {
		return err
	}

	return nil
}

func (b *file) retrieveFileInformation() error {
	result := db.Collection(values.FilesCollectionName).FindOne(ctx, bson.M{"_id": b.UniqueFileHash})
	return result.Decode(&b)
}

func (b fileChunks) fileChunkExists() bool {
	result := db.Collection(values.FileChunksCollectionName).FindOne(ctx, bson.M{"_id": b.UniqueFileHash})
	if err := result.Err(); err == nil {
		return true
	}
	return false
}

func (b fileChunks) addFileChunk() error {
	result := db.Collection(values.FileChunksCollectionName).
		FindOneAndReplace(ctx, bson.M{"_id": b.UniqueFileHash}, b, options.FindOneAndReplace().SetUpsert(true))

	// Update original file index.
	if err := result.Err(); err == nil || err == mongo.ErrNoDocuments {
		_, err := db.Collection(values.FilesCollectionName).UpdateOne(ctx,
			bson.M{"_id": b.CompressedFileHash}, bson.M{"$set": bson.M{"chunks": b.ChunkIndex}})
		return err
	}

	return result.Err()
}

func (b *fileChunks) retrieveFileChunk() error {
	result := db.Collection(values.FileChunksCollectionName).
		FindOne(ctx, bson.M{"compressedFileHash": b.CompressedFileHash, "chunkIndex": b.ChunkIndex})

	return result.Decode(&b)
}

func uploadFileGridFS(fileName string) error {
	fileBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Errorln("unable read file while uploading", err)
		return err
	}

	buc, err := gridfs.NewBucket(db)
	if err != nil {
		log.Errorln("unable GridFS bucket", err)
		return err
	}

	up, err := buc.OpenUploadStream("hhh")
	if err != nil {
		log.Errorln("unable to open upload stream", err)
		return err
	}
	defer up.Close()

	_, err = up.Write(fileBytes)
	if err != nil {
		log.Errorln("unable to write to bucket stream", err)
		return err
	}

	return nil
}

func getUser(key string, user string) interface{} {
	names := make([]struct {
		Name  string `json:"name"`
		Email string `json:"userID"`
	}, 0)

	values.MapEmailToName.Mutex.RLock()
	for email, name := range values.MapEmailToName.Mapper {
		if email == "" || email == user {
			continue
		}

		if strings.Contains(email, key) {
			names = append(names, struct {
				Name  string `json:"name"`
				Email string `json:"userID"`
			}{
				name, email,
			})
		}
	}
	values.MapEmailToName.Mutex.RUnlock()

	return names
}
