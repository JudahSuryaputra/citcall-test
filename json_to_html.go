package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Country struct {
	Name     string `json:"name"`
	DialCode string `json:"dialCode"`
	IsoCode  string `json:"isoCode"`
	Flag     string `json:"flag"`
}

const templateString = `<table>
	<thead>
		<tr>
			<th>Name</th>
			<th>Dial Code</th>
			<th>ISO Code</th>
			<th>Flag</th>
		</tr>
	</thead>
	{{range .}}
		<tr>
			<td>{{ .Name }}</td>
			<td>{{ .DialCode }}</td>
			<td>{{ .IsoCode }}</td>
			<td><img src={{ .Flag }} alt="{{ .Name }}'s flag" width="100" height="100"></td>
		</tr>
		{{end}}
</table>`

func main() {
	var countries []Country
	response, err := http.Get("https://citcall.com/test/countries.json")
	if err != nil {
		log.Fatalln(err)
	}

	decoder := json.NewDecoder(response.Body)
	decoder.DisallowUnknownFields()

	err = decoder.Decode(&countries)
	if err != nil {
		log.Fatalln("err decoding body")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		JSONtoTable(w, countries)
	})

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}

func JSONtoTable(w http.ResponseWriter, data []Country) {
	t := template.New("t")
	t, err := t.Parse(templateString)
	if err != nil {
		panic(err)
	}

	err = t.Execute(w, data)
	if err != nil {
		panic(err)
	}
}
