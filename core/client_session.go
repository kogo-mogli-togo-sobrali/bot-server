package core

import (
	"HomeServices/nlp"
	"fmt"
	"github.com/spf13/viper"
)

// ClientSession - one client session, session must provide
// category, [problem], address and phone number
type ClientSession struct {
	IsCompleted bool
	Category string
	Name string
	Problem string
	Address string
	PhoneNumber string

	processor *nlp.Processor


}

func NewClientSession(p *nlp.Processor) *ClientSession {
	return &ClientSession{IsCompleted: false, processor: p}
}

func (c *ClientSession) handleRequest(message Message) Answer {
	res, err := c.processor.ParseRequest(message.Text)
	if err != nil {
		return Answer{
			Text: viper.GetString("strings.chooseCategory"),
			Options: viper.GetStringSlice("categories"),
		}
	}

	if len(res.Categories) != 0 && c.Category == "" {
		c.Category = findMostProbably(res.Categories).Name
	}

	if len(res.Problems) != 0 && c.Problem == "" {
		c.Problem = findMostProbably(res.Problems).Name
	}

	if len(res.Addresses) != 0 && c.Address == "" {
		c.Address = findMostProbably(res.Addresses).Name
	}

	if len(res.Names) != 0 && c.Name == "" {
		c.Name = findMostProbably(res.Names).Name
	}

	if len(res.PhoneNumbers) != 0 && c.PhoneNumber == "" {
		c.PhoneNumber = findMostProbably(res.PhoneNumbers).Name
	}

	if c.Category == "" {
		return Answer{
			Text: viper.GetString("strings.chooseCategory"),
			Options: viper.GetStringSlice("common.categories"),
		}
	}

	if c.Address == "" {
		return Answer{
			Text: viper.GetString("strings.enterAddress"),
		}
	}

	if c.Name == "" {
		return Answer{
			Text: viper.GetString("strings.enterName"),
		}
	}

	if c.PhoneNumber == "" {
		return Answer{
			Text: viper.GetString("strings.enterPhoneNumber"),
		}
	}

	categoriesGiving := map[string]string {
		"Газ": "газу",
		"Отопление": "отоплению",
		"Электричество": "электричеству",
		"Канализация": "канализации",
	}

	c.IsCompleted = true
	return Answer{
		Text: fmt.Sprintf(viper.GetString("strings.finalMessage"), c.Name, categoriesGiving[c.Category], c.PhoneNumber),
	}
}

func (c *ClientSession) getData() string {
	return ""
}
