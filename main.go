package Scrapper

import (
	"github.com/gocolly/colly"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"errors"
)

var (
	ErrNameNotProvided = errors.New("No name provided")
)

// {
//  "feedId": "39689c2a-3633-511a-a402-f1e732d023c4",
//  "imageDetails": {},
//  "itemId": "198f7e0d-ce86-5a02-81c6-60d7ebad69b0",
//  "title": "New Snapdragon chips bring dual cameras to more mid-tier phones",
//  "url": "https://www.engadget.com/2018/06/26/qualcomm-snapdragon-632-439-429/"
// }

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Processing lamda %s\n", request.RequestContext)

	if len(request.Body) < 1 {
		return events.APIGatewayProxyResponse{}, ErrNameNotProvided
	}

	return events.APIGatewayProxyResponse{
		Body: "Hello " + request.Body,
		StatusCode: 200,
	}, nil
}

func main () {
	lambda.Start(Handler)
}