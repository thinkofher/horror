package horror

import (
	"log"
	"net/http"

	"github.com/thinkofher/horror/status"
)

func Example_404() {
	// Minimal example of using horror package with 404 error not found
	// status error.
	handler := horror.HandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		name := r.URL.Query().Get("name")
		if name != "john" {
			return status.Forbidden([]byte("john is forbidden from using this site"))
		}
		w.Write([]byte("hello " + name + "!"))
	})

	http.Handle("/greet", horror.WithError(handler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
