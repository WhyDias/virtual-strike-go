package monitoring

import "github.com/prometheus/client_golang/prometheus"

var (
	ErrorHandler = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "error_msg_handler",
			Help: "Number of errors occured in app runtime",
		},
		[]string{"error_message"},
	)
)

func Init() {
	prometheus.MustRegister(ErrorHandler)
}
