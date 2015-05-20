package generator

import (
	uuid "code.google.com/p/go-uuid/uuid"
	"fmt"
	"time"
)

type LogGenerator interface {
	Start(t time.Time) string
	Finish(t time.Time) string
	Fail(t time.Time) string
}

type TestLogGenerator struct {
	id uuid.UUID
}

func NewTestLogGenerator(id uuid.UUID) *TestLogGenerator {
	return &TestLogGenerator{
		id: id,
	}
}

func (g *TestLogGenerator) Start(t time.Time) string {
	return fmt.Sprintf("%s %s start\n", t, g.id)
}

func (g *TestLogGenerator) Finish(t time.Time) string {
	return fmt.Sprintf("%s %s finish\n", t, g.id)
}

func (g *TestLogGenerator) Fail(t time.Time) string {
	return fmt.Sprintf("%s %s fail\n", t, g.id)
}

func GenerateNormalDistributedStartLog(startTime time.Time, endTime time.Time, bucketSize time.Duration, distrib []float64) {

	for curr := startTime; curr.Before(endTime); curr = curr.Add(bucketSize) {
		startBucket := (curr.Sub(startTime)).Seconds()
		endBucket := startBucket + bucketSize.Seconds()
		count := GetBucketCount(distrib, startBucket, endBucket)

		output := ""
		for i := 0; i < count; i++ {
			//			go func() {
			gen := NewTestLogGenerator(uuid.NewRandom())
			output += gen.Start(curr)
			//			}()
		}

		fmt.Printf("%v", output)
	}
}
