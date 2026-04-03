package bitable

import (
	"context"
	"fmt"

	lark "github.com/larksuite/oapi-sdk-go/v3"
	larkbitable "github.com/larksuite/oapi-sdk-go/v3/service/bitable/v1"
)

// Client wraps Feishu Bitable API operations.
type Client struct {
	client   *lark.Client
	appToken string
	tableID  string
}

func NewClient(client *lark.Client, appToken, tableID string) *Client {
	return &Client{
		client:   client,
		appToken: appToken,
		tableID:  tableID,
	}
}

// CreateRecord creates a new record in the Bitable table.
func (c *Client) CreateRecord(ctx context.Context, fields map[string]interface{}) (string, error) {
	req := larkbitable.NewCreateAppTableRecordReqBuilder().
		AppToken(c.appToken).
		TableId(c.tableID).
		AppTableRecord(&larkbitable.AppTableRecord{Fields: fields}).
		Build()

	resp, err := c.client.Bitable.AppTableRecord.Create(ctx, req)
	if err != nil {
		return "", fmt.Errorf("failed to create record: %w", err)
	}

	if resp == nil || resp.Data == nil || resp.Data.Record == nil || resp.Data.Record.RecordId == nil {
		return "", fmt.Errorf("empty response from Bitable API")
	}

	return *resp.Data.Record.RecordId, nil
}

// UpdateRecord updates an existing record in the Bitable table.
func (c *Client) UpdateRecord(ctx context.Context, recordID string, fields map[string]interface{}) error {
	req := larkbitable.NewUpdateAppTableRecordReqBuilder().
		AppToken(c.appToken).
		TableId(c.tableID).
		RecordId(recordID).
		AppTableRecord(&larkbitable.AppTableRecord{Fields: fields}).
		Build()

	_, err := c.client.Bitable.AppTableRecord.Update(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to update record: %w", err)
	}

	return nil
}

// DeleteRecord deletes a record from the Bitable table.
func (c *Client) DeleteRecord(ctx context.Context, recordID string) error {
	req := larkbitable.NewDeleteAppTableRecordReqBuilder().
		AppToken(c.appToken).
		TableId(c.tableID).
		RecordId(recordID).
		Build()

	_, err := c.client.Bitable.AppTableRecord.Delete(ctx, req)
	return err
}

// ListRecords lists records from the Bitable table.
func (c *Client) ListRecords(ctx context.Context) ([]*larkbitable.AppTableRecord, error) {
	req := larkbitable.NewListAppTableRecordReqBuilder().
		AppToken(c.appToken).
		TableId(c.tableID).
		Build()

	resp, err := c.client.Bitable.AppTableRecord.List(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to list records: %w", err)
	}

	if resp == nil || resp.Data == nil {
		return nil, nil
	}

	return resp.Data.Items, nil
}
