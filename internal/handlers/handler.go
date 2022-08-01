package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/arkarhtethan/golang-web-booking/internal/config"
	"github.com/arkarhtethan/golang-web-booking/internal/forms"
	"github.com/arkarhtethan/golang-web-booking/internal/helpers"
	"github.com/arkarhtethan/golang-web-booking/internal/models"
	"github.com/arkarhtethan/golang-web-booking/internal/render"
)

type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{App: a}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "home.page.html", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "about.page.html", &models.TemplateData{})
}

func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	stringMap["test"] = "Hello, again."
	render.RenderTemplate(w, r, "generals.page.html", &models.TemplateData{StringMap: stringMap})
}

func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	stringMap["test"] = "Hello, again."
	render.RenderTemplate(w, r, "majors.page.html", &models.TemplateData{StringMap: stringMap})
}

func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	var emptyReservation models.Reservation
	data := make(map[string]interface{})
	data["reservation"] = emptyReservation
	render.RenderTemplate(w, r, "make-reservation.page.html", &models.TemplateData{Form: forms.New(nil), Data: data})
}

func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	reservation := models.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Phone:     r.Form.Get("phone"),
		Email:     r.Form.Get("email"),
	}
	form := forms.New(r.PostForm)
	form.Required("first_name", "last_name", "email")
	form.MinLength("first_name", 3)
	form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation
		render.RenderTemplate(w, r, "make-reservation.page.html", &models.TemplateData{Form: form, Data: data})
		return
	}
	m.App.Session.Put(r.Context(), "reservation", reservation)
	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)
}

func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search-availability.page.html", &models.TemplateData{})
}

func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")
	w.Write([]byte(fmt.Sprintf("Start date is %s and end date is %s", start, end)))
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK:      true,
		Message: "Available",
	}
	out, err := json.MarshalIndent(resp, "", "     ")
	if err != nil {
		helpers.ServerError(w, err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.html", &models.TemplateData{})
}

func (m *Repository) ReservationSummary(w http.ResponseWriter, r *http.Request) {
	reservation, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		m.App.ErrorLog.Println("Cannot get item from session")
		m.App.Session.Put(r.Context(), "error", "Can't get reservation from session.")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	m.App.Session.Remove(r.Context(), "reservation")
	data := make(map[string]interface{})
	data["reservation"] = reservation

	render.RenderTemplate(w, r, "reservation-summary.page.html", &models.TemplateData{Data: data})
}
