package test

import (
	"testing"
	"time"

	"github.com/barbiequeue/discokit/entity"
)

const testSnowflake = "1218325944931192904"

func initTestSnowflake() (entity.Snowflake, error) {
	return entity.ParseSnowflake(testSnowflake)
}

func TestSnowflake_Time(t *testing.T) {

	expectedTime := time.Date(2024, time.March, 15, 22, 32, 20, 673*int(time.Millisecond), time.UTC)

	snowflake, err := initTestSnowflake()
	if err != nil {
		t.Fatalf("failed to parse snowflake: %v", err)
	}

	actualTime := snowflake.Time()
	if !actualTime.Equal(expectedTime) {
		t.Errorf("expected time %v, got %v", expectedTime, actualTime.UTC())
	}
}

func TestSnowflake_WorkerID(t *testing.T) {
	expectedWorkerID := uint8(2)

	snowflake, err := initTestSnowflake()
	if err != nil {
		t.Fatalf("failed to parse snowflake: %v", err)
	}

	actualWorkerID := snowflake.WorkerID()
	if actualWorkerID != expectedWorkerID {
		t.Errorf("expected worker ID %d, got %d", expectedWorkerID, actualWorkerID)
	}
}

func TestSnowflake_FormatTime(t *testing.T) {
	expectedFormatString := "2024-03-16T02:32:20+04:00"

	snowflake, err := initTestSnowflake()
	if err != nil {
		t.Fatalf("failed to parse snowflake: %v", err)
	}

	actualFormatString := snowflake.FormatTime()
	if actualFormatString != expectedFormatString {
		t.Errorf("expected format string %s, got %s", expectedFormatString, actualFormatString)
	}
}

func TestSnowflake_ProcessID(t *testing.T) {
	expectedProcessID := uint8(1)

	snowflake, err := initTestSnowflake()
	if err != nil {
		t.Fatalf("failed to parse snowflake: %v", err)
	}

	actualProcessID := snowflake.ProcessID()
	if actualProcessID != expectedProcessID {
		t.Errorf("expected process ID %d, got %d", expectedProcessID, actualProcessID)
	}
}

func TestSnowflake_Increment(t *testing.T) {
	expectedIncrement := uint16(72)

	snowflake, err := initTestSnowflake()
	if err != nil {
		t.Fatalf("failed to parse snowflake: %v", err)
	}

	actualIncrement := snowflake.Increment()
	if actualIncrement != expectedIncrement {
		t.Errorf("expected increment %d, got %d", expectedIncrement, actualIncrement)
	}
}
