package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type PriceDB struct {
	m  sync.RWMutex
	db map[string]int
}

var listtpl = template.Must(template.New("listing").Parse(`
<html>
<body>

<table>
	<tr>
		<th> item </th>
		<th> ($) price </th>
{{range $key, $val := .}}
	<tr>
		<td> {{$key}} </td>
		<td> ($) {{$val}} </td>
	</tr>
{{end}}

</table>
</body>
</html>
`))

func (d *PriceDB) Create(w http.ResponseWriter, r *http.Request) {
	item := r.FormValue("item")
	if item == "" {
		http.Error(w, "item is expected", http.StatusBadRequest)
		return
	}

	p := r.FormValue("price")
	if p == "" {
		http.Error(w, "price is expected", http.StatusBadRequest)
		return
	}

	price, err := strconv.Atoi(p)
	if err != nil {
		http.Error(w, "price error invalid price input", http.StatusBadRequest)
		return
	}

	d.m.Lock()
	defer d.m.Unlock()

	if _, ok := d.db[item]; ok {
		http.Error(w, fmt.Sprintf("%v is already exist", item), http.StatusBadRequest)
		return
	}

	d.db[item] = price
}

func (d *PriceDB) Read(w http.ResponseWriter, r *http.Request) {
	item := r.FormValue("item")
	if item == "" {
		http.Error(w, "item is expected", http.StatusBadRequest)
		return
	}

	if _, ok := d.db[item]; !ok {
		http.Error(w, fmt.Sprintf("%v not found", item), http.StatusNotFound)
		return
	}

	d.m.RLock()
	defer d.m.RUnlock()

	fmt.Fprintf(w, "%v: %v\n", item, d.db[item])
}

func (d *PriceDB) Update(w http.ResponseWriter, r *http.Request) {
	item := r.FormValue("item")
	if item == "" {
		http.Error(w, "item is expected", http.StatusBadRequest)
		return
	}

	p := r.FormValue("price")
	price, err := strconv.Atoi(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, ok := d.db[item]; !ok {
		http.Error(w, fmt.Sprintf("%v is not found", item), http.StatusNotFound)
		return
	}

	d.m.Lock()
	defer d.m.Unlock()

	d.db[item] = price
}

func (d *PriceDB) Delete(w http.ResponseWriter, r *http.Request) {
	item := r.FormValue("item")
	if item == "" {
		http.Error(w, "item is expected", http.StatusBadRequest)
		return
	}

	if _, ok := d.db[item]; !ok {
		http.Error(w, fmt.Sprintf("%v is not found", item), http.StatusNotFound)
		return
	}

	d.m.Lock()
	defer d.m.Unlock()
	delete(d.db, item)
}

func (d *PriceDB) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	switch method {
	case http.MethodPost:
		d.Create(w, r)
	case http.MethodGet:
		d.Read(w, r)
	case http.MethodPut:
		d.Update(w, r)
	case http.MethodDelete:
		d.Delete(w, r)
	default:
		http.Error(w, fmt.Sprintf("method %v is not allowed", method), http.StatusBadRequest)
	}
}

func (d *PriceDB) List(w http.ResponseWriter, _ *http.Request) {
	d.m.RLock()
	defer d.m.RUnlock()
	if err := listtpl.Execute(w, d.db); err != nil {
		http.Error(w, fmt.Sprintf("html generate failed: %v", err.Error()), http.StatusInternalServerError)
		return
	}
}

func main() {
	db := PriceDB{}
	db.db = make(map[string]int, 0)
	db.db["shoe"] = 50
	db.db["socks"] = 5

	http.HandleFunc("/create", db.Create)
	http.HandleFunc("/read", db.Read)
	http.HandleFunc("/update", db.Update)
	http.HandleFunc("/delete", db.Delete)
	http.HandleFunc("/list", db.List)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
