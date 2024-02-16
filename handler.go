package siappkg

import (
	"database/sql"
	"strings"

	"github.com/aiteung/module/model"
)

func Handler(Pesan model.IteungMessage, db *sql.DB) (reply string) {
	mahasiswa, _ := GetMahasiswaByPhoneNumber(db, Pesan.Phone_number)
	if strings.Contains(Pesan.Message, "mahasiswa") {
		// Split pesan menjadi kata-kata
		pesanSplit := strings.Fields(Pesan.Message)
		foundSiap := false
		for i, kata := range pesanSplit {
			if kata == "siap" && i+1 < len(pesanSplit) {
				// Password baru adalah kata setelah "siap"
				PasswordBaru := pesanSplit[i+2]
				reply = ResetPassword(db, PasswordBaru, Pesan)
				foundSiap = true
				break
			}
		}
		// Jika tidak ada kata "siap" atau password baru, berikan pesan error
		if !foundSiap {
			reply = "Keyword kakak belum benar nihh, kakak harus ganti password dengan cara 'Iteung ganti password siap mahasiswa [password_baru]'. Maaciww kakakkk"
		}
	} else if strings.Contains(Pesan.Message, "dosen") {
		// Split pesan menjadi kata-kata
		pesanSplit := strings.Fields(Pesan.Message)
		foundSiap := false
		for i, kata := range pesanSplit {
			if kata == "siap" && i+1 < len(pesanSplit) {
				// Password baru adalah kata setelah "siap"
				PasswordBaru := strings.ToLower(pesanSplit[i+2])
				reply = ResetPasswordDosen(db, Pesan, PasswordBaru)
				foundSiap = true
				break
			}
		}
		// Jika tidak ada kata "siap" atau password baru, berikan pesan error
		if !foundSiap {
			reply = "Keyword kakak belum benar nihh, kakak harus ganti password dengan cara 'Iteung ganti password siap dosen [password_baru]'. Maaciww kakakkk"
		}
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
