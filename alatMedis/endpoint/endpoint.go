package endpoint

import (
	"context"

	svc "rumahsakit/alatMedis/alatMedis/server"

	kit "github.com/go-kit/kit/endpoint"
)

type AlatMedisEndpoint struct {
	AddAlatMedisEndpoint          kit.Endpoint
	ReadAlatMedisByKodeEndpoint   kit.Endpoint
	ReadAlatMedisByStatusEndpoint kit.Endpoint
	ReadAlatMedisEndpoint         kit.Endpoint
	UpdateAlatMedisEndpoint       kit.Endpoint
}

func NewAlatMedisEndpoint(service svc.AlatMedisService) AlatMedisEndpoint {
	addAlatMedisEp := makeAddAlatMedisEndpoint(service)
	readAlatMedisByKodeEp := makeReadAlatMedisByKodeEndpoint(service)
	readAlatMedisByStatusEp := makeReadAlatMedisByStatusEndpoint(service)
	readAlatMedisEp := makeReadAlatMedisEndpoint(service)
	updateAlatMedisEp := makeUpdateAlatMedisEndpoint(service)
	return AlatMedisEndpoint{AddAlatMedisEndpoint: addAlatMedisEp,
		ReadAlatMedisByKodeEndpoint:   readAlatMedisByKodeEp,
		ReadAlatMedisByStatusEndpoint: readAlatMedisByStatusEp,
		ReadAlatMedisEndpoint:         readAlatMedisEp,
		UpdateAlatMedisEndpoint:       updateAlatMedisEp,
	}
}

func makeAddAlatMedisEndpoint(service svc.AlatMedisService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.AlatMedis)
		err := service.AddAlatMedisService(ctx, req)
		return nil, err
	}
}

func makeReadAlatMedisByKodeEndpoint(service svc.AlatMedisService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.AlatMedis)
		result, err := service.ReadAlatMedisByKodeService(ctx, req.KodeAlatMedis)
		/*return svc.Customer{CustomerId: result.CustomerId, Name: result.Name,
		CustomerType: result.CustomerType, Mobile: result.Mobile, Email: result.Email,
		Gender: result.Gender, CallbackPhone: result.CallbackPhone, Status: result.Status}, err*/
		return result, err
	}
}

func makeReadAlatMedisByStatusEndpoint(service svc.AlatMedisService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.AlatMedis)
		result, err := service.ReadAlatMedisByStatusService(ctx, req.Status)
		/*return svc.Customer{CustomerId: result.CustomerId, Name: result.Name,
		CustomerType: result.CustomerType, Mobile: result.Mobile, Email: result.Email,
		Gender: result.Gender, CallbackPhone: result.CallbackPhone, Status: result.Status}, err*/
		return result, err
	}
}

func makeReadAlatMedisEndpoint(service svc.AlatMedisService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		result, err := service.ReadAlatMedisService(ctx)
		return result, err
	}
}

func makeUpdateAlatMedisEndpoint(service svc.AlatMedisService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.AlatMedis)
		err := service.UpdateAlatMedisService(ctx, req)
		return nil, err
	}
}
