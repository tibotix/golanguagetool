// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostWordsDeleteReader is a Reader for the PostWordsDelete structure.
type PostWordsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostWordsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostWordsDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /words/delete] PostWordsDelete", response, response.Code())
	}
}

// NewPostWordsDeleteOK creates a PostWordsDeleteOK with default headers values
func NewPostWordsDeleteOK() *PostWordsDeleteOK {
	return &PostWordsDeleteOK{}
}

/*
PostWordsDeleteOK describes a response with status code 200, with default header values.

the result of removing the word
*/
type PostWordsDeleteOK struct {
	Payload *PostWordsDeleteOKBody
}

// IsSuccess returns true when this post words delete o k response has a 2xx status code
func (o *PostWordsDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post words delete o k response has a 3xx status code
func (o *PostWordsDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post words delete o k response has a 4xx status code
func (o *PostWordsDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post words delete o k response has a 5xx status code
func (o *PostWordsDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post words delete o k response a status code equal to that given
func (o *PostWordsDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post words delete o k response
func (o *PostWordsDeleteOK) Code() int {
	return 200
}

func (o *PostWordsDeleteOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /words/delete][%d] postWordsDeleteOK %s", 200, payload)
}

func (o *PostWordsDeleteOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /words/delete][%d] postWordsDeleteOK %s", 200, payload)
}

func (o *PostWordsDeleteOK) GetPayload() *PostWordsDeleteOKBody {
	return o.Payload
}

func (o *PostWordsDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostWordsDeleteOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PostWordsDeleteOKBody post words delete o k body
swagger:model PostWordsDeleteOKBody
*/
type PostWordsDeleteOKBody struct {

	// true if the word has been removed. false means the word hasn't been removed because it was not in the dictionary.
	Deleted bool `json:"deleted,omitempty"`
}

// Validate validates this post words delete o k body
func (o *PostWordsDeleteOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post words delete o k body based on context it is used
func (o *PostWordsDeleteOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostWordsDeleteOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostWordsDeleteOKBody) UnmarshalBinary(b []byte) error {
	var res PostWordsDeleteOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
