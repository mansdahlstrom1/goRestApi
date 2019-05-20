package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/gobuffalo/packr/v2"
)

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func serveHTML(w http.ResponseWriter, r *http.Request) {
	box := packr.New("HTML", "./pages")

	s, err := box.FindString("index.html")
	if err != nil {
		log.Fatal(err)
		w.Write([]byte("Something went wrong..."))
		return
	}

	t, err := template.New("index").Parse(s)
	if err != nil {
		log.Fatal(err)
		w.Write([]byte("Something went wrong..."))
		return
	}

	t.Execute(w, "")
}

func restApi(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		Respond(w, Message(true, "All Good"))
	case http.MethodPost:
		ssid := r.PostFormValue("ssid")
		passphrase := r.PostFormValue("passphrase")

		message := "ssid = " + ssid + ", passphrase = " + passphrase

		Respond(w, Message(true, message))
	case http.MethodPut:
		Respond(w, Message(true, "HTTP method not supported"))
	case http.MethodDelete:
		Respond(w, Message(true, "HTTP method not supported"))
	default:
		Respond(w, Message(false, "Not a valid HTTP method "))
	}

}

func main() {
	http.HandleFunc("/", serveHTML)
	http.HandleFunc("/api", restApi)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
