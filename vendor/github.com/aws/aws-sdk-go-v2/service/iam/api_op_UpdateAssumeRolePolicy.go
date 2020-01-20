// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/UpdateAssumeRolePolicyRequest
type UpdateAssumeRolePolicyInput struct {
	_ struct{} `type:"structure"`

	// The policy that grants an entity permission to assume the role.
	//
	// You must provide policies in JSON format in IAM. However, for AWS CloudFormation
	// templates formatted in YAML, you can provide the policy in JSON or YAML format.
	// AWS CloudFormation always converts a YAML policy to JSON format before submitting
	// it to IAM.
	//
	// The regex pattern (http://wikipedia.org/wiki/regex) used to validate this
	// parameter is a string of characters consisting of the following:
	//
	//    * Any printable ASCII character ranging from the space character (\u0020)
	//    through the end of the ASCII character range
	//
	//    * The printable characters in the Basic Latin and Latin-1 Supplement character
	//    set (through \u00FF)
	//
	//    * The special characters tab (\u0009), line feed (\u000A), and carriage
	//    return (\u000D)
	//
	// PolicyDocument is a required field
	PolicyDocument *string `min:"1" type:"string" required:"true"`

	// The name of the role to update with the new policy.
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters consisting of upper and lowercase alphanumeric characters
	// with no spaces. You can also include any of the following characters: _+=,.@-
	//
	// RoleName is a required field
	RoleName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateAssumeRolePolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateAssumeRolePolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateAssumeRolePolicyInput"}

	if s.PolicyDocument == nil {
		invalidParams.Add(aws.NewErrParamRequired("PolicyDocument"))
	}
	if s.PolicyDocument != nil && len(*s.PolicyDocument) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PolicyDocument", 1))
	}

	if s.RoleName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RoleName"))
	}
	if s.RoleName != nil && len(*s.RoleName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RoleName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/UpdateAssumeRolePolicyOutput
type UpdateAssumeRolePolicyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateAssumeRolePolicyOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateAssumeRolePolicy = "UpdateAssumeRolePolicy"

// UpdateAssumeRolePolicyRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Updates the policy that grants an IAM entity permission to assume a role.
// This is typically referred to as the "role trust policy". For more information
// about roles, go to Using Roles to Delegate Permissions and Federate Identities
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/roles-toplevel.html).
//
//    // Example sending a request using UpdateAssumeRolePolicyRequest.
//    req := client.UpdateAssumeRolePolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/UpdateAssumeRolePolicy
func (c *Client) UpdateAssumeRolePolicyRequest(input *UpdateAssumeRolePolicyInput) UpdateAssumeRolePolicyRequest {
	op := &aws.Operation{
		Name:       opUpdateAssumeRolePolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateAssumeRolePolicyInput{}
	}

	req := c.newRequest(op, input, &UpdateAssumeRolePolicyOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return UpdateAssumeRolePolicyRequest{Request: req, Input: input, Copy: c.UpdateAssumeRolePolicyRequest}
}

// UpdateAssumeRolePolicyRequest is the request type for the
// UpdateAssumeRolePolicy API operation.
type UpdateAssumeRolePolicyRequest struct {
	*aws.Request
	Input *UpdateAssumeRolePolicyInput
	Copy  func(*UpdateAssumeRolePolicyInput) UpdateAssumeRolePolicyRequest
}

// Send marshals and sends the UpdateAssumeRolePolicy API request.
func (r UpdateAssumeRolePolicyRequest) Send(ctx context.Context) (*UpdateAssumeRolePolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateAssumeRolePolicyResponse{
		UpdateAssumeRolePolicyOutput: r.Request.Data.(*UpdateAssumeRolePolicyOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateAssumeRolePolicyResponse is the response type for the
// UpdateAssumeRolePolicy API operation.
type UpdateAssumeRolePolicyResponse struct {
	*UpdateAssumeRolePolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateAssumeRolePolicy request.
func (r *UpdateAssumeRolePolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}