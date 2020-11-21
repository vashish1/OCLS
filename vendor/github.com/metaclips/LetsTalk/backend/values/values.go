package values

import (
	"log"
	"sync"

	"github.com/pion/webrtc/v2"
)

// Message type supported by LetsTalk
const (
	MessageTypeFile             = "File"
	MessageTypeMessage          = "Message"
	MessageTypeInfo             = "Info"
	MessageTypeClassSession     = "ClassSession"
	MessageTypeClassSessionLink = "MessageTypeClassSessionLink"
)

// All websocket message types both clients and server
const (
	UnauthorizedAcces             = "UnauthorizedAccess"
	NewFileUploadMsgType          = "NewFileUpload"
	NewMessageMsgType             = "NewMessage"
	RequestMessages               = "RequestMessages"
	SearchUserMsgType             = "SearchUser"
	WebsocketOpenMsgType          = "WebsocketOpen"
	CreateRoomMsgType             = "CreateRoom"
	ExitRoomMsgType               = "ExitRoom"
	RequestUsersToJoinRoomMsgType = "RequestUsersToJoinRoom"
	SentRoomRequestMsgType        = "SentRoomRequest"
	JoinRoomMsgType               = "JoinRoom"
	OnlineStatusMsgType           = "OnlineStatus"

	UploadFileErrorMsgType   = "UploadFileError" // UploadFileErrorMsgType is sent to client only.
	UploadFileSuccessMsgType = "FileUploadSuccess"
	UploadFileChunkMsgType   = "UploadFileChunk"

	RequestDownloadMsgType     = "RequestDownload"
	DownloadFileChunkMsgType   = "DownloadFileChunk"
	DownloadFileErrorMsgType   = "DownloadFileError"   // DownloadFileErrorMsgType is sent to client only.
	DownloadFileSuccessMsgType = "DownloadFileSuccess" // DownloadFileSuccessMsgType is sent to client only.

	StartClassSession = "StartClassSession"
	JoinClassSession  = "JoinClassSession"
	NegotiateSDP      = "NegotiateSDP"
	RenegotiateSDP    = "RenegotiateSDP"
	ClassSessionError = "ClassSessionError"
	EndClassSession   = "EndClassSession"
)

var (
	// MapEmailToName retrieve user registered name if given an email.
	// MapEmailToName is concurrent safe as it uses a mutex lock.
	MapEmailToName = struct {
		Mapper map[string]string
		Mutex  *sync.RWMutex
	}{
		make(map[string]string),
		&sync.RWMutex{},
	}

	// PeerConnectionConfig contains peerconnection configuration
	PeerConnectionConfig = webrtc.Configuration{
		SDPSemantics: webrtc.SDPSemanticsUnifiedPlanWithFallback,
	}

	credetialType map[string]webrtc.ICECredentialType = map[string]webrtc.ICECredentialType{
		"Password": webrtc.ICECredentialTypePassword,
		"Oauth":    webrtc.ICECredentialTypePassword,
	}
)

func initIceServers() {
	if len(Config.ICEServers) == 0 {
		PeerConnectionConfig.ICEServers = []webrtc.ICEServer{
			{URLs: []string{"stun:stun.l.google.com:19302"}},
		}

		return
	}

	for _, config := range Config.ICEServers {
		if len(config.URLs) == 0 {
			log.Fatalln("User did not specify ICE server.")
		}

		credential, ok := credetialType[config.AuthType]
		if !ok {
			log.Fatalln("Invalid webrtc credential type", config.AuthType, "only AuthType Password and Oauth are allowed.")
		}

		iceServer := webrtc.ICEServer{
			URLs:           config.URLs,
			Username:       config.Username,
			Credential:     config.AuthSecret,
			CredentialType: credential,
		}

		PeerConnectionConfig.ICEServers = append(PeerConnectionConfig.ICEServers, iceServer)
	}
}
