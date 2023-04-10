// This is the name of our package
// Everything with this package name can see everything
// else inside the same package, regardless of the file they are in
package main

// These are the libraries we are going to use
// Both "fmt" and "net" are part of the Go standard library
import (
	// "fmt" has methods for formatted I/O operations (like printing to the console)
	"fmt"
	// The "net/http" library has methods to implement HTTP clients and servers
	"log"
	"net/http"

	"example.com/accounting/src"
)

func main() {
	// Declare a new router
	var r = src.NewServer()
	log.Printf("Server stopped, err: %v", http.ListenAndServe(":8050", r))
	// http.ListenAndServe(":8080", r)
	// r.Run(":8072")
}

// "handler" is our handler function. It has to follow the function signature of a ResponseWriter and Request type
// as the arguments.
func handler(w http.ResponseWriter, r *http.Request) {
	// For this case, we will always pipe "Hello World" into the response writer
	fmt.Fprintf(w, "Hello World!")
}
