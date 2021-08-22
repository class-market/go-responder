package response

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/class-market/go-responder/headers"
)

func TestSuccessFunction(t *testing.T) {
	data := struct {
		TextValue   string `json:"textValue"`
		NumberValue int    `json:"numberValue"`
	}{"test", 5}
	valueToCompare := `{"message":"error example","code":200,"data":{"textValue":"test","numberValue":5}}`

	res := Success("error example", data, nil, headers.New())
	body, err := res.GetBody()

	assert.EqualValues(t, res.StatusCode(), 200)
	assert.EqualValues(t, string(body), valueToCompare)
	assert.NotNil(t, body)
	assert.Nil(t, err)
}
