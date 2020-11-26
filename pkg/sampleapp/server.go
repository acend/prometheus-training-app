package sampleapp

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"prometheus-training-app/pkg/business"
)

func invalidMetrics(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintf(w, "sampleapp_readfile_count 0\n")
}

func HttpServer(port int) {
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/metrics/invalid", invalidMetrics)

	//business functions
	http.HandleFunc("/business/readfile", business.ReadFile)


	log.Printf("Starting sample app on port %d", port)
	log.Printf("Examine metric endpoint: curl localhost:%d/metrics", port)
	err := http.ListenAndServe(fmt.Sprint(":", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
