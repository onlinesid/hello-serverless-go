package main

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

func process(httpRequest events.APIGatewayV2HTTPRequest) ([]byte, error) {

	return json.Marshal(map[string]interface{}{
		"message": "Go Serverless v1.0! Your function executed successfully!",
		"RawPath": httpRequest.RawPath,
		"RawQueryString": httpRequest.RawQueryString,
		"Headers": httpRequest.Headers,
	})
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, httpRequest events.APIGatewayV2HTTPRequest) (Response, error) {
	var buf bytes.Buffer

	body, err := process(httpRequest)

	if err != nil {
		return Response{StatusCode: 404}, err
	}

	json.HTMLEscape(&buf, body)

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type":           "application/json",
			"X-MyCompany-Func-Reply": "hello-handler",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
