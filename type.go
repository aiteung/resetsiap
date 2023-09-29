package reset

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
