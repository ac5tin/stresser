package task

import "github.com/BurntSushi/toml"

// Task denotes a single executable HTTP request task
type Task struct {
	URL                 string
	Headers             map[string]string
	Method              string
	Timeout             uint32
	Payload             payload
	AcceptedStatusCodes []uint32
}

type payload []byte

func (p *payload) UnmarshalText(text []byte) error {
	*p = payload(text)
	return nil
}

// ParseFile parses Task definition from a TOML file
func ParseFile(filePath string) (*Task, error) {
	t := Task{}
	if _, err := toml.DecodeFile(filePath, &t); err != nil {
		return nil, err
	}
	return &t, nil
}
