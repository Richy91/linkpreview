package linkpreview

import "time"

type Option func(*LinkPreview)

func WithTitle(v bool) Option {
	return func(lp *LinkPreview) {
		lp.Title = v
	}
}

func WithDescription(v bool) Option {
	return func(lp *LinkPreview) {
		lp.Description = v
	}
}

func WithSiteName(v bool) Option {
	return func(lp *LinkPreview) {
		lp.SiteName = v
	}
}

func WithImage(v bool) Option {
	return func(lp *LinkPreview) {
		lp.Image = v
	}
}

func WithFavicon(v bool) Option {
	return func(lp *LinkPreview) {
		lp.Favicon = v
	}
}

func WithTimeout(d time.Duration) Option {
	return func(lp *LinkPreview) {
		lp.Timeout = d
	}
}

func WithUserAgent(ua string) Option {
	return func(lp *LinkPreview) {
		lp.UserAgent = ua
	}
}
