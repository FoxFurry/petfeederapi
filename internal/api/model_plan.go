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

type Plan struct {

	// Name of the plan
	PlanName string `json:"plan_name,omitempty"`

	// UUID of the pet
	PetUuid string `json:"pet_uuid,omitempty"`

	// TBD: Define data which is required for a plan
	PlanData string `json:"plan_data,omitempty"`

	// Plan UUID
	PlanUuid string `json:"plan_uuid,omitempty"`
}

// AssertPlanRequired checks if the required fields are not zero-ed
func AssertPlanRequired(obj Plan) error {
	return nil
}

// AssertRecursePlanRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of Plan (e.g. [][]Plan), otherwise ErrTypeAssertionError is thrown.
func AssertRecursePlanRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aPlan, ok := obj.(Plan)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertPlanRequired(aPlan)
	})
}
