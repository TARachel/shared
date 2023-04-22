package rpcclient

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/twitchtv/twirp"
)

// NewLoggingClientHooks logs request and errors to stdout in the client
func NewLoggingClientHooks() *twirp.ClientHooks {
	return &twirp.ClientHooks{
		RequestPrepared: func(ctx context.Context, r *http.Request) (context.Context, error) {
			fmt.Printf("Req: %s %s\n", r.Host, r.URL.Path)
			return ctx, nil
		},
		Error: func(ctx context.Context, twerr twirp.Error) {
			log.Println("Error: " + string(twerr.Code()))
		},
		ResponseReceived: func(ctx context.Context) {
			log.Println("Success")
		},
	}
}
