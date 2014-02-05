// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package diffbot

import (
	"encoding/json"
	"reflect"
	"testing"
)

func Product_type(t *testing.T) {
	var a Product
	a.Products = append(a.Products, productProductType{})
	a.Products[0].Medias = append(a.Products[0].Medias, productProductMediaType{})
	_ = a
}

func TestProduct_parseJson(t *testing.T) {
	var result1 Product
	if err := json.Unmarshal([]byte(testJsonDataProduct), &result1); err != nil {
		t.Fatal(err)
	}
	var result2 Product
	if err := json.Unmarshal([]byte(result1.String()), &result2); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(result1, result2) {
		t.Fatalf("not equal, expect = %q, got = %q", result1, result2)
	}

	if !reflect.DeepEqual(testGoldenProduct, result1) {
		t.Fatalf("not equal, expect = \n%q, got = \n%q", testGoldenProduct, result1)
	}
}

var testGoldenProduct = func() (result Product) {
	result.Url = "http://store.livrada.com/collections/all/products/before-i-go-to-sleep"
	result.Products = append(result.Products, productProductType{})
	result.Products[0].Medias = append(result.Products[0].Medias, productProductMediaType{})

	result.Products[0].Title = "Before I Go To Sleep"
	result.Products[0].Description = "Memories define us. So what if you lost yours every time you went to sleep? Your name, your identity, your past, even the people you love -- all forgotten overnight. And the one person you trust may be telling you only half the story. Before I Go To Sleep is a disturbing psychological thriller in which an amnesiac desperately tries to uncover the truth about who she is and who she can trust."

	result.Products[0].OfferPrice = "$7.99"
	result.Products[0].RegularPrice = "$9.99"
	result.Products[0].SaveAmount = "$2.00"

	result.Products[0].Medias[0].Height = 480
	result.Products[0].Medias[0].Width = 340
	result.Products[0].Medias[0].Link = "http://cdn.shopify.com/s/files/1/0184/6296/products/BeforeIGoToSleep_large.png?946"
	result.Products[0].Medias[0].Type = "image"
	result.Products[0].Medias[0].XPath = "/HTML[@class='no-js']/BODY[@id='page-product']/DIV[@class='content-frame']/DIV[@class='content']/DIV[@class='content-shop']/DIV[@class='row']/DIV[@class='span5']/DIV[@class='product-thumbs']/UL/LI[@class='first-image']/A[@class='single_image']/IMG"

	return
}()

const testJsonDataProduct = `
{
  "type": "product",
  "products": [
    {
      "title": "Before I Go To Sleep",
      "description": "Memories define us. So what if you lost yours every time you went to sleep? Your name, your identity, your past, even the people you love -- all forgotten overnight. And the one person you trust may be telling you only half the story. Before I Go To Sleep is a disturbing psychological thriller in which an amnesiac desperately tries to uncover the truth about who she is and who she can trust.",
      "offerPrice": "$7.99",
      "regularPrice": "$9.99",
      "saveAmount": "$2.00",
      "media": [
        {
          "height": 480,
          "width": 340,
          "link": "http://cdn.shopify.com/s/files/1/0184/6296/products/BeforeIGoToSleep_large.png?946",
          "type": "image",
          "xpath": "/HTML[@class='no-js']/BODY[@id='page-product']/DIV[@class='content-frame']/DIV[@class='content']/DIV[@class='content-shop']/DIV[@class='row']/DIV[@class='span5']/DIV[@class='product-thumbs']/UL/LI[@class='first-image']/A[@class='single_image']/IMG"
        }
      ]
    }
  ],
  "url": "http://store.livrada.com/collections/all/products/before-i-go-to-sleep"
}
`
