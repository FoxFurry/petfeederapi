package service

import (
	openapi "FoxFurry/petfeederapi/internal/api"
	"context"
)

type Device struct {

}

var _ openapi.DeviceApiServicer = Device{}

func (s *Device) DeviceDelete(context.Context, string) (openapi.ImplResponse, error) {

}

func (s *Device) DeviceKey(context.Context, Device) (openapi.ImplResponse, error) {

}

func (s *Device) DeviceMetadataGet(context.Context) (openapi.ImplResponse, error) {

}

func (s *Device) DeviceMetadataPatch(context.Context, DeviceMetadata) (openapi.ImplResponse, error) {}

func (s *Device) {}
