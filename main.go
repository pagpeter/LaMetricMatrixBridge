package main

import (
	"log"
	"time"

	"maunium.net/go/mautrix"
)

func main() {
	var c Config
	c.getConf()
	// fmt.Println(c)
	if c.Blacklist.Active && c.Whitelist.Active {
		log.Fatal("Only one of blacklist or whitelist can be active")
	}

	log.Println("Getting IP")
	IP := GetIPAddress()
	if IP == "" {
		log.Println("Could not get IP, trying again in 5 seconds")
		time.Sleep(time.Second * 5)
		IP = GetIPAddress()
		if IP == "" {
			log.Fatal("Could not get IP")
		}
	}
	log.Println("IP of the LaMetic Time:", IP)

	l := LaMetric{
		IP: IP,
	}
	l.Api_Key = l.EncodeApiKey(c.Api_Key)

	log.Println("Logging into", c.Homeserver, "as", c.Username)

	client, err := mautrix.NewClient(c.Homeserver, "", "")
	if err != nil {
		panic(err)
	}
	_, err = client.Login(&mautrix.ReqLogin{
		Type:             "m.login.password",
		Identifier:       mautrix.UserIdentifier{Type: mautrix.IdentifierTypeUser, User: c.Username},
		Password:         c.Password,
		StoreCredentials: true,
	})
	if err != nil {
		panic(err)
	}
	log.Println("Login successful")

	bridge_messages(l, client, c)
}
