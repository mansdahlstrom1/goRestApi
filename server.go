package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/gobuffalo/packr/v2"
)

type Request struct {
	SSID       string `json:"ssid"`
	Passphrase string `json:"passphrase"`
}

// Message Creates a JSON Object.
func Message(success bool, message string) map[string]interface{} {
	return map[string]interface{}{
		"success": success,
		"message": message,
	}
}

// Respond returns a JSON respons to the caller.
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

func restAPI(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// Response with shower id to verify the connection
		Respond(w, Message(true, "All Good"))
	case http.MethodPost:
		decoder := json.NewDecoder(r.Body)
		var req Request
		err := decoder.Decode(&req)
		if err != nil {
			panic(err)
		}

		message := "SSID=" + req.SSID + ", pass=" + req.Passphrase

		// Test to connect to WiFi
		// If it works Respond success ("the device will now reboot")
		// if it does not work ("Return error message")

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
	box := packr.New("someBoxName", "./pages/static/")

	http.HandleFunc("/", serveHTML)
	http.HandleFunc("/api", restAPI)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(box)))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
