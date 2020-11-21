package values

import "errors"

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
