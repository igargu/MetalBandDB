package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"metalbanddb/pkg/handlers"
	"metalbanddb/pkg/services"
	"github.com/gorilla/mux"
	corshandlers "github.com/gorilla/handlers"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// Cargar las variables de entorno desde el archivo .env
	err := godotenv.Load("../.env") 
	if err != nil {
		log.Fatal("Error cargando archivo .env")
	}

	// Obtener la API Key de Last.fm desde las variables de entorno
	apiKey := os.Getenv("LFM_API_KEY")
	if apiKey == "" {
		log.Fatal("La API Key de Last.fm no está configurada")
	}

	// Obtener las credenciales de la base de datos desde las variables de entorno
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSLMODE")

	if dbUser == "" || dbPassword == "" || dbName == "" || dbSSLMode == "" {
		log.Fatal("Las credenciales de la base de datos no están configuradas correctamente")
	}

	// Construir la cadena de conexión a la base de datos
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", dbUser, dbPassword, dbName, dbSSLMode)
	
	// Conectar a la base de datos
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error al conectar a la base de datos: ", err)
	}
	defer db.Close()

	// Configurar el router
	r := mux.NewRouter()

	// Definir las rutas
	r.HandleFunc("/api/bands/{name}/lastfm", handlers.GetBandFromLastFM(db, services.FetchBandInfoFromLastFM)).Methods("GET")

	// Habilitar CORS
    corsHandler := corshandlers.CORS(
        corshandlers.AllowedOrigins([]string{
			"http://localhost:8081",     // Permite localhost
        	"http://192.168.1.147:8081", // Permite la IP local	
		}), // Especifica el origen de tu frontend
        corshandlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
        corshandlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
    )(r)

	// Iniciar el servidor
	fmt.Println("Servidor corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", corsHandler))
}
