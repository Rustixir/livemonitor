package livemonitor

import (
	"context"
	"github.com/Rustixir/skyview/component"
	"github.com/Rustixir/skyview/server"
	"net/http"
	"os"
	"time"
)

func ListenAndServe(ctx context.Context, interval time.Duration, port string) {
	builder := component.NewBuilder("secret", "LiveMonitor", 100)
	builder.AddComponent("stats", true, NewStatsComp(), []string{"refresh", "mem-stats"})
	startJob(ctx, interval, builder.GetNotify())
	server.AddHandler("/logo", &LogoHandler{})
	builder.Start(port)
}

type LogoHandler struct {
	cache []byte
}

func (r *LogoHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if r.cache == nil {
		r.cache, _ = os.ReadFile("./logo.png")
	}
	writer.Header().Add("Content-Type", "image/png")
	writer.Write(r.cache)
}
