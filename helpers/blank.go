package helpers

import (
	"fmt"
	"strings"

	"github.com/byuoitav/common/log"

	"github.com/byuoitav/common/nerr"
	"github.com/byuoitav/common/status"
)

func SetBlank(address string, blank, pooled bool) *nerr.E {
	log.L.Infof("Setting blank on %s to %v", address, blank)

	var command string
	if blank {
		command = fmt.Sprintf("PATTERN=BLACK")
	} else {
		command = fmt.Sprintf("PATTERN=NONE")
	}

	return sendCommand(command, address, pooled)
}

func GetBlankStatus(address string, pooled bool) (status.Blanked, *nerr.E) {
	log.L.Infof("Querying blank status of %s", address)

	response, err := queryState("PATTERN?", address, pooled)
	if err != nil {
		return status.Blanked{}, err
	}

	var status status.Blanked
	resp := string(response)

	if strings.Contains(resp, "PATTERN:BLACK") {
		status.Blanked = true
	} else if strings.Contains(resp, "PATTERN:NONE") {
		status.Blanked = false
	}

	return status, nil
}
