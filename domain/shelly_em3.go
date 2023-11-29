package domain

type ShellyData struct {
	ID                  int      `json:"id"`                    // Id of the EM component instance
	ACurrent            float64  `json:"a_current"`             // Phase A current measurement value, [A]
	AVoltage            float64  `json:"a_voltage"`             // Phase A voltage measurement value, [V]
	AActPower           float64  `json:"a_act_power"`           // Phase A active power measurement value, [W]
	AAprtPower          float64  `json:"a_aprt_power"`          // Phase A apparent power measurement value, [VA]
	APf                 float64  `json:"a_pf"`                  // Phase A power factor measurement value
	AFreq               float64  `json:"a_freq"`                // Phase A network frequency measurement value
	AErrors             []string `json:"a_errors"`              // Phase A error conditions occurred
	BCurrent            float64  `json:"b_current"`             // Phase B current measurement value, [A]
	BVoltage            float64  `json:"b_voltage"`             // Phase B voltage measurement value, [V]
	BActPower           float64  `json:"b_act_power"`           // Phase B active power measurement value, [W]
	BAprtPower          float64  `json:"b_aprt_power"`          // Phase B apparent power measurement value, [VA]
	BPf                 float64  `json:"b_pf"`                  // Phase B power factor measurement value
	BFreq               float64  `json:"b_freq"`                // Phase B network frequency measurement value
	BErrors             []string `json:"b_errors"`              // Phase B error conditions occurred
	CCurrent            float64  `json:"c_current"`             // Phase C current measurement value, [A]
	CVoltage            float64  `json:"c_voltage"`             // Phase C voltage measurement value, [V]
	CActPower           float64  `json:"c_act_power"`           // Phase C active power measurement value, [W]
	CAprtPower          float64  `json:"c_aprt_power"`          // Phase C apparent power measurement value, [VA]
	CPf                 float64  `json:"c_pf"`                  // Phase C power factor measurement value
	CFreq               float64  `json:"c_freq"`                // Phase C network frequency measurement value
	CErrors             []string `json:"c_errors"`              // Phase C error conditions occurred
	NCurrent            float64  `json:"n_current"`             // Neutral current measurement value, [A]
	NErrors             []string `json:"n_errors"`              // Neutral error conditions occurred
	TotalCurrent        float64  `json:"total_current"`         // Sum of the current on all phases(excluding neutral readings if available)
	TotalActPower       float64  `json:"total_act_power"`       // Sum of the active power on all phases
	TotalAprtPower      float64  `json:"total_aprt_power"`      // Sum of the apparent power on all phases
	UserCalibratedPhase []string `json:"user_calibrated_phase"` // Indicates which phase was user calibrated
	Errors              []string `json:"errors"`                // EM component error conditions
}
