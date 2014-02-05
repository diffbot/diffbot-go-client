// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package diffbot

import (
	"fmt"
	"io/ioutil"
	"net/http"
	urlPkg "net/url"
)

const (
	DefaultServer = `http://api.diffbot.com/v2`
)

// Diffbot uses computer vision, natural language processing
// and machine learning to automatically recognize
// and structure specific page-types.
func Diffbot(method, token, url string, opt *Options) (body []byte, err error) {
	return DiffbotServer(DefaultServer, method, token, url, opt)
}

// DiffbotServer like Diffbot function, but support custom server.
func DiffbotServer(server, method, token, url string, opt *Options) (body []byte, err error) {
	req, err := http.NewRequest("GET", makeRequestUrl(server, method, token, url, opt), nil)
	if err != nil {
		return nil, err
	}
	if opt != nil && opt.CustomHeader != nil {
		req.Header = opt.CustomHeader
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		if len(body) != 0 {
			var apiError Error
			if err = apiError.ParseJson(string(body)); err != nil {
				err = &Error{
					ErrCode:    resp.StatusCode,
					ErrMessage: string(body),
				}
				return
			} else {
				err = &apiError
				return
			}
		} else {
			err = &Error{
				ErrCode:    resp.StatusCode,
				ErrMessage: resp.Status,
			}
			return
		}
	}
	return
}

func makeRequestUrl(server, method, token, webUrl string, opt *Options) string {
	return fmt.Sprintf("%s/%s?token=%s&url=%s%s",
		server, method, token, urlPkg.QueryEscape(webUrl), opt.MethodParamString(method),
	)
}
