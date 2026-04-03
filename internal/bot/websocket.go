package bot

import (
	"context"
	"encoding/json"
	"log"

	larkdispatcher "github.com/larksuite/oapi-sdk-go/v3/event/dispatcher"
	larkws "github.com/larksuite/oapi-sdk-go/v3/ws"
	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
)

// WebSocketHandler implements the WebSocket long-connection mode for receiving Feishu events.
type WebSocketHandler struct {
	dispatcher *Dispatcher
	appID      string
	appSecret  string
}

func NewWebSocketHandler(dispatcher *Dispatcher, appID, appSecret string) *WebSocketHandler {
	return &WebSocketHandler{
		dispatcher: dispatcher,
		appID:      appID,
		appSecret:  appSecret,
	}
}

// Start connects to Feishu via WebSocket and starts receiving events.
func (h *WebSocketHandler) Start(ctx context.Context) error {
	eventDispatcher := larkdispatcher.NewEventDispatcher("", "").
		OnP2MessageReceiveV1(func(ctx context.Context, event *larkim.P2MessageReceiveV1) error {
			return h.handleMessage(ctx, event)
		})

	cli := larkws.NewClient(h.appID, h.appSecret,
		larkws.WithEventHandler(eventDispatcher),
		larkws.WithAutoReconnect(true),
	)

	log.Println("Starting Feishu WebSocket connection...")
	return cli.Start(ctx)
}

func (h *WebSocketHandler) handleMessage(ctx context.Context, event *larkim.P2MessageReceiveV1) error {
	if event == nil || event.Event == nil || event.Event.Message == nil {
		return nil
	}

	content := *event.Event.Message.Content
	var textContent struct {
		Text string `json:"text"`
	}
	if err := json.Unmarshal([]byte(content), &textContent); err != nil {
		return err
	}

	text := stripMention(textContent.Text)

	sender := ""
	if event.Event.Sender != nil && event.Event.Sender.SenderId != nil && event.Event.Sender.SenderId.UserId != nil {
		sender = *event.Event.Sender.SenderId.UserId
	}

	response := h.dispatcher.ProcessMessage(ctx, text, sender)
	log.Printf("Bot response: %s", response)

	return nil
}
