// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package diffbot

import (
	"testing"
	"time"
)

func TestOptions(t *testing.T) {
	for i, v := range testOptionsList {
		if s := v.opt.MethodParamString(v.method); s != v.str {
			t.Fatalf("%d: expect = %q, got = %q", i, v.str, s)
		}
	}
}

var testOptionsList = []struct {
	method string
	opt    Options
	str    string
}{
	// empty
	{
		method: "",
		opt: Options{
			Fields:   "meta,querystring,images(*)",
			Timeout:  time.Second * 5,
			Callback: "abc",
		},
		str: "",
	},

	// case "article", "image", "product":
	{
		method: "article",
		opt: Options{
			Fields:   "meta,querystring,images(*)",
			Timeout:  time.Second * 5,
			Callback: "abc",
		},
		str: "&fields=meta,querystring,images(*)&timeout=5000&callback=abc",
	},
	{
		method: "image",
		opt: Options{
			Fields:   "meta,querystring,images(*)",
			Timeout:  time.Second * 5,
			Callback: "abc",
		},
		str: "&fields=meta,querystring,images(*)&timeout=5000&callback=abc",
	},
	{
		method: "product",
		opt: Options{
			Fields:   "meta,querystring,images(*)",
			Timeout:  time.Second * 5,
			Callback: "abc",
		},
		str: "&fields=meta,querystring,images(*)&timeout=5000&callback=abc",
	},

	// case "frontpage":
	{
		method: "frontpage",
		opt: Options{
			Fields:       "meta,querystring,images(*)",
			Timeout:      time.Second * 5,
			Callback:     "abc",
			FrontpageAll: "*",
		},
		str: "&timeout=5000&all=*",
	},

	// case "analyze":
	{
		method: "analyze",
		opt: Options{
			Fields:          "meta,querystring,images(*)",
			Timeout:         time.Second * 5,
			Callback:        "abc",
			FrontpageAll:    "*",
			ClassifierMode:  "frontpage",
			ClassifierStats: "abc",
		},
		str: "&mode=frontpage&fields=meta,querystring,images(*)&stats=abc",
	},

	// case "bulk":
	// case "crawl":
	// case "batch":
	// default: // Custom APIs
}
