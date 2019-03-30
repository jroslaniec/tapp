package main

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"os"
	"strconv"
)

const page = `
<body bgcolor="%s">
<h1 style="background-color:white;">%s</h1>
</body>
`

var status = 200

func main() {
	fmt.Println("Hello World!")
	http.HandleFunc("/", index)
	http.HandleFunc("/change-status", changeStatus)

	port := os.Getenv("TAPP_PORT")
	if len(port) == 0 {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	name, err := os.Hostname()
	if err != nil {
		panic(errors.Wrap(err, "failed to get hostname"))
	}

	w.WriteHeader(status)
	w.Header().Add("Content-Type", "text/html")
	_, err = w.Write([]byte(fmt.Sprintf(page, name, name)))
	if err != nil {
		panic(err)
	}
}

func changeStatus(w http.ResponseWriter, r *http.Request) {
	qStatus := r.URL.Query().Get("status")

	s, err := strconv.Atoi(qStatus)
	if err != nil {
		w.WriteHeader(500)
		_, _ = w.Write([]byte(err.Error()))
	}
	status = s
	http.Redirect(w, r, "/", 307)
}
