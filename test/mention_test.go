package test

import (
	"reflect"
	"testing"

	"github.com/barbiequeue/discokit/entity"
)

func TestParseMention(t *testing.T) {
	tests := []struct {
		input   string
		want    *entity.DiscordMention
		wantErr error
	}{
		{
			input: "<@80351110224678912>",
			want: &entity.DiscordMention{
				ID:   80351110224678912,
				Type: entity.MentionTypeUser,
				Meta: map[string]string{},
			},
			wantErr: entity.ErrIsNotMention,
		},
		{
			input: "<@!80351110224678912>",
			want: &entity.DiscordMention{
				ID:   80351110224678912,
				Type: entity.MentionTypeUser,
				Meta: map[string]string{},
			},
			wantErr: entity.ErrIsNotMention,
		},
		{
			input: "<#103735883630395392>",
			want: &entity.DiscordMention{
				ID:   103735883630395392,
				Type: entity.MentionTypeChannel,
				Meta: map[string]string{},
			},
			wantErr: entity.ErrIsNotMention,
		},
		{
			input: "<@&165511591545143296>",
			want: &entity.DiscordMention{
				ID:   165511591545143296,
				Type: entity.MentionTypeRole,
				Meta: map[string]string{},
			},
			wantErr: entity.ErrIsNotMention,
		},
		{
			input: "</airhorn:816437322781949972>",
			want: &entity.DiscordMention{
				ID:   816437322781949972,
				Type: entity.MentionTypeSlashCommand,
				Meta: map[string]string{"command": "airhorn"},
			},
			wantErr: entity.ErrIsNotMention,
		},
		{
			input: "<:mmLol:216154654256398347>",
			want: &entity.DiscordMention{
				ID:   216154654256398347,
				Type: entity.MentionTypeCustomEmoji,
				Meta: map[string]string{"emoji_name": "mmLol"},
			},
			wantErr: entity.ErrIsNotMention,
		},
		{
			input: "<a:b1nzy:392938283556143104>",
			want: &entity.DiscordMention{
				ID:   392938283556143104,
				Type: entity.MentionTypeCustomEmojiAnimated,
				Meta: map[string]string{"emoji_name": "b1nzy"},
			},
			wantErr: entity.ErrIsNotMention,
		},
		{
			input: "<t:1618953630>",
			want: &entity.DiscordMention{
				ID:   0,
				Type: entity.MentionTypeTimestamp,
				Meta: map[string]string{"timestamp": "1618953630"},
			},
			wantErr: entity.ErrIsNotMention,
		},
		{
			input: "<t:1618953630:d>",
			want: &entity.DiscordMention{
				ID:   0,
				Type: entity.MentionTypeTimestampStyled,
				Meta: map[string]string{
					"timestamp": "1618953630",
					"style":     "d",
				},
			},
			wantErr: entity.ErrIsNotMention,
		},
		{
			input: "<id:customize>",
			want: &entity.DiscordMention{
				ID:   0,
				Type: entity.MentionTypeGuildNavigation,
				Meta: map[string]string{"guild_navigation_type": "customize"},
			},
			wantErr: entity.ErrIsNotMention,
		},
		{
			input:   "<:>",
			want:    nil,
			wantErr: entity.ErrIsNotMention,
		},
		{
			input:   "<3204u0@>",
			want:    nil,
			wantErr: entity.ErrIsNotMention,
		},
		{
			input:   "<@22>",
			want:    nil,
			wantErr: entity.ErrIsNotMention,
		},
		{
			input:   "<:sdf:fsdf93>",
			want:    nil,
			wantErr: entity.ErrIsNotMention,
		},
		{
			input:   "</hello:8888991f>",
			want:    nil,
			wantErr: entity.ErrIsNotMention,
		},
	}

	for _, tt := range tests {
		got, err := entity.ParseMention(tt.input)
		if err != nil && err != tt.wantErr {
			t.Errorf("ParseMention() error = %v, wantErr %v", err, tt.wantErr)
			continue
		}

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("ParseMention() = %v, want %v", got, tt.want)
		}
	}
}
