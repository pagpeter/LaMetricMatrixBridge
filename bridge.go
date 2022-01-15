package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"maunium.net/go/mautrix"
	"maunium.net/go/mautrix/event"
)

func truncateText(s string, max int) string {
	if max > len(s) {
		return s
	}
	return s[:strings.LastIndexAny(s[:max], " .,:;-")] + "..."
}

func get_notification(text string) Notification {
	return Notification{
		Priority: "info",
		IconType: "info",
		Model: Model{
			Frames: []Frames{
				Frames{
					Text: text,
				},
			},
			Sound: Sound{
				Category: "notifications",
				ID:       "notification2",
			},
			Cycles: 1,
		},
	}
}

func bridge_messages(l LaMetric, client *mautrix.Client, c Config) {
	active := false

	syncer := client.Syncer.(*mautrix.DefaultSyncer)

	go func() {
		time.Sleep(time.Second * 5)
		active = true
		log.Println("Starting to listen for messages")
	}()

	syncer.OnEventType(event.EventMessage, func(source mautrix.EventSource, evt *event.Event) {
		if !active {
			return
		}
		if evt.Sender == client.UserID {
			log.Println("Send message")
			return
		}
		log.Printf("<%[1]s> %[4]s (%[2]s/%[3]s)\n", evt.Sender, evt.Type.String(), evt.ID, evt.Content.AsMessage().Body)

		msg := strings.TrimSpace(evt.Content.AsMessage().Body)
		text := truncateText(msg, 10)
		sender := strings.Split(string(evt.Sender), ":")[0]
		toSend := fmt.Sprint(sender+": ", text)
		n := get_notification(toSend)
		go l.SendNotification(n)
		// l.SendNotification(get_notification(fmt.Sprintf("Message from %v", evt.Sender)))
	})

	err := client.Sync()
	if err != nil {
		log.Println(err)
	}
	// syncer := client.Syncer.(*mautrix.DefaultSyncer)
	// syncer.OnEventType(event.EventMessage, func(source mautrix.EventSource, evt *event.Event) {
	// 	// fmt.Printf("<%[1]s> %[4]s (%[2]s/%[3]s)\n", evt.Sender, evt.Type.String(), evt.ID, evt.Content.AsMessage().Body)
	// 	// l.SendNotification(get_notification("Message from " + evt.Sender))
	// 	fmt.Println("Sending notification")
	// 	//  + ": " + evt.Content.AsMessage().Body
	// })
}
