package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/variableway/innate/capture/internal/bot"
	"github.com/variableway/innate/capture/internal/service"
	"github.com/variableway/innate/capture/internal/store"
)

var (
	botMode string
	botPort int
)

var botServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the Feishu Bot server",
	Long: `Start the Feishu Bot server in webhook or websocket mode.

Webhook mode: Starts an HTTP server to receive Feishu event callbacks.
  Requires a publicly accessible URL configured in Feishu developer console.

WebSocket mode: Connects to Feishu via long-lived WebSocket connection.
  No public URL needed. Recommended for local development.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		dir := getDataDir()

		dualStore, err := store.NewDualStore(dir)
		if err != nil {
			return fmt.Errorf("failed to init store: %w", err)
		}
		defer dualStore.Close()

		taskSvc := service.NewTaskService(dualStore, dir)
		dispatcher := bot.NewDispatcher(taskSvc)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

		switch botMode {
		case "webhook":
			appID := os.Getenv("FEISHU_APP_ID")
			appSecret := os.Getenv("FEISHU_APP_SECRET")
			verifyToken := os.Getenv("FEISHU_VERIFICATION_TOKEN")
			encryptKey := os.Getenv("FEISHU_ENCRYPT_KEY")

			if appID == "" || appSecret == "" {
				return fmt.Errorf("FEISHU_APP_ID and FEISHU_APP_SECRET environment variables are required")
			}

			handler := bot.NewWebhookHandler(dispatcher, verifyToken, encryptKey)
			mux := http.NewServeMux()
			mux.Handle("/webhook/feishu", handler)

			addr := fmt.Sprintf(":%d", botPort)
			server := &http.Server{Addr: addr, Handler: mux}

			go func() {
				log.Printf("Starting webhook server on %s", addr)
				if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					log.Fatalf("Server error: %v", err)
				}
			}()

			<-sigCh
			log.Println("Shutting down...")
			server.Shutdown(ctx)

		case "websocket":
			appID := os.Getenv("FEISHU_APP_ID")
			appSecret := os.Getenv("FEISHU_APP_SECRET")

			if appID == "" || appSecret == "" {
				return fmt.Errorf("FEISHU_APP_ID and FEISHU_APP_SECRET environment variables are required")
			}

			wsHandler := bot.NewWebSocketHandler(dispatcher, appID, appSecret)

			go func() {
				if err := wsHandler.Start(ctx); err != nil {
					log.Printf("WebSocket error: %v", err)
					cancel()
				}
			}()

			<-sigCh
			log.Println("Shutting down...")
			cancel()

		default:
			return fmt.Errorf("invalid mode: %s (use 'webhook' or 'websocket')", botMode)
		}

		return nil
	},
}

func init() {
	botServeCmd.Flags().StringVar(&botMode, "mode", "websocket", "Bot mode: webhook or websocket")
	botServeCmd.Flags().IntVar(&botPort, "port", 8080, "HTTP port for webhook mode")
	botCmd.AddCommand(botServeCmd)
}
