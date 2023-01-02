package routes

import (
	"fmt"
	"net/http"

	"github.com/krable55/go-server/pkg/db"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := db.GetAllCustomers()

	if err != nil {
		fmt.Println("error fetching users")
		panic(err)
	}
	fmt.Println(string(users))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(users)
}
