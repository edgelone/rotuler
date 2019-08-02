package main

import (
	"encoding/json"
	"log"
	"net/http"
	"rotuler/model"
	_ "rotuler/model"
)

var Routes []model.Route

func init() {
	Routes = model.Init()
}

func main() {

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/", IndexHandler)
	server.ListenAndServe()

}

func IndexHandler(writer http.ResponseWriter, reuqest *http.Request) {
	uri := reuqest.RequestURI

	result, _ := model.Patten(uri)

	log.Println(uri + " " + reuqest.Method)
	log.Println(result)
	writer.Header().Set("Content-Type", "application/json")

	out, _ := json.MarshalIndent(nil, "", "")

	writer.Write(out)
	return
}
