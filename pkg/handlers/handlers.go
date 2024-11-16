package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"database/sql"
	"log"
)

// Obtener información de la banda desde Last.fm
func GetBandFromLastFM(db *sql.DB, fetchBandInfo func(string) (map[string]interface{}, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Obtener el nombre de la banda de los parámetros de la URL
		params := mux.Vars(r)
		bandName := params["name"]

		// Llamar a la función que obtiene la información de Last.fm
		info, err := fetchBandInfo(bandName)
		if err != nil {
			// Logueamos el error para depurarlo
			log.Printf("Error al obtener la información de la banda '%s': %v", bandName, err)

			// Devolver el error al cliente
			http.Error(w, "Error al obtener información de la banda", http.StatusInternalServerError)
			return
		}

		// Si todo está bien, devolver la respuesta con la información
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(info)
	}
}
