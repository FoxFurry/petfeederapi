/*
 * PetFeeder Gateway
 *
 * This is PBL VI main gateway
 *
 * API version: 1.2.0
 * Contact: isacartur@gmail.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

// DeviceApiController binds http requests to an api service and writes the service results to the http response
type DeviceApiController struct {
	service DeviceApiServicer
	errorHandler ErrorHandler
}

// DeviceApiOption for how the controller is set up.
type DeviceApiOption func(*DeviceApiController)

// WithDeviceApiErrorHandler inject ErrorHandler into controller
func WithDeviceApiErrorHandler(h ErrorHandler) DeviceApiOption {
	return func(c *DeviceApiController) {
		c.errorHandler = h
	}
}

// NewDeviceApiController creates a default api controller
func NewDeviceApiController(s DeviceApiServicer, opts ...DeviceApiOption) Router {
	controller := &DeviceApiController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the DeviceApiController
func (c *DeviceApiController) Routes() Routes {
	return Routes{ 
		{
			"DeviceDelete",
			strings.ToUpper("Delete"),
			"/device/{device_uuid}",
			c.DeviceDelete,
		},
		{
			"DeviceKey",
			strings.ToUpper("Post"),
			"/device/key",
			c.DeviceKey,
		},
		{
			"DeviceMetadataGet",
			strings.ToUpper("Get"),
			"/device/metadata",
			c.DeviceMetadataGet,
		},
		{
			"DeviceMetadataPatch",
			strings.ToUpper("Patch"),
			"/device/metadata",
			c.DeviceMetadataPatch,
		},
		{
			"DeviceRegister",
			strings.ToUpper("Post"),
			"/device",
			c.DeviceRegister,
		},
	}
}

// DeviceDelete - Delete a device with specified UUID
func (c *DeviceApiController) DeviceDelete(w http.ResponseWriter, r *http.Request) {
	deviceUuidParam := chi.URLParam(r, "device_uuid")
	
	result, err := c.service.DeviceDelete(r.Context(), deviceUuidParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// DeviceKey - Generate device token
func (c *DeviceApiController) DeviceKey(w http.ResponseWriter, r *http.Request) {
	bodyParam := Device{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.DeviceKey(r.Context(), bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// DeviceMetadataGet - Get metadata for specified device
func (c *DeviceApiController) DeviceMetadataGet(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.DeviceMetadataGet(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// DeviceMetadataPatch - Partially update metadata for specified user
func (c *DeviceApiController) DeviceMetadataPatch(w http.ResponseWriter, r *http.Request) {
	bodyParam := DeviceMetadata{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.DeviceMetadataPatch(r.Context(), bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// DeviceRegister - Register a new device
func (c *DeviceApiController) DeviceRegister(w http.ResponseWriter, r *http.Request) {
	bodyParam := Device{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.DeviceRegister(r.Context(), bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}
