package battery

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
)

type Status struct {
	ChargePercent int
}

var acpiOutputRegex = regexp.MustCompile("([0-9]+)%")

func GetAcpiOutput() (string, error) {
	data, err := exec.Command("/usr/bin/acpi").CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func ParseAcpiOutput(acpiOutput string) (Status, error) {
	matches := acpiOutputRegex.FindStringSubmatch(acpiOutput)
	if len(matches) < 2 {
		return Status{}, fmt.Errorf("failed to parse acpi output: %q", acpiOutput)
	}
	charge, err := strconv.Atoi(matches[1])
	if err != nil {
		return Status{}, fmt.Errorf("failed to parse charge percentage: %q", matches[1])
	}
	return Status{ChargePercent: charge}, nil
}
