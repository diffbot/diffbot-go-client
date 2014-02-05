// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package diffbot

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestImage_type(t *testing.T) {
	var a Image
	a.Images = append(a.Images, imageImageType{})
	_ = a
}

func TestImage_parseJson(t *testing.T) {
	var result1 Image
	if err := json.Unmarshal([]byte(testJsonDataImage), &result1); err != nil {
		t.Fatal(err)
	}
	var result2 Image
	if err := json.Unmarshal([]byte(result1.String()), &result2); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(result1, result2) {
		t.Fatalf("not equal, expect = %q, got = %q", result1, result2)
	}

	images := result1.Images
	result1.Images = nil

	if !reflect.DeepEqual(testGoldenImage, result1) {
		t.Fatalf("not equal, expect = \n%q, got = \n%q", testGoldenImage, result1)
	}
	for i := 0; i < len(images) && i < len(testGoldenImageImages); i++ {
		if !reflect.DeepEqual(testGoldenImageImages[i], imageImageType(images[i])) {
			t.Fatalf("%d: not equal, expect = \n%q, got = \n%q", i, testGoldenImageImages[i], images[i])
		}
	}
}

var testGoldenImage = Image{
	Title:       "The National Flower - Rose",
	NextPage:    "",
	AlbumUrl:    "",
	Url:         "http://www.statesymbolsusa.org/National_Symbols/National_flower.html",
	ResolvedUrl: "",
	Meta:        nil,
	QueryString: "",
	Links:       nil,
}
var testGoldenImageImages = []imageImageType{
	{
		Url:           "http://www.statesymbolsusa.org/IMAGES/rose_usda-web.jpg",
		AnchorUrl:     "",
		Mime:          "",
		Caption:       "",
		AttrAlt:       "Red rose in full bloom - click to see state flowers",
		AttrTitle:     "",
		Date:          "",
		Size:          12328,
		PixelHeight:   371, // doc error
		PixelWidth:    300,
		DisplayHeight: 371,
		DisplayWidth:  300,
		Meta: []string{
			"[Jpeg] Compression Type - Baseline",
			"[Jpeg] Data Precision - 8 bits",
			"[Jpeg] Image Height - 371 pixels",
			"[Jpeg] Image Width - 300 pixels",
			"[Jpeg] Number of Components - 3",
			"[Jpeg] Component 1 - Y component: Quantization table 0, Sampling factors 2 horiz/2 vert",
			"[Jpeg] Component 2 - Cb component: Quantization table 1, Sampling factors 1 horiz/1 vert",
			"[Jpeg] Component 3 - Cr component: Quantization table 1, Sampling factors 1 horiz/1 vert",
			"[Jfif] Version - 1.2",
			"[Jfif] Resolution Units - none",
			"[Jfif] X Resolution - 100 dots",
			"[Jfif] Y Resolution - 100 dots",
			"[Adobe Jpeg] DCT Encode Version - 1",
			"[Adobe Jpeg] Flags 0 - 192",
			"[Adobe Jpeg] Flags 1 - 0",
			"[Adobe Jpeg] Color Transform - YCbCr",
		},
		Faces:  nil,
		Ocr:    "",
		Colors: "",
		XPath:  "/HTML[1]/BODY[1]/DIV[1]/TABLE[3]/TBODY[1]/TR[2]/TD[4]/DIV[1]/DIV[1]/H6[1]/SPAN[1]/A[1]/IMG[1]",
	},
	{
		Url:           "http://www.statesymbolsusa.org/IMAGES/rose_yellow-380.jpg",
		AnchorUrl:     "",
		Mime:          "",
		Caption:       "",
		AttrAlt:       "Yellow rose - click to see state flowers",
		AttrTitle:     "",
		Date:          "",
		Size:          12142,
		PixelHeight:   304, // doc error
		PixelWidth:    380,
		DisplayHeight: 304,
		DisplayWidth:  380,
		Meta: []string{
			"[Jpeg] Compression Type - Baseline",
			"[Jpeg] Data Precision - 8 bits",
			"[Jpeg] Image Height - 304 pixels",
			"[Jpeg] Image Width - 380 pixels",
			"[Jpeg] Number of Components - 3",
			"[Jpeg] Component 1 - Y component: Quantization table 0, Sampling factors 2 horiz/2 vert",
			"[Jpeg] Component 2 - Cb component: Quantization table 1, Sampling factors 1 horiz/1 vert",
			"[Jpeg] Component 3 - Cr component: Quantization table 1, Sampling factors 1 horiz/1 vert",
			"[Jfif] Version - 1.2",
			"[Jfif] Resolution Units - none",
			"[Jfif] X Resolution - 100 dots",
			"[Jfif] Y Resolution - 100 dots",
			"[Adobe Jpeg] DCT Encode Version - 1",
			"[Adobe Jpeg] Flags 0 - 192",
			"[Adobe Jpeg] Flags 1 - 0",
			"[Adobe Jpeg] Color Transform - YCbCr",
		},
		Faces:  nil,
		Ocr:    "",
		Colors: "",
		XPath:  "/HTML[1]/BODY[1]/DIV[1]/TABLE[3]/TBODY[1]/TR[2]/TD[4]/DIV[1]/DIV[1]/TABLE[1]/TBODY[1]/TR[1]/TD[1]/DIV[1]/H6[1]/SPAN[1]/A[1]/IMG[1]",
	},
}

const testJsonDataImage = `
{
  "title": "The National Flower - Rose",
  "type": "image",
  "url": "http://www.statesymbolsusa.org/National_Symbols/National_flower.html",
  "images": [
    {
      "attrAlt": "Red rose in full bloom - click to see state flowers",
      "pixelHeight": 371,
      "pixelWidth": 300,
      "displayWidth": 300,
      "meta": [
          "[Jpeg] Compression Type - Baseline",
          "[Jpeg] Data Precision - 8 bits",
          "[Jpeg] Image Height - 371 pixels",
          "[Jpeg] Image Width - 300 pixels",
          "[Jpeg] Number of Components - 3",
          "[Jpeg] Component 1 - Y component: Quantization table 0, Sampling factors 2 horiz/2 vert",
          "[Jpeg] Component 2 - Cb component: Quantization table 1, Sampling factors 1 horiz/1 vert",
          "[Jpeg] Component 3 - Cr component: Quantization table 1, Sampling factors 1 horiz/1 vert",
          "[Jfif] Version - 1.2",
          "[Jfif] Resolution Units - none",
          "[Jfif] X Resolution - 100 dots",
          "[Jfif] Y Resolution - 100 dots",
          "[Adobe Jpeg] DCT Encode Version - 1",
          "[Adobe Jpeg] Flags 0 - 192",
          "[Adobe Jpeg] Flags 1 - 0",
          "[Adobe Jpeg] Color Transform - YCbCr"
          ],
      "url": "http://www.statesymbolsusa.org/IMAGES/rose_usda-web.jpg",
      "size": 12328,
      "displayHeight": 371,
      "xpath": "/HTML[1]/BODY[1]/DIV[1]/TABLE[3]/TBODY[1]/TR[2]/TD[4]/DIV[1]/DIV[1]/H6[1]/SPAN[1]/A[1]/IMG[1]"
    },
    {
      "attrAlt": "Yellow rose - click to see state flowers",
      "pixelHeight": 304,
      "pixelWidth": 380,
      "displayWidth": 380,
      "meta": [
          "[Jpeg] Compression Type - Baseline",
          "[Jpeg] Data Precision - 8 bits",
          "[Jpeg] Image Height - 304 pixels",
          "[Jpeg] Image Width - 380 pixels",
          "[Jpeg] Number of Components - 3",
          "[Jpeg] Component 1 - Y component: Quantization table 0, Sampling factors 2 horiz/2 vert",
          "[Jpeg] Component 2 - Cb component: Quantization table 1, Sampling factors 1 horiz/1 vert",
          "[Jpeg] Component 3 - Cr component: Quantization table 1, Sampling factors 1 horiz/1 vert",
          "[Jfif] Version - 1.2",
          "[Jfif] Resolution Units - none",
          "[Jfif] X Resolution - 100 dots",
          "[Jfif] Y Resolution - 100 dots",
          "[Adobe Jpeg] DCT Encode Version - 1",
          "[Adobe Jpeg] Flags 0 - 192",
          "[Adobe Jpeg] Flags 1 - 0",
          "[Adobe Jpeg] Color Transform - YCbCr"
          ],
      "url": "http://www.statesymbolsusa.org/IMAGES/rose_yellow-380.jpg",
      "size": 12142,
      "displayHeight": 304,
      "xpath": "/HTML[1]/BODY[1]/DIV[1]/TABLE[3]/TBODY[1]/TR[2]/TD[4]/DIV[1]/DIV[1]/TABLE[1]/TBODY[1]/TR[1]/TD[1]/DIV[1]/H6[1]/SPAN[1]/A[1]/IMG[1]"
    }
  ]
}
`
