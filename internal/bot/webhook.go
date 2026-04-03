package bot

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

// WebhookHandler implements the HTTP webhook mode for receiving Feishu events.
type WebhookHandler struct {
	dispatcher        *Dispatcher
	verificationToken string
	encryptKey        string
}

func NewWebhookHandler(dispatcher *Dispatcher, verificationToken, encryptKey string) *WebhookHandler {
	return &WebhookHandler{
		dispatcher:        dispatcher,
		verificationToken: verificationToken,
		encryptKey:        encryptKey,
	}
}

// ServeHTTP handles incoming Feishu webhook requests.
func (h *WebhookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Handle URL verification challenge
	var challengeReq struct {
		Challenge string `json:"challenge"`
		Token     string `json:"token"`
		Type      string `json:"type"`
	}
	if err := json.Unmarshal(body, &challengeReq); err == nil && challengeReq.Type == "url_verification" {
		if challengeReq.Token != h.verificationToken {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"challenge": challengeReq.Challenge})
		return
	}

	// Parse event
	var event struct {
		Schema string `json:"schema"`
		Header struct {
			EventID   string `json:"event_id"`
			Token     string `json:"token"`
			EventType string `json:"event_type"`
		} `json:"header"`
		Event struct {
			Message struct {
				MessageID string `json:"message_id"`
				Content   string `json:"content"`
				ChatType  string `json:"chat_type"`
				ChatID    string `json:"chat_id"`
			} `json:"message"`
			Sender struct {
				SenderID struct {
					UnionID string `json:"union_id"`
					UserID  string `json:"user_id"`
				} `json:"sender_id"`
			} `json:"sender"`
		} `json:"event"`
	}

	if err := json.Unmarshal(body, &event); err != nil {
		log.Printf("Failed to parse event: %v", err)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	// Verify token
	if h.verificationToken != "" && event.Header.Token != h.verificationToken {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	// Extract message text
	var content struct {
		Text string `json:"text"`
	}
	if err := json.Unmarshal([]byte(event.Event.Message.Content), &content); err != nil {
		log.Printf("Failed to parse message content: %v", err)
		w.WriteHeader(http.StatusOK)
		return
	}

	text := stripMention(content.Text)
	sender := event.Event.Sender.SenderID.UserID
	response := h.dispatcher.ProcessMessage(context.Background(), text, sender)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"content": response})
}

func stripMention(text string) string {
	// Remove @user mentions
	result := text
	for strings.Contains(result, "@") {
		idx := strings.Index(result, "@")
		end := strings.Index(result[idx:], " ")
		if end == -1 {
			result = result[:idx]
		} else {
			result = result[:idx] + result[idx+end:]
		}
	}
	return strings.TrimSpace(result)
}
