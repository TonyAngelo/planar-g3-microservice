package helpers

import (
	"fmt"
	"strings"

	"github.com/byuoitav/common/log"
	"github.com/byuoitav/common/nerr"
	"github.com/byuoitav/common/status"
	"github.com/byuoitav/common/structs"
)

func SetInput(address, port string, pooled bool) *nerr.E {
	log.L.Debugf("Setting input on %s to %s", address, port)

	// validInput := false
	// for _, input := range validADCPInputs {
	// 	if strings.EqualFold(port, input) {
	// 		validInput = true
	// 		break
	// 	}
	// }

	// if !validInput {
	// 	return nerr.Create(fmt.Sprintf("error: %s is not a valid ADCP input.", port), "invalid port")
	// }

	command := fmt.Sprintf("PRESET.RECALL(%s)", port)
	return sendCommand(command, address, pooled)
}

func GetInputStatus(address string, pooled bool) (status.Input, *nerr.E) {
	log.L.Debugf("Querying input status of %s", address)

	response, err := queryState("PRESET.ACTIVE?", address, pooled)
	if err != nil {
		return status.Input{}, err.Add("Couldn't query input status")
	}

	status := status.Input{
		Input: strings.Trim(string(response), "PRESET.ACTIVE:"),
	}
	return status, nil
}

// HasActiveSignal checks to see if the projector has an active signal on the given port currently.
func HasActiveSignal(address string, port string, pooled bool) (structs.ActiveSignal, *nerr.E) {
	log.L.Debugf("Checking if %s has an active signal right now", address)

	var toReturn structs.ActiveSignal
	toReturn.Active = false

	input, err := GetInputStatus(address, false)
	if err != nil {
		return toReturn, err.Add("couldn't query the input status")
	}

	if strings.EqualFold(input.Input, port) {
		response, err := queryState("INPUT.PRESENT(VC1.IN1)?", address, pooled)
		if err != nil {
			return toReturn, err.Add("Couldn't get active signal")
		}

		//active := strings.Trim(string(response), "INPUT.PRESENT(VC1.IN1):")
		var active = string("Invalid")
		if strings.Contains(string(response), "TRUE") {
			active = string("True")
		} //else if strings.Contains(resp, "FALSE") {
			//active = false
		//}

		toReturn.Active = !strings.EqualFold(active, "Invalid")
	}

	return toReturn, nil
}
