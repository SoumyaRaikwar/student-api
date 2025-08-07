package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SoumyaRaikwar/api_students/internal/config"
)

func main() {
	// Load config
	cfg := config.MustLoad()

	// Setup router
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		w.Write([]byte("Welcome to Students API"))
	})

	// Setup server using corrected field
	server := http.Server{
		Addr:    cfg.HTTPServer.Addr, // ✅ Fix: Access nested address
		Handler: router,
	}

	fmt.Println("🚀 Server starting at", cfg.HTTPServer.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("❌ Failed to start server: ", err)
	}
}
