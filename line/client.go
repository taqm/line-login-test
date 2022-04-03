package line

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"golang.org/x/oauth2"
)

type LineApiClient struct {
	client *http.Client
}

func NewLineApiClient(ctx context.Context, tok *oauth2.Token) *LineApiClient {
	client := oauthConfig.Client(ctx, tok)
	return &LineApiClient{client}
}

func NewLineApiClientByCode(ctx context.Context, code string) (*LineApiClient, error) {
	tok, err := oauthConfig.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}
	client := oauthConfig.Client(ctx, tok)
	return &LineApiClient{client}, nil
}

type GetProfileResponse struct {
	UserID      string
	DisplayName string
	PictureURL  string
}

func (c LineApiClient) GetProfile() (GetProfileResponse, error) {
	resp, err := c.client.Get("https://api.line.me/v2/profile")
	if err != nil {
		return GetProfileResponse{}, nil
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return GetProfileResponse{}, nil
	}

	var ret GetProfileResponse
	err = json.Unmarshal(b, &ret)
	if err != nil {
		return GetProfileResponse{}, nil
	}
	return ret, nil
}
