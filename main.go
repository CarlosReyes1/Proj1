package main

import (
	"html/template" // templates
	"io/ioutil"     // file reading
	"net/http"      // web server
)

//Holds the username.
type User struct {
	UserName string
}

//Holds a loaded file.
var file string

//Loads the file's contents.
func init() {
	temp, _ := ioutil.ReadFile("templates/file1.htemplate")
	file = string(temp)
}

//Serves the executed template.
func lPage(res http.ResponseWriter, req *http.Request) {
	obj := User{"Test User"}
	t, _ := template.New("name").Parse(file)
	t.Execute(res, obj)
}

func main() {
	http.HandleFunc("/", lPage)
	http.ListenAndServe(":8080", nil)
}
