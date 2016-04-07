package main

import (
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/bmizerany/pat"
	"github.com/t-fukui/alpaca/config"
)

var db gorm.DB

func main() {
	css := http.FileServer(http.Dir("assets/css"))
	http.Handle("/assets/css/", http.StripPrefix("/assets/css/", css))
	js := http.FileServer(http.Dir("assets/js"))
	http.Handle("/assets/js/", http.StripPrefix("/assets/js/", js))

	mux := pat.New()

	// Root
	mux.Get("/", http.HandlerFunc(TopHandler))

	// Title
	mux.Get("/title/new", http.HandlerFunc(TitleNewHandler))
	mux.Post("/title/create", http.HandlerFunc(TitleCreateHandler))
	mux.Get("/title/edit/:id", http.HandlerFunc(TitleEditHandler))
	mux.Post("/title/update/:id", http.HandlerFunc(TitleUpdateHandler))

	// Message
	mux.Get("/title/:id/messages", http.HandlerFunc(MessagesIndexHandler))
	mux.Get("/title/:id/message/new", http.HandlerFunc(MessageNewHandler))
	mux.Post("/title/:id/message/create", http.HandlerFunc(MessageCreateHandler))
	mux.Get("/title/:id/message/edit/:message_id", http.HandlerFunc(MessageEditHandler))
	mux.Post("/title/:id/message/update/:message_id", http.HandlerFunc(MessageUpdateHandler))

	http.Handle("/", mux)
	// Webサーバーを起動
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal("ListenAndServe:", Log(http.DefaultServeMux))
	}
}

func init() {
	db = config.Database()
}
