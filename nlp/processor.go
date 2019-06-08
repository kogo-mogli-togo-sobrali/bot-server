package nlp

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	witai "github.com/wit-ai/wit-go"
)

// Processor of user requests
// parse request and provide some structured data about issue

type Processor struct {
	ai *witai.Client
}

type Entity struct {
	Confidence float32 `json:"confidence"`
	Name string  `json:"value"`
}

type Response struct {
	Categories   []Entity `json:"intent"`
	Problems     []Entity `json:"problem_kind"`
	Locations    []Entity `json:"interior_location"`
	PhoneNumbers []Entity `json:"phone_number"`
	Addresses    []Entity `json:"address"`
	Names        []Entity `json:"name"`
}

func NewProcessor() *Processor {

	witClient := witai.NewClient(viper.GetString("nlp.token"))

	p := &Processor{witClient}

	return p
}

// Parse user language request
func (p *Processor) ParseRequest(request string) (*Response, error) {
	msg, err := p.ai.Parse(&witai.MessageRequest{Query: request})
	if err != nil {
		return nil, err
	}


	bytes, err := json.Marshal(msg.Entities)
	if err != nil {
		return nil, fmt.Errorf("failed to read response from the server: %v", bytes)
	}

	resp := Response{}
	err = json.Unmarshal(bytes, &resp)
	if err != nil {
		return nil, fmt.Errorf("failed to read response from the server: %v", bytes)
	}

	switch msg.Text {
	case "Газ":
		resp.Categories = append(resp.Categories, Entity{Name: "Газ", Confidence: 1})
	case "Отопление":
		resp.Categories = append(resp.Categories, Entity{Name: "Отопление", Confidence: 1})
	case "Электричество":
		resp.Categories = append(resp.Categories, Entity{Name: "Электричество", Confidence: 1})
	case "Канализация":
		resp.Categories = append(resp.Categories, Entity{Name: "Канализация", Confidence: 1})
	}

	return &resp, nil
}
