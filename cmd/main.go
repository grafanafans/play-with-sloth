package main

import (
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	httpDurationsHistogram = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "http_request_duration_seconds",
		Help: "HTTP request latency distributions.",
	}, []string{"code"})

	errRate float64
)

func main() {
	r := prometheus.NewRegistry()
	r.MustRegister(httpDurationsHistogram)

	updateMetrics()

	http.HandleFunc("/errrate", func(w http.ResponseWriter, r *http.Request) {
		errRate, _ = strconv.ParseFloat(r.URL.Query().Get("value"), 64)
		w.Write([]byte("ok"))
	})

	http.Handle("/metrics", promhttp.HandlerFor(r, promhttp.HandlerOpts{
		EnableOpenMetrics: true,
		Registry:          r,
	}))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func updateMetrics() {
	go func() {
		ticker := time.NewTicker(10 * time.Millisecond)

		for {
			select {
			case <-ticker.C:
				randDuration := rand.Float64()
				code := "200"
				if randDuration < errRate {
					code = "500"
				}

				httpDurationsHistogram.WithLabelValues(code).Observe(randDuration)
			}
		}
	}()
}
