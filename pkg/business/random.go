package business

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"math/rand"
	"net/http"
	"time"
)

var (
	randomValueHistogram = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "sampleapp_random_values",
		Help:    "Generated random values",
		Buckets: prometheus.LinearBuckets(0, 2, 7),
	})

)

func init() {
	prometheus.MustRegister(randomValueHistogram)
}

func Random(w http.ResponseWriter, _ *http.Request) {
	rand.Seed(time.Now().UnixNano())
	v := rand.Intn(20)

	randomValueHistogram.(prometheus.Observer).Observe(float64(v))

	_, _ = fmt.Fprintf(w, "generated value %d", v)

}