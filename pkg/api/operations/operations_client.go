// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new operations API client with basic auth credentials.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - user: user for basic authentication header.
// - password: password for basic authentication header.
func NewClientWithBasicAuth(host, basePath, scheme, user, password string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BasicAuth(user, password)
	return &Client{transport: transport, formats: strfmt.Default}
}

// New creates a new operations API client with a bearer token for authentication.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - bearerToken: bearer token for Bearer authentication header.
func NewClientWithBearerToken(host, basePath, scheme, bearerToken string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(bearerToken)
	return &Client{transport: transport, formats: strfmt.Default}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetLanguages(params *GetLanguagesParams, opts ...ClientOption) (*GetLanguagesOK, error)

	GetWords(params *GetWordsParams, opts ...ClientOption) (*GetWordsOK, error)

	PostCheck(params *PostCheckParams, opts ...ClientOption) (*PostCheckOK, error)

	PostWordsAdd(params *PostWordsAddParams, opts ...ClientOption) (*PostWordsAddOK, error)

	PostWordsDelete(params *PostWordsDeleteParams, opts ...ClientOption) (*PostWordsDeleteOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetLanguages gets a list of supported languages
*/
func (a *Client) GetLanguages(params *GetLanguagesParams, opts ...ClientOption) (*GetLanguagesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLanguagesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetLanguages",
		Method:             "GET",
		PathPattern:        "/languages",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLanguagesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetLanguagesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetLanguages: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetWords lists words in dictionaries

List words in the user's personal dictionaries.
*/
func (a *Client) GetWords(params *GetWordsParams, opts ...ClientOption) (*GetWordsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWordsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetWords",
		Method:             "GET",
		PathPattern:        "/words",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetWordsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetWordsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetWords: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostCheck checks a text

The main feature - check a text with LanguageTool for possible style and grammar issues.
*/
func (a *Client) PostCheck(params *PostCheckParams, opts ...ClientOption) (*PostCheckOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostCheckParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostCheck",
		Method:             "POST",
		PathPattern:        "/check",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostCheckReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostCheckOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostCheck: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostWordsAdd adds word to a dictionary

Add a word to one of the user's personal dictionaries. Please note that this feature is considered to be used for personal dictionaries which must not contain more than 500 words. If this is an issue for you, please contact us.
*/
func (a *Client) PostWordsAdd(params *PostWordsAddParams, opts ...ClientOption) (*PostWordsAddOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostWordsAddParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostWordsAdd",
		Method:             "POST",
		PathPattern:        "/words/add",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostWordsAddReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostWordsAddOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostWordsAdd: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostWordsDelete removes word from a dictionary

Remove a word from one of the user's personal dictionaries.
*/
func (a *Client) PostWordsDelete(params *PostWordsDeleteParams, opts ...ClientOption) (*PostWordsDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostWordsDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostWordsDelete",
		Method:             "POST",
		PathPattern:        "/words/delete",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostWordsDeleteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostWordsDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostWordsDelete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
