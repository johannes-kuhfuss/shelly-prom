package domain

type ShellyData struct {
	ID                  int     `json:"id"`
	ACurrent            float64 `json:"a_current"`
	AVoltage            float64 `json:"a_voltage"`
	AActPower           float64 `json:"a_act_power"`
	AAprtPower          float64 `json:"a_aprt_power"`
	APf                 float64 `json:"a_pf"`
	AFreq               float64 `json:"a_freq"`
	BCurrent            float64 `json:"b_current"`
	BVoltage            float64 `json:"b_voltage"`
	BActPower           float64 `json:"b_act_power"`
	BAprtPower          float64 `json:"b_aprt_power"`
	BPf                 float64 `json:"b_pf"`
	BFreq               float64 `json:"b_freq"`
	CCurrent            float64 `json:"c_current"`
	CVoltage            float64 `json:"c_voltage"`
	CActPower           float64 `json:"c_act_power"`
	CAprtPower          float64 `json:"c_aprt_power"`
	CPf                 float64 `json:"c_pf"`
	CFreq               float64 `json:"c_freq"`
	NCurrent            float64 `json:"n_current"`
	TotalCurrent        float64 `json:"total_current"`
	TotalActPower       float64 `json:"total_act_power"`
	TotalAprtPower      float64 `json:"total_aprt_power"`
	UserCalibratedPhase []any   `json:"user_calibrated_phase"`
}
