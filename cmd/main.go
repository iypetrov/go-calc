package main

import (
	"encoding/json"
	"github.com/IliyaYavorovPetrov/go-calc/app/math"
	"io"
	"log"
	"net"
	"net/http"
)

func getNumbers(r *http.Request) []int {
	var numbers []int

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return numbers
	}

	err = json.Unmarshal(body, &numbers)
	if err != nil {
		return numbers
	}

	return numbers
}

func writeResult(w http.ResponseWriter, result int) {
	response := map[string]int{"result": result}
	responseJSON, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(responseJSON)
	if err != nil {
		return
	}
}

func AppRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/health-check", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		w.WriteHeader(http.StatusOK)
	})

	mux.HandleFunc("/sum", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		result := math.Sum(getNumbers(r))
		writeResult(w, result)
	})

	mux.HandleFunc("/sub", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		result := math.Sub(getNumbers(r))
		writeResult(w, result)
	})

	mux.HandleFunc("/mul", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		result := math.Mul(getNumbers(r))
		writeResult(w, result)
	})

	mux.HandleFunc("/div", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		result := math.Div(getNumbers(r))

		response := map[string]float64{"result": result}
		responseJSON, err := json.Marshal(response)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(responseJSON)
		if err != nil {
			return
		}
	})

	mux.HandleFunc("/and", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		result := math.And(getNumbers(r))
		writeResult(w, result)
	})

	mux.HandleFunc("/or", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		result := math.Or(getNumbers(r))
		writeResult(w, result)
	})

	mux.HandleFunc("/xor", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		result := math.Xor(getNumbers(r))
		writeResult(w, result)
	})

	return mux
}

func main() {
	addr := net.JoinHostPort("", "8080")

	s := http.Server{
		Addr:    addr,
		Handler: AppRouter(),
	}

	log.Println("go-calc started")
	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}
