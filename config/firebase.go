package config

import (
	"context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

func SetupFirebase() (*firebase.App, context.Context, *messaging.Client) {
	ctx := context.Background()
	// serviceAccountKeyFilePath, err := filepath.Abs("./serviceAccountKey.json")
	// if err != nil {
	// 	panic("Unable to load serviceAccountKeys.json file")
	// }

	// opt := option.WithCredentialsFile(serviceAccountKeyFilePath)

	opt := option.WithCredentialsJSON([]byte(`{
		"type": "service_account",
		"project_id": "test-msg-f5587",
		"private_key_id": "8d5f3aeb580d827edbff3a2bd54ecf30e79c242e",
		"private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCo2jn01LEMGQXC\nuzcXFblduuCDyJXhHm6oy7gelvitEAgZSf2TZjMOYmlBq5rbO+BxuvkwHKXL1jfb\nDTEwyQGfA4tGtFM8+78siWZKYNQ9t4yvpjTzx8flM6QQ8rYs2B74+cfFe0/FJUP0\ntsc9TRF5XKjzjecPE9uL3uJXPmjEAmY48gT6Ha7GElxWj8CvIaH7bLywHAXl6LGK\nXi4THrGeBm4bGvSEfNUK8ijs1wkWfXwSxVMht0EGdXComEhRZi/tsF9fgr0g/9//\nZ6ZRvG+agTv0KixbSTK7WjP3b+QimLJ+imXeyAf5OEa1k+ZiDnC7dgBUUQkxwsAL\nKcE+esIPAgMBAAECggEAIMcMDRzkqYaJ+b1B+nc7HpjSUIK1pZe0v4ucWEPSTjYY\nfGBCm0RxkM9Fw3u+eeRrs2gqS06kWKdi/GDZNbEvYWDcvoLuzZ3JBKo/cxzM1pY0\n9zv6d4BGq5WUub2D47yodoh2YT2IxKpxQKXik2HMjJSSL4VSXzjpm4NOyMTSrYnA\n5qunvVxlZNy4gBKcnVrVnHaXConuC0dglRjvglEbKJZSEYWN5HJLjcbML/lnROFo\niBvEAI4JWz7u71h6gwQEGOT38Up9tEnffFC1BySH+i0uvXtTZltkCM5i0+EjIdrx\nAUCHkqWgyrLL+mz2VOvKHAFv/SoYeW8BXSSACmfIEQKBgQDZGSxYvfsmDbfxrjno\nuRC3mUrTVgx+mKYrNT+qTGS15uc5Fr87U/2wqfZNcMu3DNlxOy2T0Pq6fdjnZ0Vd\nJxPs91r5SytJstDYeoYrbC4k5wjEbKTxuDzLC6cA+iDd0e0w8sp6sWDoYshrE7OP\n1PGYcgiC7afFjkTePMkAU/3K4wKBgQDHG+W9cH3BPUfqqWZKQiQDOpE9ZsI0p2u9\nGMLS6dkkxYpJWBN0aQmV8xEl15xCQXzQqEIkuhkCyP9+3cTbE9BJi0zPRnp8BvaR\n4OCl7AECKjYA6cxIz0MXJtFFDePzs6v4v5hM/WgFI9OnDCUkaHJiscCi+lwKdamx\nbaveHaS35QKBgQCYa3ATLe7yHJeUern1hlkSTfGWxscm6o1fsJbuPYxHmcGk9y9z\nu9hU/D8Mx7B+5+qR8PZi3UnrPQfYD01HxXPb3x5kAD54E69FDWC9g8ox5nLlVVHI\n04z7EUdDDFme4xAgWZWG5pofDZugcciTpvGoEgefxLMRpHz9ere0H0QD+QKBgQC/\nqA7g5yNlmwNa0mzwFfJsWTftDrjpzi5a+zatpwOp7axLJUi0yVa8zBg5gO2cqUTn\n5M6mY6wnjirBh98xskTQRhJNgt3r/RfjG7+idyYRW2hYQLSvcTJ3WpoeClzd/JkQ\n+/wyP8qdB/t31Kz7+r4AWnG/b5ahStlpPNQvHgGsMQKBgEvlFrzm+tjW/Mzejg1l\nyBky923QoGCP4KbM5eXXDipy27bJ4BE/68JLRm2eZfcBgxy9oYoD3h6LbfEDegkV\nFG7O1rEoFJHHVofRLHOwPQfkxf3X1ha77cKw5AgsDjNUEdy6t4eR4041F9FeBrj8\nxtZgmKopOTZfXigNxN1ZOfiH\n-----END PRIVATE KEY-----\n",
		"client_email": "firebase-adminsdk-wg42m@test-msg-f5587.iam.gserviceaccount.com",
		"client_id": "101693710149459905814",
		"auth_uri": "https://accounts.google.com/o/oauth2/auth",
		"token_uri": "https://oauth2.googleapis.com/token",
		"auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
		"client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/firebase-adminsdk-wg42m%40test-msg-f5587.iam.gserviceaccount.com"
	  }`))

	//Firebase admin SDK initialization
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic("Firebase load error")
	}

	//Messaging client
	client, _ := app.Messaging(ctx)

	return app, ctx, client
}
