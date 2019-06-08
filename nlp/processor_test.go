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

	req := "У меня прорвало трубу"

	resp, err := processor.ParseRequest(req)
	if err != nil {
		t.Errorf("ParseRequest with %v failed: %v", req, err)
	}

	assert.Equal(t, len(resp.Location), 0)
	assert.Equal(t, len(resp.Subcategory), 1)
	assert.Equal(t, resp.Subcategory[0].Name, "прорвало трубу")
	assert.Equal(t, len(resp.Category), 1)
	assert.Equal(t, resp.Category[0].Name, "Отопление")
}

func TestMain(m *testing.M) {
	viper.AddConfigPath("..")

	err := config.InitConfiguration()
	if err != nil {
		log.Fatalf("Failed to init configuration: %v", err)
	}

	os.Exit(m.Run())
}
