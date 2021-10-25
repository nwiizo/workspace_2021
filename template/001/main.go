package main

import (
	"math/rand"
	"os"
	"text/template"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	tmpl := template.Must(template.
		New("").
		Funcs(map[string]interface{}{
			"rand": func() int {
				return rand.Intn(100)
			},
		}).
		Parse(`Hi {{.}}, you are number {{rand}}.`))
	_ = tmpl.Execute(os.Stdout, "User")
}
