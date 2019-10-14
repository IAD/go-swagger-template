// Code generated by go-swagger; DO NOT EDIT.

package todos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/sirupsen/logrus"

	models "github.com/IAD/go-swagger-template/examples/todo/gen/server/models"
)

// FindOKCode is the HTTP code returned for type FindOK
const FindOKCode int = 200

/*FindOK OK

swagger:response findOK
*/
type FindOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Item `json:"body,omitempty"`
}

// NewFindOKFunc is declaration for func that create response
type NewFindOKFunc func() *FindOK

// NewFindOK creates FindOK with default headers values
func NewFindOK() *FindOK {

	return &FindOK{}
}

// WithPayload adds the payload to the find o k response
func (o *FindOK) WithPayload(payload []*models.Item) *FindOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the find o k response
func (o *FindOK) SetPayload(payload []*models.Item) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*models.Item, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		logrus.Panic(err) // let the recovery middleware deal with this
	}

}

/*FindDefault error

swagger:response findDefault
*/
type FindDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewFindDefaultFunc is declaration for func that create response
type NewFindDefaultFunc func() *FindDefault

// NewFindDefault creates FindDefault with default headers values
func NewFindDefault(code int) *FindDefault {
	if code <= 0 {
		code = 500
	}

	return &FindDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the find default response
func (o *FindDefault) WithStatusCode(code int) *FindDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the find default response
func (o *FindDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the find default response
func (o *FindDefault) WithPayload(payload *models.Error) *FindDefault {
	o.Payload = payload
	return o
}

// WithErr adds the Error payload with a default code to the find default response
func (o *FindDefault) WithErr(err error) *FindDefault {
	type swaggerErr interface {
		Plain() (code string, detail string, attributes map[string]string)
	}

	if swaggerErr, ok := err.(swaggerErr); ok {
		code, detail, attributes := swaggerErr.Plain()

		o.Payload = &models.Error{
			Code:       code,
			Detail:     detail,
			Attributes: attributes,
		}
		return o
	}

	return o.WithError("500", err.Error())
}

// WithError  adds the Error payload to the find default response
func (o *FindDefault) WithError(code string, message string) *FindDefault {
	o.Payload = &models.Error{
		Code:    code,
		Message: message,
	}
	return o
}

// SetPayload sets the payload to the find default response
func (o *FindDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			logrus.Panic(err) // let the recovery middleware deal with this
		}
	}
}