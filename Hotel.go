package main

import(
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"database/sql"
	"log"
	"fmt"
	"encoding/json"
)

var db *sql.DB

type room_entry struct {
	hotel_id int
	room_description string
}

func room_data_fetcher(w http.ResponseWriter, r *http.Request) {
	var rooms []room_entry

	state := r.URL.Query()["state"]	//gets state query parameter that user passed on GET request

	//gets hotel rooms available in that state
	rows, err := db.Query("SELECT r.HotelID, r.Description FROM Room r, Hotel h WHERE r.HotelID = h.HotelID and h.State=?", state)
	if err != nil {
        	log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
        	var entry room_entry
        	if err := rows.Scan(&(entry.hotel_id), &(entry.room_description)); err != nil {
                	log.Fatal(err)
        	}

		//append room_entry to the rooms array
		rooms = append(rooms, entry)
	}

	if err := rows.Err(); err != nil {
        	log.Fatal(err)
	}

	json, err := json.Marshal(rooms)
	if err != nil {
		log.Fatal(err)
	}

	//returns back json response
	fmt.Printf("At the end\n")
	
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func main() {
	//open database
	var err error
	db, err = sql.Open("mysql", "root:120490Arkadi@tcp(172.31.130.134:3306)/hulton")
	if err != nil {
		log.Fatal(err)
	}

	//serves the static files in directory static
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/room_data/", room_data_fetcher)

	//listens on port 8081
	log.Fatal(http.ListenAndServe(":8081", nil))
}
