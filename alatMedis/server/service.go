package server

import (
	"context"
)

type Status int32

const (
	//ServiceID is dispatch service ID

	//ini adalah konfigurasi sub domainnya
	ServiceID        = "rumahsakit"
	CreatedBy        = "Olgi Tawakhal"
	onAdd     Status = 1
)

type AlatMedis struct {
	KodeAlatMedis string
	NamaAlatMedis string
	Biaya         int64
	Deskripsi     string
	CreatedBy     string
	CreatedOn     string
	UpdateBy      string
	UpdateOn      string
	Status        string
}

type AlatMediss []AlatMedis

// ini interface untuk melakukan read
type ReadWriter interface {
	AddAlatMedis(AlatMedis) error
	ReadAlatMedisByKode(string) (AlatMedis, error)
	ReadAlatMedisByStatus(string) (AlatMediss, error)
	ReadAlatMedis() (AlatMediss, error)
	UpdateAlatMedis(AlatMedis) error
}

// ini interface yang mempunyai nilai return yang berupa interfase
type AlatMedisService interface {
	AddAlatMedisService(context.Context, AlatMedis) error
	ReadAlatMedisByKodeService(context.Context, string) (AlatMedis, error)
	ReadAlatMedisByStatusService(context.Context, string) (AlatMediss, error)
	ReadAlatMedisService(context.Context) (AlatMediss, error)
	UpdateAlatMedisService(context.Context, AlatMedis) error
}
