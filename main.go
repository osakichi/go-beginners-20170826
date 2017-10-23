package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"os"
	"os/signal"
	"syscall"
	//"github.com/CyberMergina/go-beginers20170826/fruit"
	"./fruit"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(
		func() {
			p := filepath.Join(os.Getenv("GOPATH"), "templates", t.filename)
			log.Println("Template path ")
			t.templ = template.Must(template.ParseFiles(p))
		})

	log.Println("Template execute")
	d := fruit.GetList()
	t.templ.Execute(w, d)
}

func main() {
	http.Handle("/", &templateHandler{filename: "index.html"})

/*
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
*/

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT)
	signal.Notify(sig, syscall.SIGTERM)
	signal.Notify(sig, syscall.SIGKILL)
        go func() {
            log.Println(<-sig)
            //listener.Close()
	    os.Exit(10)
        }()

	crt := filepath.Join(os.Getenv("GOPATH"), "crt", "server.crt")
	key := filepath.Join(os.Getenv("GOPATH"), "crt", "server.key")

        err := http.ListenAndServeTLS(":8443", crt, key, nil)
        if err != nil {
            log.Fatal("ListenAndServeTLS: ", err)
        }
}
