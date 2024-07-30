package tools

import (
	"context"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/config"
)

type Http interface {
	Get(url string) (resp *http.Response, err error)
	Delete(url string) (resp *http.Response, err error)
	Post(url, contentType string, body io.Reader) (resp *http.Response, err error)
	Put(url, contentType string, body io.Reader) (resp *http.Response, err error)
}

type HttpCreator func(cfg HttpConfig) Http

// Configuration for the requests
type HttpConfig struct {
	AWSv4Signed   bool
	Authorization *string
	Endpoint      string
}

type defaultHttpClient struct {
	Http

	config HttpConfig
}

var (
	httpCreator HttpCreator
)

// Creates a default http client
func defaultCreator(cfg HttpConfig) Http {
	return &defaultHttpClient{
		config: cfg,
	}
}

// Returns a http instance for a given configuration
func GetHttpInstance(cfg HttpConfig) Http {
	if httpCreator == nil {
		httpCreator = defaultCreator
	}

	return httpCreator(cfg)
}

// Overrides the default http creator function
func SetHttpCreator(creator HttpCreator) {
	httpCreator = creator
}

func (c defaultHttpClient) Get(url string) (resp *http.Response, err error) {
	req, err := http.NewRequest("GET", c.getUrl(url), nil)

	if err != nil {
		return nil, err
	}

	return executeRequest(req, c.config)
}

func (c defaultHttpClient) Post(url, contentType string, body io.Reader) (resp *http.Response, err error) {
	req, err := http.NewRequest("POST", c.getUrl(url), body)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", contentType)

	return executeRequest(req, c.config)
}

func (c defaultHttpClient) Put(url, contentType string, body io.Reader) (resp *http.Response, err error) {
	req, err := http.NewRequest("PUT", c.getUrl(url), body)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", contentType)

	return executeRequest(req, c.config)
}

func (c defaultHttpClient) Delete(url string) (resp *http.Response, err error) {
	req, err := http.NewRequest("DELETE", c.getUrl(url), nil)

	if err != nil {
		return nil, err
	}

	return executeRequest(req, c.config)
}

func executeRequest(request *http.Request, cfg HttpConfig) (*http.Response, error) {
	prepareGeneralHeaders(request)

	if cfg.AWSv4Signed {
		err := awsSignRequest(request)
		if err != nil {
			return nil, err
		}
	} else if cfg.Authorization != nil {
		addAuthorization(request, cfg)
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("HTTP: Error while sending a request: %s", err.Error())
	}

	return response, err
}

func addAuthorization(request *http.Request, cfg HttpConfig) {
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", *cfg.Authorization))
}

func prepareGeneralHeaders(request *http.Request) {
	request.Header.Add("Accept", "application/json")
}

func awsSignRequest(request *http.Request) error {
	ctx := context.TODO()

	awsConfig, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return err
	}

	creds, err := awsConfig.Credentials.Retrieve(context.TODO())
	if err != nil {
		return err
	}

	request = request.WithContext(context.TODO())
	signer := v4.NewSigner()
	hasher := sha256.New()

	body := []byte{}
	if request.Body != nil {
		body, err = io.ReadAll(request.Body)
		if err != nil {
			return err
		}
	}

	hasher.Write(body)
	bodyHash := fmt.Sprintf("%x", hasher.Sum(nil))

	return signer.SignHTTP(context.TODO(), creds, request, bodyHash, "execute-api", awsConfig.Region, GetTimeInstance().Now())
}

func (c defaultHttpClient) getUrl(url string) string {
	if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
		return url
	}

	return fmt.Sprintf("%s/%s", strings.TrimRight(c.config.Endpoint, "/"), strings.TrimLeft(url, "/"))
}
