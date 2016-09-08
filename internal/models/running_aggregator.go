package models

import (
	"github.com/influxdata/telegraf"
)

type RunningAggregator struct {
	Name       string
	Aggregator telegraf.Aggregator
	Config     *AggregatorConfig
}

// FilterConfig containing a name and filter
type AggregatorConfig struct {
	Name   string
	Filter Filter
}

// func (rf *RunningProcessor) Apply(in ...telegraf.Metric) []telegraf.Metric {
// 	ret := []telegraf.Metric{}

// 	for _, metric := range in {
// 		if rf.Config.Filter.IsActive() {
// 			// check if the filter should be applied to this metric
// 			if ok := rf.Config.Filter.Apply(metric.Name(), metric.Fields(), metric.Tags()); !ok {
// 				// this means filter should not be applied
// 				ret = append(ret, metric)
// 				continue
// 			}
// 		}
// 		// This metric should pass through the filter, so call the filter Apply
// 		// function and append results to the output slice.
// 		ret = append(ret, rf.Processor.Apply(metric)...)
// 	}

// 	return ret
// }
