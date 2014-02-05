// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package diffbot

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestClassification_type(t *testing.T) {
	var a Classification
	a.Stats = classificationStatsType{}
	a.Stats.Types = classificationStatsTypeType{}
	_ = a
}

func TestClassification_parseJson(t *testing.T) {
	var result1 Classification
	if err := json.Unmarshal([]byte(testJsonDataClassification), &result1); err != nil {
		t.Fatal(err)
	}
	var result2 Classification
	if err := json.Unmarshal([]byte(result1.String()), &result2); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(result1, result2) {
		t.Fatalf("not equal, expect = %q, got = %q", result1, result2)
	}
}

const testJsonDataClassification = `
{
  "type": "article",
  "stats": {
     "types": {
        "article": 0.46,
        "audio": 0.15,
        "chart": 0.01,
        "discussion": 0.03,
        "document": 0.04,
        "download": 0.01,
        "error": 0.00,
        "event": 0.00,
        "faq": 0.02,
        "frontpage": 0.12,
        "game": 0.01,
        "image": 0.02,
        "job": 0.02,
        "location": 0.08,
        "product": 0.09,
        "profile": 0.09,
        "recipe": 0.08,
        "reviewslist": 0.09,
        "serp": 0.06,
        "video": 0.01
      }
    },
  "resolved_url": "http://techcrunch.com/2012/05/31/diffbot-raises-2-million-seed-round-for-web-content-extraction-technology/",
  "url": "http://tcrn.ch/Jw7ZKw",
  "human_language": "en"
}
`
