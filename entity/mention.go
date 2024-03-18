// Package entity provides structures and functions to work with Discord entities.
package entity

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// ErrIsNotMention signifies an error when a string does not conform to a valid mention format.
var ErrIsNotMention = errors.New("is not a valid mention")

var prefixedMentionRegexp = regexp.MustCompile(`(?m)^(@|@!|#|@&)(\d+)$`)

const (
	MentionTypeRole = iota + 1
	MentionTypeUser
	MentionTypeChannel
	MentionTypeSlashCommand
	MentionTypeTimestamp
	MentionTypeTimestampStyled
	MentionTypeGuildNavigation
	MentionTypeCustomEmoji
	MentionTypeCustomEmojiAnimated
)

// MentionType defines the various types of mentions that can be recognized in Discord.
type MentionType int

var mentionTypeNames = map[MentionType]string{
	MentionTypeRole:                "role",
	MentionTypeUser:                "user",
	MentionTypeChannel:             "channel",
	MentionTypeSlashCommand:        "slash_command",
	MentionTypeTimestamp:           "timestamp",
	MentionTypeTimestampStyled:     "timestamp_styled",
	MentionTypeGuildNavigation:     "guild_navigation",
	MentionTypeCustomEmoji:         "emoji",
	MentionTypeCustomEmojiAnimated: "emoji_animate",
}

// String returns the string representation of the MentionType.
func (t MentionType) String() string {
	n, ok := mentionTypeNames[t]
	if !ok {
		return "unknown_mention_type"
	}
	return n
}

var mentionTypeMapping = map[string]MentionType{
	"@&": MentionTypeRole,
	"@":  MentionTypeUser,
	"@!": MentionTypeUser,
	"#":  MentionTypeChannel,
}

// DiscordMention represents a mention within Discord, including its type and ID.
type DiscordMention struct {
	Type MentionType
	ID   Snowflake
	Meta map[string]string
}

func ParseMention(s string) (*DiscordMention, error) {
	return detectMention(s)
}

func detectMention(s string) (m *DiscordMention, err error) {
	m = &DiscordMention{Meta: map[string]string{}}

	s = strings.TrimSpace(s)
	if !strings.HasPrefix(s, "<") || !strings.HasSuffix(s, ">") {
		return nil, ErrIsNotMention
	}
	s = strings.Trim(s, "<>")

	matches := prefixedMentionRegexp.FindAllStringSubmatch(s, -1)

	if len(matches) != 0 {
		match := matches[0]

		detectedType := match[1]
		detectedId := match[2]

		t, ok := mentionTypeMapping[detectedType]
		if ok {
			id, parseErr := ParseSnowflake(detectedId)
			if parseErr == nil {
				m.ID = id
				m.Type = t
				return
			}
		}
	}

	parts := strings.Split(s, ":")
	partsCount := len(parts)

	name := parts[0]

	if partsCount == 2 {
		if strings.HasPrefix(name, "/") {
			id, parseErr := ParseSnowflake(parts[1])
			if parseErr == nil {
				m.ID = id
				m.Type = MentionTypeSlashCommand
				m.Meta["command"] = strings.TrimLeft(name, "/")
				return
			}
		}
		if name == "t" {
			m.ID = 0
			m.Type = MentionTypeTimestamp
			m.Meta["timestamp"] = parts[1]
			return
		}
		if name == "id" {
			m.ID = 0
			m.Type = MentionTypeGuildNavigation
			m.Meta["guild_navigation_type"] = parts[1]
			return
		}
	}

	if partsCount == 3 {
		switch name {
		case "":
			id, parseErr := ParseSnowflake(parts[2])
			if parseErr == nil {
				m.ID = id
				m.Type = MentionTypeCustomEmoji
				m.Meta["emoji_name"] = parts[1]
				return
			}
		case "a":
			id, parseErr := ParseSnowflake(parts[2])
			if parseErr == nil {
				m.ID = id
				m.Type = MentionTypeCustomEmojiAnimated
				m.Meta["emoji_name"] = parts[1]
				return
			}
		case "t":
			m.ID = 0
			m.Type = MentionTypeTimestampStyled
			m.Meta["timestamp"] = parts[1]
			m.Meta["style"] = parts[2]
			return
		}
	}

	return nil, ErrIsNotMention
}

func (m DiscordMention) DiscordString() string {
	failedContent := "unknown-mention-type"

	if m.Meta == nil {
		m.Meta = make(map[string]string)
	}

	s := strings.Builder{}
	s.WriteString("<")
	switch m.Type {
	case MentionTypeRole:
		s.WriteString("@&" + m.ID.String())
	case MentionTypeUser:
		s.WriteString("@" + m.ID.String())
	case MentionTypeChannel:
		s.WriteString("#" + m.ID.String())
	case MentionTypeSlashCommand:
		cmdName, ok := m.Meta["command"]
		if !ok {
			s.WriteString(failedContent)
			break
		}
		s.WriteString("/" + cmdName + ":" + m.ID.String())
	case MentionTypeTimestamp:
		ts, ok := m.Meta["timestamp"]
		if !ok {
			s.WriteString(failedContent)
			break
		}
		s.WriteString("t:" + ts)
	case MentionTypeTimestampStyled:
		ts, tsOk := m.Meta["timestamp"]
		st, styleOk := m.Meta["style"]
		if !tsOk || !styleOk {
			s.WriteString(failedContent)
			break
		}
		s.WriteString("t:" + ts + ":" + st)
	case MentionTypeGuildNavigation:
		t, ok := m.Meta["guild_navigation_type"]
		if !ok {
			s.WriteString(failedContent)
			break
		}
		s.WriteString("id:" + t)
	case MentionTypeCustomEmoji:
		n, ok := m.Meta["emoji_name"]
		if !ok {
			s.WriteString(failedContent)
			break
		}
		s.WriteString(":" + m.ID.String() + ":" + n)
	case MentionTypeCustomEmojiAnimated:
		n, ok := m.Meta["emoji_name"]
		if !ok {
			s.WriteString(failedContent)
			break
		}
		s.WriteString("a:" + m.ID.String() + ":" + n)
	default:
		s.WriteString(failedContent)
	}
	s.WriteString(">")

	return s.String()
}

func UserMention(sf Snowflake) DiscordMention {
	return DiscordMention{
		Type: MentionTypeUser,
		ID:   sf,
	}
}

func RoleMention(sf Snowflake) DiscordMention {
	return DiscordMention{
		Type: MentionTypeRole,
		ID:   sf,
	}
}

func ChannelMention(sf Snowflake) DiscordMention {
	return DiscordMention{
		Type: MentionTypeChannel,
		ID:   sf,
	}
}

func EmojiMention(sf Snowflake, name string) DiscordMention {
	return DiscordMention{
		Type: MentionTypeCustomEmoji,
		ID:   sf,
		Meta: map[string]string{"emoji_name": name},
	}
}

func AnimatedEmojiMention(sf Snowflake, name string) DiscordMention {
	m := EmojiMention(sf, name)
	m.Type = MentionTypeCustomEmojiAnimated
	return m
}

func TimestampMention(t time.Time, style ...string) DiscordMention {
	m := DiscordMention{
		Type: MentionTypeTimestamp,
		Meta: map[string]string{"timestamp": strconv.Itoa(t.Second())},
	}
	if len(style) != 0 {
		m.Type = MentionTypeCustomEmojiAnimated
		m.Meta["style"] = style[0]
	}

	return m
}

func SlashCommandMention(sf Snowflake, command string) DiscordMention {
	return DiscordMention{
		Type: MentionTypeSlashCommand,
		ID:   sf,
		Meta: map[string]string{"command": command},
	}
}

func GuildNavigationMention(navigationType string) DiscordMention {
	return DiscordMention{
		Type: MentionTypeGuildNavigation,
		Meta: map[string]string{"guild_navigation_type": navigationType},
	}
}
