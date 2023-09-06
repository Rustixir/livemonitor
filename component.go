package livemonitor

import (
	"errors"
	"github.com/Rustixir/skyview"
	"github.com/Rustixir/skyview/component"
	"github.com/jfyne/live"
	"runtime"
)

type StatsComp struct {
	Metric   Metric
	MemStats runtime.MemStats
}

func NewStatsComp() func() component.Component {
	return func() component.Component {
		return new(StatsComp)
	}
}

func (s *StatsComp) Init() error {
	s.Metric = Collect()
	return nil
}

func (s *StatsComp) HandleEvent(evtName string, p live.Params) skyview.Response {
	switch evtName {
	case skyview.NotifyEvent, "refresh":
		switch data := p["metric"].(type) {
		case Metric:
			s.Metric = data
		case runtime.MemStats:
			s.MemStats = data
		}
	default:
		err := errors.New("unknown command")
		return skyview.RaiseError(err)
	}
	return skyview.Ok()
}

func (s *StatsComp) HandleBroadcast(_ interface{}) error {
	return nil
}

func (s *StatsComp) Render() (name string, html string) {
	return "./view.html", ""
}

func (s *StatsComp) Terminate() error {
	return nil
}
