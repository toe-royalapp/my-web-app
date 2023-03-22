package main

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/toe-royalapp/my-web-app/config"
)

func main() {
	app, _, _ := config.SetupFirebase()
	sendToToken(app)
}

func sendToToken(app *firebase.App) {
	ctx := context.Background()
	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting messaging client: %v\n", err)
	}
	registrationToken := "eHebYoXYCTytsAzWFT65G3:APA91bHcTkdqqI9e3UlZ-4ZKqtlwrV0WS6Qd8SFCuFf8fw4wz_ZZ18oVPnVhduroyt322q5cPHm4uLvwY27eLso_2uQvNDNK7q6cEZo0tC4ZSZaHt_8DosDBTxi8_VvHPwSgfumh1Y-3"
	message := &messaging.Message{
		Data: map[string]string{},
		Notification: &messaging.Notification{
			Title:    "My Testing Go Firebase Messaging from the go package",
			Body:     "This is the testing go firebase messaging application",
			ImageURL: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSiNxAxVDejAH3t6abHvKLHWedw4KAoM65j1Zr5SArSqg&s",
		},
		Token: registrationToken,
	}
	response, err := client.Send(ctx, message)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("successfully send meesage : ", response)
}
