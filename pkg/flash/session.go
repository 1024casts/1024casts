package flash

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

var store *sessions.CookieStore
var flashName = "flash-session"

func init() {
	store = sessions.NewCookieStore([]byte("flash-secret"))
}

func SetFlashMessage(w http.ResponseWriter, r *http.Request, name string, value string) {
	session, err := store.Get(r, flashName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	session.AddFlash(value, name)
	session.Save(r, w)
}

func GetFlashMessage(w http.ResponseWriter, r *http.Request, name string) []interface{} {
	session, err := store.Get(r, flashName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fm := session.Flashes(name)
	if fm == nil {
		fmt.Fprint(w, "No flash messages")
		return []interface{}{}
	}
	session.Save(r, w)
	fmt.Fprintf(w, "%v", fm[0])

	return fm
}
