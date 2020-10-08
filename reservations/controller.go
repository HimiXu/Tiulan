package reservations

import (
	"encoding/json"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/HimiXu/Tiulan/models"
	"github.com/julienschmidt/httprouter"
)

var rm *models.ReservationManager

func init() {
	rm = models.NewReservationManager()
}

func CreateReservation(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var rsv models.Reservation
	json.NewDecoder(r.Body).Decode(&rsv)
	rsv.ID = uuid.NewV4().String()
	rsv.Timestamp = time.Now()
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if !rm.CreateReservation(rsv) {
		json.NewEncoder(w).Encode("Cannot make reservation, already reserved or something")
		return
	}
	json.NewEncoder(w).Encode(rsv)
}

func GetReservation(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	rsv, ok := rm.Reservations[id]
	if ok {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(rsv)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func DeleteReservation(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	rsv, ok := rm.Reservations[id]
	rm.DeleteReservation(id)
	if ok {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(rsv)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}
