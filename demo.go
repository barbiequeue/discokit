package main

import (
	"fmt"

	"github.com/barbiequeue/discokit/entity"
)

func main() {
	mentions := []string{
		"<@80351110224678912>",
		"<@!80351110224678912>",
		"<#103735883630395392>",
		"<@&165511591545143296>",
		"</airhorn:816437322781949972>",
		"<:mmLol:216154654256398347>",
		"<a:b1nzy:392938283556143104>",
		"<t:1618953630>",
		"<t:1618953630:d>",
		"<id:customize>",
		"<@22>",
	}

	for _, rawMention := range mentions {
		m, err := entity.ParseMention(rawMention)
		if err != nil {
			fmt.Println(err, rawMention)
			return
		}
		fmt.Println("Mention:", rawMention)
		printMentionInfo(m)
	}
}

func printMentionInfo(m *entity.DiscordMention) {
	fmt.Printf("{\n  ID: %d (date: %s)\n  Type: %s\n  Meta: {", m.ID, m.ID.FormatTime(), m.Type)
	for key, value := range m.Meta {
		fmt.Printf("\n    %q: %q", key, value)
	}
	fmt.Print("\n  }\n}\n")
}
