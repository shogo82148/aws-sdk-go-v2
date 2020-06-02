// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package fms

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetAppsListInput struct {
	_ struct{} `type:"structure"`

	// Specifies whether the list to retrieve is a default list owned by AWS Firewall
	// Manager.
	DefaultList *bool `type:"boolean"`

	// The ID of the AWS Firewall Manager applications list that you want the details
	// for.
	//
	// ListId is a required field
	ListId *string `min:"36" type:"string" required:"true"`
}

// String returns the string representation
func (s GetAppsListInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetAppsListInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetAppsListInput"}

	if s.ListId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ListId"))
	}
	if s.ListId != nil && len(*s.ListId) < 36 {
		invalidParams.Add(aws.NewErrParamMinLen("ListId", 36))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetAppsListOutput struct {
	_ struct{} `type:"structure"`

	// Information about the specified AWS Firewall Manager applications list.
	AppsList *AppsListData `type:"structure"`

	// The Amazon Resource Name (ARN) of the applications list.
	AppsListArn *string `min:"1" type:"string"`
}

// String returns the string representation
func (s GetAppsListOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetAppsList = "GetAppsList"

// GetAppsListRequest returns a request value for making API operation for
// Firewall Management Service.
//
// Returns information about the specified AWS Firewall Manager applications
// list.
//
//    // Example sending a request using GetAppsListRequest.
//    req := client.GetAppsListRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/fms-2018-01-01/GetAppsList
func (c *Client) GetAppsListRequest(input *GetAppsListInput) GetAppsListRequest {
	op := &aws.Operation{
		Name:       opGetAppsList,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetAppsListInput{}
	}

	req := c.newRequest(op, input, &GetAppsListOutput{})

	return GetAppsListRequest{Request: req, Input: input, Copy: c.GetAppsListRequest}
}

// GetAppsListRequest is the request type for the
// GetAppsList API operation.
type GetAppsListRequest struct {
	*aws.Request
	Input *GetAppsListInput
	Copy  func(*GetAppsListInput) GetAppsListRequest
}

// Send marshals and sends the GetAppsList API request.
func (r GetAppsListRequest) Send(ctx context.Context) (*GetAppsListResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetAppsListResponse{
		GetAppsListOutput: r.Request.Data.(*GetAppsListOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetAppsListResponse is the response type for the
// GetAppsList API operation.
type GetAppsListResponse struct {
	*GetAppsListOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetAppsList request.
func (r *GetAppsListResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}