package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
if r.URL.Path != "/hello"{
	http.Error(w,"404 not found"http.StatusNotFound)
	return
}
if r.Method !="GET"{
	http.Error(w,"Method isnot Supported",http.StatusNotFound)
	return
}
fmt.Fprintf(w,"hello")
}

func formHandler (w http.ResponseWriter,r *http.Request){
	if err:= r.ParseForm();error != nil{
		fmt.Fprintf(w,"ParseForm() err: : %v",err)
		return
	}
	Fprintf(w,"POST req is succsesfull")
	name := r.FormValue("name")
	adress := r.FormValue("adress")
	fmt.Fprintf(w,"Name = %s\n",name)
	fmt.Fprintf(w,"Adress = %s\n",adress)
}


func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Print("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
