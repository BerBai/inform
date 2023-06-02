```go
	package main

	import (
	    "context"
		"github.com/berbai/inform/bark"
	)

	func main() {
	    // Create a bark service. `device key` is generated when you install the application. You can use the
	    // `bark.NewWithServers` function to create a service with a custom server.
	    barkService := bark.NewWithServers("your bark device key", bark.DefaultServerURL)

	    // Or use `bark.New` to create a service with the default server.
	    barkService = bark.New("your bark device key")

	    // Send a test message.
	    err := barkService.Send(
	        context.Background(),
	        "Subject/Title",
	        "The actual message - Hello, you awesome gophers! :)",
	    )
		if err != nil {
			print(err.Error())
        }
	}
```