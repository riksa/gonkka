// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/riksa/gonkka/pkg/gocd/models"
)

// HistoryReader is a Reader for the History structure.
type HistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHistoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewHistoryOK creates a HistoryOK with default headers values
func NewHistoryOK() *HistoryOK {
	return &HistoryOK{}
}

/*HistoryOK handles this case with default header values.

History request success
*/
type HistoryOK struct {
	Payload *models.HistoryResponse
}

func (o *HistoryOK) Error() string {
	return fmt.Sprintf("[GET /pipelines/{app}/history/{id}][%d] historyOK  %+v", 200, o.Payload)
}

func (o *HistoryOK) GetPayload() *models.HistoryResponse {
	return o.Payload
}

func (o *HistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HistoryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
