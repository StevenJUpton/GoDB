// project in order to learn go

package main

import (
	"container/list"
	"io"
	"net/http"
	"os"
	"strings"
)

//create global tables map
var gtab = make(map[string]*list.List)

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
	ptabData := list.New()
	//add column headers
	ptabData.PushFront("ID|NAME")
	//add new table to global map
	gtab[ptableName[1]] = ptabData

	io.WriteString(w, ptableName[1])
	io.WriteString(w, "ID|NAME")

}

//insertdata handler
func insertData(w http.ResponseWriter, r *http.Request) {
	//get data to insert
	pinsertData := strings.Split(r.URL.RequestURI(), "|")
	//get tablename
	ptableNameTemp := strings.Split(r.URL.RequestURI(), "?")
	//get get rid of the actual data
	ptableName := strings.Split(ptableNameTemp[1], "|")
	//get table from global map
	ptabData := gtab[ptableName[0]]
	//add data to table
	ptabData.PushBack(pinsertData)

}

func selectTable(w http.ResponseWriter, r *http.Request) {
	//get tablename from querystring
	ptableName := strings.Split(r.URL.RequestURI(), "?")
	io.WriteString(w, ptableName[1])
	//return table to select from global map
	ptabData := gtab[ptableName[1]]
	presponse := ""

	for e := ptabData.Front(); e != nil; e = e.Next() {
		presponse += e.Value.(string)
	}

	io.WriteString(w, presponse)

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
