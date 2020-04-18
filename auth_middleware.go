package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
)

func AuthorizationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("We are accepting.")

		for header := range r.Header {
			fmt.Printf("%s: %s\n", header, r.Header.Get(header))
		}

		token := r.Header.Get("Authorization")
		splitToken := strings.Split(token, ":")

		username := ""

		if len(splitToken) > 0 {
			usernameBase64 := splitToken[0]
			usernameByte, _ := base64.StdEncoding.DecodeString(usernameBase64)
			username = string(usernameByte)

			fmt.Printf("Username: %s\n", username)
		}

		fmt.Printf("Accept: %s\n", r.Header.Get("Accept"))
		fmt.Printf("Connection: %s\n", r.Header.Get("Connection"))
		fmt.Printf("Token: %s\n", token)
		fmt.Printf("Username: %s\n", username)

		var apiKey APIKey

		if db.Where("api_key = ?", token).First(&apiKey).RecordNotFound() {
			fmt.Println("API key not found!")
		} else {
			fmt.Println("Api key found!")
		}

		fmt.Println(apiKey)

		next.ServeHTTP(w, r)
	})
}
