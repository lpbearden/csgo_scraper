package services

import (
    "log"

    "github.com/dghubble/go-twitter/twitter"
    "github.com/dghubble/oauth1"
)

func GetTwitterClient() *twitter.Client {
    consumerKey := API_KEY
    consumerSecret := API_SECRET_KEY
    accessToken := ACCESS_TOKEN
    accessSecret := ACCESS_TOKEN_SECRET

    if consumerKey == "" || consumerSecret == "" || accessToken == "" || accessSecret == "" {
        log.Fatal("Consumer key/secret and Access token/secret required")
    }

    config := oauth1.NewConfig(consumerKey, consumerSecret)
    token := oauth1.NewToken(accessToken, accessSecret)
    // OAuth1 http.Client will automatically authorize Requests
    httpClient := config.Client(oauth1.NoContext, token)

    // Twitter client
    client := twitter.NewClient(httpClient)
    return client
}