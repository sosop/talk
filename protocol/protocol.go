package protocol

import "encoding/json"

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
		From string      `json:"from"`
		To   string      `json:"to"`
		Comm Communicate `json:"comm"`
		Type int         `json:"type"`
		Data `json:"data"`
	}
	Data struct {
		Text   []byte            `json:"text"`
		Images map[string][]byte `json:"images"`
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
	err := json.Unmarshal(data, &p)
	if err != nil {
		// TODO log
	}
	return p
}
