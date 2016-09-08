package min

import (
	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/aggregators"
)

type Min struct {
}

var sampleConfig = `
`

func (m *Min) SampleConfig() string {
	return sampleConfig
}

func (m *Min) Description() string {
	return "Aggregate the minimum value of each numerical field."
}

func (m *Min) Apply(in telegraf.Metric) {
}

func (m *Min) Start(acc telegraf.Accumulator) {
}

func (m *Min) Stop() {
}

func init() {
	aggregators.Add("min", func() telegraf.Aggregator {
		return &Min{}
	})
}
