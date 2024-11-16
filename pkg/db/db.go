package db

import (
	"database/sql"
	"log"
	"metalbanddb/pkg/models"
)

// Obtener todas las bandas
func GetBands(db *sql.DB) ([]models.Band, error) {
	rows, err := db.Query("SELECT id, name, genre, formed FROM bands")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()

	var bands []models.Band
	for rows.Next() {
		var band models.Band
		err := rows.Scan(&band.ID, &band.Name, &band.Genre, &band.Formed)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		bands = append(bands, band)
	}
	return bands, nil
}
