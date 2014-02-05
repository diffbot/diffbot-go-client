// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package diffbot

import (
	"encoding/json"
)

// Product represents a shopping or e-commerce product information.
//
// See http://diffbot.com/dev/docs/product/
type Product struct {
	Url         string                 `json:"url"`
	ResolvedUrl string                 `json:"resolved_url"`
	Meta        map[string]interface{} `json:"meta,omitempty"`        // Returned with fields.
	QueryString string                 `json:"querystring,omitempty"` // Returned with fields.
	Links       []string               `json:"links,omitempty"`       // Returned with fields.
	Breadcrumb  []string               `json:"breadcrumb"`
	Products    []struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Brand       string `json:"brand,omitempty"` // Returned with fields.
		Medias      []struct {
			Type    string `json:"type"`
			Link    string `json:"link"`
			Height  int    `json:"height"`
			Width   int    `json:"width"`
			Caption string `json:"caption"`
			Primary string `json:"primary"`
			XPath   string `json:"xpath"`
		} `json:"media"`
		OfferPrice     string `json:"offerPrice"`
		RegularPrice   string `json:"regularPrice"`
		SaveAmount     string `json:"saveAmount"`
		ShippingAmount string `json:"shippingAmount"`
		ProductId      string `json:"productId"`
		Upc            string `json:"upc"`
		PrefixCode     string `json:"prefixCode"`
		ProductOrigin  string `json:"productOrigin"`
		Isbn           string `json:"isbn"`
		Sku            string `json:"sku,omitempty"` // Returned with fields.
		Mpn            string `json:"mpn,omitempty"` // Returned with fields.
	} `json:"products"`
}

// type of Product.Products[?]
type productProductType struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Brand       string `json:"brand,omitempty"` // Returned with fields.
	Medias      []struct {
		Type    string `json:"type"`
		Link    string `json:"link"`
		Height  int    `json:"height"`
		Width   int    `json:"width"`
		Caption string `json:"caption"`
		Primary string `json:"primary"`
		XPath   string `json:"xpath"`
	} `json:"media"`
	OfferPrice     string `json:"offerPrice"`
	RegularPrice   string `json:"regularPrice"`
	SaveAmount     string `json:"saveAmount"`
	ShippingAmount string `json:"shippingAmount"`
	ProductId      string `json:"productId"`
	Upc            string `json:"upc"`
	PrefixCode     string `json:"prefixCode"`
	ProductOrigin  string `json:"productOrigin"`
	Isbn           string `json:"isbn"`
	Sku            string `json:"sku,omitempty"` // Returned with fields.
	Mpn            string `json:"mpn,omitempty"` // Returned with fields.
}

// type of Product.Products[?].Medias[?]
type productProductMediaType struct {
	Type    string `json:"type"`
	Link    string `json:"link"`
	Height  int    `json:"height"`
	Width   int    `json:"width"`
	Caption string `json:"caption"`
	Primary string `json:"primary"`
	XPath   string `json:"xpath"`
}

// ParseProduct parse a shopping or e-commerce product page and returns information on the product.
//
// Request
//
// To use the Product API, perform a HTTP GET request on the following endpoint:
//
//	http://api.diffbot.com/v2/product
//
// Provide the following arguments:
//
//	+----------+-------------------------------------------------------------------------+
//	| ARGUMENT | DESCRIPTION                                                             |
//	+----------+-------------------------------------------------------------------------+
//	| token    | Developer token                                                         |
//	| url      | Product URL to process (URL encoded)                                    |
//	+----------+-------------------------------------------------------------------------+
//	| Optional arguments                                                                 |
//	+----------+-------------------------------------------------------------------------+
//	| fields   | Used to control which fields are returned by the API.                   |
//	|          | See the Response section below.                                         |
//	| timeout  | Set a value in milliseconds to terminate the response.                  |
//	|          | By default the Product API has no timeout.                              |
//	| callback | Use for jsonp requests. Needed for cross-domain ajax.                   |
//	+----------+-------------------------------------------------------------------------+
//	| Basic authentication                                                               |
//	| To access pages that require a login/password (using basic access authentication), |
//	| include the username and password in your url parameter,                           |
//	| e.g.: url=http%3A%2F%2FUSERNAME:PASSWORD@www.diffbot.com                           |
//	+------------------------------------------------------------------------------------+
//
// Response
//
// The Product API returns product details in the products array.
// Currently extracted data will only be returned from a single product.
// In the future the API will return information from multiple products,
// if multiple items are available on the same page.
//
// Use the fields query parameter to limit or expand which fields
// are returned in the JSON response. For product-specific content your
// desired fields should be contained within the 'products' parentheses:
//
//	http://api.diffbot.com/v2/product...&fields=products(offerPrice,sku)
//
// Response fields:
//
//	+----------------+-------------------------------------------------------------------+
//	| FIELD          | DESCRIPTION                                                       |
//	+----------------+-------------------------------------------------------------------+
//	| *              | Returns all fields available.                                     |
//	| url            | URL submitted. Returned by default.                               |
//	| resolved_url   | Returned if the resolving URL is different from the               |
//	|                | submitted URL (e.g., link shortening services).                   |
//	|                | Returned by default.                                              |
//	| meta           | Returns the full contents of page meta tags,                      |
//	|                | including sub-arrays for OpenGraph tags,                          |
//	|                | Twitter Card metadata, schema.org microdata,                      |
//	|                | and -- if available -- oEmbed metadata.                           |
//	|                | Returned with fields.                                             |
//	| querystring    | Returns the key/value pairs of the URL querystring, if present.   |
//	|                | Items without a value will be returned as "true."                 |
//	|                | Returned with fields.                                             |
//	| links          | Returns all links (anchor tag href values) found on the page.     |
//	|                | Returned with fields.                                             |
//	| breadcrumb     | If available, an array of link URLs and link text                 |
//	|                | from page breadcrumbs. Returned by default.                       |
//	+----------------+-------------------------------------------------------------------+
//	| For each item in the products array:                                               |
//	+----------------+-------------------------------------------------------------------+
//	| title          | Name of the product. Returned by default.                         |
//	| description    | Description, if available, of the product.                        |
//	|                | Returned by default.                                              |
//	| brand          | Experimental Brand, if available, of the product.                 |
//	|                | Returned with fields.                                             |
//	| media          | Array of media items (images or videos) of the product.           |
//	|  |             | Returned by default.                                              |
//	|  +- type       | Type of media identified (image or video).                        |
//	|  +- link       | Direct (fully resolved) link to image or video content.           |
//	|  +- height     | Image height, in pixels.                                          |
//	|  +- width      | Image width, in pixels.                                           |
//	|  +- caption    | Diffbot-determined best caption for the image.                    |
//	|  +- primary    | Only images. Returns "True" if image is identified                |
//	|  |             | as primary in terms of size or positioning.                       |
//	|  +- xpath      | Full document Xpath to the media item.                            |
//	|                                                                                    |
//	| offerPrice     | Identified offer or actual/'final' price of the product.          |
//	|                | Returned by default.                                              |
//	| regularPrice   | Regular or original price of the product, if available.           |
//	|                | Returned by default.                                              |
//	| saveAmount     | Discount or amount saved, if available. Returned by default.      |
//	| shippingAmount | Shipping price, if available. Returned by default.                |
//	| productId      | A Diffbot-determined unique product ID.                           |
//	|                | If upc, isbn, mpn or sku are identified on the page,              |
//	|                | productId will select from these values in the above order.       |
//	|                | Otherwise Diffbot will attempt to derive the best unique          |
//	|                | value for the product. Returned by default.                       |
//	| upc            | Universal Product Code (UPC/EAN), if available.                   |
//	|                | Returned by default.                                              |
//	| prefixCode     | GTIN prefix code, typically the country of origin                 |
//	|                | as identified by UPC/ISBN. Returned by default.                   |
//	| productOrigin  | If available, the two-character ISO country code where            |
//	|                | the product was produced. Returned by default.                    |
//	| isbn           | International Standard Book Number (ISBN), if available.          |
//	|                | Returned by default.                                              |
//	| sku            | Stock Keeping Unit -- store/vendor inventory                      |
//	|                | number -- if available. Returned with fields.                     |
//	| mpn            | Manufacturer's Product Number, if available.                      |
//	|                | Returned with fields.                                             |
//	+----------------+-------------------------------------------------------------------+
//	| The following fields are in an early beta stage:                                   |
//	+----------------+-------------------------------------------------------------------+
//	| availability   | Item's availability, either true or false. Returned by default.   |
//	| brand          | The item brand, if identified. Returned with fields.              |
//	| quantityPrices | If a product page includes discounts for quantity purchases,      |
//	|                | quantityPrices will return an array of quantity and price values. |
//	|                | Returned with fields.                                             |
//	+----------------+-------------------------------------------------------------------|
//
// Example Response
//
// This is a simple response:
//
//	{
//	  "type": "product",
//	  "products": [
//	    {
//	      "title": "Before I Go To Sleep",
//	      "description": "Memories define us...",
//	      "offerPrice": "$7.99",
//	      "regularPrice": "$9.99",
//	      "saveAmount": "$2.00",
//	      "media": [
//	        {
//	          "height": 480,
//	          "width": 340,
//	          "link": "http://cdn.shopify.com/s/files/1/0184/6296/products/BeforeIGoToSleep_large.png?946",
//	          "type": "image",
//	          "xpath": "/HTML[@class='no-js']/BODY[@id='page-product']..."
//	        }
//	      ]
//	    }
//	  ],
//	  "url": "http://store.livrada.com/collections/all/products/before-i-go-to-sleep"
//	}
func ParseProduct(token, url string, opt *Options) (*Product, error) {
	body, err := Diffbot("product", token, url, opt)
	if err != nil {
		return nil, err
	}
	var result Product
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (p *Product) String() string {
	d, _ := json.Marshal(p)
	return string(d)
}
