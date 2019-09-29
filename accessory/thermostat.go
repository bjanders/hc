package accessory

import (
	"github.com/brutella/hc/service"
)

type Thermostat struct {
	*Accessory

	Thermostat *service.Thermostat
}

// NewThermostat returns a Thermostat which implements model.Thermostat.
//func NewThermostat(info Info, temp, minTemp, maxTemp, stepsTemp, setTHCS, minTHCS, maxTHCS, stepsTHCS float64) *Thermostat {
//func NewThermostat(info Info, temp, minTemp, maxTemp, stepsTemp float64)  *Thermostat {
func NewThermostat(info Info, temp, minTemp, maxTemp, stepsTemp float64)  *Thermostat {
	acc := Thermostat{}
	acc.Accessory = New(info, TypeThermostat)
	acc.Thermostat = service.NewThermostat()

	acc.Thermostat.CurrentTemperature.SetValue(temp)
	acc.Thermostat.CurrentTemperature.SetMinValue(minTemp)
	acc.Thermostat.CurrentTemperature.SetMaxValue(maxTemp)
	acc.Thermostat.CurrentTemperature.SetStepValue(stepsTemp)

	acc.Thermostat.TargetTemperature.SetValue(temp)
	acc.Thermostat.TargetTemperature.SetMinValue(minTemp)
	acc.Thermostat.TargetTemperature.SetMaxValue(maxTemp)
	acc.Thermostat.TargetTemperature.SetStepValue(stepsTemp)

//	acc.Thermostat.TargetTempearture.SetValue(temp)
//	acc.Thermostat.TargetHeatingCoolingState.SetMinValue(minTemp)
//	acc.Thermostat.TargetHeatingCoolingState.SetMaxValue(maxTemp)
//	acc.Thermostat.TargetHeatingCoolingState.SetStepValue(stepsTemp)

	acc.AddService(acc.Thermostat.Service)

	return &acc
}
