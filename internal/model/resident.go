package model

import "time"

type Resident struct {
	ID                 int       `bson:"id,empty"`
	Nama               string    `bson:"nama,empty"`
	Alamat             string    `bson:"alamat,empty"`
	Difabel            string    `bson:"difabel,empty"`
	Ektp               string    `bson:"ektp,empty"`
	Email              string    `bson:"email,empty"`
	JenisKelamin       string    `bson:"jenis_kelamin,empty"`
	Kawin              string    `bson:"kawin,empty"`
	KabID              int       `bson:"kab_id,empty"`
	KecID              int       `bson:"kec_id,empty"`
	KelID              int       `bson:"kel_id,empty"`
	NamaKabupaten      string    `bson:"nama_kabupaten,empty"`
	NamaKecamatan      string    `bson:"nama_kecamatan,empty"`
	NamaKelurahan      string    `bson:"nama_kelurahan,empty"`
	Nik                string    `bson:"nik,empty"`
	Nkk                string    `bson:"nkk,empty"`
	NoKtp              string    `bson:"no_ktp,empty"`
	Rt                 string    `bson:"rt,empty"`
	Rw                 string    `bson:"rw,empty"`
	SaringanID         string    `bson:"saringan_Id,empty"`
	Status             string    `bson:"status,empty"`
	StatusTpsLabel     string    `bson:"status_tps_label,empty"`
	TanggalLahir       string    `bson:"tanggal_lahir,empty"`
	Usia               int       `bson:"usia", empty`
	IsTrue             int       `bson:"is_true", empty`
	IsFalse            int       `bson:"is_false", empty`
	TempatLahir        string    `bson:"tempat_lahir,empty"`
	TanggalLahirString string    `bson:"tanggal_lahir_string,empty"`
	Telp               string    `bson:"telp,empty"`
	Tps                string    `bson:"tps,empty"`
	CreatedAt          time.Time `bson:"created_at,empty"`
	UpdateAt           time.Time `bson:"updated_at,empty"`
}
