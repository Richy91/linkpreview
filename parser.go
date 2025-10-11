package linkpreview

import (
	"encoding/json"
	"io"

	"github.com/PuerkitoBio/goquery"
)

func (l *LinkPreview) parseResponseBody(body io.Reader) ([]byte, error) {
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return nil, err
	}

	data := make(map[string]string)

	if l.Title {
		title := extractMetaContent(doc, "og:title")

		data["title"] = title
	}

	if l.Description {
		description := extractMetaContent(doc, "og:description")

		data["description"] = description
	}

	if l.Image {
		image := extractMetaContent(doc, "og:image")

		data["image"] = image
	}

	if l.Favicon {
		favicon := doc.Find("link[rel='icon']").AttrOr("href", "")

		data["favicon"] = favicon
	}

	if l.SiteName {
		SiteName := extractMetaContent(doc, "og:site_name")

		data["site_name"] = SiteName
	}

	result, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func extractMetaContent(doc *goquery.Document, property string) string {
	return doc.Find("meta[property='"+property+"']").AttrOr("content", "")
}
