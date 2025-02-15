package middlewares

import (
    "strconv"
    "time"

    "github.com/labstack/echo/v4"
    "github.com/prometheus/client_golang/prometheus"
)

var (
    httpRequestsTotal = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "http_requests_total",
            Help: "Total number of HTTP requests",
        },
        []string{"method", "endpoint", "status"},
    )

    httpRequestDuration = prometheus.NewHistogramVec(
        prometheus.HistogramOpts{
            Name:    "http_request_duration_seconds",
            Help:    "HTTP request duration in seconds",
            Buckets: []float64{0.1, 0.5, 1, 2, 5},
        },
        []string{"method", "endpoint"},
    )
)

func MetricsMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        start := time.Now()

        err := next(c)

        duration := time.Since(start).Seconds()
        status := c.Response().Status
        method := c.Request().Method
        endpoint := c.Path()

        httpRequestsTotal.WithLabelValues(method, endpoint, strconv.Itoa(status)).Inc()
        httpRequestDuration.WithLabelValues(method, endpoint).Observe(duration)

        return err
    }
}

func RegisterMetrics(e *echo.Echo) {
    prometheus.MustRegister(httpRequestsTotal)
    prometheus.MustRegister(httpRequestDuration)

    e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
}