// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package diffbot

import (
	"testing"
)

func TestError(t *testing.T) {
	var e Error
	for i := 0; i < len(testErrors); i++ {
		if err := e.ParseJson(testErrors[i].str); err != nil {
			t.Fatalf("%d: %v", i, err)
		}
		if a, b := testErrors[i].err.ErrCode, e.ErrCode; a != b {
			t.Fatalf("%d: expect = %v, got = %v", i, a, b)
		}
		if a, b := testErrors[i].err.ErrMessage, e.ErrMessage; a != b {
			t.Fatalf("%d: expect = %v, got = %v", i, a, b)
		}
		if a, b := testErrors[i].str, e.RawString; a != b {
			t.Fatalf("%d: expect = %v, got = %v", i, a, b)
		}
	}
}

var testErrors = []struct {
	str string
	err Error
}{
	{
		str: `
{"error":"Not authorized API token.","errorCode":401}
`,
		err: Error{
			ErrCode:    401,
			ErrMessage: "Not authorized API token.",
		},
	},
	{
		str: `
{
	"errorCode": 404,
	"error": "Something went crazy wrong.",
	"errorAnalysis": {
		"analyzed": "True",
		"blame": "human",
		"reasonCode": 817221,
		"reason": "inherent fallibility"
	},
	"suggestedImprovement": "Use a robot."
}
`,
		err: Error{
			ErrCode:    404,
			ErrMessage: "Something went crazy wrong.",
		},
	},
}
