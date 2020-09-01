// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package datasync

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateLocationFsxWindowsInput struct {
	_ struct{} `type:"structure"`

	// The name of the Windows domain that the FSx for Windows server belongs to.
	Domain *string `type:"string"`

	// The Amazon Resource Name (ARN) for the FSx for Windows file system.
	//
	// FsxFilesystemArn is a required field
	FsxFilesystemArn *string `type:"string" required:"true"`

	// The password of the user who has the permissions to access files and folders
	// in the FSx for Windows file system.
	//
	// Password is a required field
	Password *string `type:"string" required:"true" sensitive:"true"`

	// The Amazon Resource Names (ARNs) of the security groups that are to use to
	// configure the FSx for Windows file system.
	//
	// SecurityGroupArns is a required field
	SecurityGroupArns []string `min:"1" type:"list" required:"true"`

	// A subdirectory in the location’s path. This subdirectory in the Amazon
	// FSx for Windows file system is used to read data from the Amazon FSx for
	// Windows source location or write data to the FSx for Windows destination.
	Subdirectory *string `type:"string"`

	// The key-value pair that represents a tag that you want to add to the resource.
	// The value can be an empty string. This value helps you manage, filter, and
	// search for your resources. We recommend that you create a name tag for your
	// location.
	Tags []TagListEntry `type:"list"`

	// The user who has the permissions to access files and folders in the FSx for
	// Windows file system.
	//
	// User is a required field
	User *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateLocationFsxWindowsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateLocationFsxWindowsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateLocationFsxWindowsInput"}

	if s.FsxFilesystemArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("FsxFilesystemArn"))
	}

	if s.Password == nil {
		invalidParams.Add(aws.NewErrParamRequired("Password"))
	}

	if s.SecurityGroupArns == nil {
		invalidParams.Add(aws.NewErrParamRequired("SecurityGroupArns"))
	}
	if s.SecurityGroupArns != nil && len(s.SecurityGroupArns) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SecurityGroupArns", 1))
	}

	if s.User == nil {
		invalidParams.Add(aws.NewErrParamRequired("User"))
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateLocationFsxWindowsOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the FSx for Windows file system location
	// that is created.
	LocationArn *string `type:"string"`
}

// String returns the string representation
func (s CreateLocationFsxWindowsOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateLocationFsxWindows = "CreateLocationFsxWindows"

// CreateLocationFsxWindowsRequest returns a request value for making API operation for
// AWS DataSync.
//
// Creates an endpoint for an Amazon FSx for Windows file system.
//
//    // Example sending a request using CreateLocationFsxWindowsRequest.
//    req := client.CreateLocationFsxWindowsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/datasync-2018-11-09/CreateLocationFsxWindows
func (c *Client) CreateLocationFsxWindowsRequest(input *CreateLocationFsxWindowsInput) CreateLocationFsxWindowsRequest {
	op := &aws.Operation{
		Name:       opCreateLocationFsxWindows,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateLocationFsxWindowsInput{}
	}

	req := c.newRequest(op, input, &CreateLocationFsxWindowsOutput{})

	return CreateLocationFsxWindowsRequest{Request: req, Input: input, Copy: c.CreateLocationFsxWindowsRequest}
}

// CreateLocationFsxWindowsRequest is the request type for the
// CreateLocationFsxWindows API operation.
type CreateLocationFsxWindowsRequest struct {
	*aws.Request
	Input *CreateLocationFsxWindowsInput
	Copy  func(*CreateLocationFsxWindowsInput) CreateLocationFsxWindowsRequest
}

// Send marshals and sends the CreateLocationFsxWindows API request.
func (r CreateLocationFsxWindowsRequest) Send(ctx context.Context) (*CreateLocationFsxWindowsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateLocationFsxWindowsResponse{
		CreateLocationFsxWindowsOutput: r.Request.Data.(*CreateLocationFsxWindowsOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateLocationFsxWindowsResponse is the response type for the
// CreateLocationFsxWindows API operation.
type CreateLocationFsxWindowsResponse struct {
	*CreateLocationFsxWindowsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateLocationFsxWindows request.
func (r *CreateLocationFsxWindowsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}