package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type DeleteUserRequest struct {
	Username string `json:"username"`
}

type DeleteUserResponse struct {
	Message string `json:"message"`
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	var req DeleteUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	db, err := sql.Open("mysql", "root:sasdoP123@tcp(127.0.0.1:3306)/data")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// trying transaction style
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	var id int
	err = db.QueryRow("SELECT id FROM users WHERE username = ?", req.Username).Scan(&id)
	if err != nil {
		// new stff - rollbacks
		tx.Rollback()
		if err == sql.ErrNoRows {
			json.NewEncoder(w).Encode(DeleteUserResponse{Message: "No such entry"})
		} else {
			log.Fatal(err)
		}
		return
	}
	_, err = tx.Exec("DELETE FROM roles WHERE user_id = ?", id)
	if err != nil {
		// If an error is returned, rollback the transaction
		tx.Rollback()
		log.Fatal(err)
	}

	_, err = tx.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("User deleted successfully.")
	json.NewEncoder(w).Encode(DeleteUserResponse{Message: "User deleted successfully"})
}
