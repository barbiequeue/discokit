package embed

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func ExampleNewEmbedBuilder() {
	s := discordgo.Session{}

	eb := NewEmbedBuilder().
		Image("https://example.com/img/fh982h93.jpg", "", 0, 0).
		FooterSimple("Look at this awesome picture!")

	if _, err := s.ChannelMessageSendEmbed("92963946923497253235", eb.Build()); err != nil {
		fmt.Println("Failed to send message with embed:", err)
	}
}

func ExampleNewEmbedBuilderWithLimitsCheck() {
	s := discordgo.Session{}

	eb := NewEmbedBuilder().
		Image("https://example.com/img/fh982h93.jpg", "", 0, 0).
		FooterSimple("Look at this awesome picture!")

	if !eb.InLimits() {
		fmt.Println("Failed to build discord embed")
		for _, exceed := range eb.LimitExceeds() {
			fmt.Println("Embed exceeded limit:", exceed)
		}
	}

	if _, err := s.ChannelMessageSendEmbed("92963946923497253235", eb.Build()); err != nil {
		fmt.Println("Failed to send message with embed:", err)
	}
}
