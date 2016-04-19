package protocol

import (
	"encoding/json"
)

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

func Serializer(p Protocol) []byte {
	data, err := json.Marshal(p)
	if err != nil {
		// TODO log
		return nil
	}
	return data
}

func UnSerializer(data []byte) Protocol {
	p := Protocol{}
	err := json.Unmarshal(data, p)
	if err != nil {
		// TODO log
	}
	return p
}
