package provider

import (
	"context"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GoogleProvider struct{}

func (g GoogleProvider) Provider() OAuth2Provider {
	return "google"
}

func (g GoogleProvider) NewConfig(ClientID, ClientSecret string) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     ClientID,
		ClientSecret: ClientSecret,
		Scopes:       []string{"profile"},
		Endpoint:     google.Endpoint,
	}
}

func (g GoogleProvider) GetUserInfo(ctx context.Context, config *oauth2.Config, code string) (*UserInfo, error) {
	oauth2Token, err := config.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}
	client := config.Client(ctx, oauth2Token)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://www.googleapis.com/oauth2/v2/userinfo", nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return nil, FormatErrNotImplemented("google")
}

func init() {
	RegisterProvider(GoogleProvider{})
}
