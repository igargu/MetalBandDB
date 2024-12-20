package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// Esta función consulta la API de Last.fm para obtener información de la banda
func FetchBandInfoFromLastFM(bandName string) (map[string]interface{}, error) {
	// Obtener la clave de API de Last.fm desde las variables de entorno
	apiKey := os.Getenv("LFM_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("API key de Last.fm no configurada")
	}

	// Construir la URL para la API de Last.fm
	url := fmt.Sprintf("http://ws.audioscrobbler.com/2.0/?method=artist.getinfo&artist=%s&api_key=%s&format=json", bandName, apiKey)

	// Hacer la solicitud HTTP a la API
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error al realizar la solicitud HTTP: %v", err)
	}
	defer resp.Body.Close()

	// Verificar si la respuesta es exitosa
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error en la respuesta de la API: status code %d", resp.StatusCode)
	}

	// Parsear la respuesta JSON
	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("error al decodificar la respuesta JSON: %v", err)
	}

	// Devolver los datos obtenidos de Last.fm
	return result, nil
}
