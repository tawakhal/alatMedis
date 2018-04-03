package server

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	addAlatMedis            = `INSERT INTO alat_medis(kode_alat_medis,nama_alat_medis,biaya,deskripsi,createdby,createdon,status) VALUES (?,?,?,?,?,?,?)`
	selectAlatMedisByKode   = `SELECT nama_alat_medis,biaya,deskripsi,createdby,createdon,status FROM dokter WHERE kode_alat_medis = ?`
	selectAlatMedisByStatus = `SELECT alat_medis(kode_alat_medis,nama_alat_medis,biaya,deskripsi,createdby,createdon FROM dokter WHERE status = ?`
	selectAlatMedis         = `SELECT kode_alat_medis,nama_alat_medis,biaya,deskripsi,createdby,createdon,status FROM dokter`
	updateAlatMedis         = `UPDATE alat SET nama_alat_medis =?, biaya =?, deskripsi = ?, updateby = ?, updateon =?, status = ? WHERE kode_dokter =?`
)

type dbReadWriter struct {
	db *sql.DB
}

func NewDBReadWriter(url string, schema string, user string, password string) ReadWriter {
	schemaURL := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, url, schema)
	db, err := sql.Open("mysql", schemaURL)
	if err != nil {
		panic(err)
	}
	return &dbReadWriter{db: db}
}

func (rw *dbReadWriter) AddAlatMedis(alatmedis AlatMedis) error {
	fmt.Println("add")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(alatmedis.KodeAlatMedis, alatmedis.NamaAlatMedis, alatmedis.Biaya, alatmedis.Deskripsi, alatmedis.CreatedBy, alatmedis.CreatedOn, alatmedis.Status)
	fmt.Println(err)
	if err != nil {
		return err

	}
	return tx.Commit()
}

func (rw *dbReadWriter) ReadAlatMedisByKode(kode string) (AlatMedis, error) {
	alatmedis := AlatMedis{KodeAlatMedis: kode}
	err := rw.db.QueryRow(selectAlatMedisByKode, kode).Scan(alatmedis.NamaAlatMedis, alatmedis.Biaya, alatmedis.Deskripsi, alatmedis.CreatedBy, alatmedis.CreatedOn, alatmedis.Status)

	if err != nil {
		return AlatMedis{}, err
	}

	return alatmedis, nil
}

func (rw *dbReadWriter) ReadAlatMedisByStatus(status string) (AlatMediss, error) {
	alatmedis := AlatMediss{}
	rows, _ := rw.db.Query(selectAlatMedisByStatus, status)
	defer rows.Close()
	for rows.Next() {
		var am AlatMedis
		err := rows.Scan(am.KodeAlatMedis, am.NamaAlatMedis, am.Biaya, am.Deskripsi, am.CreatedBy, am.CreatedOn, am.Status)
		if err != nil {
			fmt.Println("error query:", err)
			return AlatMediss{}, err
		}
		alatmedis = append(alatmedis, am)
	}
	//fmt.Println("db nya:", customer)
	return alatmedis, nil
}

func (rw *dbReadWriter) ReadAlatMedis() (AlatMediss, error) {
	alatmedis := AlatMediss{}
	rows, _ := rw.db.Query(selectAlatMedis)
	defer rows.Close()
	for rows.Next() {
		var am AlatMedis
		err := rows.Scan(am.KodeAlatMedis, am.NamaAlatMedis, am.Biaya, am.Deskripsi, am.CreatedBy, am.CreatedOn, am.Status)
		if err != nil {
			fmt.Println("error query:", err)
			return AlatMediss{}, err
		}
		alatmedis = append(alatmedis, am)
	}
	//fmt.Println("db nya:", customer)
	return alatmedis, nil
}

func (rw *dbReadWriter) UpdateAlatMedis(am AlatMedis) error {
	//fmt.Println("update")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(updateAlatMedis, am.NamaAlatMedis, am.Biaya, am.Deskripsi, am.UpdateBy, am.UpdateOn, am.Status, am.KodeAlatMedis)

	//fmt.Println("name:", cus.Name, cus.CustomerId)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return tx.Commit()
}
