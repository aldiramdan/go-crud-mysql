package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/aldiramdan/go-crud-mysql/config"
	"github.com/aldiramdan/go-crud-mysql/entities"
)

type NasabahModel struct {
	conn *sql.DB
}

func NewNasabahModel() *NasabahModel {

	conn, err := config.DBConnection()
	if err != nil {
		panic(err)
	}

	return &NasabahModel{
		conn: conn,
	}

}

func (n *NasabahModel) FindAll() ([]entities.Nasabah, error) {

	rows, err := n.conn.Query("select * from nasabah")
	if err != nil {
		return []entities.Nasabah{}, err
	}
	defer rows.Close()

	var dataNasabah []entities.Nasabah
	for rows.Next() {
		var nasabah entities.Nasabah
		rows.Scan(
			&nasabah.Id,
			&nasabah.NamaLengkap,
			&nasabah.Alamat,
			&nasabah.TempatLahir,
			&nasabah.TanggalLahir,
			&nasabah.RataPenghasilan,
			&nasabah.JenisNasabah,
			&nasabah.IdCustomer)

		if nasabah.JenisNasabah == "1" {
			nasabah.JenisNasabah = "Perorangan"

		} else if nasabah.JenisNasabah == "2" {
			nasabah.JenisNasabah = "Perusahaan"
		} else if nasabah.JenisNasabah == "3" {
			nasabah.JenisNasabah = "WIC"
		}

		tgl_lahir, _ := time.Parse("2006-01-02", nasabah.TanggalLahir)

		nasabah.TanggalLahir = tgl_lahir.Format("02-01-2006")

		dataNasabah = append(dataNasabah, nasabah)
	}

	return dataNasabah, nil

}

func (n *NasabahModel) Create(nasabah entities.Nasabah) bool {

	result, err := n.conn.Exec("insert into nasabah (nama_lengkap, alamat, tempat_lahir, tanggal_lahir, avg_penghasilan, jenis_nasabah, id_customer) values(?,?,?,?,?,?,?)",
		nasabah.NamaLengkap, nasabah.Alamat, nasabah.TempatLahir, nasabah.TanggalLahir, nasabah.RataPenghasilan, nasabah.JenisNasabah, nasabah.IdCustomer)

	if err != nil {
		fmt.Println(err)
		return false
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId > 0

}

func (n *NasabahModel) Find(id int64, nasabah *entities.Nasabah) error {

	return n.conn.QueryRow("select * from nasabah where id = ?", id).Scan(
		&nasabah.Id,
		&nasabah.NamaLengkap,
		&nasabah.Alamat,
		&nasabah.TempatLahir,
		&nasabah.TanggalLahir,
		&nasabah.RataPenghasilan,
		&nasabah.JenisNasabah,
		&nasabah.IdCustomer)

}

func (n *NasabahModel) Update(nasabah entities.Nasabah) error {

	_, err := n.conn.Exec(
		"update nasabah set nama_lengkap = ?, alamat = ? , tempat_lahir = ? , tanggal_lahir = ?,avg_penghasilan = ? , jenis_nasabah = ? , id_customer = ? where id = ?",
		nasabah.NamaLengkap, nasabah.Alamat, nasabah.TanggalLahir, nasabah.TanggalLahir, nasabah.RataPenghasilan, nasabah.JenisNasabah, nasabah.IdCustomer, nasabah.Id)

	if err != nil {
		return err
	}

	return nil

}

func (n *NasabahModel) Delete(id int64) {

	n.conn.Exec("delete from nasabah where id = ?", id)

}
