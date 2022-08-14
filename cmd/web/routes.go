package main

import (
	"net/http"

	"github.com/arkarhtethan/golang-web-booking/internal/config"
	"github.com/arkarhtethan/golang-web-booking/internal/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	mux.Get("/generals-quarters", http.HandlerFunc(handlers.Repo.Generals))
	mux.Get("/majors-suite", http.HandlerFunc(handlers.Repo.Majors))

	mux.Get("/make-reservation", http.HandlerFunc(handlers.Repo.Reservation))
	mux.Post("/make-reservation", http.HandlerFunc(handlers.Repo.PostReservation))
	mux.Get("/reservation-summary", http.HandlerFunc(handlers.Repo.ReservationSummary))

	mux.Get("/search-availability", http.HandlerFunc(handlers.Repo.Availability))
	mux.Post("/search-availability", http.HandlerFunc(handlers.Repo.PostAvailability))
	mux.Post("/search-availability-json", http.HandlerFunc(handlers.Repo.AvailabilityJSON))
	mux.Get("/choose-room/{id}", http.HandlerFunc(handlers.Repo.ChooseRoom))
	mux.Get("/book-room", http.HandlerFunc(handlers.Repo.BookRoom))

	mux.Get("/user/login", http.HandlerFunc(handlers.Repo.Login))
	mux.Post("/user/login", http.HandlerFunc(handlers.Repo.PostLogin))
	mux.Get("/user/logout", http.HandlerFunc(handlers.Repo.Logout))

	mux.Get("/contact", http.HandlerFunc(handlers.Repo.Contact))
	mux.Route("/admin", func(r chi.Router) {
		// r.Use(Auth)
		r.Get("/dashboard", handlers.Repo.AdminDashboard)

		r.Get("/reservations-new", handlers.Repo.AdminNewReservations)
		r.Get("/reservations-all", handlers.Repo.AdminAllReservations)

		r.Get("/reservations-calendar", handlers.Repo.AdminReservationsCalendar)
		r.Post("/reservations-calendar", handlers.Repo.AdminPostReservationsCalendar)

		r.Get("/process-reservation/{src}/{id}/do", handlers.Repo.AdminProcessReservation)
		r.Get("/delete-reservation/{src}/{id}/do", handlers.Repo.AdminDeleteReservation)

		r.Get("/reservations/{src}/{id}/show", handlers.Repo.AdminShowReservations)
		r.Post("/reservations/{src}/{id}", handlers.Repo.AdminPostShowReservations)
	})

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
