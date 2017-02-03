// project in order to learn go

package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"strings"
)

//create global tables map
var gtab = make(map[string][]string)

//simple handler to accept basic commands
func pingHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Pong")
}

//exit program
func stopServer(w http.ResponseWriter, r *http.Request) {
	os.Exit(0)
}

//create table handler
func createTable(w http.ResponseWriter, r *http.Request) {

	//get tablename from querystring
	ptableName := strings.Split(r.URL.RequestURI(), "?")
	//create linked list to hold data
	ptabData := []string{"ID|NAME"}
	//add new table to global map
	gtab[ptableName[1]] = ptabData

	io.WriteString(w, ptableName[1]+"\n")
	io.WriteString(w, ptabData[0])

}

//insertdata handler
func insertData(w http.ResponseWriter, r *http.Request) {
	//get data to insert
	pinsertData := strings.Split(r.URL.RequestURI(), "@")
	//get tablename
	ptableNameTemp := strings.Split(r.URL.RequestURI(), "?")
	//get get rid of the actual data
	ptableName := strings.Split(ptableNameTemp[1], "@")
	//get table from global map
	ptabData := gtab[ptableName[0]]
	//add data to table
	ptabData = append(ptabData, pinsertData[1:]...)
	//write back to global map
	gtab[ptableName[0]] = ptabData

	//output what was inserted
	io.WriteString(w, pinsertData[1])

}

func selectTable(w http.ResponseWriter, r *http.Request) {
	//get tablename from querystring
	ptableName := strings.Split(r.URL.RequestURI(), "?")
	io.WriteString(w, ptableName[1])
	//return table to select from global map
	ptabData := gtab[ptableName[1]]

	var buffer bytes.Buffer

	for _, e := range ptabData {
		buffer.WriteString(e)
		buffer.WriteString("\n")
	}

	io.WriteString(w, buffer.String())

}

func main() {

	//test server is running
	http.HandleFunc("/ping", pingHandler)
	//stop the server
	http.HandleFunc("/stopserver", stopServer)
	//create table
	http.HandleFunc("/createtable", createTable)
	//select table
	http.HandleFunc("/selecttable", selectTable)
	//insert data
	http.HandleFunc("/insertdata", insertData)

	//start http listener
	http.ListenAndServe(":8000", nil)

}
