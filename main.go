package main

import (
	//"container/list"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"io"
	"net/http"
	"os"
	//"strings"
)

var db *sql.DB
var config map[string]interface{}

type Page struct {
	Title        string
	Body         string
	ProjectName  string
	PageName     string
}

func main() {
	loadConfig()
	db = connectToSql()
	defer db.Close()
	loadServer(":" + config["Port"].(string))
}

func connectToSql() *sql.DB {
	var err error
	db, err = sql.Open("sqlite3", config["SQLiteFile"].(string))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return db
}

func loadConfig() {
	file, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var dat []byte
	data := make([]byte, 100)
	for {
		count, err := file.Read(data)
		if err == io.EOF {
			break
		}
		dat = append(dat, data[:count]...)
	}

	if err := json.Unmarshal(dat, &config); err != nil {
		panic(err)
	}
}

func loadServer(port string) {
	fmt.Print("Starting web server on localhost", port, "...\n")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		/*var err error
		_, err = db.Exec("insert into admins(ServerName, Name, Password, Vhost) values('tesztszerver', 'tesztName', 'tesztpassword', 'tesztVhost')")
		if err != nil {
			fmt.Println(err)
			return
		}*/

		/*rows, err := db.Query("select Name from admins")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer rows.Close()

		var namelist = list.New()
		for rows.Next() {
			var name string
			rows.Scan(&name)
			namelist.PushBack(name)
		}

		var names []string

		for e := namelist.Front(); e != nil; e = e.Next() {
			names = append(names, e.Value.(string))
		}*/

		HandleDefaultFunc(w, r, "index.tpl", "index.tpl", "Home", "home")
	})
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		HandleDefaultFunc(w, r, "login.tpl", "login.tpl", "Login", "login")
	})
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		HandleDefaultFunc(w, r, "about.tpl", "about.tpl", "About", "about")
	})
	http.HandleFunc("/stats", func(w http.ResponseWriter, r *http.Request) {
		HandleDefaultFunc(w, r, "stats.tpl", "stats.tpl", "Public Stats", "stats")
	})
	fmt.Print("Done. Serving...\n")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("www/static/"))))
	http.ListenAndServe(port, nil)
}

func HandleDefaultFunc(w http.ResponseWriter, r *http.Request, filename string, filelocation string, title string, pagename string) {
	t, _ := template.New(filename).Funcs(
			template.FuncMap { 
                       		"eq": func(a, b string) bool { 
                               		return a == b 
                       		},
		}).ParseFiles(config["WebDir"].(string) + "/template/header.tpl", config["WebDir"].(string) + "/template/menu.tpl", config["WebDir"].(string) + "/" + filelocation, config["WebDir"].(string) + "/template/footer.tpl")
	p := PageSettings(title, pagename)
	t.Execute(w, p)
}

func HandleFunc(w http.ResponseWriter, r *http.Request, page Page, filename string, filelocation string, title string, pagename string) {
	t, _ := template.New(filename).Funcs(
			template.FuncMap { 
                       		"eq": func(a, b string) bool { 
                               		return a == b 
                       		},
		}).ParseFiles(config["WebDir"].(string) + "/template/header.tpl", config["WebDir"].(string) + "/template/menu.tpl", config["WebDir"].(string) + "/" + filelocation, config["WebDir"].(string) + "/template/footer.tpl")
	p := page
	t.Execute(w, p)
}

func PageSettings(title string, pagename string) Page {
	return Page{Title: title + " - " + config["Title"].(string), Body: "works", ProjectName: config["ProjectName"].(string), PageName: pagename}
}
