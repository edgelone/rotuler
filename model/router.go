package model

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	"log"
	"strings"
)

var Db *sql.DB
var Routes []Route

type Route struct {
	Id          string
	Api_name    string
	Path        string
	Url         string
	Priority    int
	Strip       bool
	Create_time string
	Update_time string
}

func Init() []Route {
	var err error
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=dev dbname=rotuler password=231 sslmode=disable")

	//Db, err = sql.Open("postgres", "dbname=rotuler sslmode=disable user=dev password=231")
	if err != nil {
		log.Fatal(err)
	}
	//rows, err := Db.Query("select id,api_name,path,url,priority,strip,create_time,update_time from routes")
	db.Find(&Routes)
	if err != nil {
		log.Fatalln(err)
	}
	//
	//
	//for rows.Next() {
	//	route := Route{}
	//
	//	if err = rows.Scan(&route.Id, &route.Api_name, &route.Path,
	//		&route.Url, &route.Priority, &route.Strip, &route.Create_time, &route.Update_time); err != nil {
	//		log.Fatalln(err)
	//	}
	//
	//	Routes = append(Routes, route)
	//	log.Println(Routes)
	//
	//}
	return Routes
}

func Patten(path string) (string, bool) {
	for _, r := range Routes {
		p := r.Path

		if match(p, path) {
			if r.Strip {
				return r.Url + p[strings.Index(p, "*"):], true
			} else {
				return r.Url + p, true
			}
		}
	}
	return "", false
}

func match(pattern string, str string) bool {
	pattern = moveHeadTail(pattern)
	str = moveHeadTail(str)

	if strings.Contains(pattern, "**") {
		//TODO
		return true
	} else {
		patterns := strings.Split(pattern, "/")
		strs := strings.Split(str, "/")
		if len(patterns) != len(strs) {
			return false
		}
		for i := 0; i < len(strs); i++ {
			if patterns[i] == "*" {
				continue
			} else if patterns[i] == strs[i] {
				continue
			} else {
				return false
			}
		}
		return true
	}
}

func moveHeadTail(s string) string {
	if strings.HasPrefix(s, "/") {
		s = s[1:]
	}
	if strings.HasSuffix(s, "/") {
		s = s[:len(s)-1]
	}
	return s
}
