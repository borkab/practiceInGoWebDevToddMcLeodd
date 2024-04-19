package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.html"))
	//you have to define fm ( with Funcs) before youe Parse
}

func mothDayYear(t time.Time) string {
	return t.Format("01-02-2006")
}

var fm = template.FuncMap{
	"fdateMDY": mothDayYear,
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.html", time.Now())
	if err != nil {
		log.Fatalln(err)
	}

}
