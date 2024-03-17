// Package entity provides structures and functions to work with Discord entities.
package entity

import (
	"errors"
	"strconv"
	"time"
)

// DiscordEpochShift represents the number of milliseconds that
// have elapsed since the Discord Epoch (2015-01-01T00:00:00Z).
const DiscordEpochShift = 1420070400000
const defaultDiscordTimeFormat = time.RFC3339

// ErrSnowflakeFormat signals an error when parsing a string
// into a Snowflake fails due to the string not adhering to the expected format.
var ErrSnowflakeFormat = errors.New("discord snowflake format error")

// Snowflake represents a Discord snowflake ID,
// a unique identifier used by Discord for various entities.
type Snowflake uint64

// ParseSnowflake converts a string representation of a snowflake ID to a Snowflake type.
// It returns an error if the string is not a valid numeric value or represents a time before the Discord Epoch.
func ParseSnowflake(s string) (Snowflake, error) {
	v, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, ErrSnowflakeFormat
	}
	sf := Snowflake(v)
	if sf.Time().UnixMilli() <= DiscordEpochShift {
		return 0, ErrSnowflakeFormat
	}
	return sf, nil
}

// Time converts a Snowflake to a time.Time value, representing the time when the snowflake was created.
func (sf Snowflake) Time() time.Time {
	ts := (sf >> 22) + DiscordEpochShift
	return time.UnixMilli(int64(ts))
}

// FormatTime returns a string representation of the time when the snowflake was created,
// formatted according to the default Discord time format.
func (sf Snowflake) FormatTime() string {
	return sf.Time().Format(defaultDiscordTimeFormat)
}

// DescribeTime provides a human-readable description of the snowflake's creation time.
// If the snowflake value is 0, it returns a message indicating the start of the Discord Epoch.
func (sf Snowflake) DescribeTime() string {
	if sf == 0 {
		return "id value is 0 (discord epoch start)"
	}
	return sf.FormatTime()
}

// WorkerID extracts and returns the internal worker ID component of the Snowflake.
func (sf Snowflake) WorkerID() uint8 {
	return uint8((sf & 0x3E0000) >> 17)
}

// ProcessID extracts and returns the internal process ID component of the Snowflake.
func (sf Snowflake) ProcessID() uint8 {
	return uint8((sf & 0x1F000) >> 12)
}

// Increment extracts and returns the internal increment component of the Snowflake,
// which is part of its unique identifier properties.
func (sf Snowflake) Increment() uint16 {
	return uint16(sf & 0xFFF)
}
