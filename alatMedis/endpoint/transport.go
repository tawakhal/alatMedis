package endpoint

import (
	"context"

	scv "rumahsakit/alatMedis/alatMedis/server"

	pb "rumahsakit/alatMedis/alatMedis/grpc"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	stdopentracing "github.com/opentracing/opentracing-go"
	oldcontext "golang.org/x/net/context"
)

type grpcAlatMedisServer struct {
	addAlatMedis          grpctransport.Handler
	readAlatMedisByKode   grpctransport.Handler
	readAlatMedisByStatus grpctransport.Handler
	readAlatMedis         grpctransport.Handler
	updateAlatMedis       grpctransport.Handler
}

func NewGRPCAlatMedisServer(endpoints AlatMedisEndpoint, tracer stdopentracing.Tracer,
	logger log.Logger) pb.AlatMedisServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}
	return &grpcAlatMedisServer{
		addAlatMedis: grpctransport.NewServer(endpoints.AddAlatMedisEndpoint,
			decodeAddAlatMedisRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "AddAlatMedis", logger)))...),
		readAlatMedisByKode: grpctransport.NewServer(endpoints.ReadAlatMedisByKodeEndpoint,
			decodeReadAlatMedisByKodeRequest,
			encodeReadAlatMedisByKodeResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadAlatMedisByStatus", logger)))...),
		readAlatMedisByStatus: grpctransport.NewServer(endpoints.ReadAlatMedisByStatusEndpoint,
			decodeReadAlatMedisByStatusRequest,
			encodeReadAlatMedisByStatusResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadAlatMedisByKode", logger)))...),
		readAlatMedis: grpctransport.NewServer(endpoints.ReadAlatMedisEndpoint,
			decodeReadAlatMedisRequest,
			encodeReadAlatMedisResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadAlatMedis", logger)))...),
		updateAlatMedis: grpctransport.NewServer(endpoints.UpdateAlatMedisEndpoint,
			decodeUpdateAlatMedisRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "UpdateAlatMedis", logger)))...),
	}
}

func decodeAddAlatMedisRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.AddAlatMedisReq)
	return scv.AlatMedis{KodeAlatMedis: req.KodeAlatMedis, NamaAlatMedis: req.NamaAlatMedis, Biaya: req.Biaya, Deskripsi: req.Deskripsi, CreatedBy: req.CreatedBy,
		CreatedOn: req.CreatedOn, Status: req.Status}, nil
}

func decodeReadAlatMedisByKodeRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ReadAlatMedisByKodeReq)
	return scv.AlatMedis{KodeAlatMedis: req.Kode}, nil
}
func decodeReadAlatMedisByStatusRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ReadAlatMedisByStatusReq)
	return scv.AlatMedis{Status: req.Status}, nil
}
func decodeReadAlatMedisRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, nil
}

func decodeUpdateAlatMedisRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdateAlatMedisReq)
	return scv.AlatMedis{KodeAlatMedis: req.KodeAlatMedis, NamaAlatMedis: req.NamaAlatMedis, Biaya: req.Biaya, Deskripsi: req.Deskripsi,
		UpdateBy: req.UpdateBy, UpdateOn: req.UpdateOn, Status: req.Status}, nil
}

func encodeEmptyResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeReadAlatMedisByKodeResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.AlatMedis)
	return &pb.ReadAlatMedisByKodeResp{KodeAlatMedis: resp.KodeAlatMedis, NamaAlatMedis: resp.NamaAlatMedis, Biaya: resp.Biaya, Deskripsi: resp.Deskripsi, CreatedBy: resp.CreatedBy,
		CreatedOn: resp.CreatedOn, UpdateBy: resp.UpdateBy, UpdateOn: resp.UpdateOn, Status: resp.Status}, nil
}
func encodeReadAlatMedisByStatusResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.AlatMedis)
	return &pb.ReadAlatMedisByStatusResp{KodeAlatMedis: resp.KodeAlatMedis, NamaAlatMedis: resp.NamaAlatMedis, Biaya: resp.Biaya, Deskripsi: resp.Deskripsi, CreatedBy: resp.CreatedBy,
		CreatedOn: resp.CreatedOn, UpdateBy: resp.UpdateBy, UpdateOn: resp.UpdateOn, Status: resp.Status}, nil
}
func encodeReadAlatMedisResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.AlatMediss)

	rsp := &pb.ReadAlatMedisResp{}

	for _, v := range resp {
		itm := &pb.ReadAlatMedisByKodeResp{
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
		rsp.Allkode = append(rsp.Allkode, itm)
	}
	return rsp, nil
}

func (s *grpcAlatMedisServer) AddAlatMedis(ctx oldcontext.Context, alatMedis *pb.AddAlatMedisReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.addAlatMedis.ServeGRPC(ctx, alatMedis)
	if err != nil {
		return nil, err
	}
	return resp.(*google_protobuf.Empty), nil
}

func (s *grpcAlatMedisServer) ReadAlatMedisByKode(ctx oldcontext.Context, kode *pb.ReadAlatMedisByKodeReq) (*pb.ReadAlatMedisByKodeResp, error) {
	_, resp, err := s.readAlatMedisByKode.ServeGRPC(ctx, kode)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadAlatMedisByKodeResp), nil
}
func (s *grpcAlatMedisServer) ReadAlatMedisByStatus(ctx oldcontext.Context, status *pb.ReadAlatMedisByStatusReq) (*pb.ReadAlatMedisByStatusResp, error) {
	_, resp, err := s.readAlatMedisByStatus.ServeGRPC(ctx, status)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadAlatMedisByStatusResp), nil
}
func (s *grpcAlatMedisServer) ReadAlatMedis(ctx oldcontext.Context, e *google_protobuf.Empty) (*pb.ReadAlatMedisResp, error) {
	_, resp, err := s.readAlatMedis.ServeGRPC(ctx, e)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadAlatMedisResp), nil
}

func (s *grpcAlatMedisServer) UpdateAlatMedis(ctx oldcontext.Context, cus *pb.UpdateAlatMedisReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.updateAlatMedis.ServeGRPC(ctx, cus)
	if err != nil {
		return &google_protobuf.Empty{}, err
	}
	return resp.(*google_protobuf.Empty), nil
}
