package battery_test

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jm96441n/gotools/battery"
)

func TestParseAcpiOutput(t *testing.T) {
	t.Parallel()
	data, err := os.ReadFile("testdata/acpi.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := battery.Status{ChargePercent: 13}
	got, err := battery.ParseAcpiOutput(string(data))
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetAcpiOutput(t *testing.T) {
	t.Parallel()
	runIntegrationTest := os.Getenv("INTEGRATION")
	if runIntegrationTest == "" {
		t.Skip("Set the env variable 'INTEGRATION' to run this test")
	}

	text, err := battery.GetAcpiOutput()
	if err != nil {
		t.Fatal(err)
	}

	status, err := battery.ParseAcpiOutput(text)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Charge: %d%%", status.ChargePercent)
}
