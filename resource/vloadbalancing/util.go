package vloadbalancing

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

var (
	ErrorCodeLoadBalancerNotReady = "LoadBalancerNotReady"
)

func checkErrorResponse(httpResponse *http.Response) bool {
	if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		return true
	}
	return false
}

func getResponseBody(httpResponse *http.Response) string {
	localVarBody, _ := io.ReadAll(httpResponse.Body)
	responseMessage := string(localVarBody)
	if httpResponse.StatusCode == http.StatusForbidden {
		responseMessage = "You don't have permission to do this action"
	}
	return fmt.Sprint("StatusCode: ", httpResponse.StatusCode, ", ", responseMessage)
}

type ResponseError struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	ErrorCode  string `json:"errorCode"`
}

func (e *ResponseError) Error() string {
	return e.Message
}

func (e *ResponseError) ErrorMessage() string {
	return fmt.Sprint("StatusCode: ", e.StatusCode, ": ", e.ErrorCode, ": ", e.Message)
}

func parseErrorResponse(httpResponse *http.Response) error {
	localVarBody, _ := io.ReadAll(httpResponse.Body)
	var responseError ResponseError
	_ = json.Unmarshal(localVarBody, &responseError)

	if responseError.ErrorCode == "" && httpResponse.StatusCode != http.StatusForbidden {
		return &ResponseError{
			StatusCode: httpResponse.StatusCode,
			Message:    string(localVarBody),
			ErrorCode:  "Unknown",
		}
	}

	if httpResponse.StatusCode == http.StatusForbidden {
		log.Printf("You don't have permission to do this action %s\n", httpResponse.Request.URL.Path)
		return &ResponseError{
			StatusCode: httpResponse.StatusCode,
			Message:    "You don't have permission to do this action",
			ErrorCode:  "Forbidden",
		}
	}

	return &responseError
}

func errorCodeEquals(err error, code string) bool {
	if responseError, ok := err.(*ResponseError); ok {
		return responseError.ErrorCode == code
	}
	return false
}
