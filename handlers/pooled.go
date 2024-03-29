package handlers

import (
	"github.com/labstack/echo"
)

// PowerOnPooled turns the power on when pooled
func PowerOnPooled(context echo.Context) error {
	return powerOn(context, true)
}

// PowerOn turns the power on
func PowerOn(context echo.Context) error {
	return powerOn(context, false)
}

// PowerStandbyPooled sets the power state to standby when pooled
func PowerStandbyPooled(context echo.Context) error {
	return powerStandby(context, true)
}

// PowerStandby sets the power state to standby
func PowerStandby(context echo.Context) error {
	return powerStandby(context, false)
}

// DisplayBlankPooled blanks the display when pooled
func DisplayBlankPooled(context echo.Context) error {
	return displayBlank(context, true)
}

// DisplayBlank blanks the display
func DisplayBlank(context echo.Context) error {
	return displayBlank(context, false)
}

// DisplayUnBlankPooled unblanks the display when pooled
func DisplayUnBlankPooled(context echo.Context) error {
	return displayUnBlank(context, true)
}

// DisplayUnBlank unblanks the display
func DisplayUnBlank(context echo.Context) error {
	return displayUnBlank(context, false)
}

// SetInputPortPooled sets the input port when pooled
func SetInputPortPooled(context echo.Context) error {
	return setInputPort(context, true)
}

// SetInputPort sets the input port
func SetInputPort(context echo.Context) error {
	return setInputPort(context, false)
}

// PowerStatusPooled checks the power status when pooled
func PowerStatusPooled(context echo.Context) error {
	return powerStatus(context, true)
}

// PowerStatus checks the power status
func PowerStatus(context echo.Context) error {
	return powerStatus(context, false)
}

// BlankedStatusPooled checks if the display is blanked when pooled
func BlankedStatusPooled(context echo.Context) error {
	return blankedStatus(context, true)
}

// BlankedStatus checks if the display is blanked
func BlankedStatus(context echo.Context) error {
	return blankedStatus(context, false)
}

// InputStatusPooled checks what input the display is on when pooled
func InputStatusPooled(context echo.Context) error {
	return inputStatus(context, true)
}

// InputStatus checks what input the display is on
func InputStatus(context echo.Context) error {
	return inputStatus(context, false)
}

// HasActiveSignalPooled checks to see if the display has an active input when pooled
func HasActiveSignalPooled(context echo.Context) error {
	return hasActiveSignal(context, true)
}

// HasActiveSignal checks to see if the display has an active input
func HasActiveSignal(context echo.Context) error {
	return hasActiveSignal(context, false)
}

// GetHardwareInfoPooled obtains the hardware information when pooled
func GetHardwareInfoPooled(context echo.Context) error {
	return getHardwareInfo(context, true)
}

// GetHardwareInfo obtains the hardware information
func GetHardwareInfo(context echo.Context) error {
	return getHardwareInfo(context, false)
}
