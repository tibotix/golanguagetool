package golanguagetool

import (
	"fmt"
	"net/url"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/tibotix/golanguagetool/pkg/api"
)

type Client struct {
	ApiClient *api.LanguageToolAPI
	ApiKey    string
	Username  string
}

func NewClient() *Client {
	return &Client{
		ApiClient: api.Default,
		ApiKey:    "",
		Username:  "",
	}
}

func NewClientWithApiUrl(apiUrl string) (*Client, error) {
	u, err := url.Parse(apiUrl)
	if err != nil {
		return nil, fmt.Errorf("invalid API Url: %s", apiUrl)
	}

	transport := httptransport.New(u.Host, u.Path, []string{u.Scheme})
	apiClient := api.New(transport, strfmt.Default)
	return &Client{
		ApiClient: apiClient,
		ApiKey:    "",
		Username:  "",
	}, nil
}
