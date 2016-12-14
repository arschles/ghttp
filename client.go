package grest

import (
	"context"
	"net/http"
)

type Client struct {
	cl *http.Client
}

func (c *Client) Do(ctx context.Context, r Request) (*http.Response, error) {
	return ctxDo(ctx, c.cl, r.ToHTTP())
}
