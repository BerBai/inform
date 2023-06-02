package bark

import (
	"context"
	"testing"
)

const (
	deviceKey = "4SZWHSrYXLQw5Z6GAvUZ2V"
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
