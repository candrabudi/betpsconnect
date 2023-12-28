package model

type Resident struct {
	ID                 int    `bson:"id,omitempty"`
	Nama               string `bson:"nama,omitempty"`
	Alamat             string `bson:"alamat,omitempty"`
	Difabel            string `bson:"difabel,omitempty"`
	Ektp               string `bson:"ektp,omitempty"`
	Email              string `bson:"email,omitempty"`
	JenisKelamin       string `bson:"jenis_kelamin,omitempty"`
	Kawin              string `bson:"kawin,omitempty"`
	KabID              int    `bson:"kab_id,omitempty"`
	KecID              int    `bson:"kec_id,omitempty"`
	KelID              int    `bson:"kel_id,omitempty"`
	NamaKabupaten      string `bson:"nama_kabupaten,omitempty"`
	NamaKecamatan      string `bson:"nama_kecamatan,omitempty"`
	NamaKelurahan      string `bson:"nama_kelurahan,omitempty"`
	Nik                string `bson:"nik,omitempty"`
	Nkk                string `bson:"nkk,omitempty"`
	NoKtp              string `bson:"no_ktp,omitempty"`
	Rt                 string `bson:"rt,omitempty"`
	Rw                 string `bson:"rw,omitempty"`
	SaringanID         string `bson:"saringan_Id,omitempty"`
	Status             string `bson:"status,omitempty"`
	StatusTpsLabel     string `bson:"status_tps_label,omitempty"`
	TanggalLahir       string `bson:"tanggal_lahir,omitempty"`
	TempatLahir        string `bson:"tempat_lahir,omitempty"`
	TanggalLahirString string `bson:"tanggal_lahir_string,omitempty"`
	Telp               string `bson:"telp,omitempty"`
	Tps                string `bson:"tps,omitempty"`
}
