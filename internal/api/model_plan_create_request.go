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

type PlanCreateRequest struct {

	// TBD: Define data which is required for a plan
	PlanData string `json:"plan_data,omitempty"`
}

// AssertPlanCreateRequestRequired checks if the required fields are not zero-ed
func AssertPlanCreateRequestRequired(obj PlanCreateRequest) error {
	return nil
}

// AssertRecursePlanCreateRequestRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of PlanCreateRequest (e.g. [][]PlanCreateRequest), otherwise ErrTypeAssertionError is thrown.
func AssertRecursePlanCreateRequestRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aPlanCreateRequest, ok := obj.(PlanCreateRequest)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertPlanCreateRequestRequired(aPlanCreateRequest)
	})
}
