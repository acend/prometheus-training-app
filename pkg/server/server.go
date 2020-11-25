package server

import (
	"fmt"
	"net/http"
)

func metrics(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintf(w, "http_request_count{foo=\"bar\"} 0\n")
}

func invalidMetrics(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintf(w, "http_request_counter{foo=\"bar\"} 0\n")
}

func HttpServer(_ []string) {
	http.HandleFunc("/metrics", metrics)
	http.HandleFunc("/metrics/invalid", invalidMetrics)

	_ = http.ListenAndServe(":8080", nil)
}
