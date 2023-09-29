package core

type Config struct {
	Log struct {
		Level string
	}
	IP struct {
		IPV4 string
		IPV6 string
	}
	Service struct {
		Interval uint32
	}
	Services map[string]map[string]any
}
