package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
	"github.com/arkarhtethan/golang-web-booking/internal/models"
)

// AppConfig holds the application config
type AppConfig struct {
	UseCache     bool
	TempateCache map[string]*template.Template
	InfoLog      *log.Logger
	ErrorLog     *log.Logger
	InProduction bool
	Session      *scs.SessionManager
	MailChan     chan models.MailData
}
