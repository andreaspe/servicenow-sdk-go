package tableapi

import "github.com/michaeldcanady/servicenow-sdk-go/core"

type TableItemDeleteRequestConfiguration struct {
	Header          interface{}
	QueryParameters *TableItemRequestBuilderDeleteQueryParameters
	Data            interface{}
	ErrorMapping    core.ErrorMapping
	response        *TableItemResponse
}

func (rC *TableItemDeleteRequestConfiguration) toConfiguration() *core.RequestConfiguration {
	return &core.RequestConfiguration{
		Header:          rC.Header,
		QueryParameters: rC.QueryParameters,
		Data:            rC.Data,
		ErrorMapping:    rC.ErrorMapping,
		Response:        rC.response,
	}
}
