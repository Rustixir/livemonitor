package livemonitor

import (
	"context"
	"github.com/Rustixir/skyview/notify"
	"github.com/jfyne/live"
	"runtime"
	"time"
)

func startJob(ctx context.Context, interval time.Duration, notify *notify.Local) {
	go func() {
		ticker := time.NewTicker(interval)
		memstats := time.NewTicker(interval * 3)
		for {
			select {
			case <-ticker.C:
				metric := Collect()
				params := live.Params{}
				params["metric"] = metric
				notify.Send(ctx, params)
			case <-memstats.C:
				memStats := runtime.MemStats{}
				runtime.ReadMemStats(&memStats)

				params := live.Params{}
				params["metric"] = memStats
				notify.Send(ctx, params)
			case <-ctx.Done():
				return
			}
		}
	}()
}
