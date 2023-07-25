package main

import (
	"encoding/json"
	"net/http"

	"github.com/carlosgenuino/fieldservice/internal/entity"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Get("/order", Order)
	http.ListenAndServe(":8080", r)
}

func Order(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Method not allowed")
		return
	}	

	order, err := entity.NewOrder("123", 19.9, 0.1)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	err = order.CalculateFinalPrice()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(order)
}
