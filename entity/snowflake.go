package entity

import (
	"errors"
	"strconv"
	"time"
)

const DiscordEpochShift = 1420070400000
const defaultDiscordTimeFormat = time.RFC3339

var ErrSnowflakeFormat = errors.New("discord snowflake format error")

type Snowflake uint64

func ParseSnowflake(s string) (Snowflake, error) {
	v, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, ErrSnowflakeFormat
	}
	sf := Snowflake(v)
	if sf.Time().UnixMilli() <= DiscordEpochShift {
		return 0, ErrIsNotMention
	}
	return sf, nil
}

func (sf Snowflake) Time() time.Time {
	ts := (sf >> 22) + DiscordEpochShift
	return time.UnixMilli(int64(ts))
}

func (sf Snowflake) FormatTime() string {
	return sf.Time().Format(defaultDiscordTimeFormat)
}

func (sf Snowflake) DescribeTime() string {
	if sf == 0 {
		return "id value is 0 (discord epoch start)"
	}
	return sf.FormatTime()
}

func (sf Snowflake) WorkerID() uint8 {
	return uint8((sf & 0x3E0000) >> 17)
}

func (sf Snowflake) ProcessID() uint8 {
	return uint8((sf & 0x1F000) >> 12)
}

func (sf Snowflake) Increment() uint16 {
	return uint16(sf & 0xFFF)
}
