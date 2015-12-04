package client

import (
	"encoding/json"
	"fmt"
	"runtime"
)

type Usage struct {
	Data    []UsageData
	Product string `json:"-"`
}

// Create an instance of Usage and automatically include some usage information
// about the runtime memory stats.
func NewUsage(product string, data []UsageData) *Usage {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	d := []UsageData{
		{
			Tags: Tags{
				"os":   runtime.GOOS,
				"arch": runtime.GOARCH,
			},
			Values: Values{
				// The values here are an arbitrary selection of fields from MemStats
				// that look like they could be interesting when viewed in aggregate over time.
				"alloc":             m.Alloc,
				"heap_objects":      m.HeapObjects,
				"num_gc":            m.NumGC,
				"gc_pause_total_ns": m.PauseTotalNs,
				"gc_cpu_fraction":   m.GCCPUFraction,
			},
		},
	}

	return &Usage{
		Product: product,
		Data:    append(d, data...),
	}
}

func (u Usage) Path() string {
	return fmt.Sprintf("/usage/%s", u.Product)
}

func (u Usage) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.Data)
}

type UsageData struct {
	Tags   Tags   `json:"tags"`
	Values Values `json:"values"`
}
