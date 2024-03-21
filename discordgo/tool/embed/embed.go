package embed

import (
	"strconv"
	"time"
	"unicode/utf8"

	"github.com/barbiequeue/discokit/discordgo/tool"
	"github.com/bwmarrin/discordgo"
)

func InlineEmbed(message string, color int) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title: message,
		Color: color,
	}
}

const (
	limitTotal       = 6000
	limitTitle       = 256
	limitDescription = 4096
	limitFieldName   = 256
	limitFieldValue  = 1024
	limitFooterText  = 2048
	limitAuthorName  = 256
	limitFieldsCount = 25
)

// Builder using for build embed and check embed field limits
//
// ------------------------------
//
// # You can use Example 1 if you sure about data length which will use for builder
//
// Example 1:
//
//	eb := NewEmbedBuilder().
//		Image("https://example.com/img/fh982h93.jpg", "", 0, 0).
//		FooterSimple("Look at this awesome picture!")
//
//	if _, err := s.ChannelMessageSendEmbed("92963946923497253235", eb.Build()); err != nil {
//		fmt.Println("Failed to send message with embed:", err)
//	}
//
// ------------------------------
//
// # If you not sure about data length, which you use for building embed, you can use builder like in Example 2
//
// Example 2:
//
//	eb := NewEmbedBuilder().
//		Image("https://example.com/img/fh982h93.jpg", "", 0, 0).
//		FooterSimple("Look at this awesome picture!")
//	if !eb.InLimits() {
//		fmt.Println("Failed to build discord embed")
//		for _, exceed := range eb.LimitExceeds() {
//			fmt.Println("Embed exceeded limit:", exceed)
//		}
//	}
//
//	em := eb.Build()
//
//	if _, err := s.ChannelMessageSendEmbed("92963946923497253235", em); err != nil {
//		fmt.Println("Failed to send message with embed:", err)
//	}
type Builder struct {
	limitsExceeds []string
	embed         *discordgo.MessageEmbed
}

func (b *Builder) describeExceed(field string, number ...int) {
	s := field
	if len(number) > 0 {
		s += ":" + strconv.Itoa(number[0])
	}
	b.limitsExceeds = append(b.limitsExceeds, s)
}

const (
	TypeRich    = "rich"
	TypeImage   = "image"
	TypeVideo   = "video"
	TypeGIFV    = "gifv"
	TypeArticle = "article"
	TypeLink    = "link"
)

// NewEmbedBuilder creates new embed builder with underlying embed and limits
// Default underlying embed has default type `rich` and color `#ffffff`
func NewEmbedBuilder() *Builder {
	return &Builder{
		embed: &discordgo.MessageEmbed{
			Type:  TypeRich,
			Color: tool.ColorMaterialWhite,
		},
	}
}

// Reset renew the underlying embed for avoid cases when you don't need to create new builder
func (b *Builder) Reset() *Builder {
	b.embed = &discordgo.MessageEmbed{}
	return b
}

func (b *Builder) Type(embedType string) *Builder {
	b.embed.Type = discordgo.EmbedType(embedType)
	return b
}

func (b *Builder) Title(title string) *Builder {
	b.embed.Title = title
	return b
}

func (b *Builder) Description(desc string) *Builder {
	b.embed.Description = desc
	return b
}

func (b *Builder) URL(url string) *Builder {
	b.embed.URL = url
	return b
}

func (b *Builder) Timestamp(t *time.Time) *Builder {
	b.embed.Timestamp = t.Format(time.RFC3339)
	return b
}

func (b *Builder) Color(color int) *Builder {
	b.embed.Color = color
	return b
}

func (b *Builder) ImageSimple(url string) *Builder {
	b.embed.Image = &discordgo.MessageEmbedImage{URL: url}
	return b
}

func (b *Builder) Image(url, proxyURL string, height, width int) *Builder {
	b.ImageSimple(url)
	b.embed.Image.ProxyURL = proxyURL
	b.embed.Image.Height = height
	b.embed.Image.Width = width

	return b
}

func (b *Builder) ThumbnailSimple(url string) *Builder {
	b.embed.Thumbnail = &discordgo.MessageEmbedThumbnail{URL: url}
	return b
}

func (b *Builder) Thumbnail(url, proxyURL string, height, width int) *Builder {
	b.ThumbnailSimple(url)
	b.embed.Thumbnail.ProxyURL = proxyURL
	b.embed.Thumbnail.Height = height
	b.embed.Thumbnail.Width = width

	return b
}

func (b *Builder) VideoSimple(url string) *Builder {
	b.embed.Video = &discordgo.MessageEmbedVideo{URL: url}
	return b
}

func (b *Builder) Video(url, proxyURL string, height, width int) *Builder {
	b.VideoSimple(url)
	b.embed.Video.Height = height
	b.embed.Video.Width = width

	return b
}

func (b *Builder) Provider(name, url string) *Builder {
	b.embed.Provider = &discordgo.MessageEmbedProvider{
		Name: name,
		URL:  url,
	}

	return b
}

func (b *Builder) AuthorSimple(name string) *Builder {
	b.embed.Author = &discordgo.MessageEmbedAuthor{Name: name}
	return b
}

func (b *Builder) Author(name, url, iconURL, proxyIconURL string) *Builder {
	b.AuthorSimple(name)
	b.embed.Author.URL = url
	b.embed.Author.IconURL = iconURL
	b.embed.Author.ProxyIconURL = proxyIconURL

	return b
}

func (b *Builder) FooterSimple(text string) *Builder {
	b.embed.Footer = &discordgo.MessageEmbedFooter{Text: text}
	return b
}

func (b *Builder) Footer(text, iconURL, iconProxyURL string) *Builder {
	b.FooterSimple(text)
	b.embed.Footer.IconURL = iconURL
	b.embed.Footer.ProxyIconURL = iconProxyURL

	return b
}

func (b *Builder) AddField(name, value string, inline bool) *Builder {
	b.embed.Fields = append(b.embed.Fields, &discordgo.MessageEmbedField{
		Name:   name,
		Value:  value,
		Inline: inline,
	})

	return b
}

func (b *Builder) InLimits() bool {
	ok := true
	totalLen := 0

	titleLen := utf8.RuneCountInString(b.embed.Title)
	totalLen += titleLen
	if titleLen > limitTitle {
		ok = false
		b.describeExceed("title")
	}
	descriptionLen := utf8.RuneCountInString(b.embed.Description)
	totalLen += descriptionLen
	if descriptionLen > limitDescription {
		ok = false
		b.describeExceed("description")
	}
	if b.embed.Footer != nil {
		footerTextLen := utf8.RuneCountInString(b.embed.Footer.Text)
		totalLen += footerTextLen
		if footerTextLen > limitFooterText {
			ok = false
			b.describeExceed("footer.text")
		}
	}
	if b.embed.Author != nil {
		authorNameLen := utf8.RuneCountInString(b.embed.Author.Name)
		totalLen += authorNameLen
		if authorNameLen > limitAuthorName {
			ok = false
			b.describeExceed("author.name")
		}
	}
	if len(b.embed.Fields) > limitFieldsCount {
		ok = false
		b.describeExceed("fieldsCount")
	}
	for i, field := range b.embed.Fields {
		nameLen := utf8.RuneCountInString(field.Name)
		totalLen += nameLen
		if nameLen > limitFieldName {
			ok = false
			b.describeExceed("field.name", i+1)
		}
		valueLen := utf8.RuneCountInString(field.Value)
		totalLen += valueLen
		if valueLen > limitFieldValue {
			ok = false
			b.describeExceed("field.value", i+1)
		}
	}

	if totalLen > limitTotal {
		ok = false
		b.describeExceed("total")
	}

	return ok
}

func (b *Builder) LimitExceeds() []string {
	return b.limitsExceeds
}

func (b *Builder) Build() *discordgo.MessageEmbed { return b.embed }
