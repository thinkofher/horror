package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/thinkofher/horror"
	"github.com/thinkofher/horror/status"
)

// below example describes simple scenario when user tries to
// login into service with http Basic Auth protocol
//
// you can use horror/status package to return Unauthorized http status code
// with some basic response or define your own horror.Error as below

type unauthorized struct {
	username string
	password string
}

func (u unauthorized) Error() string {
	// if you will not implement ServeHTTP method for unauthorized type
	// below message will be shown to the client with http status code 500.
	//
	// you can overwrite this default behaviour with InternalHandler function
	// or Adapter type.
	return fmt.Sprintf("user %s tried to login with password %s",
		u.username, u.password)
}

func (u unauthorized) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusUnauthorized)
	fmt.Fprintf(w, "there is no such user as %s", u.username)
}

// handler accepts map with usernames as keys and passwords as values in the
// form of regular strings. It returns Handler that will use basic auth to
// authorize users registered in given map.
func handler(users map[string]string) horror.Handler {
	return horror.HandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		username, password, ok := r.BasicAuth()
		if !ok {
			return status.Unauthorized([]byte("invalid auth method"))
		}

		correctPassword, ok := users[username]
		if !ok {
			return unauthorized{
				username: username,
				password: password,
			}
		}

		if password != correctPassword {
			return status.Unauthorized([]byte("invalid password"))
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("super secret resources"))
		return nil
	})
}

func main() {
	// define dummy map with users and their passwords
	users := map[string]string{
		// having such admin account of course is not a good idea
		// but this is only example :)
		"admin": "admin",
		"user1": "1user",
	}

	// initialize new http.ServeMux
	mux := http.NewServeMux()

	// mount adapted handler with horror.WithError function
	mux.Handle("/secret", horror.WithError(handler(users)))

	// serve and listen with mux
	log.Fatal(http.ListenAndServe(":8080", mux))
}
