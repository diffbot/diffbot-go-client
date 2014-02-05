// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package diffbot

import (
	"encoding/json"
)

// Article represents an clean article text.
//
// See http://diffbot.com/dev/docs/analyze/
type Classification struct {
	Type          string `json:"type"`
	Title         string `json:"title"`
	Url           string `json:"url"`
	ResolvedUrl   string `json:"resolved_url,omitempty"` // Returned with fields.
	HumanLanguage string `json:"humanLanguage"`
	Stats         struct {
		Types struct {
			Article     float64 `json:"article"`
			Audio       float64 `json:"audio"`
			Chart       float64 `json:"chart"`
			Discussion  float64 `json:"discussion"`
			Document    float64 `json:"document"`
			Download    float64 `json:"download"`
			Error       float64 `json:"error"`
			Event       float64 `json:"event"`
			Faq         float64 `json:"faq"`
			Frontpage   float64 `json:"frontpage"`
			Game        float64 `json:"game"`
			Image       float64 `json:"image"`
			Job         float64 `json:"job"`
			Location    float64 `json:"location"`
			Other       float64 `json:"other"`
			Product     float64 `json:"product"`
			Profile     float64 `json:"profile"`
			Recipe      float64 `json:"recipe"`
			ReviewsList float64 `json:"reviewslist"`
			Serp        float64 `json:"serp"`
			Video       float64 `json:"video"`
		} `json:"types"`
	} `json:"stats"`
}

// type of Classification.Stats
type classificationStatsType struct {
	Types struct {
		Article     float64 `json:"article"`
		Audio       float64 `json:"audio"`
		Chart       float64 `json:"chart"`
		Discussion  float64 `json:"discussion"`
		Document    float64 `json:"document"`
		Download    float64 `json:"download"`
		Error       float64 `json:"error"`
		Event       float64 `json:"event"`
		Faq         float64 `json:"faq"`
		Frontpage   float64 `json:"frontpage"`
		Game        float64 `json:"game"`
		Image       float64 `json:"image"`
		Job         float64 `json:"job"`
		Location    float64 `json:"location"`
		Other       float64 `json:"other"`
		Product     float64 `json:"product"`
		Profile     float64 `json:"profile"`
		Recipe      float64 `json:"recipe"`
		ReviewsList float64 `json:"reviewslist"`
		Serp        float64 `json:"serp"`
		Video       float64 `json:"video"`
	} `json:"types"`
}

// type of Classification.Stats.Types
type classificationStatsTypeType struct {
	Article     float64 `json:"article"`
	Audio       float64 `json:"audio"`
	Chart       float64 `json:"chart"`
	Discussion  float64 `json:"discussion"`
	Document    float64 `json:"document"`
	Download    float64 `json:"download"`
	Error       float64 `json:"error"`
	Event       float64 `json:"event"`
	Faq         float64 `json:"faq"`
	Frontpage   float64 `json:"frontpage"`
	Game        float64 `json:"game"`
	Image       float64 `json:"image"`
	Job         float64 `json:"job"`
	Location    float64 `json:"location"`
	Other       float64 `json:"other"`
	Product     float64 `json:"product"`
	Profile     float64 `json:"profile"`
	Recipe      float64 `json:"recipe"`
	ReviewsList float64 `json:"reviewslist"`
	Serp        float64 `json:"serp"`
	Video       float64 `json:"video"`
}

// ParseClassification analyzes a web page's layout, structure, markup,
// text and other components and classifies the page as a particular "type."
// It also fully extracts the page contents if the page matches an existing
// Diffbot extraction API.
//
// Please note: The Page Classifier API is currently in beta.
//
// Request
//
// To use the Classifier API, perform a HTTP GET request on the following endpoint:
//
//	http://api.diffbot.com/v2/analyze?token=...&url=...
//
// Provide the following arguments:
//
//	+----------+----------------------------------------------------------------------------------------------+
//	| ARGUMENT | DESCRIPTION                                                                                  |
//	+----------+----------------------------------------------------------------------------------------------+
//	| token    | Developer token                                                                              |
//	| url      | URL to classify (URLEncoded)                                                                 |
//	+----------+----------------------------------------------------------------------------------------------+
//	| Optional arguments                                                                                      |
//	+----------+----------------------------------------------------------------------------------------------+
//	| mode     | By default the Page Classifier API will fully extract                                        |
//	|          | pages that match an existing Diffbot Automatic API.                                          |
//	|          | Set mode to a specific page-type (e.g., mode=article)                                        |
//	|          | to extract content only from that particular page-type.                                      |
//	|          | All others will simply return the page classification information.                           |
//	| fields   | You can choose the fields to be returned                                                     |
//	|          | by the Diffbot extraction API by supplying a comma-separated                                 |
//	|          | list of fields, e.g.:                                                                        |
//	|          | http://api.diffbot.com/v2/analyze?token=...&url=http://diffbot.com/company&fields=meta,tags. |
//	| stats    | Returns statistics on page classification and extraction,                                    |
//	|          | including an array of individual page-types and                                              |
//	|          | the Diffbot-determined score (likelihood) for each type.                                     |
//	+----------+----------------------------------------------------------------------------------------------+
//	| Basic authentication                                                                                    |
//	+---------------------------------------------------------------------------------------------------------+
//	| To access pages that require a login/password                                                           |
//	| (using basic access authentication), include the username and password                                  |
//	| in your url parameter,                                                                                  |
//	| e.g.: url=http%3A%2F%2FUSERNAME:PASSWORD@www.diffbot.com                                                |
//	+---------------------------------------------------------------------------------------------------------+
//
// Response
//
// The Classifier API returns, depending on parameters, the following:
//
//	+----------------+------------------------------------------------------------------+
//	| FIELD          | DESCRIPTION                                                      |
//	+----------------+------------------------------------------------------------------+
//	| type           | Page-type of the submitted URL (from the below enumerated list). |
//	|                | Always returned.                                                 |
//	| title          | Page title. Returned by default.                                 |
//	| url            | Submitted URL. Always returned.                                  |
//	| resolved_url   | Returned if the resolving URL is different from                  |
//	|                | the submitted URL (e.g., link shortening services).              |
//	|                | Returned by default, configurable with fields                    |
//	| human_language | Returns the (spoken/human) language of the submitted URL,        |
//	|                | using two-letter ISO 639-1 nomenclature.                         |
//	|                | Returned by default.                                             |
//	+----------------+------------------------------------------------------------------+
//
// Example Response
//
// This is a simple response:
//
//	{
//	  "type": "article"
//	  "stats": {
//	     "types": {
//	        "article": 0.46,
//	        "audio": 0.15,
//	        "chart": 0.01,
//	        "discussion": 0.03,
//	        "document": 0.04,
//	        "download": 0.01,
//	        "error": 0.00,
//	        "event": 0.00,
//	        "faq": 0.02,
//	        "frontpage": 0.12,
//	        "game": 0.01,
//	        "image": 0.02,
//	        "job": 0.02,
//	        "location": 0.08,
//	        "product": 0.09,
//	        "profile": 0.09,
//	        "recipe": 0.08,
//	        "reviewslist": 0.09,
//	        "serp": 0.06,
//	        "video": 0.01
//	      }
//	    },
//	  "resolved_url": "http://techcrunch.com/2012/05/31/diffbot-raises-2-million-seed-round-for-web-content-extraction-technology/",
//	  "url": "http://tcrn.ch/Jw7ZKw",
//	  "human_language": "en"
//	}
//
// Page Types
//
// Diffbot currently classifies pages into the following types.
// Please note this list will evolve over time to include additional page types.
//
//	+-------------+-----------------------------------------------------------------------------------+
//	| PAGE TYPE   | DESCRIPTION                                                                       |
//	+-------------+-----------------------------------------------------------------------------------+
//	| None        | Returned if Diffbot confidence in the page classification is low.                 |
//	|             | Use of the stats field will give you the individual scores for each page-type.    |
//	| article     | A news article, blog post or other primarily-text page.                           |
//	| audio       | A music or audio player.                                                          |
//	| chart       | A graph or chart, typically financial.                                            |
//	| discussion  | Specific forum, group or discussion topic.                                        |
//	| document    | An embedded or downloadable document or slideshow.                                |
//	| download    | A downloadable file.                                                              |
//	| error       | Error page, e.g. 404.                                                             |
//	| event       | A page detailing specific event information,                                      |
//	|             | e.g. time/date/location.                                                          |
//	| faq         | A page of multiple frequently asked questions, or a single FAQ entry.             |
//	| frontpage   | A news- or blog-style home page, with links to myriad sections and items.         |
//	| game        | A playable game.                                                                  |
//	| image       | An image or photo page.                                                           |
//	| job         | A job posting.                                                                    |
//	| location    | A page detailing location information, typically including an address and/or map. |
//	| other       | Returned if the result is below a certain confidence threshold.                   |
//	| product     | A product page, typically of a product for purchase.                              |
//	| profile     | A person or user profile page.                                                    |
//	| recipe      | Page detailing recipe instructions and ingredients.                               |
//	| reviewslist | A list of user reviews.                                                           |
//	| serp        | A Search Engine Results Page                                                      |
//	| video       | An individual video.                                                              |
//	+-------------+-----------------------------------------------------------------------------------+
//
func ParseClassification(token, url string, opt *Options) (*Classification, error) {
	body, err := Diffbot("analyze", token, url, opt)
	if err != nil {
		return nil, err
	}
	var result Classification
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (p *Classification) String() string {
	d, _ := json.Marshal(p)
	return string(d)
}
