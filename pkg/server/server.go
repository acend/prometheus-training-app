package server

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"prometheus-sample-app/pkg/business"
)

func invalidMetrics(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintf(w, "sampleapp_readfile_count 0\n")
}

func HttpServer(_ []string) {
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/metrics/invalid", invalidMetrics)

	//business functions
	http.HandleFunc("/business/readfile", business.ReadFile)

	_ = http.ListenAndServe(":8080", nil)
}
