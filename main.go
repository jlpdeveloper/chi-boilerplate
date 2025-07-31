package main

import (
	"chi-boilerplate/router"
	"log"
	"net/http"
	"os"
)

func main() {
	r := router.SetupRouter()
	port := getEnvVarValue("PORT", "8080")
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func getEnvVarValue(key string, defaultValue string) string {
	val, found := os.LookupEnv(key)
	if !found {
		log.Println("Environment variable " + key + " not found")
		return defaultValue
	}
	return val
}
