package book2Lesson1

import (
	"app/src/book2/lesson1/data"
	"errors"
	"fmt"
	"net/http"
)

type Configuration struct {
	Address      string
	ReadTimeout  int64
	WriteTimeout int64
	Static       string
}

var config Configuration

func p(a ...interface{}) {
	fmt.Println(a)
}

func session(w http.ResponseWriter, r *http.Request) (sess data.Session, err error) {
	cookie, err := r.Cookie("_cookie")

	if err == nil {
		sess = data.Session{Uuid: cookie.Value}

		if ok, _ := sess.Check(); !ok {
			err = errors.New("Invalid session")
		}
	}
	return
}
