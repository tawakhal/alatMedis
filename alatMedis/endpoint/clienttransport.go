package endpoint

import (
	"context"
	"time"

	svc "rumahsakit/alatMedis/alatMedis/server"

	pb "rumahsakit/alatMedis/alatMedis/grpc"

	util "rumahsakit/alatMedis/util/grpc"
	disc "rumahsakit/alatMedis/util/microservice"

	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/lb"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	stdopentracing "github.com/opentracing/opentracing-go"
	"github.com/sony/gobreaker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	grpcName = "grpc.AlatMedisService"
)

func NewGRPCAlatMedisClient(nodes []string, creds credentials.TransportCredentials, option util.ClientOption,
	tracer stdopentracing.Tracer, logger log.Logger) (svc.AlatMedisService, error) {

	instancer, err := disc.ServiceDiscovery(nodes, svc.ServiceID, logger)
	if err != nil {
		return nil, err
	}

	retryMax := option.Retry
	retryTimeout := option.RetryTimeout
	timeout := option.Timeout

	var addAlatMedisEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientAddAlatMedisEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		addAlatMedisEp = retry
	}

	var readAlatMedisByKodeEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadAlatMedisByKodeEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readAlatMedisByKodeEp = retry
	}

	var readAlatMedisByStatusEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadAlatMedisByStatusEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readAlatMedisByStatusEp = retry
	}
	var readAlatMedisEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadAlatMedisEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readAlatMedisEp = retry
	}

	var updateAlatMedisEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientUpdateAlatMedis, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		updateAlatMedisEp = retry
	}

	return AlatMedisEndpoint{AddAlatMedisEndpoint: addAlatMedisEp, ReadAlatMedisByKodeEndpoint: readAlatMedisByKodeEp, ReadAlatMedisByStatusEndpoint: readAlatMedisByStatusEp,
		ReadAlatMedisEndpoint: readAlatMedisEp, UpdateAlatMedisEndpoint: updateAlatMedisEp}, nil
}

func encodeAddAlatMedisRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.AlatMedis)
	return &pb.AddAlatMedisReq{
		KodeAlatMedis: req.KodeAlatMedis,
		NamaAlatMedis: req.NamaAlatMedis,
		Biaya:         req.Biaya,
		Deskripsi:     req.Deskripsi,
		CreatedBy:     req.CreatedBy,
		CreatedOn:     req.CreatedOn,
		Status:        req.Status,
	}, nil
}

func encodeReadAlatMedisByKodeRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.AlatMedis)
	return &pb.ReadAlatMedisByKodeReq{Kode: req.KodeAlatMedis}, nil
}

func encodeReadAlatMedisByStatusRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.AlatMedis)
	return &pb.ReadAlatMedisByStatusReq{Status: req.Status}, nil
}

func encodeReadAlatMedisRequest(_ context.Context, request interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeUpdateAlatMedisRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.AlatMedis)
	return &pb.UpdateAlatMedisReq{
		KodeAlatMedis: req.KodeAlatMedis,
		NamaAlatMedis: req.NamaAlatMedis,
		Biaya:         req.Biaya,
		Deskripsi:     req.Deskripsi,
		UpdateBy:      req.UpdateBy,
		UpdateOn:      req.UpdateOn,
		Status:        req.Status,
	}, nil
}

func decodeAlatMedisResponse(_ context.Context, response interface{}) (interface{}, error) {
	return nil, nil
}

func decodeReadAlatMedisByKodeRespones(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadAlatMedisByKodeResp)
	return svc.AlatMedis{
		KodeAlatMedis: resp.KodeAlatMedis,
		NamaAlatMedis: resp.NamaAlatMedis,
		Biaya:         resp.Biaya,
		Deskripsi:     resp.Deskripsi,
		CreatedBy:     resp.CreatedBy,
		CreatedOn:     resp.CreatedOn,
		UpdateBy:      resp.UpdateBy,
		UpdateOn:      resp.UpdateOn,
		Status:        resp.Status,
	}, nil
}

func decodeReadAlatMedisByStatusRespones(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadAlatMedisResp)
	var rsp svc.AlatMediss

	for _, v := range resp.Allkode {
		itm := svc.AlatMedis{
			KodeAlatMedis: v.KodeAlatMedis,
			NamaAlatMedis: v.NamaAlatMedis,
			Biaya:         v.Biaya,
			Deskripsi:     v.Deskripsi,
			CreatedBy:     v.CreatedBy,
			CreatedOn:     v.CreatedOn,
			UpdateBy:      v.UpdateBy,
			UpdateOn:      v.UpdateOn,
			Status:        v.Status,
		}
		rsp = append(rsp, itm)
	}
	return rsp, nil
}

func decodeReadAlatMedisResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadAlatMedisResp)
	var rsp svc.AlatMediss

	for _, v := range resp.Allkode {
		itm := svc.AlatMedis{
			KodeAlatMedis: v.KodeAlatMedis,
			NamaAlatMedis: v.NamaAlatMedis,
			Biaya:         v.Biaya,
			Deskripsi:     v.Deskripsi,
			CreatedBy:     v.CreatedBy,
			CreatedOn:     v.CreatedOn,
			UpdateBy:      v.UpdateBy,
			UpdateOn:      v.UpdateOn,
			Status:        v.Status,
		}
		rsp = append(rsp, itm)
	}
	return rsp, nil
}

func makeClientAddAlatMedisEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn,
		grpcName,
		"AddAlatMedis",
		encodeAddAlatMedisRequest,
		decodeAlatMedisResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "AddAlatMedis")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "AddAlatMedis",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadAlatMedisByKodeEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadAlatMedisByKode",
		encodeReadAlatMedisByKodeRequest,
		decodeReadAlatMedisByKodeRespones,
		pb.ReadAlatMedisByKodeResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadAlatMedisByKode")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadAlatMedisByKode",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadAlatMedisByStatusEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadAlatMedisByStatus",
		encodeReadAlatMedisByStatusRequest,
		decodeReadAlatMedisByStatusRespones,
		pb.ReadAlatMedisByStatusResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadAlatMedisByStatus")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadAlatMedisByStatus",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadAlatMedisEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadAlatMedis",
		encodeReadAlatMedisRequest,
		decodeReadAlatMedisResponse,
		pb.ReadAlatMedisResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadAlatMedis")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadAlatMedis",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientUpdateAlatMedis(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {
	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"UpdateAlatMedis",
		encodeUpdateAlatMedisRequest,
		decodeAlatMedisResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "UpdateAlatMedis")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "UpdateAlatMedis",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}
