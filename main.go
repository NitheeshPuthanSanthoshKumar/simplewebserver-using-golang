package main

import (
	"fmt"
	"log"
	"net/http"

)
func testFunc(w http.ResponseWriter,r *http.Request){
	if r.URL.Path != "/test" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	

	fmt.Fprintf(w, "test!")

}
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello!")
}
func formHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	//fmt.Print(r.ParseForm())
	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
	
	
}
func main(){
	
	fmt.Print("hello world")
	filserver:=http.FileServer(http.Dir("./static"))
	http.Handle("/",filserver)
	http.HandleFunc("/f",formHandler)
	http.HandleFunc("/hello",helloHandler)
	http.HandleFunc("/test",testFunc)
	fmt.Printf("listening at 8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}

	
}
	



