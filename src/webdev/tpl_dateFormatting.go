package webdev

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpldateformat.gohtml"))
}

func monthDayYear(t time.Time) string {	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))

	return t.Format("01-02-2006")
}

var fm = template.FuncMap{
	"fdateMDY": monthDayYear,
}

func main() {

	err:= tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", time.Now())
	if err != nil {
		log.Fatalln(err)
	}
}