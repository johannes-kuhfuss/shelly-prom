package app

import "github.com/prometheus/client_golang/prometheus"

func initMetrics() {
	cfg.Metrics.VoltageGauge = *prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "Shelly",
		Subsystem: "Energy",
		Name:      "Voltage",
		Help:      "Voltage measured in Volts (V)",
	}, []string{
		"Device",
		"Phase",
	})
	cfg.Metrics.CurrentGauge = *prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "Shelly",
		Subsystem: "Energy",
		Name:      "Current",
		Help:      "Current measured in Ampere (A)",
	}, []string{
		"Device",
		"Phase",
	})
	cfg.Metrics.ActivePowerGauge = *prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "Shelly",
		Subsystem: "Energy",
		Name:      "ActivePower",
		Help:      "Active Power measured in Watts (W)",
	}, []string{
		"Device",
		"Phase",
	})
	cfg.Metrics.ApparentPowerGauge = *prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "Shelly",
		Subsystem: "Energy",
		Name:      "ApparentPower",
		Help:      "Apparent Power measured in Volt-Ampere (VA)",
	}, []string{
		"Device",
		"Phase",
	})
	cfg.Metrics.PowerFactorGauge = *prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "Shelly",
		Subsystem: "Energy",
		Name:      "PowerFactor",
		Help:      "Power Factor",
	}, []string{
		"Device",
		"Phase",
	})
	cfg.Metrics.FrequencyGauge = *prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "Shelly",
		Subsystem: "Energy",
		Name:      "Frequency",
		Help:      "Freqency measured in Hertz (Hz)",
	}, []string{
		"Device",
		"Phase",
	})
	prometheus.MustRegister(cfg.Metrics.VoltageGauge)
	prometheus.MustRegister(cfg.Metrics.CurrentGauge)
	prometheus.MustRegister(cfg.Metrics.ActivePowerGauge)
	prometheus.MustRegister(cfg.Metrics.ApparentPowerGauge)
	prometheus.MustRegister(cfg.Metrics.PowerFactorGauge)
	prometheus.MustRegister(cfg.Metrics.FrequencyGauge)
}
