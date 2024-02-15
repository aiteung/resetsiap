package siappkg

import (
	"fmt"
	"testing"

	"github.com/aiteung/module/model"
	_ "github.com/denisenkom/go-mssqldb"
)

func TestResetDosen(t *testing.T) {
	pesan := model.IteungMessage{
		Message:      "Iteung ganti password siap dosen nisaaja",
		Phone_number: "628112311504",
	}
	ganti := Handler(pesan, Sqlconn())
	fmt.Println(ganti)
}

func TestResetMahasiswa(t *testing.T) {
	pesan := model.IteungMessage{
		Message:      "Iteung ganti password siap mahasiswa palenpalen",
		Phone_number: "6289522910966",
	}
	ganti := Handler(pesan, Sqlconn())
	fmt.Println(ganti)
}
