package main

import (
	"chat_service/templates/server/wshandlers"
	"net/http"
	"path"
	"text/template"
)

func main() {

	ws := wshandlers.WsService
	wshandlers.InitWsServise(ws)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fp := path.Join("templates", "index.html")
		tmpl, _ := template.ParseFiles(fp)

		tmpl.Execute(w, nil)
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ws.HandleRequest(w, r)
	})

	err := http.ListenAndServe("localhost:1113", nil)
	if err != nil {
		println(err)
	}

}
