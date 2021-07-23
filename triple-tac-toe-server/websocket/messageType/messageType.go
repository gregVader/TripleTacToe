package messageType

type MessageType string

const (
	Wait          MessageType = "WAIT"
	UsernameTaken MessageType = "USERNAME_TAKEN"
	StartGame     MessageType = "START_GAME"
	Left          MessageType = "LEFT"
)

func (messageType MessageType) String() string {
	return string(messageType)
}
