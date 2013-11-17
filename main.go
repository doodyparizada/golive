package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type BlockResponse struct {
	Result string `json:"result"`
}
type CreatePrivateRoomResponse struct {
	PrivateID string `json:"private_id"`
}
type JoinRoomResponse struct {
	ConferenceID string `json:"conference_id"`
	Host         string `json:"host"`
	Port         int    `json:"port"`
}

func RoomView(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	room_url := req.FormValue("url")
	token := req.FormValue("token")
	log.Printf("url: %s, token: %s", room_url, token)
	if req.Method == "GET" {
		private_id := req.FormValue("private_id")
		log.Printf("private_id: %s", private_id)
		// join room
		res := &JoinRoomResponse{
			ConferenceID: "23423423ABSDFSDFR/SDF",
			Host:         "live.rounds.com",
			Port:         50000,
		}
		jres, _ := json.Marshal(res)
		io.WriteString(w, string(jres))
	} else if req.Method == "POST" {
		// create private room
		res := &CreatePrivateRoomResponse{
			PrivateID: "23423423ABSDFSDFR/SDF",
		}
		jres, _ := json.Marshal(res)
		io.WriteString(w, string(jres))
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

	res := &BlockResponse{
		Result: "ok",
	}
	jres, _ := json.Marshal(res)
	io.WriteString(w, string(jres))
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
