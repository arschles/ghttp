// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// this file has been modified from its original package --  golang.org/x/net/context/ctxhttp --
// to work with the Go 1.7 build-in context package

// +build go1.7

package grest

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// Do sends an HTTP request with the provided http.Client and returns
// an HTTP response.
//
// If the client is nil, http.DefaultClient is used.
//
// The provided ctx must be non-nil. If it is canceled or times out,
// ctx.Err() will be returned.
func ctxDo(ctx context.Context, client *http.Client, req *http.Request) (*http.Response, error) {
	if client == nil {
		client = http.DefaultClient
	}
	resp, err := client.Do(req.WithContext(ctx))
	// If we got an error, and the context has been canceled,
	// the context's error is probably more useful.
	if err != nil {
		select {
		case <-ctx.Done():
			err = ctx.Err()
		default:
		}
	}
	return resp, err
}

// Get issues a GET request via the Do function.
func ctxGet(ctx context.Context, client *http.Client, url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return Do(ctx, client, req)
}

// Head issues a HEAD request via the Do function.
func ctxHead(ctx context.Context, client *http.Client, url string) (*http.Response, error) {
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return nil, err
	}
	return Do(ctx, client, req)
}

// Post issues a POST request via the Do function.
func ctxPost(ctx context.Context, client *http.Client, url string, bodyType string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", bodyType)
	return Do(ctx, client, req)
}

// PostForm issues a POST request via the Do function.
func ctxPostForm(ctx context.Context, client *http.Client, url string, data url.Values) (*http.Response, error) {
	return Post(ctx, client, url, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
}
