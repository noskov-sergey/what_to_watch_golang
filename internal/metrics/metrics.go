package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

const (
	CreateHandler    = "create"
	GetHandler       = "Get"
	GetRandomHandler = "GetRandom"
)

type Metrics struct {
	Success  *prometheus.CounterVec
	Failures *prometheus.CounterVec
}

type Met struct {
	Err     error
	Handler string
}

func NewMetrics(reg prometheus.Registerer) *Metrics {
	m := &Metrics{
		Success: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "success",
				Help: "Number of success.",
			},
			[]string{"handler"}),
		Failures: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "failure",
				Help: "Number of errors.",
			},
			[]string{"handler"},
		),
	}
	reg.MustRegister(m.Success)
	reg.MustRegister(m.Failures)
	return m
}

func (m *Metrics) Add(met Met) {
	if met.Err != nil {
		m.Failures.With(prometheus.Labels{"handler": met.Handler}).Inc()
		return
	}
	m.Success.With(prometheus.Labels{"handler": met.Handler}).Inc()
}
