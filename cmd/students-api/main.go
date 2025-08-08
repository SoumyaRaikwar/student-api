package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/SoumyaRaikwar/api_students/internal/config"
	"github.com/SoumyaRaikwar/api_students/internal/http/handlers/student"
)

func main() {
	// Load config
	cfg := config.MustLoad()

	// Setup router
	router := http.NewServeMux()
	router.HandleFunc("POST /api/students", student.New()) // Use the student handler

	// Setup server using corrected field
	server := http.Server{
		Addr:    cfg.HTTPServer.Addr, // ✅ Fix: Access nested address
		Handler: router,
	}

	fmt.Println("🚀 Server starting at", cfg.HTTPServer.Addr)

	done := make(chan os.Signal, 1)
    signal.Notify(done, os.Interrupt,syscall.SIGINT, syscall.SIGTERM)
	go func(){
err := server.ListenAndServe()
	if err != nil {
		log.Fatal("❌ Failed to start server: ", err)
	}
	}()
<-done
slog.Info("🔔 Received shutdown signal, shutting down server gracefully...")
ctx,cancel :=context.WithTimeout(context.Background(), 5 * time.Second)
defer cancel()


err := server.Shutdown(ctx)
if err != nil {
	slog.Error("❌ Error during server shutdown: ",slog.String("error",err.Error()))
} else {
	slog.Info("✅ Server shutdown gracefully")
}
fmt.Println("🛑 Shutting down server...")	
}
