package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Define a counter metric to count requests.
var requests = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "http_requests_total",
	Help: "Total number of HTTP requests.",
})

// Define a histogram metric to measure request duration.
var requestDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
	Name:    "http_request_duration_seconds",
	Help:    "HTTP request duration in seconds.",
	Buckets: prometheus.DefBuckets,
})

// Define a function to create a Prometheus middleware.
func prometheusMiddleware() gin.HandlerFunc {
	// Create a new Prometheus handler.
	handler := promhttp.Handler()

	return func(c *gin.Context) {
		// If the request is for the metrics endpoint, handle it with the Prometheus handler.
		if c.Request.URL.Path == "/metrics" {
			handler.ServeHTTP(c.Writer, c.Request)
			return
		}

		// Otherwise, call the next middleware handler.
		c.Next()
	}
}
