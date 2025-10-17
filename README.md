# linkpreview

Small, opinionated Go library for generating link previews (title, description, image, favicon, site name) using `goquery`.

-   ✅ Open Graph + standard HTML fallbacks
-   ✅ Functional options (pick only what you need)
-   ✅ Sensible defaults (timeout, UA)
-   ✅ Tiny surface area

---

## Install

```bash
go get github.com/Richy91/linkpreview
```

---

## Use

```
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Richy91/linkpreview"
)

func main() {
	lp := linkpreview.New(
		"https://github.com/PuerkitoBio/goquery",
		linkpreview.WithTitle(true),
		linkpreview.WithDescription(true),
		linkpreview.WithImage(true),
		linkpreview.WithFavicon(true),
		linkpreview.WithSiteName(true),
		linkpreview.WithTimeout(10*time.Second),
	)

	// GenerateLinkPreview returns JSON ([]byte) with the requested fields
	data, err := lp.GenerateLinkPreview()
	if err != nil {
		log.Fatalf("preview error: %v", err)
	}

	fmt.Printf("%s\n", data)
}
```

## output

```
{
  "url": "https://github.com/PuerkitoBio/goquery",
  "title": "PuerkitoBio/goquery",
  "description": "A little like that jQuery you know...",
  "image": "https://opengraph.githubassets.com/...",
  "favicon": "https://github.githubassets.com/favicons/favicon.svg",
  "site_name": "GitHub"
}
```
