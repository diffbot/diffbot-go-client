// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package diffbot

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestArticle_type(t *testing.T) {
	var a Article
	a.Images = append(a.Images, articleImageType{})
	a.Videos = append(a.Videos, articleVideoType{})
	_ = a
}

func TestArticle_parseJson(t *testing.T) {
	var result1 Article
	if err := json.Unmarshal([]byte(testJsonDataArticle), &result1); err != nil {
		t.Fatal(err)
	}
	var result2 Article
	if err := json.Unmarshal([]byte(result1.String()), &result2); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(result1, result2) {
		t.Fatalf("not equal, expect = %q, got = %q", result1, result2)
	}

	result1.Text = ""
	result1.Html = ""
	if !reflect.DeepEqual(testArticleValue, result1) {
		t.Fatalf("not equal, expect = \n%q, got = \n%q", testArticleValue, result1)
	}
}

var testArticleValue = Article{
	Url:         "http://blog.diffbot.com/diffbots-new-product-api-teaches-robots-to-shop-online/",
	ResolvedUrl: "",
	Icon:        "",
	Meta: map[string]interface{}{
		"viewport":  "width=device-width",
		"microdata": map[string]interface{}{"author": "John Davi"},
		"title":     "Diffbot\u2019s New Product API Teaches Robots to Shop Online | Diffblog",
		"generator": "WordPress 3.5",
	},
	Links: []string{
		"http://diffbot.com",
		"http://blog.diffbot.com/author/johndavi/",
		"http://www.diffbot.com/products/crawlbot",
		"http://blog.diffbot.com/category/api-features/",
		"http://www.diffbot.com/products/automatic/product",
		"http://twitter.com/diffbot",
		"http://www.diffbot.com/pricing",
		"http://blog.diffbot.com",
		"http://blog.diffbot.com/feed/",
		"http://wordpress.org",
		"http://blog.diffbot.com/crawlbot-updates-webhooks-and-preventing-duplicate-content/",
		"http://blog.diffbot.com/diffbots-new-product-api-teaches-robots-to-shop-online/",
		"http://blog.diffbot.com/announcing-crawlbot-smart-site-spidering-and-extraction/",
		"http://blog.diffbot.com/diffbot-apis-now-with-more-meta/",
		"http://blog.diffbot.com/crawlbot-enhancements-api-parameters-product-crawl-csvs/",
	},
	Type:          "article",
	Title:         "Diffbot\u2019s New Product API Teaches Robots to Shop Online",
	Text:          "", // too large, ingore
	Html:          "", // too large, ingore
	NumPages:      "",
	Date:          "Wed, 31 Jul 2013 07:00:00 GMT",
	Author:        "John Davi",
	Tags:          []string{"Data model", "Product (chemistry)", "Intelligence", "Technology"},
	HumanLanguage: "en",
	Images:        nil,
	Videos: []struct {
		Url         string `json:"url"`
		PixelHeight int    `json:"pixelHeight"`
		PixelWidth  int    `json:"pixelWidth"`
		Primary     string `json:"primary"`
	}{
		{
			Url:     "http://www.youtube.com/embed/lfcri5ungRo?feature=oembed",
			Primary: "true",
		},
	},
}

const testJsonDataArticle = `
{
	"tags":["Data model","Product (chemistry)","Intelligence","Technology"],
	"summary":"Diffbot\u2019s human wranglers are proud today to announce the release of our newest product: an API for\u2026 products! The Product API can be used for extracting clean, structured data from any e-commerce product page.  It automatically makes available all the product data you\u2019d expect: price, discount/savings amount, shipping cost, product description, any relevant product images, SKU and/or other product IDs.",
	"text":"Diffbot\u2019s human wranglers are proud today to announce the release of our newest product: an API for\u2026 products!\nThe Product API can be used for extracting clean, structured data from any e-commerce product page. It automatically makes available all the product data you\u2019d expect: price, discount/savings amount, shipping cost, product description, any relevant product images, SKU and/or other product IDs.\nEven cooler: pair the Product API with Crawlbot , our intelligent site-spidering tool, and let Diffbot determine which pages are products, then automatically structure the entire catalog. Here\u2019s a quick demonstration of Crawlbot at work:\nWe\u2019ve developed the Product API over the course of two years, building upon our core vision technology that\u2019s extracted structured data from billions of web pages, and training our machine learning systems using data from tens of thousands of unique shopping sites. We can\u2019t wait for you to try it out.\nWhat are you waiting for? Check out the Product API documentation and dive on in! If you need a token, check out our pricing and plans (including our Free plan).\nQuestions? Hit us up at support@diffbot.com .",
	"stats":{"confidence":"0.800"},
	"videos":[
		{
			"primary":"true",
			"url":"http://www.youtube.com/embed/lfcri5ungRo?feature=oembed"
		}
	],
	"humanLanguage":"en",
	"links":[
		"http://diffbot.com","http://blog.diffbot.com/author/johndavi/",
		"http://www.diffbot.com/products/crawlbot","http://blog.diffbot.com/category/api-features/",
		"http://www.diffbot.com/products/automatic/product","http://twitter.com/diffbot",
		"http://www.diffbot.com/pricing","http://blog.diffbot.com","http://blog.diffbot.com/feed/",
		"http://wordpress.org",
		"http://blog.diffbot.com/crawlbot-updates-webhooks-and-preventing-duplicate-content/",
		"http://blog.diffbot.com/diffbots-new-product-api-teaches-robots-to-shop-online/",
		"http://blog.diffbot.com/announcing-crawlbot-smart-site-spidering-and-extraction/",
		"http://blog.diffbot.com/diffbot-apis-now-with-more-meta/",
		"http://blog.diffbot.com/crawlbot-enhancements-api-parameters-product-crawl-csvs/"
	],
	"type":"article",
	"date":"Wed, 31 Jul 2013 07:00:00 GMT",
	"url":"http://blog.diffbot.com/diffbots-new-product-api-teaches-robots-to-shop-online/",
	"meta":{
		"viewport":"width=device-width",
		"microdata":{"author":"John Davi"},
		"title":"Diffbot\u2019s New Product API Teaches Robots to Shop Online | Diffblog",
		"generator":"WordPress 3.5"
	},
	"author":"John Davi",
	"title":"Diffbot\u2019s New Product API Teaches Robots to Shop Online",
	"html":"<div><div class=\"video_frame\"><iframe allowfullscreen=\"\" filter_annotation=\"DO_NOT_USE\" frameborder=\"0\" height=\"352\" src=\"http://www.youtube.com/embed/lfcri5ungRo?feature=oembed\" width=\"625\"><\/iframe><\/div><div><div>\n\t\t\t<p>Diffbot&rsquo;s human wranglers are proud today to announce the release of our newest product: an API for&hellip; products!<\/p>\n<p>The&nbsp;<a href=\"http://www.diffbot.com/products/automatic/product\" title=\"Diffbot's Product API\">Product API<\/a>&nbsp;can be used for extracting clean, structured data from any e-commerce product page. It&nbsp;automatically makes available all the product data you&rsquo;d expect: price, discount/savings amount, shipping cost, product description, any relevant product images, SKU and/or other product IDs.<\/p>\n\n<p>Even cooler: pair the Product API with <a href=\"http://www.diffbot.com/products/crawlbot\" title=\"Crawlbot from Diffbot\">Crawlbot<\/a>, our intelligent site-spidering tool, and let Diffbot determine which pages are products, then automatically structure the entire catalog. Here&rsquo;s a quick demonstration of Crawlbot at work:<\/p>\n\n<p>We&rsquo;ve developed the Product API over the course of two years, building upon our core vision technology that&rsquo;s extracted structured data from billions of web pages, and training our machine learning systems using data from tens of thousands of unique shopping sites. We can&rsquo;t wait for you to try it out.<\/p>\n<p>What are you waiting for? Check out the <a href=\"http://www.diffbot.com/products/automatic/product\" title=\"Diffbot's Product API\">Product API documentation<\/a>&nbsp;and dive on in! If you need a token, check out our <a href=\"http://www.diffbot.com/pricing\">pricing and plans<\/a> (including our Free plan).<\/p>\n<p>Questions? Hit us up at <a href=\"mailto:support@diffbot.com\">support@diffbot.com<\/a>.<\/p>\n\t\t\t\t\t<\/div><\/div><\/div>",
	"supertags":[
		{
			"id":82871,
			"positions":[[161,176],[764,779]],
			"name":"Data model",
			"score":0.5415940109802859,
			"contentMatch":0.44444444444444453,
			"categories":{"1116481":"Data modeling"},
			"type":1,
			"senseRank":1,
			"variety":0.7011494252873562,
			"depth":0.5882352941176471
		},
		{
			"id":540448,
			"positions":[[80,87],[115,122],[197,204],[252,259],[326,333],[360,367],[393,400],[428,435],[535,543],[664,671],[987,994]],
			"scope":"chemistry",
			"name":"Product (chemistry)",
			"score":0.5192867889085311,
			"contentMatch":0.11111111111111113,
			"categories":{"995738":"Chemical reactions"},
			"type":1,
			"senseRank":1,
			"variety":0.8275862068965517,
			"depth":0.5294117647058824
		},
		{
			"id":519280,
			"positions":[[460,471]],
			"name":"Intelligence",
			"score":0.516535415482424,
			"contentMatch":0.4920634920634922,
			"categories":{
				"2871217":"Educational psychology",
				"4942617":"Intelligence",
				"15993560":"Psychological testing",
				"2870476":"Developmental psychology"
			},
			"type":1,
			"senseRank":1,
			"variety":0.5517241379310345,
			"depth":0.7058823529411764
		},
		{
			"id":29816,
			"positions":[[736,746]],
			"name":"Technology",
			"score":0.500092248978243,
			"contentMatch":0.9365079365079365,
			"categories":{
				"696648":"Technology","10958482":"Technology systems"
			},
			"type":1,
			"senseRank":1,
			"variety":0.2183908045977011,
			"depth":0.7647058823529411
		}
	]
}
`
