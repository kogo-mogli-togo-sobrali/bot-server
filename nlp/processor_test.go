package nlp

import (
	"HomeServices/config"
	"github.com/magiconair/properties/assert"
	"github.com/spf13/viper"
	"log"
	"os"
	"testing"
)

func TestProcessor_ParseRequest(t *testing.T) {
	processor := NewProcessor()

	getResponse := func(req string) *Response {
		resp, err := processor.ParseRequest(req)
		if err != nil {
			t.Errorf("ParseRequest with %v failed: %v", req, err)
		}

		return resp
	}

	resp := getResponse("У меня прорвало трубу")
	assert.Equal(t, len(resp.Locations), 0)
	assert.Equal(t, len(resp.Problems), 1)
	assert.Equal(t, resp.Problems[0].Name, "прорвало трубу")
	assert.Equal(t, len(resp.Categories), 1)
	assert.Equal(t, resp.Categories[0].Name, "Отопление")

	resp = getResponse("88005553535")
	assert.Equal(t, len(resp.PhoneNumbers), 1)

	resp = getResponse("Добрый вечер, меня зовут Александр Стешенко и у меня нет света")
	assert.Equal(t, len(resp.Locations), 0)
	assert.Equal(t, len(resp.Problems), 1)
	assert.Equal(t, resp.Problems[0].Name, "нет света")
	assert.Equal(t, len(resp.Categories), 1)
	assert.Equal(t, resp.Categories[0].Name, "Электричество")

	resp = getResponse("Газ")
	assert.Equal(t, len(resp.Categories), 1)
	assert.Equal(t, resp.Categories[0].Name, "Газ")
}

func TestMain(m *testing.M) {
	viper.AddConfigPath("..")

	err := config.InitConfiguration()
	if err != nil {
		log.Fatalf("Failed to init configuration: %v", err)
	}

	os.Exit(m.Run())
}
