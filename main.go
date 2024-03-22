package main

import (
	"crypto/rand"
	"encoding/json"
	"math/big"
	"net/http"

	"github.com/9bany/kgs/pkg/key0"
)

func keyGen(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result, _ := rand.Int(rand.Reader, big.NewInt(7))
	// make sure the lenght always greater than 3 and lessthan 10
	len := result.Int64() + int64(3)
	gen := key0.NewGen(key0.WithLen(int8(len)))
	k, _ := gen.New()
	resp := map[string]interface{}{
		"key": k,
	}
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/key", keyGen)
	http.ListenAndServe(":8080", nil)
}
