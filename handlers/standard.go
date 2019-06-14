package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"../helpers"
	"github.com/byuoitav/common/log"
	se "github.com/byuoitav/common/status"
	"github.com/labstack/echo"
)

func powerOn(context echo.Context, pooled bool) error {
	address := context.Param("address")

	err := helpers.PowerOn(address, pooled)
	if err != nil {
		log.L.Warnf(err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, se.Power{"on"})

}

func powerStandby(context echo.Context, pooled bool) error {
	address := context.Param("address")

	err := helpers.PowerStandby(address, pooled)
	if err != nil {
		log.L.Warnf(err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}
	return context.JSON(http.StatusOK, se.Power{"standby"})

}

func displayBlank(context echo.Context, pooled bool) error {
	log.L.Debugf("Blanking Display..")

	address := context.Param("address")

	err := helpers.SetBlank(address, true, pooled)
	if err != nil {
		log.L.Warnf(err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, se.Blanked{true})
}

func displayUnBlank(context echo.Context, pooled bool) error {
	log.L.Debugf("Unblanking Display..")

	address := context.Param("address")

	err := helpers.SetBlank(address, false, pooled)
	if err != nil {
		log.L.Warnf(err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, se.Blanked{false})
}

func setInputPort(context echo.Context, pooled bool) error {
	log.L.Debugf("Setting input...")

	port := context.Param("port")
	address := context.Param("address")

	err := helpers.SetInput(address, port, pooled)
	if err != nil {
		log.L.Warnf(err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, se.Input{port})
}

func powerStatus(context echo.Context, pooled bool) error {
	address := context.Param("address")

	status, err := helpers.GetPower(address, pooled)
	if err != nil {
		log.L.Warnf(err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, status)
}

func blankedStatus(context echo.Context, pooled bool) error {
	address := context.Param("address")

	status, err := helpers.GetBlankStatus(address, pooled)
	if err != nil {
		log.L.Warnf(err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, status)
}

func inputStatus(context echo.Context, pooled bool) error {
	address := context.Param("address")

	status, err := helpers.GetInputStatus(address, pooled)
	if err != nil {
		log.L.Warnf(err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, status)
}

func hasActiveSignal(context echo.Context, pooled bool) error {
	address := context.Param("address")
	port := context.Param("port")

	active, err := helpers.HasActiveSignal(address, port, pooled)
	if err != nil {
		log.L.Warnf(err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, active)
}

func getHardwareInfo(context echo.Context, pooled bool) error {
	address := context.Param("address")

	hardware, err := helpers.GetHardwareInfo(address, pooled)
	if err != nil {
		log.L.Warnf(err.Error())
		return context.JSON(http.StatusInternalServerError, err.Error())
	}

	return context.JSON(http.StatusOK, hardware)
}
