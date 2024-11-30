package api

import (
	"net/http"

	"github.com/lowerzedo/gofolder/types"
)

func HandleGetUser(w http.ResponseWriter, r *http.Request) {
	var user types.User
	_ = user
}
