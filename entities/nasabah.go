package entities

type Nasabah struct {
	Id              int64
	NamaLengkap     string `validate:"required" label:"Nama Lengkap"`
	Alamat          string `validate:"required"`
	TempatLahir     string `validate:"required" label:"Tempat Lahir"`
	TanggalLahir    string `validate:"required" label:"Tanggal Lahir"`
	RataPenghasilan string `validate:"required" label:"Rata-Rata Penghasilan"`
	JenisNasabah    string `validate:"required" label:"Jenis Nsabah"`
	IdCustomer      string `validate:"required" label:"ID Customer"`
}
