package delivery

import (
	"chatbotmrrobotdumbazz/service"
	"net/http"
	"text/template"

	"golang.org/x/exp/slog"
)

func (h *Handler) Homepage(w http.ResponseWriter, r *http.Request) {
	const op = "delivery.homepage"
	log := h.log.With(
		slog.String("op", op),
	)
	switch r.Method {
	case "GET":
		temp, err := template.ParseFiles("template/homepage.html")
		if err != nil {
			log.Error(err.Error())
			h.Errors(w, http.StatusInternalServerError, err.Error())
			return
		}
		if err = temp.Execute(w, temp); err != nil {
			log.Error(err.Error())
			h.Errors(w, http.StatusInternalServerError, err.Error())
			return
		}
	case "POST":
		temp, err := template.ParseFiles("template/homepage.html")
		if err != nil {
			log.Error(err.Error())
			h.Errors(w, http.StatusInternalServerError, err.Error())
			return
		}
		if err = r.ParseForm(); err != nil {
			log.Error(err.Error())
			h.Errors(w, http.StatusBadRequest, "Error parsing")
			return
		}
		message, ok := r.Form["message"]
		if !ok {
			h.Errors(w, http.StatusBadRequest, "Bad typing title")
			return
		}
		messages, err := service.WriteMessageChatGPT(message[0])
		if err != nil {
			log.Error(err.Error())
			h.Errors(w, http.StatusBadRequest, "Error parsing")
			return
		}
		if err = temp.Execute(w, messages); err != nil {
			log.Error(err.Error())
			h.Errors(w, http.StatusInternalServerError, err.Error())
			return
		}
	default:
		log.Error(http.StatusText(http.StatusMethodNotAllowed))
		h.Errors(w, http.StatusMethodNotAllowed, "")
		return
	}
}
