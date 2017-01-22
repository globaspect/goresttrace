package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	// "github.com/gorilla/mux"
	"io/ioutil"
	"io"

	"github.com/schilli69/src/gatewayComTracer/types"
	"time"
)

//  curl -v -X GET -H 'Content-Type: application/json' http://localhost:14001/health

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "resp:", time.Now(), "\n")
}

//  curl -v -X POST -H 'Content-Type: application/json' -d '{"traceMsg":"trace test message"}' http://localhost:14001/trace

func traceInputHandler(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	var traceInfo types.TraceInfo
	if err := json.Unmarshal(body, &traceInfo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	traceInfo.TimeStamp = time.Now()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(traceInfo); err != nil {
		panic(err)
	}
	return
}

// curl -v -X POST -H 'Content-Type: application/json' -d '{"traceMsg":"trace test message"}' http://localhost:14001/string

func traceInputHandlerString(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	s := string(body[:])
	println("Body:", s)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		panic(err)
	}
	return
}
