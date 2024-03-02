package main

import (
	"errors"
	"io"
	"net/http"
)

func keyValuePutHandler(w http.ResponseWriter, r *http.Request) {

	key := r.PathValue("key")

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	if string(bytes) == "" {
		http.Error(w, errors.New("please provide the value in the body").Error(), http.StatusBadRequest)
	}

	err = Put(key, string(bytes))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func keyValueGetHandler(w http.ResponseWriter, r *http.Request) {

	key := r.PathValue("key")

	str, err := Get(key)
	if errors.Is(err, ErrorNoSuchKey) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	_, err = w.Write([]byte(str))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
