package siappkg

import "time"

type TblMhs struct {
	Nim     string `json:"Nim"`
	NamaMhs string `json:"Nama_Mhs"`
	TlpMhs  string `json:"Tlp_Mhs"`
}

type Perwalian struct {
	Nim           string `json:"Nim"`
	TahunAkademik string `json:"Thn_Akademik"`
	AppDosenWali  int    `json:"AppDosenWali"`
}

type TblSkmk struct {
	NamaMhs         string `json:"Nama_Mhs"`
	TempatTglLahir  string `json:"ttl"`
	IDAgama         string `json:"id_agama"`
	NamaAgama       string `json:"nama_agama"`
	AlamatMhs       string `json:"alamat_mhs"`
	KodeJp          string `json:"Kode_Jp"`
	Prodi           string `json:"prodi"`
	Nim             string `json:"Nim"`
	NamaAyah        string `json:"Nama_Ayah"`
	IDPekerjaanAyah int    `json:"id_pekerjaan_ayah"`
	NamaPekerjaan   string `json:"nama_pekerjaan"`
	AlamatOrangTua  string `json:"AlamatOrangTua"`
	KotaKodePos     string `json:"kota_kodepos"`
	TlpMhs          string `json:"Tlp_Mhs"`
	Email           string `json:"Email"`
	ThnAkademik     string
}

type AcademicYear struct {
	Nim              string
	ThnAkademik      string
	Periode          int
	TanggalPerwalian time.Time
}

type Email struct {
	From        string  `json:"from"`
	To          string  `json:"to"`
	Subject     string  `json:"subject"`
	Body        string  `json:"body"`
	Attachments []Files `json:"attachments"`
}

type Files struct {
	Mimetype string `json:"mimetype"`
	Name     string `json:"name"`
	Base64   string `json:"base64"`
}

type StatusEmail struct {
	Status string `json:"status"`
}

type SKMKResponse struct {
	Nama     string `json:"nama"`
	Npm      string `json:"npm"`
	Email    string `json:"email"`
	Filename string `json:"filename"`
}
