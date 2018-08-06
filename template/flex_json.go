package template

import (
	"encoding/json"

	"github.com/line/line-bot-sdk-go/linebot"
)

type BubbleContainer struct {
	Type   string
	Header BoxComponent
	Hero   ImageComponent
	Body   BoxComponent
	Footer BoxComponent
	Styles BubbleStyle
}

type FlexComponent interface {
	FlexComponent()
}

type BoxComponent struct {
	Type     string          `json:"type"`
	Layout   string          `json:"layout"`
	Contents []FlexComponent `json:"contents"`
	Flex     int             `json:"flex,omitempty"`
	Spacing  string          `json:"spacing,omitempty"`
	Margin   string          `json:"margin,omitempty"`
}

type ImageComponent struct {
	Type            string                 `json:"type"`
	URL             string                 `json:"url"`
	Flex            int                    `json:"flex,omitempty"`
	Margin          string                 `json:"margin,omitempty"`
	Align           string                 `json:"align,omitempty"`
	Gravity         string                 `json:"gravity,omitempty"`
	Size            string                 `json:"size,omitempty"`
	AspectRatio     string                 `json:"aspectRatio,omitempty"`
	AspectMode      string                 `json:"aspectMode,omitempty"`
	BackgroundColor string                 `json:"backgroundColor,omitempty"`
	Action          linebot.TemplateAction `json:"action,omitempty"`
}

type TextComponent struct {
	Type    string                 `json:"type"`
	Text    string                 `json:"text"`
	Flex    int                    `json:"flex,omitempty"`
	Margin  string                 `json:"margin,omitempty"`
	Size    string                 `json:"size,omitempty"`
	Align   string                 `json:"align,omitempty"`
	Gravity string                 `json:"gravity,omitempty"`
	Wrap    bool                   `json:"wrap,omitempty"`
	Weight  string                 `json:"weight,omitempty"`
	Color   string                 `json:"color,omitempty"`
	Action  linebot.TemplateAction `json:"action,omitempty"`
}

type ButtonComponent struct {
	Type    string                 `json:"type"`
	Action  linebot.TemplateAction `json:"action"`
	Flex    int                    `json:"flex,omitempty"`
	Margin  string                 `json:"margin,omitempty"`
	Height  string                 `json:"height,omitempty"`
	Style   string                 `json:"style,omitempty"`
	Color   string                 `json:"color,omitempty"`
	Gravity string                 `json:"gravity,omitempty"`
}

type IconComponent struct {
	Type        string `json:"type"`
	URL         string `json:"url"`
	Margin      string `json:"margin,omitempty"`
	Size        string `json:"size,omitempty"`
	AspectRatio string `json:"aspectRatio,omitempty"`
}

type BubbleStyle struct {
	Header BlockStyle `json:"header,omitempty"`
	Hero   BlockStyle `json:"hero,omitempty"`
	Body   BlockStyle `json:"body,omitempty"`
	Footer BlockStyle `json:"footer,omitempty"`
}

type BlockStyle struct {
	BackgroundColor string `json:"backgroundColor,omitempty"`
	Separator       bool   `json:"separator,omitempty"`
	SeparatorColor  string `json:"separatorColor,omitempty"`
}

// FlexComponent implements FlexComponent interface
func (*BoxComponent) FlexComponent() {}

// FlexComponent implements FlexComponent interface
func (*ImageComponent) FlexComponent() {}

// FlexComponent implements FlexComponent interface
func (*TextComponent) FlexComponent() {}

// FlexComponent implements FlexComponent interface
func (*IconComponent) FlexComponent() {}

func NewFlexMessage() *BubbleContainer {
	return &BubbleContainer{
		Type: "buble",
	}
}

func (bc *BubbleContainer) SetHeader(box BoxComponent) error {
	bc.Header = box
	return nil
}

func (bc *BubbleContainer) SetHero(img ImageComponent) error {
	bc.Hero = img
	return nil
}

func (bc *BubbleContainer) SetBody(box BoxComponent) error {
	bc.Body = box
	return nil
}

func (bc *BubbleContainer) SetFooter(box BoxComponent) error {
	bc.Footer = box
	return nil
}

func (bc *BubbleContainer) SetStyles(styles BubbleStyle) error {
	bc.Styles = styles
	return nil
}

func (bc *BubbleContainer) FlexTemplate() (*linebot.FlexMessage, error) {
	data, err := json.Marshal(bc)
	if err != nil {
		return nil, err
	}

	contents, err := linebot.UnmarshalFlexMessageJSON(data)
	if err != nil {
		return nil, err
	}

	return linebot.NewFlexMessage(
		"Flex Message",
		contents,
	), nil
}
