// project in order to learn go

package main

import (
	"io"
	"net/http"
	"os"
)

//simple handler to accept basic commands
func pingHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Pong")
}

//exit program
func stopServer(w http.ResponseWriter, r *http.Request) {
	os.Exit(0)
}

func main() {

	//test server is running
	http.HandleFunc("/ping", pingHandler)
	//stop the server
	http.HandleFunc("/stopserver", stopServer)

	//start http listener
	http.ListenAndServe(":8000", nil)

}
