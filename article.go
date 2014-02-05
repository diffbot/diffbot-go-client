// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package diffbot

import (
	"encoding/json"
)

// Article represents an clean article text.
//
// See http://diffbot.com/dev/docs/article/
type Article struct {
	Url           string                 `json:"url"`
	ResolvedUrl   string                 `json:"resolved_url"`
	Icon          string                 `json:"icon"`
	Meta          map[string]interface{} `json:"meta,omitempty"`        // Returned with fields.
	QueryString   string                 `json:"querystring,omitempty"` // Returned with fields.
	Links         []string               `json:"links,omitempty"`       // Returned with fields.
	Type          string                 `json:"type"`
	Title         string                 `json:"title"`
	Text          string                 `json:"text"`
	Html          string                 `json:"html"`
	NumPages      string                 `json:"numPages"`
	Date          string                 `json:"date"`
	Author        string                 `json:"author"`
	Tags          []string               `json:"tags,omitempty"`          // Returned with fields.
	HumanLanguage string                 `json:"humanLanguage,omitempty"` // Returned with fields.
	Images        []struct {
		Url         string `json:"url"`
		PixelHeight int    `json:"pixelHeight"`
		PixelWidth  int    `json:"pixelWidth"`
		Caption     string `json:"caption"`
		Primary     string `json:"primary"`
	} `json:"images"`
	Videos []struct {
		Url         string `json:"url"`
		PixelHeight int    `json:"pixelHeight"`
		PixelWidth  int    `json:"pixelWidth"`
		Primary     string `json:"primary"`
	} `json:"videos"`
}

// type of Article.Images[?]
type articleImageType struct {
	Url         string `json:"url"`
	PixelHeight int    `json:"pixelHeight"`
	PixelWidth  int    `json:"pixelWidth"`
	Caption     string `json:"caption"`
	Primary     string `json:"primary"`
}

// type of Article.Videos[?]
type articleVideoType struct {
	Url         string `json:"url"`
	PixelHeight int    `json:"pixelHeight"`
	PixelWidth  int    `json:"pixelWidth"`
	Primary     string `json:"primary"`
}

// ParseArticle parse the clean article text from news article web pages.
//
// Request
//
// To use the Article API, perform a HTTP GET request on the following endpoint:
//
//	http://api.diffbot.com/v2/article
//
// Provide the following arguments:
//
//	+----------+-----------------------------------------------------------------+
//	| ARGUMENT | DESCRIPTION                                                     |
//	+----------+-----------------------------------------------------------------+
//	| token    | Developer token                                                 |
//	| url      | Article URL to process (URL encoded).                           |
//	|          | If you wish to POST content, please see POSTing Content, below. |
//	+----------+-----------------------------------------------------------------+
//	| Optional arguments                                                         |
//	+----------+-----------------------------------------------------------------+
//	| fields   | Used to control which fields are returned by the API.           |
//	|          | See the Response section below.                                 |
//	| timeout  | Set a value in milliseconds to terminate the response.          |
//	|          | By default the Article API has a five second timeout.           |
//	| callback | Use for jsonp requests. Needed for cross-domain ajax.           |
//	+----------+-----------------------------------------------------------------+
//
// Response
//
// The Article API returns information about the primary article content on the submitted page.
//
// Use the fields query parameter to limit or expand which fields are returned in the JSON response.
// For nested arrays, use parentheses to retrieve specific fields, or * to return all sub-fields.
//
//	http://api.diffbot.com/v2/article?...&fields=meta,querystring,images(*)
//
// Example Response
//
// This is a simple response:
//
//	{
//	  "type": "article",
//	  "icon": "http://www.diffbot.com/favicon.ico",
//	  "title": "Diffbot's New Product API Teaches Robots to Shop Online",
//	  "author": "John Davi",
//	  "date": "Wed, 31 Jul 2013 08:00:00 GMT",
//	  "videos": [
//	    {
//	      "primary": "true",
//	      "url": "http://www.youtube.com/embed/lfcri5ungRo?feature=oembed",
//	    }
//	  ],
//	  "tags": [
//	    "e-commerce",
//	    "SaaS"
//	  ]
//	  "url": "http://blog.diffbot.com/diffbots-new-product-api-teaches-robots-to-shop-online/",
//	  "humanLanguage": "en",
//	  "text": "Diffbot's human wranglers are proud today to announce the release of our newest product..."
//	}
//
// Authentication and Custom Headers
//
// You can supply Diffbot with custom headers, or basic authentication credentials,
// in order to access intranet pages or other sites that require a login.
//
//	Basic Authentication
//	To access pages that require a login/password (using basic access authentication),
//	include the username and password in your url parameter, e.g.: url=http%3A%2F%2FUSERNAME:PASSWORD@www.diffbot.com.
//
//	Custom Headers
//	You can supply the Article API with custom values for the user-agent, referer,
//	or cookie values in the HTTP request. These will be used in place of the Diffbot default values.
//
// To provide custom headers, pass in the following values in your own headers when calling the Diffbot API:
//
//	+----------------------+-----------------------------------------------------------------------+
//	| HEADER               | DESCRIPTION                                                           |
//	+----------------------+-----------------------------------------------------------------------+
//	| X-Forward-User-Agent | Will be used as Diffbot's User-Agent header when making your request. |
//	| X-Forward-Referer    | Will be used as Diffbot's Referer header when making your request.    |
//	| X-Forward-Cookie     | Will be used as Diffbot's Cookie header when making your request.     |
//	+----------------------+-----------------------------------------------------------------------+
//
// Posting Content
//
// If your content is not publicly available (e.g., behind a firewall),
// you can POST markup directly to the Article API endpoint for analysis:
//
//	http://api.diffbot.com/v2/article?token=...&url=...
//
// Please note that the url parameter is still required in the endpoint,
// and will be used to resolve any relative links contained in the markup.
//
// Provide markup to analyze as your POST body, and specify the Content-Type header as text/html.
//
// The following call submits a sample to the API:
//
//	curl
//	    -H "Content-Type:text/html"
//	    -d 'Now is the time for all good robots to come to the aid of their-- oh never mind, run!'
//	    http://api.diffbot.com/v2/article?token=...&url=http://www.diffbot.com/products/automatic/article
//
// See http://diffbot.com/dev/docs/article/.
//
func ParseArticle(token, url string, opt *Options) (*Article, error) {
	body, err := Diffbot("article", token, url, opt)
	if err != nil {
		return nil, err
	}
	var result Article
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (p *Article) String() string {
	d, _ := json.Marshal(p)
	return string(d)
}
