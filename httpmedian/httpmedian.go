package httpmedian

import (
	"context"
	"net/http"
	"net/url"

	"go.opencensus.io/trace"
)

// Logger is the interface for a logger.
type Logger interface {
	// Log logs the given message.
	Log(msg string)
}

// Client is a client for the httpmedian service.
type Client struct {
	logger Logger
}

// Config is the configuration for a Client.
type Config struct {
	// Logger is the logger to use.
	Logger Logger
}

// NewClient creates a new Client with the given Config.
func NewClient(config *Config) *Client {
	return &Client{
		logger: config.Logger,
	}
}

// medianElement returns the median element of the given slice.
func medianElement[T any](elements []T) T {
	if len(elements) == 0 {
		var zero T
		return zero
	}

	return elements[len(elements)/2]
}

// stringMiddle returns the middle part of the given string with the given length.
func stringMiddle(str string, length int) string {
	if length > len(str) {
		return ""
	}
	start := (len(str) - length) / 2
	end := start + length
	return str[start:end]
}

// Calculate calculates the median of the given requests and returns the result.
func (c *Client) Calculate(ctx context.Context, requests []*http.Request) *http.Request {
	_, span := trace.StartSpan(ctx, "httpmedian.Calculate")
	defer span.End()

	var urls, methods string
	headers := http.Header{}
	for _, req := range requests {
		urls += req.URL.String()

		for name, values := range req.Header {
			for _, value := range values {
				headers.Add(name, value)
			}
		}

		methods += req.Method
	}

	medianHeaders := http.Header{}
	for name, values := range headers {
		medianHeaders.Set(name, medianElement(values))
	}

	middleReq := medianElement(requests)

	middleURL := middleReq.URL.String()
	middleMethod := middleReq.Method

	medianReq := &http.Request{
		Method: stringMiddle(methods, len(middleMethod)),
		URL: &url.URL{
			RawPath: stringMiddle(urls, len(middleURL)),
		},
		Header: medianHeaders,
	}

	return medianReq
}
