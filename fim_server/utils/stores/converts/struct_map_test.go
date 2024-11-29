package converts

import (
	"fim_server/utils/stores/logs"
	"testing"
)

type ContactInfo struct {
	PhoneNumber string `json:"phoneNumber"`
	Email       string `form:"email"`
}

func TestName(t *testing.T) {
	company := []ContactInfo{
		{PhoneNumber: "123-456-7890", Email: "contact@techcorp.com"},
		{PhoneNumber: "987-654-3210", Email: "support@techcorp.com"},
	}
	logs.Error(StructJsonMap[[]map[string]interface{}](company))
}
