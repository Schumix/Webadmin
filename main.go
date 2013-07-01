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
	Title string
	Name  string
	Body  string
}

func main() {
	loadConfig()
	db = connectToSql()
	defer db.Close()
	loadServer(":" + config["port"].(string))
}

func connectToSql() *sql.DB {
	var err error
	db, err = sql.Open("sqlite3", "Schumix.db3")
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
	funcMap := template.FuncMap{
		"equal": equal,
	}
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

		t, err := template.New("index.tpl").Funcs(funcMap).ParseFiles("www/template/menu.tpl", "www/template/header.tpl", "www/template/footer.tpl", "www/index.tpl")
		if err != nil {
			fmt.Println(err)
		}
		p := Page{Title: config["webtitle"].(string), Name: "index", Body: "works"}
		err = t.Execute(w, p)
		if err != nil {
			fmt.Println(err)
		}
	})
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.New("signin.tpl").Funcs(funcMap).ParseFiles("www/template/menu.tpl", "www/template/header.tpl", "www/template/footer.tpl", "www/signin.tpl")
		if err != nil {
			fmt.Println(err)
		}
		p := Page{Title: "Login - " + config["webtitle"].(string), Name: "login", Body: "works"}
		err = t.Execute(w, p)
		if err != nil {
			fmt.Println(err)
		}
	})
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.New("about.tpl").Funcs(funcMap).ParseFiles("www/template/menu.tpl", "www/template/header.tpl", "www/template/footer.tpl", "www/about.tpl")
		if err != nil {
			fmt.Println(err)
		}
		p := Page{Title: "About - " + config["webtitle"].(string), Name: "about", Body: "works"}
		err = t.Execute(w, p)
		if err != nil {
			fmt.Println(err)
		}
	})
	http.HandleFunc("/stats", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.New("stats.tpl").Funcs(funcMap).ParseFiles("www/template/menu.tpl", "www/template/header.tpl", "www/template/footer.tpl", "www/stats.tpl")
		if err != nil {
			fmt.Println(err)
		}
		p := Page{Title: "Public Stats - " + config["webtitle"].(string), Name: "stats", Body: "works"}
		err = t.Execute(w, p)
		if err != nil {
			fmt.Println(err)
		}
	})
	fmt.Print("Done. Serving...\n")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("www/static/"))))
	http.ListenAndServe(port, nil)
}

// In .tpl files equality check
func equal(a, b interface{}) bool {
	if a == b {
		return true
	}
	return false
}
