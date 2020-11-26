package server

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func invalidMetrics(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintf(w, "http_request_counter{foo=\"bar\"} 0\n")
}

func HttpServer(_ []string) {
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/metrics/invalid", invalidMetrics)

	_ = http.ListenAndServe(":8080", nil)
}
