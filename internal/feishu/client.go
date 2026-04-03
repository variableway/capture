package feishu

import (
	lark "github.com/larksuite/oapi-sdk-go/v3"
)

// NewClient creates a Feishu API client with app credentials.
func NewClient(appID, appSecret string) *lark.Client {
	return lark.NewClient(appID, appSecret)
}
