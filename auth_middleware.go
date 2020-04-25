package main

import (
	"encoding/base64"
	"net/http"
	"strings"
)

var noAuthUris = []string{"/generate", "/hello", "/login", "/check-api-key"}

func AuthorizationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestUri := r.RequestURI
		// fmt.Printf("We are accepting. %s\n", requestUri)

		for _, uri := range noAuthUris {
			if strings.HasPrefix(requestUri, uri) == true {
				next.ServeHTTP(w, r)
				return
			}
		}

		// for header := range r.Header {
		// 	fmt.Printf("%s: %s\n", header, r.Header.Get(header))
		// }

		token := r.Header.Get("Authorization")
		splitToken := strings.Split(token, ":")

		username := ""

		if len(splitToken) > 0 {
			usernameBase64 := splitToken[0]
			usernameByte, _ := base64.StdEncoding.DecodeString(usernameBase64)
			username = string(usernameByte)

			// fmt.Printf("Username: %s\n", username)
		}

		// fmt.Printf("Accept: %s\n", r.Header.Get("Accept"))
		// fmt.Printf("Connection: %s\n", r.Header.Get("Connection"))
		// fmt.Printf("Token: %s\n", token)
		// fmt.Printf("Username: %s\n", username)

		var apiKey APIKey

		if db.Where("api_key = ?", token).First(&apiKey).RecordNotFound() {
			// fmt.Println("API key not found!")
			w.WriteHeader(http.StatusUnauthorized)
			return
		} else {
			// fmt.Println("Api key found!")
		}

		// fmt.Println(apiKey)
		r.Header.Set("Username", username)

		next.ServeHTTP(w, r)
	})
}
