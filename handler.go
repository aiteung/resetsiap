package siappkg

import (
	"database/sql"
	"strings"

	"github.com/aiteung/module/model"
)

func Handler(Pesan model.IteungMessage, db *sql.DB) (reply string) {
	mahasiswa, _ := GetMahasiswaByPhoneNumber(db, Pesan.Phone_number)
	if (strings.Contains(Pesan.Message, "ganti") || strings.Contains(Pesan.Message, "reset")) && strings.Contains(Pesan.Message, "mahasiswa") {
		// Split pesan menjadi kata-kata
		pesanSplit := strings.Fields(Pesan.Message)
		foundSiap := false
		var PasswordBaru string
		for i := 0; i < len(pesanSplit); i++ {
			kata := pesanSplit[i]
			if kata == "siap" && i+2 < len(pesanSplit) {
				// Pastikan bahwa indeks i+2 tidak melebihi panjang pesanSplit
				PasswordBaru = pesanSplit[i+2]
				foundSiap = true
				break
			}
		}
		// Jika tidak ada kata "siap" atau password baru, berikan pesan error
		if !foundSiap {
			PasswordBaru = "sariasih54"
		}
		reply = ResetPassword(db, PasswordBaru, Pesan)
	} else if (strings.Contains(Pesan.Message, "ganti") || strings.Contains(Pesan.Message, "reset")) && strings.Contains(Pesan.Message, "dosen") {
		// Split pesan menjadi kata-kata
		pesanSplit := strings.Fields(Pesan.Message)
		foundSiap := false
		var PasswordBaru string
		for i := 0; i < len(pesanSplit); i++ {
			kata := pesanSplit[i]
			if kata == "siap" && i+2 < len(pesanSplit) {
				// Pastikan bahwa indeks i+2 tidak melebihi panjang pesanSplit
				PasswordBaru = pesanSplit[i+2]
				foundSiap = true
				break
			}
		}
		// Jika tidak ada kata "siap" atau password baru, berikan pesan error
		if !foundSiap {
			PasswordBaru = "sariasih54"
		}
		reply = ResetPasswordDosen(db, Pesan, PasswordBaru)
	} else if strings.Contains(Pesan.Message, "cara") {
		reply = CaraResetPassword(mahasiswa)
	} else if strings.Contains(Pesan.Message, "approval") {
		if !IsPhoneNumberExist(db, Pesan) {
			reply = MessageUpdateNomorDiSiap()
		} else {
			// Periksa persetujuan dosen wali
			tahun := GetCurrentAcademicYear()
			_, prw, _ := CheckMahasiswaApproval(db, Pesan, tahun)
			if prw.AppDosenWali == 2 {
				reply = MessageSudahApproval(mahasiswa, prw)
			} else if prw.AppDosenWali == 0 {
				reply = MessageBelumApproval(mahasiswa, prw)
			}
		}
	} else {
		return "Terjadi Error"
	}
	return
}
