package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	eval "github.com/hatobus/go-training/ch07/ex7_14"
)

type Answer struct {
	Answer float64 `json:"answer"`
}

// localhost:8080?expression=x+y&x=1&y=2
func getCalclatorHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rq := strings.Split(r.URL.RawQuery, "&")
		rawQuery := make(map[string]string, len(rq))

		var expression string
		for _, q := range rq {
			qs := strings.SplitN(q, "=", 2)
			key, value := qs[0], qs[1]
			value, _ = url.QueryUnescape(value)
			if key == "expression" {
				expression = value
			} else {
				rawQuery[key] = value
			}
		}

		if expression == "" {
			http.Error(w, "expression expected", http.StatusBadRequest)
			return
		}

		expr, err := eval.Parse(expression)
		if err != nil {
			http.Error(w, fmt.Sprintf("bad expression: %v", err), http.StatusBadRequest)
			return
		}

		env := eval.Env{}

		for key, val := range rawQuery {
			f64v, err := strconv.ParseFloat(val, 64)
			if err != nil {
				http.Error(w, fmt.Sprintf("invalid value %v parse error %v", val, err), http.StatusBadRequest)
				return
			}
			env[eval.Var(key)] = f64v
		}

		ans := expr.Eval(env)

		b, err := json.Marshal(Answer{ans})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(b)
	}
}

func main() {
	http.HandleFunc("/", getCalclatorHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
