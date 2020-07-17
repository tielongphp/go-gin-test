package tool

func FormatElapse(elapse int64) float64 {
	latency := elapse / 1000
	floatLatency := float64(latency)
	return floatLatency / 1000.0
}
