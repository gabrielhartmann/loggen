package generator

import (
	uuid "code.google.com/p/go-uuid/uuid"
	"fmt"
	"testing"
	"time"
)

func TestGenerateTaskStart(t *testing.T) {
	const longForm = "Jan 2, 2006 at 3:04pm (MST)"
	t0, _ := time.Parse(longForm, "Feb 3, 2013 at 7:54pm (PST)")

	id := uuid.NewRandom()
	gen := NewTestLogGenerator(id)

	startMsg := gen.Start(t0)
	expectedMsg := fmt.Sprintf("2013-02-03 19:54:00 -0800 PST %v start", id)

	if expectedMsg != startMsg {
		t.Errorf("Expected '%v', returned: '%v'", expectedMsg, startMsg)
	}
}

func TestGenerateNormalDistribution(t *testing.T) {
	mean := 0.0
	stdDev := 10.0
	count := 1000
	list := GetNormalDistribution(mean, stdDev, count)

	//PrintDistribution(list, mean, stdDev, 1.0)

	if len(list) != count {
		t.Errorf("Expected length %v, returned %v", count, len(list))
	}
}

func TestGenerateNormalDistributionLog(t *testing.T) {
	mean := 0.0
	stdDev := 30.0
	count := 1000
	list := GetNormalDistribution(mean, stdDev, count)

	const longForm = "Jan 2, 2006 at 3:04pm (MST)"
	t0, _ := time.Parse(longForm, "Feb 3, 2013 at 7:54pm (PST)")
	t1, _ := time.Parse(longForm, "Feb 3, 2013 at 7:55pm (PST)")

	GenerateNormalDistributedStartLog(t0, t1, time.Second, list)
}
