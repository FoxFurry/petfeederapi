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

// UserApiController binds http requests to an api service and writes the service results to the http response
type UserApiController struct {
	service UserApiServicer
	errorHandler ErrorHandler
}

// UserApiOption for how the controller is set up.
type UserApiOption func(*UserApiController)

// WithUserApiErrorHandler inject ErrorHandler into controller
func WithUserApiErrorHandler(h ErrorHandler) UserApiOption {
	return func(c *UserApiController) {
		c.errorHandler = h
	}
}

// NewUserApiController creates a default api controller
func NewUserApiController(s UserApiServicer, opts ...UserApiOption) Router {
	controller := &UserApiController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the UserApiController
func (c *UserApiController) Routes() Routes {
	return Routes{ 
		{
			"UserDeviceGet",
			strings.ToUpper("Get"),
			"/user/devices",
			c.UserDeviceGet,
		},
		{
			"UserLogin",
			strings.ToUpper("Post"),
			"/user/login",
			c.UserLogin,
		},
		{
			"UserMetadataGet",
			strings.ToUpper("Get"),
			"/user/metadata",
			c.UserMetadataGet,
		},
		{
			"UserMetadataPatch",
			strings.ToUpper("Patch"),
			"/user/metadata",
			c.UserMetadataPatch,
		},
		{
			"UserRegister",
			strings.ToUpper("Post"),
			"/user/register",
			c.UserRegister,
		},
	}
}

// UserDeviceGet - Get all devices for specified user
func (c *UserApiController) UserDeviceGet(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.UserDeviceGet(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// UserLogin - Login a new user
func (c *UserApiController) UserLogin(w http.ResponseWriter, r *http.Request) {
	bodyParam := UserCredentials{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.UserLogin(r.Context(), bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// UserMetadataGet - Get metadata for specified user
func (c *UserApiController) UserMetadataGet(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.UserMetadataGet(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// UserMetadataPatch - Partially update metadata for specified user
func (c *UserApiController) UserMetadataPatch(w http.ResponseWriter, r *http.Request) {
	bodyParam := UserMetadata{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.UserMetadataPatch(r.Context(), bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// UserRegister - Register a new user
func (c *UserApiController) UserRegister(w http.ResponseWriter, r *http.Request) {
	bodyParam := UserCredentials{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.UserRegister(r.Context(), bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}
