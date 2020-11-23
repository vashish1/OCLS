package models

import (
	"errors"
	"log"
	"sync"
	"time"

	"github.com/pion/webrtc/v2"
)

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

const (
	AdminCollectionName      = "administrators"
	UsersCollectionName      = "users"
	RoomsCollectionName      = "rooms"
	MessageCollectionName    = "messages"
	FilesCollectionName      = "files"
	FileChunksCollectionName = "fileChunks"

	AdminCookieName = "Admin"
	UserCookieName  = "User"
	TimeLayout      = "Monday, 02-Jan-06"
)

const (
	DefaultCost     = 10
	RtcpPLIInterval = time.Second * 3
)

var (
	ErrIncorrectUUID           = errors.New("incorrect UUID")
	ErrInvalidExpiryTime       = errors.New("invalid expiry time")
	ErrCookieExpired           = errors.New("generated cookie has expired")
	ErrInvalidUser             = errors.New("invalid user")
	ErrInvalidDetails          = errors.New("invalid signin details")
	ErrRetrieveUUID            = errors.New("could not retrieve UUID")
	ErrMarshal                 = errors.New("could not marshal content")
	ErrWrite                   = errors.New("error while sending content")
	ErrAuthentication          = errors.New("Authentication error")
	ErrIllicitJoinRequest      = errors.New("User was not originally requested to join")
	ErrUserExistInRoom         = errors.New("User already exist in room")
	ErrUserAlreadyRequested    = errors.New("User already requested to join room")
	ErrUserNotRegisteredToRoom = errors.New("User was not registered to room")
	ErrFileUpload              = errors.New("error while uploading file to server")
	ErrPeerConnectionNotFound  = errors.New("PeerConnection not found")
	ErrFileUploadLink          = errors.New("could not generate file upload link")
	ErrTokenNotSpecified       = errors.New("Token was not specified")
	ErrInvalidToken            = errors.New("Token specified is invalid")
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
