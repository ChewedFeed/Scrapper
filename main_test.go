package Scrapper_test

import (
	"testing"
	"github.com/chewedfeed/Scrapper"
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		request events.APIGatewayProxyRequest
		expect string
		err error
	}{
		{
			request: events.APIGatewayProxyRequest{Body: "Tester"},
			expect: "Hello Tester",
			err: nil,
		},
		{
			request: events.APIGatewayProxyRequest{Body: ""},
			expect: "",
			err: Scrapper.ErrNameNotProvided,
		},
	}

	for _, test := range tests {
		response, err := Scrapper.Handler(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.expect, response.Body)
	}
}