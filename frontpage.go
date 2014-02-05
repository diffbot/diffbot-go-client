// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package diffbot

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// Frontpage represents a frontpage information.
//
// See http://diffbot.com/dev/docs/frontpage/
type Frontpage struct {
	Id        int64  `json:"id,string"`
	Title     string `json:"title"`
	SourceURL string `json:"sourceURL"`
	Icon      string `json:"icon"`
	NumItems  int    `json:"numItems"`
	Items     []struct {
		Id          int     `json:"id"`
		Title       string  `json:"title"`
		Description string  `json:"description"`
		XRoot       string  `json:"xroot"`
		PubDate     string  `json:"pubDate"`
		Link        string  `json:"link"`
		Type        string  `json:"type"` // STORY/LINK/...
		Img         string  `json:"img"`
		TextSummary string  `json:"textSummary"`
		Sp          float64 `json:"sp"`
		Sr          float64 `json:"sr"`
		Fresh       float64 `json:"fresh"`
	} `json:"items,omitempty"`
}

// type of Frontpage.Items[?]
type frontpageItemType struct {
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	XRoot       string  `json:"xroot"`
	PubDate     string  `json:"pubDate"`
	Link        string  `json:"link"`
	Type        string  `json:"type"` // STORY/LINK/...
	Img         string  `json:"img"`  // ?
	TextSummary string  `json:"textSummary"`
	Sp          float64 `json:"sp"`
	Sr          float64 `json:"sr"`
	Fresh       float64 `json:"fresh"`
}

// FrontpageDML (Diffbot Markup Language) is an XML format for encoding
// the extracted structural information from the page.
// A DML consists of a single info section and a list of items.
//
// See http://diffbot.com/products/automatic/frontpage/
type FrontpageDML struct {
	Id         int64  `json:"id,string"`
	TagName    string `json:"tagName"` // dml
	ChildNodes []struct {
		TagName          string `json:"tagName"`             // info/item/...
		ItemId           int64  `json:"id,string"`           // item.id, eg. "180194704"
		ItemSp           string `json:"sp"`                  // item.sp, eg. "0.000"
		ItemFresh        string `json:"fresh"`               // item.fresh, eg. "1.000"
		ItemSr           string `json:"sr"`                  // item.sr, eg. "4.000"
		ItemCluster      string `json:"cluster"`             // item.cluster, eg. "/HTML[1]/BODY[1]/DIV[4]/..."
		ItemCommentCount int64  `json:"commentCount,string"` // item.commentCount, eg. "34"
		ItemType         string `json:"type"`                // item.type, eg. "STORY"
		ItemXRoot        string `json:"xroot"`               // item.xroot, eg. "/HTML[1]/BODY[1]/DIV[4]/..."
		ChildNodes       []struct {
			TagName    string   `json:"tagName"` // title/sourceType/...
			ChildNodes []string `json:"childNodes"`
		} `json:"childNodes"`
	} `json:"childNodes"`
}

func (p *FrontpageDML) ParseJson(data []byte) error {
	if err := json.Unmarshal(data, p); err != nil {
		return err
	}
	return nil
}

func (p *FrontpageDML) String() string {
	d, _ := json.Marshal(p)
	return string(d)
}

// ParseFrontpage parse a multifaceted "homepage" and returns individual page elements.
//
// Request
//
// To use the Frontpage API, perform a HTTP GET request on the following endpoint:
//
//	http://api.diffbot.com/v2/frontpage
//
// Provide the following arguments:
//
//	+----------+----------------------------------------------------------------------------------------------------------+
//	| ARGUMENT | DESCRIPTION                                                                                              |
//	+----------+----------------------------------------------------------------------------------------------------------+
//	| token    | Developer token                                                                                          |
//	| url      | Frontpage URL from which to extract items                                                                |
//	+----------+----------------------------------------------------------------------------------------------------------+
//	| Optional arguments                                                                                                  |
//	+----------+----------------------------------------------------------------------------------------------------------+
//	| timeout  | Specify a value in milliseconds (e.g., &timeout=15000) to override the default API timeout of 5000ms.    |
//	| format   | Format the response output in xml (default) or json                                                      |
//	| all      | Returns all content from page, including navigation and similar links that the Diffbot visual processing |
//	|          | engine considers less important / non-core.                                                              |
//	+----------+----------------------------------------------------------------------------------------------------------+
//	| Basic authentication                                                                                                |
//	+---------------------------------------------------------------------------------------------------------------------+
//	| To access pages that require a login/password (using basic access authentication),                                  |
//	| include the username and password in your url parameter, e.g.: url=http%3A%2F%2FUSERNAME:PASSWORD@www.diffbot.com   |
//	+---------------------------------------------------------------------------------------------------------------------+
//
// Alternatively, you can POST the content to analyze directly to the same endpoint.
// Specify the Content-Type header as either text/plain or text/html.
//
// Response
//
// DML (Diffbot Markup Language) is an XML format for encoding
// the extracted structural information from the page. A DML consists of a single
// info section and a list of items.
//
//	+-------------+--------+------------------------------------------------------+
//	| INFO FIELD  | TYPE   | DESCRIPTION                                          |
//	+-------------+--------+------------------------------------------------------+
//	| id          | long   | DMLID of the URL                                     |
//	| title       | string | Extracted title of the page                          |
//	| sourceURL   | url    | the URL this was extracted from                      |
//	| icon        | url    | A link to a small icon/favicon representing the page |
//	| numItems    | int    | The number of items in this DML document             |
//	+-------------+--------+------------------------------------------------------+
//
// Some of the fields found in Items:
//
//	+-------------+--------------------------+--------------------------------------------------------------------+
//	| ITEM FIELD  | TYPE                     | DESCRIPTION                                                        |
//	+-------------+--------------------------+--------------------------------------------------------------------+
//	| id          | long                     | Unique hashcode/id of item                                         |
//	| title       | string                   | Title of item                                                      |
//	| description | string                   | innerHTML content of item                                          |
//	| xroot       | xpath                    | XPATH of where item was found on the page                          |
//	| pubDate     | timestamp                | Timestamp when item was detected on page                           |
//	| link        | URL                      | Extracted permalink (if applicable) of item                        |
//	| type        | {IMAGE,LINK,STORY,CHUNK} | Extracted type of the item, whether the item represents an image,  |
//	|             |                          | permalink, story (image+summary), or html chunk.                   |
//	| img         | URL                      | Extracted image from item                                          |
//	| textSummary | string                   | A plain-text summary of the item                                   |
//	| sp          | double<-[0,1]            | Spam score - the probability that the item is spam/ad              |
//	| sr          | double<-[1,5]            | Static rank - the quality score of the item on a 1 to 5 scale      |
//	| fresh       | double<-[0,1]            | Fresh score - the percentage of the item that has changed          |
//	|             |                          | compared to the previous crawl                                     |
//	+-------------+--------------------------+--------------------------------------------------------------------+
//
// See http://diffbot.com/dev/docs/frontpage/.
//
func ParseFrontpage(token, url string, opt *Options) (*Frontpage, error) {
	body, err := Diffbot("frontpage", token, url, opt)
	if err != nil {
		return nil, err
	}
	var dml FrontpageDML
	if err := json.Unmarshal(body, &dml); err != nil {
		return nil, err
	}
	var page Frontpage
	if err = page.ParseDML(&dml); err != nil {
		return nil, err
	}
	return &page, nil
}

func (p *Frontpage) ParseDML(dml *FrontpageDML) error {
	if dml.TagName != "dml" {
		return fmt.Errorf("diffbot: invalid FrontpageDML.")
	}
	*p = Frontpage{Id: dml.Id}
	for _, node := range dml.ChildNodes {
		switch node.TagName {
		case "info":
			for _, node := range node.ChildNodes {
				switch node.TagName {
				case "title":
					if len(node.ChildNodes) != 0 {
						p.Title = node.ChildNodes[0]
					}
				case "sourceURL":
					if len(node.ChildNodes) != 0 {
						p.SourceURL = node.ChildNodes[0]
					}
				case "icon":
					if len(node.ChildNodes) != 0 {
						p.Icon = node.ChildNodes[0]
					}
				case "numItems":
					if len(node.ChildNodes) != 0 {
						if v, err := strconv.Atoi(node.ChildNodes[0]); err == nil {
							p.NumItems = v
						}
					}
				default:
					// Unknown Fileds
				}
			}
		case "item":
			item := frontpageItemType{
				Id:    int(node.ItemId),
				Sp:    atof(node.ItemSp),
				Sr:    atof(node.ItemSr),
				Fresh: atof(node.ItemFresh),
				Type:  node.ItemType,
				XRoot: node.ItemXRoot,
			}
			for _, node := range node.ChildNodes {
				switch node.TagName {
				case "title":
					if len(node.ChildNodes) != 0 {
						item.Title = node.ChildNodes[0]
					}
				case "link":
					if len(node.ChildNodes) != 0 {
						item.Link = node.ChildNodes[0]
					}
				case "pubDate":
					if len(node.ChildNodes) != 0 {
						item.PubDate = node.ChildNodes[0]
					}
				case "textSummary":
					if len(node.ChildNodes) != 0 {
						item.TextSummary = node.ChildNodes[0]
					}
				case "description":
					if len(node.ChildNodes) != 0 {
						item.Description = node.ChildNodes[0]
					}
				default:
					// Unknown Fileds
				}
			}
			p.Items = append(p.Items, item)
		default:
			// Unknown Fileds
		}
	}
	return nil
}

func (p *Frontpage) String() string {
	d, _ := json.Marshal(p)
	return string(d)
}
