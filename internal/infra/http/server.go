package http

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func Server(ctx context.Context, router http.Handler) error {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", 8080),
		Handler: router,
	}

	errServeCh := make(chan error)
	go func() {
		fmt.Println("Starting server on :8080") // Повідомлення про запуск сервера
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			errServeCh <- err
		}
	}()

	select {
	case <-ctx.Done():
		fmt.Println("Server is shutting down...") // Повідомлення про зупинку сервера
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
		defer cancel()
		if err := srv.Shutdown(shutdownCtx); err != nil {
			return fmt.Errorf("error during server shutdown: %w", err)
		}
		fmt.Println("Server gracefully stopped") // Повідомлення про успішну зупинку сервера
	case err := <-errServeCh:
		return fmt.Errorf("error during server execution: %w", err)
	}
	return nil
}
