package core

import (
	"encoding/json"
	"io"
	"net/http"
	"reflect"
	"strings"

	"github.com/hetiansu5/urlquery"
	"github.com/yosida95/uritemplate/v3"
)

// ToQueryMap converts a struct to query parameter map
func ToQueryMap(source interface{}) (map[string]string, error) {
	if source == nil {
		return nil, ErrNilSource
	}

	queryBytes, err := urlquery.Marshal(source)
	if err != nil {
		return nil, err
	}

	var queryMap map[string]string
	err = urlquery.Unmarshal(queryBytes, &queryMap)
	if err != nil {
		return nil, err
	}

	return queryMap, nil
}

// normalizeVarNames normalizes variable names for URI template expansion.
func normalizeVarNames(varNames []string) map[string]string {
	normalizedNames := make(map[string]string)
	for _, varName := range varNames {
		normalizedNames[strings.ToLower(varName)] = varName
	}
	return normalizedNames
}

func addParametersWithOriginalNames(params map[string]string, normalizedNames map[string]string, values uritemplate.Values) uritemplate.Values {

	if values == nil {
		values = uritemplate.Values{}
	}

	for key, value := range params {
		values.Set(getKeyWithOriginalName(key, normalizedNames), uritemplate.String(value))
	}
	return values
}

func getKeyWithOriginalName(key string, normalizedNames map[string]string) string {
	originalName, exists := normalizedNames[key]
	if exists {
		return originalName
	}
	return key
}

func IsPointer(value interface{}) bool {

	if value == nil {
		return false
	}

	valueKind := reflect.ValueOf(value).Kind()

	return valueKind == reflect.Ptr
}

func FromJson[T any](response *http.Response, v *T) error {
	if response == nil {
		return ErrNilResponse
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if err := json.Unmarshal(body, v); err != nil {
		return err
	}

	return nil
}

// ParseResponse parses the HTTP Response to the provided type
func ParseResponse[T Response](response *http.Response, value *T) error {

	err := FromJson(response, &value)
	if err != nil {
		return err
	}

	(*value).ParseHeaders(response.Header)

	return nil
}

// Deprecated: deprecated in version {version}. Please use SendGet2.
func sendGet[T Response](requestBuilder *RequestBuilder, params interface{}, errorMapping ErrorMapping, value *T) error {

	requestInfo, err := requestBuilder.ToGetRequestInformation(params)
	if err != nil {
		return err
	}

	response, err := requestBuilder.Client.Send(requestInfo, errorMapping)
	if err != nil {
		return err
	}

	return ParseResponse(response, value)
}

func SendGet2(requestBuilder *RequestBuilder, config *RequestConfiguration) error {

	requestInfo, err := requestBuilder.ToGetRequestInformation2(config)
	if err != nil {
		return err
	}

	response, err := requestBuilder.Client.Send(requestInfo, config.ErrorMapping)
	if err != nil {
		return err
	}

	return ParseResponse(response, &config.Response)
}

// Deprecated: deprecated in version {version}. Please use SendPost2.
func sendPost[T Response](requestBuilder *RequestBuilder, data interface{}, params interface{}, errorMapping ErrorMapping, value *T) error {

	requestInfo, err := requestBuilder.ToPostRequestInformation2(data, params)
	if err != nil {
		return err
	}

	response, err := requestBuilder.Client.Send(requestInfo, errorMapping)
	if err != nil {
		return err
	}

	return ParseResponse(response, value)
}

func SendPost2(requestBuilder *RequestBuilder, config *RequestConfiguration) error {
	requestInfo, err := requestBuilder.ToPostRequestInformation3(config)
	if err != nil {
		return err
	}

	response, err := requestBuilder.Client.Send(requestInfo, config.ErrorMapping)
	if err != nil {
		return err
	}

	return ParseResponse(response, &config.Response)
}

// Deprecated: deprecated in version {version}. Please use sendDelete2.
func sendDelete(requestBuilder *RequestBuilder, params interface{}, errorMapping ErrorMapping) error {
	requestInfo, err := requestBuilder.ToDeleteRequestInformation(params)
	if err != nil {
		return err
	}

	_, err = requestBuilder.Client.Send(requestInfo, errorMapping)
	if err != nil {
		return err
	}

	return nil
}

func sendDelete2(requestBuilder *RequestBuilder, config *RequestConfiguration) error {
	requestInfo, err := requestBuilder.ToDeleteRequestInformation2(config)
	if err != nil {
		return err
	}

	_, err = requestBuilder.Client.Send(requestInfo, config.ErrorMapping)
	if err != nil {
		return err
	}

	return nil
}

// Deprecated: deprecated in version {version}. Please use sendPut2.
func sendPut[T Response](requestBuilder *RequestBuilder, data map[string]string, params interface{}, errorMapping ErrorMapping, value *T) error {
	requestInfo, err := requestBuilder.ToPutRequestInformation(data, params)
	if err != nil {
		return err
	}

	response, err := requestBuilder.Client.Send(requestInfo, errorMapping)
	if err != nil {
		return err
	}

	return ParseResponse(response, value)
}

func sendPut2(requestBuilder *RequestBuilder, config *RequestConfiguration) error {
	requestInfo, err := requestBuilder.ToPutRequestInformation2(config)
	if err != nil {
		return err
	}

	response, err := requestBuilder.Client.Send(requestInfo, config.ErrorMapping)
	if err != nil {
		return err
	}

	return ParseResponse(response, &config.Response)
}
