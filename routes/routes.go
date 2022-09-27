package routes

import (
	"html/template"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/hasirm/goauth/controllers"
)

var tmpl *template.Template

func Setup(app *fiber.App) {

	app.Post("/register", controllers.Register)
	app.Post("login", controllers.Login)
	app.Get("/user", controllers.User)
	app.Post("/logout", controllers.Logout)

	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))
	http.HandleFunc("/register", controllers.HomeHandler)
	http.HandleFunc("/login", controllers.LoginHandler)
	http.HandleFunc("/resetpass", controllers.ResetHandler)
	http.HandleFunc("/process", controllers.Processor)

	tmpl = template.Must(template.ParseGlob("templates/*.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		// do something with details
		// _ = details

		// tmpl.Execute(w, struct{ Success bool }{true})
	})
}
