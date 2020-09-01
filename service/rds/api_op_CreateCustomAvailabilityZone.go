// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateCustomAvailabilityZoneInput struct {
	_ struct{} `type:"structure"`

	// The name of the custom Availability Zone (AZ).
	//
	// CustomAvailabilityZoneName is a required field
	CustomAvailabilityZoneName *string `type:"string" required:"true"`

	// The ID of an existing virtual private network (VPN) between the Amazon RDS
	// website and the VMware vSphere cluster.
	ExistingVpnId *string `type:"string"`

	// The name of a new VPN tunnel between the Amazon RDS website and the VMware
	// vSphere cluster.
	//
	// Specify this parameter only if ExistingVpnId isn't specified.
	NewVpnTunnelName *string `type:"string"`

	// The IP address of network traffic from your on-premises data center. A custom
	// AZ receives the network traffic.
	//
	// Specify this parameter only if ExistingVpnId isn't specified.
	VpnTunnelOriginatorIP *string `type:"string"`
}

// String returns the string representation
func (s CreateCustomAvailabilityZoneInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateCustomAvailabilityZoneInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateCustomAvailabilityZoneInput"}

	if s.CustomAvailabilityZoneName == nil {
		invalidParams.Add(aws.NewErrParamRequired("CustomAvailabilityZoneName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateCustomAvailabilityZoneOutput struct {
	_ struct{} `type:"structure"`

	// A custom Availability Zone (AZ) is an on-premises AZ that is integrated with
	// a VMware vSphere cluster.
	//
	// For more information about RDS on VMware, see the RDS on VMware User Guide.
	// (https://docs.aws.amazon.com/AmazonRDS/latest/RDSonVMwareUserGuide/rds-on-vmware.html)
	CustomAvailabilityZone *CustomAvailabilityZone `type:"structure"`
}

// String returns the string representation
func (s CreateCustomAvailabilityZoneOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateCustomAvailabilityZone = "CreateCustomAvailabilityZone"

// CreateCustomAvailabilityZoneRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Creates a custom Availability Zone (AZ).
//
// A custom AZ is an on-premises AZ that is integrated with a VMware vSphere
// cluster.
//
// For more information about RDS on VMware, see the RDS on VMware User Guide.
// (https://docs.aws.amazon.com/AmazonRDS/latest/RDSonVMwareUserGuide/rds-on-vmware.html)
//
//    // Example sending a request using CreateCustomAvailabilityZoneRequest.
//    req := client.CreateCustomAvailabilityZoneRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/CreateCustomAvailabilityZone
func (c *Client) CreateCustomAvailabilityZoneRequest(input *CreateCustomAvailabilityZoneInput) CreateCustomAvailabilityZoneRequest {
	op := &aws.Operation{
		Name:       opCreateCustomAvailabilityZone,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateCustomAvailabilityZoneInput{}
	}

	req := c.newRequest(op, input, &CreateCustomAvailabilityZoneOutput{})

	return CreateCustomAvailabilityZoneRequest{Request: req, Input: input, Copy: c.CreateCustomAvailabilityZoneRequest}
}

// CreateCustomAvailabilityZoneRequest is the request type for the
// CreateCustomAvailabilityZone API operation.
type CreateCustomAvailabilityZoneRequest struct {
	*aws.Request
	Input *CreateCustomAvailabilityZoneInput
	Copy  func(*CreateCustomAvailabilityZoneInput) CreateCustomAvailabilityZoneRequest
}

// Send marshals and sends the CreateCustomAvailabilityZone API request.
func (r CreateCustomAvailabilityZoneRequest) Send(ctx context.Context) (*CreateCustomAvailabilityZoneResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateCustomAvailabilityZoneResponse{
		CreateCustomAvailabilityZoneOutput: r.Request.Data.(*CreateCustomAvailabilityZoneOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateCustomAvailabilityZoneResponse is the response type for the
// CreateCustomAvailabilityZone API operation.
type CreateCustomAvailabilityZoneResponse struct {
	*CreateCustomAvailabilityZoneOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateCustomAvailabilityZone request.
func (r *CreateCustomAvailabilityZoneResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}