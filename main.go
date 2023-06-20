/*
main.go

This Go program implements a simple HTTP server that handles incomingrequests to
the root path ("/") and generates an ASCII art banner with the server's hostname
and version. It serves as a basic "Moo World" service, providing a fun response
to client requests. The server listens on the specified port, which can be
configured through the "PORT" environment variable. It also includes error
handling for invalid access and logs relevant information for debugging and
monitoring purposes.

Copyright (c) 2023 Eduardo Toro
*/
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// printMooWorld handles the incoming HTTP requests for the root path "/"
func printMooWorld(response http.ResponseWriter, request *http.Request) {
	// Check if the requested URL path is not "/"
	if request.URL.Path != "/" {
		log.Printf("Invalid Access: %s", request.URL.Path)
		fmt.Fprintf(response, "404 Not Found ")
		return
	}
	log.Printf("Request: %s", request.URL.Path)
	// Get the hostname of the server
	hostname, err := os.Hostname()
	if err != nil {
		log.Println("Error getting hostname:", err)
		return
	}
	// Truncate the hostname if it exceeds 15 characters and print last 12
	if len(hostname) > 15 {
		hostname = "..." + hostname[len(hostname)-12:]
	}
	version := "1.0.0"
	// Generate the ASCII art banner with the hostname and version
	banner := `
	    -----------------------	
	  /  Mooooo, World!        /
	 /  From ` + hostname + `  /
        /  Version ` + version + `         /
         -----------------------
                     \
	              \
		        ^__^
		        (oo)\\_______ 
		        (__)\\       )\/\ 
		            ||----w | 
	                    ||     ||   `

	// Send the banner as the response
	fmt.Fprintf(response, "%s", banner)
	log.Printf("Request complete: %s", request.URL.Path)
}

func main() {
	// Get the desired listening port from the environment variable "PORT"
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	// Set up the HTTP printMooWorld function for the root path "/"
	http.HandleFunc("/", printMooWorld)
	// Start the HTTP server
	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
