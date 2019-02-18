package rotuler

import (
	_ "github.com/lib/pq"
	"database/sql"
	"log"
	"strings"
)

var Db *sql.DB
var Routes []Route

type Route struct {
	Id       string
	Api_name string
	Path     string
	Url      string
	Priority int
	Strip    bool
}

func init() {
	var err error
	Db, err = sql.Open("postgres", "dbname=chitchat sslmode=disable user=lk password=231")
	if err != nil {
		log.Fatal(err)
	}
	rows, err := Db.Query("select id,api_name,path,url,priority,strip from routes")
	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		route := Route{}

		if err = rows.Scan(&route.Id, &route.Api_name, &route.Path,
			&route.Url, &route.Priority, &route.Strip); err != nil {
			log.Fatalln(err)
		}

		append(Routes, route)
	}
	return
}

func Patten(path string) (string, bool) {
	for r := range Routes {
		rule := Routes[r].Path
		rule = rule[strings.Index(rule, "*"):-1]
		if path == rule {
			if Routes[r].Strip {
				return Routes[r].Url + rule, true
			} else {
				return Routes[r].Url + Routes[r].Path, true
			}
		}
		return "", false
	}
}
