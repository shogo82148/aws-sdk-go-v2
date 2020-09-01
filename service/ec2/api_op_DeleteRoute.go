// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
)

type DeleteRouteInput struct {
	_ struct{} `type:"structure"`

	// The IPv4 CIDR range for the route. The value you specify must match the CIDR
	// for the route exactly.
	DestinationCidrBlock *string `locationName:"destinationCidrBlock" type:"string"`

	// The IPv6 CIDR range for the route. The value you specify must match the CIDR
	// for the route exactly.
	DestinationIpv6CidrBlock *string `locationName:"destinationIpv6CidrBlock" type:"string"`

	// The ID of the prefix list for the route.
	DestinationPrefixListId *string `type:"string"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The ID of the route table.
	//
	// RouteTableId is a required field
	RouteTableId *string `locationName:"routeTableId" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteRouteInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteRouteInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteRouteInput"}

	if s.RouteTableId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RouteTableId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteRouteOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteRouteOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteRoute = "DeleteRoute"

// DeleteRouteRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Deletes the specified route from the specified route table.
//
//    // Example sending a request using DeleteRouteRequest.
//    req := client.DeleteRouteRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DeleteRoute
func (c *Client) DeleteRouteRequest(input *DeleteRouteInput) DeleteRouteRequest {
	op := &aws.Operation{
		Name:       opDeleteRoute,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteRouteInput{}
	}

	req := c.newRequest(op, input, &DeleteRouteOutput{})
	req.Handlers.Unmarshal.Remove(ec2query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteRouteRequest{Request: req, Input: input, Copy: c.DeleteRouteRequest}
}

// DeleteRouteRequest is the request type for the
// DeleteRoute API operation.
type DeleteRouteRequest struct {
	*aws.Request
	Input *DeleteRouteInput
	Copy  func(*DeleteRouteInput) DeleteRouteRequest
}

// Send marshals and sends the DeleteRoute API request.
func (r DeleteRouteRequest) Send(ctx context.Context) (*DeleteRouteResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteRouteResponse{
		DeleteRouteOutput: r.Request.Data.(*DeleteRouteOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteRouteResponse is the response type for the
// DeleteRoute API operation.
type DeleteRouteResponse struct {
	*DeleteRouteOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteRoute request.
func (r *DeleteRouteResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}