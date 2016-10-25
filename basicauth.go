package basicauth

import (
	"fmt"
	"net/http"
)

var (
	// DEBUG is debug mode.
	// Set to true to enable debug messages.
	DEBUG = false
)

const (
	defBasicRealmStr = "Please input username and password" // Default "Basic realm"
)

// Arguments is used to customize Basic Auth UI(Ex: "Basic realm" string
type Arguments struct {
	BasicRealmStr string // "Basic realm" string
}

// BasicAuth provides function to process HTTP Basic Auth
type BasicAuth struct {
	UserName string    // User Name
	Password string    // Password
	Args     Arguments // Arguments to customize Basic Auth UI(Ex: "Basic realm" string).
}

// New creates a new BasicAuth with default arguments.
func New(username, password string) (ba *BasicAuth) {
	return &BasicAuth{username, password, Arguments{defBasicRealmStr}}
}

// NewWithArgs creates a new BasicAuth with user's arguments(Ex: "Basic realm" string).
//
// See Arguments for more.
func NewWithArgs(username, password string, args Arguments) (ba *BasicAuth) {
	return &BasicAuth{username, password, args}
}

// IsOK checks HTTP Basic Auth.
func (ba *BasicAuth) IsOK(w http.ResponseWriter, r *http.Request) bool {
	username, password, ok := r.BasicAuth()
	if !ok {
		if DEBUG {
			fmt.Printf("r.BasicAuth() failed.\n")
		}
		goto end
	}

	if username != ba.UserName || password != ba.Password {
		ok = false
		goto end
	}

end:
	if !ok {
		w.Header().Set("WWW-Authenticate", fmt.Sprintf(`Basic realm="%v"`, ba.Args.BasicRealmStr))
		w.WriteHeader(401)
	}
	return ok
}
