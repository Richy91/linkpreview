package linkpreview

import (
	"context"
	"net/http"
	"time"
)

type LinkPreview struct {
	URL         string
	Title       bool
	Description bool
	Image       bool
	SiteName    bool
	Favicon     bool
	Timeout     time.Duration
	UserAgent   string
}

func New(url string, opts ...Option) *LinkPreview {
	lp := &LinkPreview{
		URL: url,
	}

	for _, opt := range opts {
		opt(lp)
	}

	return lp
}

func (l *LinkPreview) GenerateLinkPreview() ([]byte, error) {
	ctx := context.Background()

	req, err := http.NewRequestWithContext(ctx, "GET", l.URL, nil)
	if err != nil {
		return nil, err
	}

	timeout := l.Timeout
	if timeout == 0 {
		timeout = 10 * time.Second
	}

	client := &http.Client{
		Timeout: timeout,
	}

	if l.UserAgent != "" {
		req.Header.Set("User-Agent", l.UserAgent)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	result, err := l.parseResponseBody(resp.Body)
	if err != nil {
		return nil, err
	}

	return result, nil
}
