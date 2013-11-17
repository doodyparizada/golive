package main

import (
	"encoding/json"
	"log"
    "fmt"
	"net/http"
)

type JsonResponse map[string]interface{}

func (r JsonResponse) String() (s string) {
    dumps, err := json.Marshal(r)
        if err != nil {
            s = ""
                return
        }
    s = string(dumps)
        return
}

func RoomView(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	room_url := req.FormValue("url")
	token := req.FormValue("token")
	log.Printf("url: %s, token: %s", room_url, token)
	if req.Method == "GET" {
		// join room
		private_id := req.FormValue("private_id")
		log.Printf("private_id: %s", private_id)
        fmt.Fprint(w, JsonResponse{
            "conference_id": "23423423ABSDFSDFR/SDF",
            "host": "live.rounds.com",
            "port": 50000,
            })
	} else if req.Method == "POST" {
		// create private room
        fmt.Fprint(w, JsonResponse{
            "private_id": "23423423ABSDFSDFR/SDF",
            })
	} else {
		http.NotFound(w, req)
	}
}

func BlockView(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	token := req.FormValue("token")
	offender := req.FormValue("offender")
	conference_id := req.FormValue("conference_id")
	log.Printf("offender: %s, token: %s, conference_id", offender, token, conference_id)
    fmt.Fprint(w, JsonResponse{
        "result": "ok",
        })
}

func HideView(w http.ResponseWriter, req *http.Request) {
	BlockView(w, req)
}

func main() {
	http.HandleFunc("/room", RoomView)
	http.HandleFunc("/block", BlockView)
	http.HandleFunc("/hide", HideView)

	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
