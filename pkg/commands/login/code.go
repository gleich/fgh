package login

import (
	"context"
	"net/http"

	"github.com/Matt-Gleich/statuser/v2"
)

func GetToken(port string) string {
	// Create a channel to carry the token
	c := make(chan string)

	server := http.Server{
		Addr: ":" + port,
	}

	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		code := (r.URL.Query().Get("code"))
		if code == "" {
			w.WriteHeader(400)
			w.Write([]byte("Error: no code parameter"))
			return
		}

		// Attempt to exchange the code for an access token
		token, err := Exchange(code)
		if err != nil {
			statuser.Error("Failed to retrieve GitHub access token", err, 1)
			return
		}

		// Post the token to the channel
		c <- token

		// Tell the user the good news!
		_, err = w.Write([]byte("<h1>You've been successfully logged in!</h1><h2>Please head on back to your terminal.</h2>"))
		if err != nil {
			statuser.Error("Failed to send the HTTP response", err, 1)
		}

		// Shut down the server in a goroutine to provide time for the response to be sent
		go func() {
			err := server.Shutdown(context.Background())
			if err != nil {
				statuser.Error("Failed to shut down the OAuth code server", err, 1)
			}
		}()
	})

	// Start the server in a goroutine so we can listen on the channel in the main thread
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			statuser.Error("Failed to start the OAuth code server", err, 1)
		}
	}()

	return <-c
}
