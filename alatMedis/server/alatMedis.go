package server

import (
	"context"
)

// langka ke lima
type alatmedis struct {
	writer ReadWriter
}

func NewAlatMedis(writer ReadWriter) AlatMedisService {
	return &alatmedis{writer: writer}
}

//Methode pada interface CustomerService di service.go
func (a *alatmedis) AddAlatMedisService(ctx context.Context, alatmedis AlatMedis) error {
	//fmt.Println("Alat Medis")
	err := a.writer.AddAlatMedis(alatmedis)
	if err != nil {
		return err
	}

	return nil
}

func (a *alatmedis) ReadAlatMedisByKodeService(ctx context.Context, kode string) (AlatMedis, error) {
	am, err := a.writer.ReadAlatMedisByKode(kode)
	//fmt.Println(cus)
	if err != nil {
		return am, err
	}
	return am, nil
}

func (a *alatmedis) ReadAlatMedisByStatusService(ctx context.Context, status string) (AlatMediss, error) {
	am, err := a.writer.ReadAlatMedisByStatus(status)
	//fmt.Println(cus)
	if err != nil {
		return am, err
	}
	return am, nil
}

func (a *alatmedis) ReadAlatMedisService(ctx context.Context) (AlatMediss, error) {
	am, err := a.writer.ReadAlatMedis()
	//fmt.Println("customer", cus)
	if err != nil {
		return am, err
	}
	return am, nil
}

func (a *alatmedis) UpdateAlatMedisService(ctx context.Context, am AlatMedis) error {
	err := a.writer.UpdateAlatMedis(am)
	if err != nil {
		return err
	}
	return nil
}
