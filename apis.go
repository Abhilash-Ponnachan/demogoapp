package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const defQtsLimit uint = 15

// handler funcs for diff req
type reqHandler struct {
	repo quoteRepo
}

func (rh reqHandler) init() {
	rh.repo.init()
}

func (rh reqHandler) finalize() {
	rh.repo.close()
}

// handler func for /hello[?name=xxx]
func (rh reqHandler) hello(w http.ResponseWriter, r *http.Request) {
	q, ok := r.URL.Query()["name"]
	if ok {
		fmt.Fprintf(w, "<h1>Salut, Bonjour  %s!</h1>", string(q[0]))
	} else {
		fmt.Fprint(w, "<h1>Salut, Bonjour!</h1>")
	}
}

// handler func for /datetime => JSON
func (rh reqHandler) datetime(w http.ResponseWriter, r *http.Request) {
	n := time.Now()
	dt := struct {
		Date string
		Time string
	}{
		n.Format("2006 Jan 02"),
		n.Format("03:04:05 PM"),
	}
	js, err := json.Marshal(dt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// retrieve listof quotes from 'repository'
func (rh reqHandler) listquotes(w http.ResponseWriter, r *http.Request) {
	qts := []quote{}
	// populate quotes from repository
	rh.repo.listQuotes(defQtsLimit, &qts)
	js, err := json.Marshal(qts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// add quote to 'repository'
func (rh reqHandler) addquote(w http.ResponseWriter, r *http.Request) {
	// handle only POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// placeholder quote obj
	var qt quote
	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&qt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//log.Printf("Person: %+v", qt)

	// add quote to repository
	lastRowid := rh.repo.addQuote(&qt)
	rsp := struct {
		LastRowId uint
	}{
		lastRowid,
	}
	js, err := json.Marshal(rsp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	// if success, return last inserted Id
}

// delete quote from 'repository'
func (rh reqHandler) deleteqoute(w http.ResponseWriter, r *http.Request) {
	// handle only POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// dummy obj for req data
	rd := struct {
		RowId int
	}{
		-1,
	}

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&rd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//log.Printf("ReqObj: %+v", rd)
	if rd.RowId > -1 {
		rowsDel := rh.repo.deleteQuote(uint(rd.RowId))
		rsp := struct {
			RowsDeleted uint
		}{
			rowsDel,
		}
		js, err := json.Marshal(rsp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
		// if success, return rows deleted
	}

}
