package models

import (
	"os"
	"sync"
	"time"

	"github.com/at-wat/ebml-go/webm"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/files"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/sharing"
	"github.com/fasthttp/websocket"
	"github.com/pion/webrtc/v2"
	"github.com/pion/webrtc/v2/pkg/media/samplebuilder"
)

type CookieDetail struct {
	Email      string
	Collection string
	CookieName string
	Path       string
	Data       CookieData
}

type CookieData struct {
	ExitTime time.Time
	UUID     string
	Email    string
}

type User struct {
	Email string `bson:"_id" json:"userID"`
	Name  string `bson:"name" json:"name"`
	DOB   string `bson:"age" json:"age"`
	Class string `bson:"class" json:"class"`
	// ID should either be users matric or leading email stripping @....
	Password         []byte        `bson:"password" json:"-"`
	PasswordInString string        `bson:"-" json:"password"`
	RoomsJoined      []RoomsJoined `bson:"roomsJoined" json:"roomsJoined"`
	JoinRequest      []JoinRequest `bson:"joinRequest" json:"joinRequest"`
	UUID             string        `bson:"loginUUID" json:"uuid"`
}

type RoomsJoined struct {
	RoomID   string `bson:"roomID" json:"roomID"`
	RoomName string `bson:"roomName" json:"roomName"`
}

type JoinRequest struct {
	RoomID             string   `bson:"_id" json:"roomID"`
	RoomName           string   `bson:"roomName" json:"roomName"`
	RequestingUserName string   `bson:"requestingUserName" json:"requestingUserName"`
	RequestingUserID   string   `bson:"requestingUserID" json:"requestingUserID"`
	Users              []string `bson:"-" json:"users"` // Users whom join requests are to be sent to, if client wants to send request to multiple users.
}

// Room struct defineds content details for users room.
// Message count is used to track amount of messages sent by users in room,
// this helps with partitioning messages on retrieval, messages are retrieved in 20s on request.
// FirstLoad is specified if user initially wants to retrieve room messages from frontend.
type Room struct {
	RoomID          string    `bson:"_id" json:"roomID,omitempty"`
	RoomName        string    `bson:"roomName" json:"roomName,omitempty"`
	RoomIcon        string    `bson:"roomIcon" json:"roomIcon"`
	RegisteredUsers []string  `bson:"registeredUsers" json:"registeredUsers,omitempty"`
	MessageCount    int       `bson:"messageCount" json:"messageCount,omitempty"`
	Messages        []Message `bson:"-" json:"messages,omitempty"`
	FirstLoad       bool      `bson:"-" json:"firstLoad,omitempty"`
}

type AssociateStatus struct {
	Name     string `json:"name"`
	IsOnline bool   `json:"isOnline"`
}

// Message struct defines user message contents, size and hash is defined if user is sending files.
// Index is used to track message count as to Rooms messages, this should help with partitioning if we
// are to retrieve message of a particular count.
type Message struct {
	RoomID      string `bson:"roomID" json:"roomID,omitempty"`
	Message     string `bson:"message" json:"message,omitempty"`
	UserID      string `bson:"userID" json:"userID,omitempty"`
	Name        string `bson:"name" json:"name,omitempty"`
	Time        string `bson:"time" json:"time,omitempty"`
	Type        string `bson:"type" json:"type,omitempty"`
	Size        string `bson:"size,omitempty" json:"size,omitempty"`
	Hash        string `bson:"hash,omitempty" json:"hash,omitempty"`
	Index       int    `bson:"index" json:"index,omitempty"`
	MessageType string `bson:"-" json:"msgType,omitempty"`
}

type Joined struct {
	RoomID      string `json:"roomID"`
	RoomName    string `json:"roomName"`
	Email       string `json:"userID"`
	Name        string `json:"name"`
	RequesterID string `json:"requesterID"`
	Joined      bool   `json:"joined"`
	MessageType string `bson:"-" json:"msgType"`
}

type NewRoomRequest struct {
	Email       string `json:"userID"`
	RoomName    string `json:"roomName"`
	MessageType string `bson:"-" json:"msgType"`
}

// File save files making sure they are distinct.
type File struct {
	MsgType        string `bson:"-" json:"msgType,omitempty"`
	UniqueFileHash string `bson:"_id" json:"fileHash"`
	FileName       string `bson:"fileName" json:"fileName"`
	User           string `bson:"userID" json:"userID"`
	FileSize       string `bson:"fileSize" json:"fileSize"`
	FileType       string `bson:"fileType" json:"fileType"`
	Chunks         int    `bson:"chunks,omitempty" json:"chunks"`
}

type FileChunks struct {
	MsgType            string `bson:"-" json:"msgType,omitempty"`
	FileName           string `bson:"-" json:"fileName,omitempty"`
	UniqueFileHash     string `bson:"_id" json:"fileHash,omitempty"`
	CompressedFileHash string `bson:"compressedFileHash" json:"compressedFileHash,omitempty"`
	FileBinary         string `bson:"fileChunk" json:"fileChunk,omitempty"`
	ChunkIndex         int    `bson:"chunkIndex" json:"chunkIndex"`
}

type WsMessage struct {
	Data []byte
	User string
}

// Connection is an middleman between the websocket connection and the hub.
type Connection struct {
	Ws *websocket.Conn

	Send chan []byte
}

type Subscription struct {
	Conn *Connection
	User string
}

// Hub maintains the set of active connections and broadcasts messages to the
// connections.
type Hub struct {
	// Registered connections.
	Users WsUsers

	// Inbound messages from the connections.
	Broadcast chan WsMessage

	// Register requests from the connections.
	Register chan Subscription

	// Unregister requests from connections.
	UnRegister chan Subscription
}

type WsUsers struct {
	Users map[string]map[*Connection]bool
	Mutex *sync.RWMutex
}

// classSessionPeerConnections allows class session video calls where a publisher
// broadcasts audio and video to other subscribers. Other subscribers broadcasts
// audio to each other.
// Video is mapped to unique ID which is to be generated by server for current class session
// to video and audio track.
// If UUID to current VIDEO track is nil. It indicates the room session is over.
// If publisher logs off, all peer connections related to that room is closed.
// Mutexes is integrated with video, audio and peerconnections to ensure data race free.
type ClassSessionPeerConnections struct {
	api *webrtc.API

	publisherVideoTracks  map[string]*webrtc.Track // mapped sessionID to track
	publisherTrackMutexes *sync.RWMutex

	audioTrack        map[string]*webrtc.Track // mapped userID to track
	audioTrackSender  map[*webrtc.Track][]RtpSenderData
	audioTrackMutexes *sync.RWMutex

	peerConnection        map[string]*webrtc.PeerConnection // peerConnection is mapped user to peerconnection.
	peerConnectionMutexes *sync.RWMutex

	connectedUsers      map[string][]string // publisher is mapped sessionID to all connected users.
	connectedUsersMutex *sync.RWMutex
}

// rtpSenderData saves user RTPSender.
// On remove track, users can easily map all audio track to its senders.
// userID maps sender to it's peerconnection.
type RtpSenderData struct {
	UserID string
	Sender *webrtc.RTPSender
}

type WebmWriter struct {
	FileName                       string
	AudioWriter, videoWriter       webm.BlockWriteCloser
	AudioBuilder, videoBuilder     *samplebuilder.SampleBuilder
	AudioTimestamp, videoTimestamp uint32
}

type SdpConstruct struct {
	MsgType        string `json:"msgType"`
	ClassSessionID string `json:"sessionID"`
	AuthorName     string `json:"name"`
	PublisherID    string `json:"publisherID"`
	UserID         string `json:"userID"`
	RoomID         string `json:"roomID"`
	SDP            string `json:"sdp"`

	PeerConnection *webrtc.PeerConnection
}

type DropboxUploader struct {
	UploadClient       files.Client
	SharableLinkClient sharing.Client
	FileUploadInfo     *files.CommitInfo

	File         *os.File
	FileFullPath string
	FileSize     int64
}
