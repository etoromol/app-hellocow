package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	host, _ := os.Hostname()
	banner := `
	    ----------------------
	  (  Mooooo, World!        )
	 (  From ` + host + `  )
        (  Version 1.0.0         )
          ----------------------
		      \   ^__^ 
		       \  (oo)\\_______ 
		          (__)\\       )\/\ 
		              ||----w | 
	                      ||     ||   `
	fmt.Fprintf(w, "%s", banner)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
