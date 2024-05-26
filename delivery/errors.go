package delivery

import (
	"log"
	"net/http"
	"strconv"
	"text/template"
)

func (h *Handler) Errors(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	t, err := template.ParseFiles("templates/errors.html")
	if err != nil {
		log.Print(err)
		http.Error(w, strconv.Itoa(http.StatusInternalServerError)+" "+"Error parsing file", http.StatusInternalServerError)
		log.Print(err)
		return
	}
	// type Error struct {
	// 	Statusint        int
	// 	StatusText       string
	// 	StatusTextandInt string
	// 	MessageError     string
	// }
	messageerr := struct {
		StatusInt        int
		StatusText       string
		StatusTextandInt string
		MessageError     string
	}{
		StatusInt:        status,
		StatusText:       http.StatusText(status),
		StatusTextandInt: strconv.Itoa(status) + http.StatusText(status),
		MessageError:     message,
	}
	if err := t.Execute(w, messageerr); err != nil {
		log.Print(err)
		http.Error(w, strconv.Itoa(http.StatusInternalServerError)+" "+"Error executing file", http.StatusInternalServerError)
		return
	}
}
