// Code generated by go-swagger; DO NOT EDIT.

package metadata

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti-test-kitchen/zrok/rest_model_zrok"
)

// VersionOKCode is the HTTP code returned for type VersionOK
const VersionOKCode int = 200

/*
VersionOK retrieve the current server version

swagger:response versionOK
*/
type VersionOK struct {

	/*
	  In: Body
	*/
	Payload rest_model_zrok.Version `json:"body,omitempty"`
}

// NewVersionOK creates VersionOK with default headers values
func NewVersionOK() *VersionOK {

	return &VersionOK{}
}

// WithPayload adds the payload to the version o k response
func (o *VersionOK) WithPayload(payload rest_model_zrok.Version) *VersionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the version o k response
func (o *VersionOK) SetPayload(payload rest_model_zrok.Version) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *VersionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
