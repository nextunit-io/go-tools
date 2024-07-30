package tools

import "time"

type Time interface {
	Now() time.Time
}

type defaultTimeClient struct {
	Time
}

var timeInstane Time

func GetTimeInstance() Time {
	if timeInstane == nil {
		timeInstane = &defaultTimeClient{}
	}

	return timeInstane
}

func SetTimeInstance(client Time) {
	timeInstane = client
}

func (defaultTimeClient) Now() time.Time {
	return time.Now()
}
