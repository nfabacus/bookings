package handlers

import (
	"encoding/json"
	"github.com/nfabacus/bookings/internal/config"
	"github.com/nfabacus/bookings/internal/forms"
	"github.com/nfabacus/bookings/internal/models"
	"github.com/nfabacus/bookings/internal/render"
	"log"
	"net/http"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	var emptyEnquiry models.Enquiry
	data := make(map[string]interface{})
	data["enquiry"] = emptyEnquiry

	// make forms available in template below
	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}

// Post form handler
func (m *Repository) PostForm(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	enquiry := models.Enquiry{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Email:     r.Form.Get("email"),
		Phone:     r.Form.Get("phone"),
	}

	form := forms.New(r.PostForm)

	form.Required("first_name", "last_name", "email")
	form.MinLength("first_name", 3, r)
	form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["enquiry"] = enquiry

		render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

	m.App.Session.Put(r.Context(), "enquiry", enquiry)

	http.Redirect(w, r, "/submission-summary", http.StatusSeeOther)

	//resp := jsonResponse{
	//	OK:      true,
	//	Message: "Available!",
	//}
	//
	//out, err := json.MarshalIndent(resp, "", "     ")
	//if err != nil {
	//	log.Println(err)
	//}
	//w.Header().Set("Content-Type", "application/json")
	//w.Write(out)
}

func (m *Repository) SubmissionSummary(w http.ResponseWriter, r *http.Request) {
	enquiry, ok := m.App.Session.Get(r.Context(), "enquiry").(models.Enquiry)
	if !ok {
		log.Println("cannot get item from session")
		return
	}
	data := make(map[string]interface{})
	data["enquiry"] = enquiry

	render.RenderTemplate(w, r, "submissionSummary.page.tmpl", &models.TemplateData{
		Data: data,
	})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, world!"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

// Example get json handler
func (m *Repository) GetExampleJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK:      true,
		Message: "Available!",
	}

	out, err := json.MarshalIndent(resp, "", "     ")
	if err != nil {
		log.Println(err)
	}

	log.Println(string(out))
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}
