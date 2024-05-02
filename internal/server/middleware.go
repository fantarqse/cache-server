package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	emptyString string = ""
	fourSpaces  string = "    "
)

type Log struct {
	Method string
	URL    string
	Time   string
}

func (l *Log) Format() string {
	bytes, err := json.MarshalIndent(l, emptyString, fourSpaces)
	if err != nil {
		log.Println("failed to format log data to json") // should never happen
	}

	return string(bytes)
}

func logMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		msg := Log{
			Method: r.Method,
			URL:    r.URL.Path,
			Time:   time.Now().Format(time.DateTime),
		}

		fmt.Printf("%s\n", msg.Format())
	}
}
