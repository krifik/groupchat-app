package app

import (
	"os"

	"github.com/pusher/pusher-http-go"
)

func InitializedPusher() {
	pusherClient := pusher.Client{
		AppID:   os.Getenv("PUSHER_APP_ID"),
		Key:     os.Getenv("PUSHER_KEY"),
		Secret:  os.Getenv("PUSHER_SECRET"),
		Cluster: os.Getenv("PUSHER_CLUSTER"),
		Secure:  true,
	}

	data := map[string]string{"message": "hello world"}
	pusherClient.Trigger("groupchat-channel", "message", data)
}
