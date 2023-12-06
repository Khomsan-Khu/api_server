package main

import (
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/hello-world",handleHelloWorld)
	http.HandleFunc("/health",handleHealth)
	http.HandleFunc("/new-endpoint",handleNewEndpoint)

	addr :="localhost:8000"
	log.Printf("Listening on %s ...",addr)

	err:= http.ListenAndServe(addr,nil)
	if err != nil{
		log.Fatal((err))
	}
}

func handleHelloWorld(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	writeResponse(writer, "Hello World!")
}	

func handleHealth(writer http.ResponseWriter,request *http.Request){
	if(request.Method != "GET"){
		http.Error(writer,http.StatusText(http.StatusMethodNotAllowed),http.StatusMethodNotAllowed)
	}
	writeResponse(writer,"Ok!")
}
func handleNewEndpoint(writer http.ResponseWriter,request *http.Request){
	if(request.Method != "GET"){
		http.Error(writer,http.StatusText(http.StatusMethodNotAllowed),http.StatusMethodNotAllowed)
	}
	writeResponse(writer,"Ok!")
}


func writeResponse(writer http.ResponseWriter, responseString string) {
	writer.Header().Set("Content-Type", "text/plain")
	response := []byte(responseString)
	// fmt.Println(string(response))
	_, err := writer.Write(response)
	if err != nil {
		log.Printf("Error writing response: %v", err)
		return
	}
}

failed
