package main

import "net/http"
import "inquisition/db"
import "encoding/json"
import "fmt"


/*
func getUserId (w http.ResponseWriter, r *http.Request) string {
    session, err := store.Get(r, "session")
    if err != nil {
        http.Redirect(w, r, "/", http.StatusFound)
        return ""
    }
    
    useridVal := session.Values["UserId"]

    if useridVal == nil {
        http.Redirect(w, r, "/", http.StatusFound)
        return ""
    }

    id := useridVal.(string)
    
    if len(id) == 0 {
        http.Redirect(w, r, "/", http.StatusFound)
        return ""
    }
    
    return id
}
*/

func sendJson (w http.ResponseWriter, v interface{}) {
    w.Header ().Add ("Content-Type", "application/json")
    data, err := json.Marshal (v)

    if err != nil {
        panic (err)
    }

    fmt.Fprintf (w, "%s", data)
}


func loginHandler (w http.ResponseWriter, r *http.Request) {
    login := r.FormValue("username")
    password := r.FormValue("password")

    type JsonData struct
    {
      Authenticated bool
    }

    _, err := db.AuthenticateTaker (login, password)
    var ret JsonData
    ret.Authenticated = (err == nil)

    fmt.Println (err)
    sendJson (w, ret)
}

func initWebServer() {
    http.HandleFunc("/login", loginHandler)
    http.Handle("/view/", http.StripPrefix("/view/", http.FileServer(http.Dir("../src/inquisition/www"))))
    http.ListenAndServe(":8080", nil)
}

func main() {
    initWebServer()
}
