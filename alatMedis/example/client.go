package main

import (
	"context"
	"time"

	cli "rumahsakit/alatMedis/alatMedis/endpoint"
	svc "rumahsakit/alatMedis/alatMedis/server"
	opt "rumahsakit/alatMedis/util/grpc"
	util "rumahsakit/alatMedis/util/microservice"

	tr "github.com/opentracing/opentracing-go"
)

func main() {
	logger := util.Logger()
	tracer := tr.GlobalTracer()
	// dicobanya 3x, timeoutnya, dan balancingnya
	option := opt.ClientOption{Retry: 3, RetryTimeout: 500 * time.Millisecond, Timeout: 30 * time.Second}

	client, err := cli.NewGRPCAlatMedisClient([]string{"127.0.0.1:2181"}, nil, option, tracer, logger)
	if err != nil {
		logger.Log("error", err)
	}

	now := time.Now()
	//Add AlatMedis
	client.AddAlatMedisService(context.Background(), svc.AlatMedis{KodeAlatMedis: "KAM004", NamaAlatMedis: "Pisau L2", Biaya: 10000, Deskripsi: "Untuk Memotong selah 2MM", CreatedBy: "Olgi Tawakhal", CreatedOn: now.String(), Status: "1"})

	//Get AlatMedis By Kode No
	// cusMobile, _ := client.ReadDokterByKodeService(context.Background(), "DOK001")
	// fmt.Println("alatMedis based on Kode:", cusMobile)

	//List AlatMedis
	//  cuss, _ := client.ReadDokterService(context.Background())
	//  fmt.Println("All alatMedis:", cuss)

	//Update AlatMedis
	// client.UpdateDokterService(context.Background(), svc.Dokter{NamaDokter: "Olgi Tawakhal", JenisKelamin: "L", Alamat: "oakdowkdoawkdoawkd", NomorTelepon: "01231", KodeKategoriDokter: "KKD001", Biaya: 5000, KodeDokter: "DOK001"})

	//List AlatMedis
	// cuss, _ := client.ReadAlatMedisService(context.Background())
	// fmt.Println("All alatMedis:", cuss)

}
