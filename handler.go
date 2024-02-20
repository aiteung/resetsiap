package siappkg

import (
	"database/sql"
	"strings"

	"github.com/aiteung/module/model"
)

func Handler(Pesan model.IteungMessage, db *sql.DB) (reply string) {
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
	} else {
		return MessageSalahKeyword()
	}
	return
}
