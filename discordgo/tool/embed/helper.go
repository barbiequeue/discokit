package embed

import (
	"github.com/barbiequeue/discokit/discordgo/tool"
	"github.com/bwmarrin/discordgo"
)

func Inline(message string, color ...int) *discordgo.MessageEmbed {
	c := tool.ColorMaterialWhite
	if len(color) != 0 {
		c = color[0]
	}

	return NewBuilder().Title(message).Color(c).Build()
}
