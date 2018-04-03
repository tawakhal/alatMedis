package endpoint

import (
	"context"
	"fmt"

	sv "rumahsakit/alatMedis/alatMedis/server"
)

func (de AlatMedisEndpoint) AddAlatMedisService(ctx context.Context, alatmedis sv.AlatMedis) error {
	_, err := de.AddAlatMedisEndpoint(ctx, alatmedis)
	return err
}

func (de AlatMedisEndpoint) ReadAlatMedisByKodeService(ctx context.Context, Kode string) (sv.AlatMedis, error) {
	req := sv.AlatMedis{KodeAlatMedis: Kode}
	fmt.Println(req)
	resp, err := de.ReadAlatMedisByKodeEndpoint(ctx, req)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	am := resp.(sv.AlatMedis)
	return am, err
}

func (de AlatMedisEndpoint) ReadAlatMedisByStatusService(ctx context.Context, status string) (sv.AlatMediss, error) {
	req := sv.AlatMedis{Status: status}
	fmt.Println(req)
	resp, err := de.ReadAlatMedisByStatusEndpoint(ctx, req)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	am := resp.(sv.AlatMediss)
	return am, err
}

func (de AlatMedisEndpoint) ReadAlatMedisService(ctx context.Context) (sv.AlatMediss, error) {
	resp, err := de.ReadAlatMedisEndpoint(ctx, nil)
	fmt.Println("ce resp", resp)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	return resp.(sv.AlatMediss), err
}

func (de AlatMedisEndpoint) UpdateAlatMedisService(ctx context.Context, am sv.AlatMedis) error {
	_, err := de.UpdateAlatMedisEndpoint(ctx, am)
	if err != nil {
		fmt.Println("error pada endpoint:", err)
	}
	return err
}
