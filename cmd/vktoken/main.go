package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	var id int
	flag.IntVar(&id, "id", 0, "Сlient ID")
	var scope string
	flag.StringVar(&scope, "scope", "all", "Scope")
	flag.Parse()

	fmt.Println("open http://localhost:10101/token in your browser")

	http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		url := fmt.Sprintf("https://oauth.vk.com/authorize?client_id=%d&display=page&redirect_uri=%s&scope=%s&response_type=token&v=5.101&revoke=1", id, "https://oauth.vk.com/blank.html", scope)
		http.Redirect(w, r, url, http.StatusSeeOther)
	})
	if err := http.ListenAndServe(":10101", nil); err != nil {
		panic(err)
	}
}
