package main
import (
  "encoding/json"
  "html/template"
  "net/http"
  "log"
)

func Message(status bool, message string) (map[string]interface{}) {
	return map[string]interface{} {"status" : status, "message" : message}
}

func Respond(w http.ResponseWriter, data map[string] interface{})  {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func serverHtml(w http.ResponseWriter, r *http.Request) {
  tmpl, err := template.ParseFiles("./pages/index.html")
  if err != nil {
    log.Fatal(err)
    w.Write([]byte("Something went wrong..."))
    return
  }
  tmpl.Execute(w, "")
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
  http.HandleFunc("/", serverHtml)
  http.HandleFunc("/api", restApi)
  if err := http.ListenAndServe(":8080", nil); err != nil {
    panic(err)
  }
}