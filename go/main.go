package main

import (
	"gin-test/go/database"
	"gin-test/go/routes"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

func HandlerItemForm(w http.ResponseWriter , r *http.Request){
	tpl := template.Must(template.ParseFiles("frontend"))
	values := map[string]string{}
	if err := tpl.ExecuteTemplate(w, "post.html" , values); err != nil {
		log.Fatal(err)
	}
}

func main() {
	database.ConnectDatabase(1)

	router := routes.NewRoutes()
	router.Run()

	http.HandleFunc("/post.html",HandlerItemForm)
	http.ListenAndServe(":8080",nil)
}

