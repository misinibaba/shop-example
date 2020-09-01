package jobs

import (
	"context"
	"time"
)

func Jobs(callBackFunc func(), dur int, ctx context.Context) {
	ticker := time.NewTicker(time.Second * time.Duration(dur))

	for range ticker.C {
		select {
		case <-ctx.Done():
			return
		default:
			callBackFunc()
		}
	}
}

