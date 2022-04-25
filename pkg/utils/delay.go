package utils

import (
	"context"
	"time"
)

func After(ctx context.Context, afterTime time.Duration, do func()) {
	go func() {

		for {
			select {
			case <-time.After(afterTime):
				do()
			case <-ctx.Done():
				return
			}
		}
	}()
}
