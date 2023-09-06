package livemonitor

import "runtime"

type Metric struct {
	GOOS         string
	GOARCH       string
	NumCpu       int
	GoVersion    string
	MaxProcess   int
	NumGoroutine int
	NumCgoCall   int64
}

func Collect() Metric {
	return Metric{
		NumCgoCall:   runtime.NumCgoCall(),
		NumCpu:       runtime.NumCPU(),
		NumGoroutine: runtime.NumGoroutine(),
		MaxProcess:   runtime.GOMAXPROCS(0),
		GOARCH:       runtime.GOARCH,
		GOOS:         runtime.GOOS,
		GoVersion:    runtime.Version(),
	}
}
