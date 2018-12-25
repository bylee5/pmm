// Code generated by go-swagger; DO NOT EDIT.

package agents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// EnableAgentReader is a Reader for the EnableAgent structure.
type EnableAgentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EnableAgentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewEnableAgentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewEnableAgentOK creates a EnableAgentOK with default headers values
func NewEnableAgentOK() *EnableAgentOK {
	return &EnableAgentOK{}
}

/*EnableAgentOK handles this case with default header values.

(empty)
*/
type EnableAgentOK struct {
	Payload interface{}
}

func (o *EnableAgentOK) Error() string {
	return fmt.Sprintf("[POST /v0/inventory/Agents/Enable][%d] enableAgentOK  %+v", 200, o.Payload)
}

func (o *EnableAgentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*EnableAgentBody enable agent body
swagger:model EnableAgentBody
*/
type EnableAgentBody struct {

	// Unique Agent identifier.
	ID string `json:"id,omitempty"`
}

// Validate validates this enable agent body
func (o *EnableAgentBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *EnableAgentBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *EnableAgentBody) UnmarshalBinary(b []byte) error {
	var res EnableAgentBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
