package model

import "time"

type User struct {
	ID           int       `gorm:"primaryKey" json:"id"`
	Nama         string    `json:"nama"`
	Role         string    `json:"role"`
	Email        string    `json:"email"`
	Username     string    `json:"username"`
	Password     string    `json:"password"`
	NoHP         string    `json:"no_hp"`
	JenisKelamin string    `json:"jenis_kelamin"`
	Alamat       string    `json:"alamat"`
	CreatedAt    time.Time `json:"create_at"`
	UpdatedAt    time.Time `json:"update_at"`
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRegister struct {
	Nama         string `json:"nama"`
	Role         string `json:"role"`
	Email        string `json:"email"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	NoHP         string `json:"no_hp"`
	JenisKelamin string `json:"jenis_kelamin"`
	Alamat       string `json:"alamat"`
}

type DataMakam struct {
	ID     int    `gorm:"primaryKey" json:"id"`
	Alamat string `json:"alamat"`
	NoHP   string `json:"no_hp"`
	Email  string `json:"email"`
}

type DataJenazah struct {
	ID               int    `gorm:"primaryKey" json:"id"`
	Nama             string `json:"nama"`
	JenisKelamin     string `json:"jenis_kelamin"`
	TanggalLahir     string `json:"tanggal_lahir"`
	TanggalMeninggal string `json:"tanggal_meninggal"`
	Agama            string `json:"agama"`
	NamaAyah         string `json:"nama_ayah"`
	Alamat           string `json:"alamat"`
	NoMakam          string `json:"no_makam"`
}

type Pemesanan struct {
	ID            int    `json:"id"`
	NamaPemesan   string `json:"nama_pemesan"`
	NamaPadaMakam string `json:"nama_pada_makam"`
	NoMakam       string `json:"no_makam"`
	Keterangan    string `json:"keterangan"`
}

type Session struct {
	ID       int       `gorm:"primaryKey" json:"id"`
	Token    string    `json:"token"`
	Username string    `json:"username"`
	Expiry   time.Time `json:"expiry"`
}

type Credential struct {
	Host         string
	Username     string
	Password     string
	DatabaseName string
	Port         int
	Schema       string
}
