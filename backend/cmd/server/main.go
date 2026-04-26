package main

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"aiwiki/backend/internal/article"
	"aiwiki/backend/internal/demo"
	"aiwiki/backend/internal/hosted"
	httpapi "aiwiki/backend/internal/http"
	"aiwiki/backend/internal/ollama"
)

func main() {
	addr := listenAddr()

	generator := generatorFromEnv()
	service := article.NewService(generator, 10*time.Minute)
	handler := httpapi.NewHandler(service)

	provider := "model"
	if named, ok := generator.(article.ProviderNamer); ok {
		provider = named.Provider()
	}
	log.Printf("AIWIKI backend listening on %s using %s model %s", addr, provider, generator.Model())
	if err := http.ListenAndServe(addr, handler.Routes()); err != nil {
		log.Fatal(err)
	}
}

func generatorFromEnv() article.Generator {
	switch strings.ToLower(strings.TrimSpace(os.Getenv("AIWIKI_MODEL_PROVIDER"))) {
	case "demo", "template":
		return demo.NewClient()
	case "hosted", "openai", "openai-compatible", "openrouter":
		return hosted.NewClient()
	default:
		return ollama.NewClient()
	}
}

func listenAddr() string {
	if port := strings.TrimSpace(os.Getenv("PORT")); port != "" {
		if strings.Contains(port, ":") {
			return port
		}
		return ":" + port
	}

	if addr := strings.TrimSpace(os.Getenv("AIWIKI_ADDR")); addr != "" {
		return addr
	}

	return ":8080"
}
