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
	if max >= len(s) {
		return s
	}
	return s[:max] + "..."
	// return s[:strings.LastIndexAny(s[:max], " .,:;-")] + "..."
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
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
		time.Sleep(time.Second * 3)
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

		if c.Blacklist.Active {
			if contains(c.Blacklist.Rooms, fmt.Sprintf("%v", evt.RoomID)) {
				// log.Println("Message in blacklisted room", evt.RoomID)
				return
			}
		}
		if c.Whitelist.Active {
			if !contains(c.Whitelist.Rooms, fmt.Sprintf("%v", evt.RoomID)) {
				// log.Println("Message not in whitelisted room", evt.RoomID)
				return
			}
		}

		log.Printf("<%[1]s> %[4]s (%[2]s/%[3]s)\n", evt.Sender, evt.Type.String(), evt.ID, evt.Content.AsMessage().Body)

		msg := strings.TrimSpace(evt.Content.AsMessage().Body)
		text := truncateText(msg, 10)
		parts := strings.Split(fmt.Sprintf("%v", evt.Sender), ":")
		fmt.Println(parts)
		sender := fmt.Sprintf("%v", evt.Sender)
		// fmt.Println("0")

		if len(parts) >= 1 {
			sender = parts[0]
			sender = strings.ReplaceAll(sender, "@", "")
		}
		// fmt.Println("1")

		toSend := fmt.Sprint(sender+": ", text)
		n := get_notification(toSend)
		// fmt.Println("2", toSend, n)
		// fmt.Println(n)
		_, err := l.SendNotification(n)
		if err != nil {
			log.Println("Error sending notification:", err)
		}
		// fmt.Println("3", res, err)
		log.Println("Sent notification:", toSend)
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
