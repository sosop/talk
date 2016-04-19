package protocol

type Communicate string

const (
	CONN Communicate = "conn"
	TALK Communicate = "talk"

	// 8M
	MaxDataSize = 1024 * 1024 * 8

	TOUSER = iota
	TOGROUP
)

type (
	Protocol struct {
		From string
		To   string
		Comm Communicate
		Type int
		Data
	}
	Data struct {
		Text   []byte
		Images map[string][]byte
	}
)
