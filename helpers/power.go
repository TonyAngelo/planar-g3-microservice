package helpers

import (
	"strings"

	"github.com/byuoitav/common/log"

	"github.com/byuoitav/common/nerr"
	"github.com/byuoitav/common/status"
	"github.com/fatih/color"
)

func PowerOn(address string, pooled bool) *nerr.E {
	log.L.Infof("Setting power of %v to on", address)
	command := "SYSTEM.POWER=ON"

	return sendCommand(command, address, pooled)
}

func PowerStandby(address string, pooled bool) *nerr.E {
	log.L.Infof("Seting power of %v to off", address)
	command := "SYSTEM.POWER=OFF"

	return sendCommand(command, address, pooled)
}

func GetPower(address string, pooled bool) (status.Power, *nerr.E) {

	log.L.Infof("%s", color.HiCyanString("[helpers] querying power state of %v", address))

	response, err := queryState("SYSTEM.POWER?", address, pooled)
	if err != nil {
		return status.Power{}, err
	}

	var status status.Power
	responseString := string(response)

	if strings.Contains(responseString, "OFF") {
		status.Power = "standby"
	} else if strings.Contains(responseString, "ON") {
		status.Power = "on"
	}

	return status, nil
}
