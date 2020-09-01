// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package networkmanager

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetTransitGatewayRegistrationsInput struct {
	_ struct{} `type:"structure"`

	// The ID of the global network.
	//
	// GlobalNetworkId is a required field
	GlobalNetworkId *string `location:"uri" locationName:"globalNetworkId" type:"string" required:"true"`

	// The maximum number of results to return.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// The token for the next page of results.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`

	// The Amazon Resource Names (ARNs) of one or more transit gateways. The maximum
	// is 10.
	TransitGatewayArns []string `location:"querystring" locationName:"transitGatewayArns" type:"list"`
}

// String returns the string representation
func (s GetTransitGatewayRegistrationsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetTransitGatewayRegistrationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetTransitGatewayRegistrationsInput"}

	if s.GlobalNetworkId == nil {
		invalidParams.Add(aws.NewErrParamRequired("GlobalNetworkId"))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetTransitGatewayRegistrationsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.GlobalNetworkId != nil {
		v := *s.GlobalNetworkId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "globalNetworkId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.TransitGatewayArns != nil {
		v := s.TransitGatewayArns

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.QueryTarget, "transitGatewayArns", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

type GetTransitGatewayRegistrationsOutput struct {
	_ struct{} `type:"structure"`

	// The token for the next page of results.
	NextToken *string `type:"string"`

	// The transit gateway registrations.
	TransitGatewayRegistrations []TransitGatewayRegistration `type:"list"`
}

// String returns the string representation
func (s GetTransitGatewayRegistrationsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetTransitGatewayRegistrationsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.TransitGatewayRegistrations != nil {
		v := s.TransitGatewayRegistrations

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "TransitGatewayRegistrations", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opGetTransitGatewayRegistrations = "GetTransitGatewayRegistrations"

// GetTransitGatewayRegistrationsRequest returns a request value for making API operation for
// AWS Network Manager.
//
// Gets information about the transit gateway registrations in a specified global
// network.
//
//    // Example sending a request using GetTransitGatewayRegistrationsRequest.
//    req := client.GetTransitGatewayRegistrationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/networkmanager-2019-07-05/GetTransitGatewayRegistrations
func (c *Client) GetTransitGatewayRegistrationsRequest(input *GetTransitGatewayRegistrationsInput) GetTransitGatewayRegistrationsRequest {
	op := &aws.Operation{
		Name:       opGetTransitGatewayRegistrations,
		HTTPMethod: "GET",
		HTTPPath:   "/global-networks/{globalNetworkId}/transit-gateway-registrations",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &GetTransitGatewayRegistrationsInput{}
	}

	req := c.newRequest(op, input, &GetTransitGatewayRegistrationsOutput{})

	return GetTransitGatewayRegistrationsRequest{Request: req, Input: input, Copy: c.GetTransitGatewayRegistrationsRequest}
}

// GetTransitGatewayRegistrationsRequest is the request type for the
// GetTransitGatewayRegistrations API operation.
type GetTransitGatewayRegistrationsRequest struct {
	*aws.Request
	Input *GetTransitGatewayRegistrationsInput
	Copy  func(*GetTransitGatewayRegistrationsInput) GetTransitGatewayRegistrationsRequest
}

// Send marshals and sends the GetTransitGatewayRegistrations API request.
func (r GetTransitGatewayRegistrationsRequest) Send(ctx context.Context) (*GetTransitGatewayRegistrationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetTransitGatewayRegistrationsResponse{
		GetTransitGatewayRegistrationsOutput: r.Request.Data.(*GetTransitGatewayRegistrationsOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetTransitGatewayRegistrationsRequestPaginator returns a paginator for GetTransitGatewayRegistrations.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetTransitGatewayRegistrationsRequest(input)
//   p := networkmanager.NewGetTransitGatewayRegistrationsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetTransitGatewayRegistrationsPaginator(req GetTransitGatewayRegistrationsRequest) GetTransitGatewayRegistrationsPaginator {
	return GetTransitGatewayRegistrationsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetTransitGatewayRegistrationsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// GetTransitGatewayRegistrationsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetTransitGatewayRegistrationsPaginator struct {
	aws.Pager
}

func (p *GetTransitGatewayRegistrationsPaginator) CurrentPage() *GetTransitGatewayRegistrationsOutput {
	return p.Pager.CurrentPage().(*GetTransitGatewayRegistrationsOutput)
}

// GetTransitGatewayRegistrationsResponse is the response type for the
// GetTransitGatewayRegistrations API operation.
type GetTransitGatewayRegistrationsResponse struct {
	*GetTransitGatewayRegistrationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetTransitGatewayRegistrations request.
func (r *GetTransitGatewayRegistrationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}