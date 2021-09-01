package models

import (
	"fiber-go/db"
	"net/http"
)

type Pegawai struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Alamat  string `json:"alamat"`
	Telepon string `json:"telepon"`
}

func FethAllPegawai() (Response, error) {
	var obj Pegawai
	var arrobj []Pegawai
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM pegawai"

	rows, err := con.Query(sqlStatement)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Name, &obj.Alamat, &obj.Telepon)
		if err != nil {
			return res, err
		}
		arrobj = append(arrobj, obj)
	}
	res.Status = http.StatusOK
	res.Message = "sukses"
	res.Data = arrobj

	return res, nil
}

func UpdatePegawai(id int, name string, alamat string, telephone string) (Response, error) {
	var res Response
	con := db.CreateCon()
	sqlStatement := "UPDATE pegawai SET name=?, alamat=?, telephone=? WHERE id=?"
	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(name, alamat, telephone, id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rowsAffected": rowsAffected,
	}

	return res, nil
}
