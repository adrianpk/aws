package async

import "time"

type Envelope struct {
	Key   string
	Type  string
	Body  []byte
	Attrs map[string]string
	Delay time.Duration
}
