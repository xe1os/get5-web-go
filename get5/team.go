package get5

import (
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/sessions"
	_ "github.com/solovev/steam_go"
	"html/template"
	"net/http"
	_ "strconv"
	_ "time"
)

type TeamCreatePageData struct {
	LoggedIn bool
	Edit     bool
	Content  interface{} // should be template
}

func TeamCreateHandler(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("get5/templates/layout.html", "get5/templates/team_create.html")) // template
	session, _ := SessionStore.Get(r, SessionData)
	m := &TeamCreatePageData{
		Edit:    false,
		Content: tpl,
	}
	if _, ok := session.Values["Loggedin"]; ok {
		m.LoggedIn = true
	} else {
		http.Redirect(w, r, "/login", 302)
	}

	// テンプレートを描画
	tpl.Execute(w, m)
}

func TeamHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("TeamHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}

func TeamEditHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("TeamEditHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}

func TeamDeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("TeamDeleteHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}

func TeamsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //パスパラメータ取得
	fmt.Printf("TeamsHandler\nvars : %v", vars)
	w.WriteHeader(http.StatusOK)
}

func MyTeamsHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := SessionStore.Get(r, SessionData)
	if _, ok := session.Values["Loggedin"]; ok {
		if session.Values["Loggedin"] == true {

		}
	}
	w.WriteHeader(http.StatusOK)
}
