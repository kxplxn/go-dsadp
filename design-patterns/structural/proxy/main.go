package main

import "fmt"

// server is the Service Interface. It declares the interface of the Service.
// The proxy must follow this interface to be able to disguise itself as a
// service object.
type server interface {
	handleRequest(string, string) (int, string)
}

// Application is the Service. It is a class that provides some useful business
// logic.
type Application struct {
}

const (
	urlAppStatus  = "/app/status"
	urlCreateUser = "/create/user"
)

func (a *Application) handleRequest(url, method string) (int, string) {
	if url == urlAppStatus && method == "GET" {
		return 200, "Ok"
	}
	if url == urlCreateUser && method == "POST" {
		return 201, "User Created"
	}
	return 404, "Not Ok"
}

// Ngnix is the Proxy class. It has a reference field that points to a service
// object. After the proxy finishes its processing, it passes the request to
// the service object. Usually, proxies manage the full lifecycle of their
// service objects.
type Ngnix struct {
	application        *Application
	maxAllowedRequests int
	rateLimiter        map[string]int
}

func NewNgnixServer() *Ngnix {
	return &Ngnix{
		application:        &Application{},
		maxAllowedRequests: 2,
		rateLimiter:        make(map[string]int),
	}
}

func (n *Ngnix) handleRequest(url, method string) (int, string) {
	allowed := n.checkRateLimiting(url)
	if !allowed {
		return 403, "Not Allowed"
	}
	return n.application.handleRequest(url, method)
}

func (n *Ngnix) checkRateLimiting(url string) bool {
	if n.rateLimiter[url] >= n.maxAllowedRequests {
		return false
	}

	n.rateLimiter[url]++
	return true
}

func main() {
	server := NewNgnixServer()

	url := urlAppStatus
	for i := 0; i < 3; i++ {
		code, body := server.handleRequest(url, "GET")
		fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", url, code, body)
	}

	url = urlCreateUser
	for _, method := range []string{"POST", "GET"} {
		code, body := server.handleRequest(url, method)
		fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", url, code, body)
	}
}
