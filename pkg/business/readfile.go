package business

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"math/rand"
	"net/http"
	"time"
)

var (
	//define some metrics
	readFileBytesTotal = promauto.NewCounter(prometheus.CounterOpts{
		Name: "sampleapp_readfile_bytes_total",
		Help: "The total number of read bytes",
	})

	readFileTotal = promauto.NewCounter(prometheus.CounterOpts{
		Name: "sampleapp_readfile_total",
		Help: "The total number of read files",
	})
)

func ReadFile(w http.ResponseWriter, _ *http.Request) {
	rand.Seed(time.Now().UnixNano())
	v := rand.Intn(10000)

	//increment the counter by 1
	readFileTotal.Inc()

	//increment the counter by a specific value
	readFileBytesTotal.Add(float64(v))

	_, _ = fmt.Fprintf(w, "processed %d bytes", v)
}
