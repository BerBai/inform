package bark

import (
	"context"
	"testing"
)

const (
	deviceKey = "your bark key code"
)

func TestBark(t *testing.T) {
	barkService := NewWithServers(deviceKey, DefaultServerURL)

	// Or use `New` to create a service with the default server.
	// barkService = New(deviceKey)
	err := barkService.Send(context.Background(), "title", "content", WithURL("https://github.com/berbai/inform"))
	if err != nil {
		print(err.Error())
	}
}
