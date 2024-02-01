package utils

import "net/http"

func HandleError(writer http.ResponseWriter, err error) {
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}
