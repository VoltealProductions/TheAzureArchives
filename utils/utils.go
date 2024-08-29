package utils

import (
	"encoding/json"
	"errors"
	"log"
	"math/rand"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/sqids/sqids-go"
)

var Validate = validator.New()

func ParseJSON(r *http.Request, payload any) error {
	if r.Body == nil {
		return errors.New("missing request body")
	}

	return json.NewDecoder(r.Body).Decode(payload)
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, errs error) {
	err := WriteJSON(w, status, map[string]string{"error": errs.Error()})
	if err != nil {
		log.Fatal(err)
	}
}

func GenerateURLId() string {
	s, _ := sqids.New()
	id, _ := s.Encode([]uint64{randomNumber(), randomNumber(), randomNumber()})
	return id
}

func randomNumber() uint64 {
	return uint64(rand.Intn(9999))
}
