package count_test

import (
	"bytes"
	"fmt"
	"io"
	"testing"
	"time"

	"github.com/jm96441n/gotools/hello/count"
)

func TestNextIncreasesTheValueOfTheCounter(t *testing.T) {
	c := count.NewCounter()
	want := c.CurVal
	got := c.Next()
	if want != got {
		t.Errorf("Want %d, got %d", want, got)
	}
	want += 1
	got = c.Next()
	if want != got {
		t.Errorf("Want %d, got %d", want, got)
	}
}

func TestRunPrintsAllValuesUntilStop(t *testing.T) {
	fakeTerminal := bytes.NewBuffer([]byte{})
	c := count.NewCounter()
	c.Output = io.Writer(fakeTerminal)
	c.StopAt = 10
	c.DelayTime = time.Duration(0)
	expectedOutput := ""
	for i := 0; i <= c.StopAt; i++ {
		expectedOutput += fmt.Sprintf("%d\n", i)
	}
	c.Run()
	got := fakeTerminal.String()
	if expectedOutput != got {
		t.Errorf("Expected %q, got %q", expectedOutput, got)
	}
}
