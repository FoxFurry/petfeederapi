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

// UserCredentials - test
type UserCredentials struct {

	// Must be a valid email
	Mail string `json:"mail"`

	// Must be 8 characters long. Include at least 1 special character, at least 1 digit and at least one upper case character
	Password string `json:"password"`
}

// AssertUserCredentialsRequired checks if the required fields are not zero-ed
func AssertUserCredentialsRequired(obj UserCredentials) error {
	elements := map[string]interface{}{
		"mail": obj.Mail,
		"password": obj.Password,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseUserCredentialsRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of UserCredentials (e.g. [][]UserCredentials), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseUserCredentialsRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aUserCredentials, ok := obj.(UserCredentials)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertUserCredentialsRequired(aUserCredentials)
	})
}