package download

import (
	"context"
	"time"
)

func Downloading(ctx context.Context) error {
	ticker := time.NewTicker(time.Minute)
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			//downloading
		}
	}
}
