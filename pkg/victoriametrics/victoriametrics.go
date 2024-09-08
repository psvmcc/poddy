package victoriametrics

import (
	"log"
	"net/http"

	"github.com/VictoriaMetrics/metrics"
)

func ListenMetricsServer(bind string) error {
	metricsSRV := http.NewServeMux()
	metricsSRV.HandleFunc("/metrics", func(w http.ResponseWriter, _ *http.Request) {
		metrics.WritePrometheus(w, true)
	})
	err := http.ListenAndServe(bind, metricsSRV)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
