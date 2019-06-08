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
	Category []Entity    `json:"intent"`
	Subcategory []Entity `json:"problem_kind"`
	Location []Entity `json:"interior_location"`
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

	return &resp, nil
}
