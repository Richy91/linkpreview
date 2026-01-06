package linkpreview

import (
	"encoding/json"
	"io"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

func (l *LinkPreview) parseResponseBody(body io.Reader) ([]byte, error) {
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return nil, err
	}

	data := make(map[string]any)

	data["type"] = "opengraph"
	data["url"] = l.URL

	if l.Title {
		title := extractMetaContent(doc, "property", "og:title")

		if title == "" {
			title = doc.Find("title").Text()

			if title != "" {
				data["type"] = "html"
			}
		}

		data["title"] = title
	}

	if l.Description {
		description := extractMetaContent(doc, "property", "og:description")

		if description == "" {
			description = extractMetaContent(doc, "name", "description")
		}

		data["description"] = description
	}

	if l.Image {
		imageURL := extractMetaContent(doc, "property", "og:image")
		width := extractMetaContent(doc, "property", "og:image:width")
		height := extractMetaContent(doc, "property", "og:image:height")

		widthInt, err := strconv.Atoi(width)
		if err != nil {
			widthInt = 0
		}

		heightInt, err := strconv.Atoi(height)
		if err != nil {
			heightInt = 0
		}

		data["image"] = map[string]any{
			"url":    imageURL,
			"width":  widthInt,
			"height": heightInt,
		}
	}

	if l.Favicon {
		favicon := doc.Find("link[rel='icon']").AttrOr("href", "")
		if favicon == "" {
			favicon = doc.Find("link[rel='shortcut icon']").AttrOr("href", "")

			if favicon == "" {
				//apple-touch-icon is often used as favicon
				favicon = doc.Find("link[rel='apple-touch-icon']").AttrOr("href", "")
			}
		}

		data["favicon"] = favicon
	}

	if l.SiteName {
		SiteName := extractMetaContent(doc, "property", "og:site_name")

		data["site_name"] = SiteName
	}

	result, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func extractMetaContent(doc *goquery.Document, key, value string) string {
	return doc.Find("meta["+key+"='"+value+"']").AttrOr("content", "")
}
