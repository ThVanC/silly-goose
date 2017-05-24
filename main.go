package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/module1", module1)
	http.HandleFunc("/module2", module2)
	fmt.Println("Silly goose is running on http://localhost:3016. If you have OWASP ZAP running, configure your browser to connect to the ZAP Proxy.")
	http.ListenAndServe(":3016", nil)

}

type Page struct {
	Title string
	Body  []byte
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	p := Page{"Welcome to Silly goose!", []byte("Choose a module. <br /><ol><li><a href='/module1'>Module 1</a></li><li><a href='/module2'>Module 2</a></li></ol>")}
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

type LoginRequest struct{ username, password string }

func module1(w http.ResponseWriter, r *http.Request) {
	if r.Method == "" || r.Method == "GET" {
		cookie, err := r.Cookie("session")
		if err != nil {
			p := Page{"Module 1: Log-in to Silly Goose", []byte("Please log-in to silly goose! " +
				"<form action='/module1' method='post'> " +
				"<input name='username' value=''><input name='password' value=''><button>Log-in</button></form>")}
			fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
			return
		} else {
			fmt.Fprintf(w, "You are logged in as "+cookie.Value)
		}
	} else if r.Method == "POST" {
		r.ParseForm()
		log.Println(r.Form)
		loginRequest := LoginRequest{strings.Join(r.Form["username"], ""), strings.Join(r.Form["password"], "")}

		if loginRequest.username == "admin" && loginRequest.password == "admin" {
			expire := time.Now().AddDate(0, 0, 1)
			cookie := http.Cookie{"session", "admin", "/", "localhost:3017", expire, expire.Format(time.UnixDate), 86400, false, true, "username=Admin", []string{"username=Admin"}}
			http.SetCookie(w, &cookie)
			fmt.Fprintf(w, "You are now logged in as admin!")
		} else {
			p := Page{"Module 1: Log-in to Silly Goose", []byte("Wrong password. You entered: " + loginRequest.username +
				"<form action='/module1' method='post'> " +
				"<input name='username' value=''><input name='password' value=''><button>Log-in</button></form>")}
			fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
		}
	}
}

func module2(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		p := Page{"Module 2: Upload an XML document", []byte("" +
			"<form action='/module2' method='post'> " +
			"<textarea name='xml'>&lt;xml&gt;&lt;Animal&gt;A camel&lt;/Animal&gt;&lt;/xml&gt;</textarea> <button>Submit XML</button></form>")}
		fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
	} else if r.Method == "POST" {
		r.ParseForm()
		xmlReq := r.Form["xml"]
		log.Println(xmlReq)
		domainObject := struct{ Animal string }{"Silly Goose"}
		xml.Unmarshal([]byte(strings.Join(xmlReq, "")), &domainObject)
		prefix := ""
		cookie, err := r.Cookie("session")
		if err != nil && cookie != nil {
			prefix = "Hello, " + cookie.Value
		}
		p := Page{"Module 2: Upload an XML document", []byte(prefix + "You are the following animal: " + domainObject.Animal +
			"<form action='/module2' method='post'> " +
			"<textarea name='xml'>" + strings.Join(xmlReq, "") + "</textarea> <button>Submit XML</button></form>")}
		fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
	}
}
