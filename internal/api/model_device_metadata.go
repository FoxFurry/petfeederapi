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

type DeviceMetadata struct {

	// Name of the device
	DeviceName string `json:"device_name,omitempty"`

	AssignedPet Pet `json:"assignedPet,omitempty"`

	AssignedPlan Plan `json:"assignedPlan,omitempty"`
}

// AssertDeviceMetadataRequired checks if the required fields are not zero-ed
func AssertDeviceMetadataRequired(obj DeviceMetadata) error {
	if err := AssertPetRequired(obj.AssignedPet); err != nil {
		return err
	}
	if err := AssertPlanRequired(obj.AssignedPlan); err != nil {
		return err
	}
	return nil
}

// AssertRecurseDeviceMetadataRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of DeviceMetadata (e.g. [][]DeviceMetadata), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseDeviceMetadataRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aDeviceMetadata, ok := obj.(DeviceMetadata)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertDeviceMetadataRequired(aDeviceMetadata)
	})
}