package discord

import "time"

// EpochBeginning represents the number of milliseconds that
// have elapsed since the Discord Epoch (2015-01-01T00:00:00Z).
const EpochBeginning = 1420070400000

// DefaultTimeFormat used by Discord as default time format
const DefaultTimeFormat = time.RFC3339

const (
	TimestampMentionStyleAbbrShortTime     = "t"
	TimestampMentionStyleAbbrLongTime      = "T"
	TimestampMentionStyleAbbrShortDate     = "d"
	TimestampMentionStyleAbbrLongDate      = "D"
	TimestampMentionStyleAbbrShortDateTime = "f"
	TimestampMentionStyleAbbrLongDateTime  = "F"
	TimestampMentionStyleAbbrRelativeTime  = "R"
)

const (
	GuildNavigationTypeCustomize = "customize"
	GuildNavigationTypeBrowse    = "browse"
	GuildNavigationTypeGuide     = "guide"
)
