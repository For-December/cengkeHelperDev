package models

import (
	"cengkeHelperDev/src/utils/logger"
	"github.com/bytedance/sonic"
)

type PostMeta struct {
	Type string `json:"type"`
	Text string `json:"text"`
	Url  string `json:"url"`
}

type PostMetaBuilder struct {
	PostMetas []PostMeta
}

func (receiver *PostMetaBuilder) BuildText(text string) *PostMetaBuilder {

	receiver.PostMetas = append(receiver.PostMetas, PostMeta{
		Type: "text",
		Text: text,
	})

	return receiver
}

func (receiver *PostMetaBuilder) BuildImage(image string) *PostMetaBuilder {

	receiver.PostMetas = append(receiver.PostMetas, PostMeta{
		Type: "image",
		Url:  image,
	})

	return receiver
}

func (receiver *PostMetaBuilder) BuildJson() string {
	marshalString, err := sonic.MarshalString(receiver.PostMetas)
	if err != nil {
		logger.Warning(receiver.PostMetas)
		logger.Warning(err)
		return "[]"
	}
	return marshalString
}
