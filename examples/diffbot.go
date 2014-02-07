// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Diffbot Client
package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/diffbot/diffbot-go-client"
)

var (
	token  = flag.String("token", "0123456789abcdef0123456789abcdef", "diffbot token")
	server = flag.String("server", diffbot.DefaultServer, "diffbot server")
	method = flag.String("method", "article", "method: article/frontpage/")
	url    = flag.String("url",
		`http://blog.diffbot.com/diffbots-new-product-api-teaches-robots-to-shop-online/`,
		"url",
	)
)

func main() {
	flag.Parse()
	fmt.Printf("token: %s\n", *token)
	fmt.Printf("server: %s\n", *server)
	fmt.Printf("method: %s\n", *method)
	fmt.Printf("url: %s\n", *url)

	switch *method {
	case "article":
		callArticle()
	case "frontpage":
		callFrontpage()
	default:
		log.Fatalf("Unknown method: %s\n", *method)
	}
}

func callArticle() {
	opt := &diffbot.Options{Fields: "*"}
	article, err := diffbot.ParseArticle(*token, *url, opt)
	if err != nil {
		if apiErr, ok := err.(*diffbot.Error); ok {
			// ApiError, e.g. {"error":"Not authorized API token.","errorCode":401}
			log.Fatal(apiErr)
		}
		log.Fatal(err)
	}
	fmt.Println(article)
}

func callFrontpage() {
	opt := &diffbot.Options{Fields: "*"}
	article, err := diffbot.ParseArticle(*token, *url, opt)
	if err != nil {
		if apiErr, ok := err.(*diffbot.Error); ok {
			// ApiError, e.g. {"error":"Not authorized API token.","errorCode":401}
			log.Fatal(apiErr)
		}
		log.Fatal(err)
	}
	fmt.Println(article)
}
