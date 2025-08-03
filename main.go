package main

import (
	"chi-boilerplate/router"
	"log"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	r := router.SetupRouter()
	port := getEnvVarValue("PORT", "8080")
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func getEnvVarValue(key string, defaultValue string) string {
	val, found := os.LookupEnv(key)
	if !found {
		slog.Info("Environment variable " + key + " not found")
		return defaultValue
	}
	return val
}
