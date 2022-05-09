package server

import (
	openapi "FoxFurry/petfeederapi/internal/api"
	"context"
	"net/http"
)

type PetFeeder interface {
	Start(ctx context.Context) error
}

type petfeeder struct {
	device openapi.DeviceApiServicer
	pet    openapi.PetApiServicer
	plan   openapi.PlanApiServicer
	user   openapi.UserApiServicer
}

func New(
	deviceService openapi.DeviceApiServicer,
	petService openapi.PetApiServicer,
	planService openapi.PlanApiServicer,
	userService openapi.UserApiServicer,
) PetFeeder {

	return &petfeeder{
		device: deviceService,
		pet:    petService,
		plan:   planService,
		user:   userService,
	}
}

func (p *petfeeder) Start(ctx context.Context) error {
	var (
		DeviceController = openapi.NewDeviceApiController(p.device)
		PetController    = openapi.NewPetApiController(p.pet)
		PlanController   = openapi.NewPlanApiController(p.plan)
		UserController   = openapi.NewUserApiController(p.user)
	)

	mux := openapi.NewRouter(DeviceController, PetController, PlanController, UserController)

	return http.ListenAndServe(":8080", mux)
}
