// Package client provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/oapi-codegen/runtime"
)

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// AnalyzeDependencies request
	AnalyzeDependencies(ctx context.Context, params *AnalyzeDependenciesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// HealthCheck request
	HealthCheck(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetArtifactDeps request
	GetArtifactDeps(ctx context.Context, digest string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetArtifactVulns request
	GetArtifactVulns(ctx context.Context, digest string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetPackagePurls request
	GetPackagePurls(ctx context.Context, purl string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetPackageDeps request
	GetPackageDeps(ctx context.Context, purl string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetPackageVulns request
	GetPackageVulns(ctx context.Context, purl string, params *GetPackageVulnsParams, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) AnalyzeDependencies(ctx context.Context, params *AnalyzeDependenciesParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAnalyzeDependenciesRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) HealthCheck(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewHealthCheckRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetArtifactDeps(ctx context.Context, digest string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetArtifactDepsRequest(c.Server, digest)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetArtifactVulns(ctx context.Context, digest string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetArtifactVulnsRequest(c.Server, digest)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetPackagePurls(ctx context.Context, purl string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetPackagePurlsRequest(c.Server, purl)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetPackageDeps(ctx context.Context, purl string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetPackageDepsRequest(c.Server, purl)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetPackageVulns(ctx context.Context, purl string, params *GetPackageVulnsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetPackageVulnsRequest(c.Server, purl, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewAnalyzeDependenciesRequest generates requests for AnalyzeDependencies
func NewAnalyzeDependenciesRequest(server string, params *AnalyzeDependenciesParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/analysis/dependencies")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if params.PaginationSpec != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "paginationSpec", runtime.ParamLocationQuery, *params.PaginationSpec); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "sort", runtime.ParamLocationQuery, params.Sort); err != nil {
			return nil, err
		} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
			return nil, err
		} else {
			for k, v := range parsed {
				for _, v2 := range v {
					queryValues.Add(k, v2)
				}
			}
		}

		queryURL.RawQuery = queryValues.Encode()
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewHealthCheckRequest generates requests for HealthCheck
func NewHealthCheckRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/healthz")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetArtifactDepsRequest generates requests for GetArtifactDeps
func NewGetArtifactDepsRequest(server string, digest string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "digest", runtime.ParamLocationPath, digest)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v0/artifact/%s/dependencies", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetArtifactVulnsRequest generates requests for GetArtifactVulns
func NewGetArtifactVulnsRequest(server string, digest string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "digest", runtime.ParamLocationPath, digest)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v0/artifact/%s/vulns", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetPackagePurlsRequest generates requests for GetPackagePurls
func NewGetPackagePurlsRequest(server string, purl string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "purl", runtime.ParamLocationPath, purl)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v0/package/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetPackageDepsRequest generates requests for GetPackageDeps
func NewGetPackageDepsRequest(server string, purl string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "purl", runtime.ParamLocationPath, purl)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v0/package/%s/dependencies", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetPackageVulnsRequest generates requests for GetPackageVulns
func NewGetPackageVulnsRequest(server string, purl string, params *GetPackageVulnsParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "purl", runtime.ParamLocationPath, purl)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v0/package/%s/vulns", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if params.IncludeDependencies != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "includeDependencies", runtime.ParamLocationQuery, *params.IncludeDependencies); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		queryURL.RawQuery = queryValues.Encode()
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// AnalyzeDependenciesWithResponse request
	AnalyzeDependenciesWithResponse(ctx context.Context, params *AnalyzeDependenciesParams, reqEditors ...RequestEditorFn) (*AnalyzeDependenciesResponse, error)

	// HealthCheckWithResponse request
	HealthCheckWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*HealthCheckResponse, error)

	// GetArtifactDepsWithResponse request
	GetArtifactDepsWithResponse(ctx context.Context, digest string, reqEditors ...RequestEditorFn) (*GetArtifactDepsResponse, error)

	// GetArtifactVulnsWithResponse request
	GetArtifactVulnsWithResponse(ctx context.Context, digest string, reqEditors ...RequestEditorFn) (*GetArtifactVulnsResponse, error)

	// GetPackagePurlsWithResponse request
	GetPackagePurlsWithResponse(ctx context.Context, purl string, reqEditors ...RequestEditorFn) (*GetPackagePurlsResponse, error)

	// GetPackageDepsWithResponse request
	GetPackageDepsWithResponse(ctx context.Context, purl string, reqEditors ...RequestEditorFn) (*GetPackageDepsResponse, error)

	// GetPackageVulnsWithResponse request
	GetPackageVulnsWithResponse(ctx context.Context, purl string, params *GetPackageVulnsParams, reqEditors ...RequestEditorFn) (*GetPackageVulnsResponse, error)
}

type AnalyzeDependenciesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *PackageNameList
	JSON400      *BadRequest
	JSON500      *InternalServerError
	JSON502      *BadGateway
}

// Status returns HTTPResponse.Status
func (r AnalyzeDependenciesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r AnalyzeDependenciesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type HealthCheckResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *string
}

// Status returns HTTPResponse.Status
func (r HealthCheckResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r HealthCheckResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetArtifactDepsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *PurlList
	JSON400      *BadRequest
	JSON500      *InternalServerError
	JSON502      *BadGateway
}

// Status returns HTTPResponse.Status
func (r GetArtifactDepsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetArtifactDepsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetArtifactVulnsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *VulnerabilityList
	JSON400      *BadRequest
	JSON500      *InternalServerError
	JSON502      *BadGateway
}

// Status returns HTTPResponse.Status
func (r GetArtifactVulnsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetArtifactVulnsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetPackagePurlsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *PurlList
	JSON400      *BadRequest
	JSON500      *InternalServerError
	JSON502      *BadGateway
}

// Status returns HTTPResponse.Status
func (r GetPackagePurlsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetPackagePurlsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetPackageDepsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *PurlList
	JSON400      *BadRequest
	JSON500      *InternalServerError
	JSON502      *BadGateway
}

// Status returns HTTPResponse.Status
func (r GetPackageDepsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetPackageDepsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetPackageVulnsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *VulnerabilityList
	JSON400      *BadRequest
	JSON500      *InternalServerError
	JSON502      *BadGateway
}

// Status returns HTTPResponse.Status
func (r GetPackageVulnsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetPackageVulnsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// AnalyzeDependenciesWithResponse request returning *AnalyzeDependenciesResponse
func (c *ClientWithResponses) AnalyzeDependenciesWithResponse(ctx context.Context, params *AnalyzeDependenciesParams, reqEditors ...RequestEditorFn) (*AnalyzeDependenciesResponse, error) {
	rsp, err := c.AnalyzeDependencies(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAnalyzeDependenciesResponse(rsp)
}

// HealthCheckWithResponse request returning *HealthCheckResponse
func (c *ClientWithResponses) HealthCheckWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*HealthCheckResponse, error) {
	rsp, err := c.HealthCheck(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseHealthCheckResponse(rsp)
}

// GetArtifactDepsWithResponse request returning *GetArtifactDepsResponse
func (c *ClientWithResponses) GetArtifactDepsWithResponse(ctx context.Context, digest string, reqEditors ...RequestEditorFn) (*GetArtifactDepsResponse, error) {
	rsp, err := c.GetArtifactDeps(ctx, digest, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetArtifactDepsResponse(rsp)
}

// GetArtifactVulnsWithResponse request returning *GetArtifactVulnsResponse
func (c *ClientWithResponses) GetArtifactVulnsWithResponse(ctx context.Context, digest string, reqEditors ...RequestEditorFn) (*GetArtifactVulnsResponse, error) {
	rsp, err := c.GetArtifactVulns(ctx, digest, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetArtifactVulnsResponse(rsp)
}

// GetPackagePurlsWithResponse request returning *GetPackagePurlsResponse
func (c *ClientWithResponses) GetPackagePurlsWithResponse(ctx context.Context, purl string, reqEditors ...RequestEditorFn) (*GetPackagePurlsResponse, error) {
	rsp, err := c.GetPackagePurls(ctx, purl, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetPackagePurlsResponse(rsp)
}

// GetPackageDepsWithResponse request returning *GetPackageDepsResponse
func (c *ClientWithResponses) GetPackageDepsWithResponse(ctx context.Context, purl string, reqEditors ...RequestEditorFn) (*GetPackageDepsResponse, error) {
	rsp, err := c.GetPackageDeps(ctx, purl, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetPackageDepsResponse(rsp)
}

// GetPackageVulnsWithResponse request returning *GetPackageVulnsResponse
func (c *ClientWithResponses) GetPackageVulnsWithResponse(ctx context.Context, purl string, params *GetPackageVulnsParams, reqEditors ...RequestEditorFn) (*GetPackageVulnsResponse, error) {
	rsp, err := c.GetPackageVulns(ctx, purl, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetPackageVulnsResponse(rsp)
}

// ParseAnalyzeDependenciesResponse parses an HTTP response from a AnalyzeDependenciesWithResponse call
func ParseAnalyzeDependenciesResponse(rsp *http.Response) (*AnalyzeDependenciesResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &AnalyzeDependenciesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest PackageNameList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequest
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalServerError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 502:
		var dest BadGateway
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON502 = &dest

	}

	return response, nil
}

// ParseHealthCheckResponse parses an HTTP response from a HealthCheckWithResponse call
func ParseHealthCheckResponse(rsp *http.Response) (*HealthCheckResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &HealthCheckResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest string
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseGetArtifactDepsResponse parses an HTTP response from a GetArtifactDepsWithResponse call
func ParseGetArtifactDepsResponse(rsp *http.Response) (*GetArtifactDepsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetArtifactDepsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest PurlList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequest
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalServerError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 502:
		var dest BadGateway
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON502 = &dest

	}

	return response, nil
}

// ParseGetArtifactVulnsResponse parses an HTTP response from a GetArtifactVulnsWithResponse call
func ParseGetArtifactVulnsResponse(rsp *http.Response) (*GetArtifactVulnsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetArtifactVulnsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest VulnerabilityList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequest
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalServerError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 502:
		var dest BadGateway
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON502 = &dest

	}

	return response, nil
}

// ParseGetPackagePurlsResponse parses an HTTP response from a GetPackagePurlsWithResponse call
func ParseGetPackagePurlsResponse(rsp *http.Response) (*GetPackagePurlsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetPackagePurlsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest PurlList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequest
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalServerError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 502:
		var dest BadGateway
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON502 = &dest

	}

	return response, nil
}

// ParseGetPackageDepsResponse parses an HTTP response from a GetPackageDepsWithResponse call
func ParseGetPackageDepsResponse(rsp *http.Response) (*GetPackageDepsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetPackageDepsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest PurlList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequest
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalServerError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 502:
		var dest BadGateway
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON502 = &dest

	}

	return response, nil
}

// ParseGetPackageVulnsResponse parses an HTTP response from a GetPackageVulnsWithResponse call
func ParseGetPackageVulnsResponse(rsp *http.Response) (*GetPackageVulnsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetPackageVulnsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest VulnerabilityList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequest
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalServerError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 502:
		var dest BadGateway
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON502 = &dest

	}

	return response, nil
}
