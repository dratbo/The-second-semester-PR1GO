package authclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	baseURL    string
	httpClient *http.Client
}

func NewClient(baseURL string, timeout time.Duration) *Client {
	return &Client{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: timeout,
		},
	}
}

type verifyResponse struct {
	Valid   bool   `json:"valid"`
	Subject string `json:"subject"`
	Error   string `json:"error"`
}

var (
	ErrUnauthorized           = fmt.Errorf("unauthorized")
	ErrAuthServiceUnavailable = fmt.Errorf("auth service unavailable")
)

func (c *Client) VerifyToken(ctx context.Context, token, requestID string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL+"/v1/auth/verify", nil)
	if err != nil {
		return "", fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+token)
	if requestID != "" {
		req.Header.Set("X-Request-ID", requestID) // Пробрасываем request-id для сквозной трассировки.
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("%w: %v", ErrAuthServiceUnavailable, err)
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		var vResp verifyResponse
		if err := json.NewDecoder(resp.Body).Decode(&vResp); err != nil {
			return "", fmt.Errorf("decode response: %w", err)
		}
		if !vResp.Valid {
			return "", ErrUnauthorized
		}
		return vResp.Subject, nil
	case http.StatusUnauthorized, http.StatusForbidden:
		return "", ErrUnauthorized
	default:
		return "", fmt.Errorf("%w: status %d", ErrAuthServiceUnavailable, resp.StatusCode)
	}
}
