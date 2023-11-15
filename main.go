package main

import (
	"fmt"
	// "log"
	"net/http"
	// "virtual_machine/config"
	// "virtual_machine/server"

	// "net/http"
	//"os"

	_ "github.com/lib/pq"
)

// func main() {
// 	log.Println("starting restapi")

// 	log.Println("initializing configuration")
// 	config := config.InitConfig(getConfigFileName())

// 	log.Println("initializing database")
// 	dbHandler := server.InitDatabase(config)
// 	log.Println(dbHandler)

// 	log.Println("Initializing Http Server")
// 	httpServer := server.InitHttpServer(config, dbHandler)

// 	httpServer.Start()
//}

//func getConfigFileName() string {
//	env := os.Getenv("ENV")

//	if env != "" {
//		return "db_virtualMachine" + env
//	}
//	return "db_virtualMachine"
//}

func authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Authentication logic here")
		next.ServeHTTP(w, r)
	})
}

func authorize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Authorization logic here")
		next.ServeHTTP(w, r)
	})
}

func employeeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Employee service logic here")
	w.Write([]byte("Employee service response"))
}

func main() {
	router := http.NewServeMux()

	apiGatewayHandler := authenticate(authorize(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		employeeHandler(w, r)
	})))

	router.Handle("/api", apiGatewayHandler)

	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", router)
}
