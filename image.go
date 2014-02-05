// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package diffbot

import (
	"encoding/json"
)

// Image represents a page image information.
//
// See http://diffbot.com/dev/docs/image/
type Image struct {
	Title       string                 `json:"title"`
	NextPage    string                 `json:"nextPage"`
	AlbumUrl    string                 `json:"albumUrl"`
	Url         string                 `json:"url"`
	ResolvedUrl string                 `json:"resolved_url"`
	Meta        map[string]interface{} `json:"meta,omitempty"`        // Returned with fields.
	QueryString string                 `json:"querystring,omitempty"` // Returned with fields.
	Links       []string               `json:"links,omitempty"`       // Returned with fields.
	Images      []struct {
		Url           string   `json:"url"`
		AnchorUrl     string   `json:"anchorUrl"`
		Mime          string   `json:"mime,omitempty"` // Returned with fields.
		Caption       string   `json:"caption"`
		AttrAlt       string   `json:"attrAlt,omitempty"`   // Returned with fields.
		AttrTitle     string   `json:"attrTitle,omitempty"` // Returned with fields.
		Date          string   `json:"date"`
		Size          int      `json:"size"`
		PixelHeight   int      `json:"pixelHeight"`
		PixelWidth    int      `json:"pixelWidth"`
		DisplayHeight int      `json:"displayHeight,omitempty"` // Returned with fields.
		DisplayWidth  int      `json:"displayWidth",omitempty`  // Returned with fields.
		Meta          []string `json:"meta"`
		Faces         []string `json:"faces,omitempty"`  // Returned with fields.
		Ocr           string   `json:"ocr,omitempty"`    // Returned with fields.
		Colors        string   `json:"colors,omitempty"` // Returned with fields.
		XPath         string   `json:"xpath"`
	} `json:"images"`
}

// type of Image.Images[?]
type imageImageType struct {
	Url           string   `json:"url"`
	AnchorUrl     string   `json:"anchorUrl"`
	Mime          string   `json:"mime,omitempty"` // Returned with fields.
	Caption       string   `json:"caption"`
	AttrAlt       string   `json:"attrAlt,omitempty"`   // Returned with fields.
	AttrTitle     string   `json:"attrTitle,omitempty"` // Returned with fields.
	Date          string   `json:"date"`
	Size          int      `json:"size"`
	PixelHeight   int      `json:"pixelHeight"`
	PixelWidth    int      `json:"pixelWidth"`
	DisplayHeight int      `json:"displayHeight,omitempty"` // Returned with fields.
	DisplayWidth  int      `json:"displayWidth",omitempty`  // Returned with fields.
	Meta          []string `json:"meta"`
	Faces         []string `json:"faces,omitempty"`  // Returned with fields.
	Ocr           string   `json:"ocr,omitempty"`    // Returned with fields.
	Colors        string   `json:"colors,omitempty"` // Returned with fields.
	XPath         string   `json:"xpath"`
}

// ParseImage parse a web page and returns its primary image(s).
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
//	+------------------------------------------------------------------------------------+
//	| To access pages that require a login/password (using basic access authentication), |
//	| include the username and password in your url parameter,                           |
//	| e.g.: url=http%3A%2F%2FUSERNAME:PASSWORD@www.diffbot.com                           |
//	+------------------------------------------------------------------------------------+
//
// Response
//
// The Image API returns basic info about the page submitted,
// and its primary image(s) in the images array.
//
//Use the fields query parameter to limit or expand which fields are
// returned in the JSON response. To control the fields returned for images,
// your desired fields should be contained within the 'images' parentheses:
//
//	http://api.diffbot.com/v2/image...&fields=images(mime,pixelWidth)
//
// Response fields:
//
//	+---------------+------------------------------------------------------------------------+
//	| FIELD         | DESCRIPTION                                                            |
//	+---------------+------------------------------------------------------------------------+
//	| *             | Returns all fields available.                                          |
//	| title         | Title of the submitted page. Returned by default.                      |
//	| nextPage      | Link to next page (if within a gallery or paginated list of images).   |
//	|               | Returned by default.                                                   |
//	| albumUrl      | Link to containing album (if image is within an album).                |
//	|               | Returned by default.                                                   |
//	| url           | URL submitted. Returned by default.                                    |
//	| resolved_url  | Returned if the resolving URL is different from the submitted URL      |
//	|               | (e.g., link shortening services).                                      |
//	| meta          | Returns the full contents of page meta tags,                           |
//	|               | including sub-arrays for OpenGraph tags, Twitter Card metadata,        |
//	|               | schema.org microdata, and -- if available -- oEmbed metadata.          |
//	|               | Returned with fields.                                                  |
//	| querystring   | Returns the key/value pairs of the URL querystring, if present.        |
//	|               | Items without a value will be returned as "true".                      |
//	|               | Returned with fields.                                                  |
//	| links         | Returns all links (anchor tag href values) found on the page.          |
//	|               | Returned with fields.                                                  |
//	| images        | An array of image(s) contained on the page.                            |
//	+---------------+------------------------------------------------------------------------+
//	| For each item in the images array:                                                     |
//	+---------------+------------------------------------------------------------------------+
//	| url           | Direct link to image file. Returned by default.                        |
//	| anchorUrl     | If the image is wrapped by an anchor a tag, the anchor location        |
//	|               | as defined by the href attribute. Returned by default.                 |
//	| mime          | MIME type, if available, as specified by "Content-Type" of the image.  |
//	|               | Returned with fields.                                                  |
//	| caption       | The best caption for this image. Returned by default.                  |
//	| attrAlt       | Contents of the alt attribute, if available within the HTML IMG tag.   |
//	|               | Returned with fields.                                                  |
//	| attrTitle     | Contents of the title attribute, if available within the HTML IMG tag. |
//	|               | Returned with fields.                                                  |
//	| date          | Date of image upload or creation if available in page metadata.        |
//	|               | Returned by default.                                                   |
//	| size          | Size in bytes of image file. Returned by default.                      |
//	| pixelHeight   | Actual height, in pixels, of image file. Returned by default.          |
//	| pixelWidth    | Actual width, in pixels, of image file. Returned by default.           |
//	| displayHeight | Height of image as rendered on page, if different from actual          |
//	|               | (pixel) height. Returned with fields.                                  |
//	| displayWidth  | Width of image as rendered on page, if different from actual           |
//	|               | (pixel) width. Returned with fields.                                   |
//	| meta          | Comma-separated list of image-embedded metadata                        |
//	|               | (e.g., EXIF, XMP, ICC Profile), if available within the image file.    |
//	|               | Returned with fields.                                                  |
//	| faces         | The x, y, height, width of coordinates of human faces.                 |
//	|               | Null, if no faces were found. Returned with fields.                    |
//	| ocr           | If text is identified within the image, we will attempt to recognize   |
//	|               | the text string. Returned with fields.                                 |
//	| colors        | Returns an array of hex values of the dominant colors                  |
//	|               | within the image. Returned with fields.                                |
//	| xpath         | XPath expression identifying the node containing the image.            |
//	|               | Returned by default.                                                   |
//	+---------------+------------------------------------------------------------------------+
//
// Example Response
//
// This is a simple response:
//
//	{
//	  "title": "The National Flower - Rose",
//	  "type": "image",
//	  "url": "http://www.statesymbolsusa.org/National_Symbols/National_flower.html",
//	  "images": [
//	    {
//	      "attrAlt": "Red rose in full bloom - click to see state flowers",
//	      "height": 371,
//	      "width": 300,
//	      "displayWidth": 300,
//	      "meta": [
//	          "[Jpeg] Compression Type - Baseline",
//	          "[Jpeg] Data Precision - 8 bits",
//	          "[Jpeg] Image Height - 371 pixels",
//	          "[Jpeg] Image Width - 300 pixels",
//	          "[Jpeg] Number of Components - 3",
//	          "[Jpeg] Component 1 - Y component: Quantization table 0, Sampling factors 2 horiz/2 vert",
//	          "[Jpeg] Component 2 - Cb component: Quantization table 1, Sampling factors 1 horiz/1 vert",
//	          "[Jpeg] Component 3 - Cr component: Quantization table 1, Sampling factors 1 horiz/1 vert",
//	          "[Jfif] Version - 1.2",
//	          "[Jfif] Resolution Units - none",
//	          "[Jfif] X Resolution - 100 dots",
//	          "[Jfif] Y Resolution - 100 dots",
//	          "[Adobe Jpeg] DCT Encode Version - 1",
//	          "[Adobe Jpeg] Flags 0 - 192",
//	          "[Adobe Jpeg] Flags 1 - 0",
//	          "[Adobe Jpeg] Color Transform - YCbCr"
//	          ],
//	      "url": "http://www.statesymbolsusa.org/IMAGES/rose_usda-web.jpg",
//	      "size": 12328,
//	      "displayHeight": 371,
//	      "xpath": "/HTML[1]/BODY[1]/DIV[1]/TABLE[3]/TBODY[1]/TR[2]/..."
//	    },
//	    {
//	      "attrAlt": "Yellow rose - click to see state flowers",
//	      "pixelHeight": 304,
//	      "pixelWidth": 380,
//	      "displayWidth": 380,
//	      "meta": [
//	          "[Jpeg] Compression Type - Baseline",
//	          "[Jpeg] Data Precision - 8 bits",
//	          "[Jpeg] Image Height - 304 pixels",
//	          "[Jpeg] Image Width - 380 pixels",
//	          "[Jpeg] Number of Components - 3",
//	          "[Jpeg] Component 1 - Y component: Quantization table 0, Sampling factors 2 horiz/2 vert",
//	          "[Jpeg] Component 2 - Cb component: Quantization table 1, Sampling factors 1 horiz/1 vert",
//	          "[Jpeg] Component 3 - Cr component: Quantization table 1, Sampling factors 1 horiz/1 vert",
//	          "[Jfif] Version - 1.2",
//	          "[Jfif] Resolution Units - none",
//	          "[Jfif] X Resolution - 100 dots",
//	          "[Jfif] Y Resolution - 100 dots",
//	          "[Adobe Jpeg] DCT Encode Version - 1",
//	          "[Adobe Jpeg] Flags 0 - 192",
//	          "[Adobe Jpeg] Flags 1 - 0",
//	          "[Adobe Jpeg] Color Transform - YCbCr"
//	          ],
//	      "url": "http://www.statesymbolsusa.org/IMAGES/rose_yellow-380.jpg",
//	      "size": 12142,
//	      "displayHeight": 304,
//	      "xpath": "/HTML[1]/BODY[1]/DIV[1]/TABLE[3]/TBODY[1]/TR[2]/..."
//	    }
//	  ]
//	}
func ParseImage(token, url string, opt *Options) (*Image, error) {
	body, err := Diffbot("image", token, url, opt)
	if err != nil {
		return nil, err
	}
	var result Image
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (p *Image) String() string {
	d, _ := json.Marshal(p)
	return string(d)
}
