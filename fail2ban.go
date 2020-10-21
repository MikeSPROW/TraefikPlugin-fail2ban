package fail2ban

import (
	"context"
	"net"
	"net/http"
)

// Config the plugin configuration.
type Config struct {
	// ...
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		// ...
	}
}

// Example a plugin.
type Fail2ban struct {
	next http.Handler
	name string
	// ...
}

// New created a new plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	// ...
	return &Example{
		// ...
	}, nil
}

func (e *Fail2ban) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	// ...
	e.next.ServeHTTP(rw, req)
}

func checkfail() net.IPAddr {
	return ip
}

func checkban() bool {
	return ban
}

func banIP() net.IPAddr {
	return ip
}
