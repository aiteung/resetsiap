package reset

import (
	"database/sql"
	"strings"

	"github.com/aiteung/module/model"
)

func Handler(Pesan model.IteungMessage, db *sql.DB) (reply string) {
	if strings.Contains(Pesan.Message, "ganti") {
		// Ekstrak password baru dari pesan
		pesanSplit := strings.Split(Pesan.Message, " ")
		if len(pesanSplit) == 5 {
			PasswordBaru := pesanSplit[4]

			// Panggil fungsi ResetPassword untuk mengganti password
			reply = ResetPassword(db, Pesan.Phone_number, PasswordBaru)
		} else {
			// Jika pesan tidak sesuai format, berikan pesan error
			reply = "Format perintah salah. Gunakan format: Iteung ganti password siap [password_baru]"
		}
	} else if strings.Contains(Pesan.Message, "cara") {
		reply = CaraResetPassword(TblMhs{})
	}
	return
}

func ResetPassword(db *sql.DB, NomorHp string, PasswordBaru string) (reply string) {
	// Lakukan perintah SQL untuk mengganti password
	_, err := db.Exec("UPDATE tblMHS SET Password = ? WHERE Tlp_Mhs = ?", PasswordBaru, NomorHp)
	if err != nil {
		return MessageGagalReset(TblMhs{})
	}

	return MessageBerhasilReset(TblMhs{})
}

func MessageBerhasilReset(mhs TblMhs) string {
	msg := "*Reset Password*\n"
	msg = msg + "Hai kak _*" + mhs.NamaMhs + "*_,\ndengan nomor telepon *" + mhs.TlpMhs + "*,\nNPM *" + mhs.Nim + "*, \npassword kakak berhasil di reset.\nSilahkan kakak coba login lagi di https://siapmhs.ulbi.ac.id/login"
	return msg
}

func MessageGagalReset(mhs TblMhs) string {
	msg := "*Gagal Reset Password*\n"
	msg = msg + "Hai kak _*" + mhs.NamaMhs + "*_,\ndengan nomor telepon *" + mhs.TlpMhs + "*,\nNPM *" + mhs.Nim + "*, \nmaaf kak, password kakak gagal di reset :(.\nSilahkan kakak coba lagi yawww....."
	return msg
}

func CaraResetPassword(mhs TblMhs) string {
	msg := "*Reset Password*\n"
	msg = msg + "Hai kak _*" + mhs.NamaMhs + "*_,\ndengan nomor telepon *" + mhs.TlpMhs + "*,\nNPM *" + mhs.Nim + "*, \nKalo kakak mau ganti password SIAP kakak, kakak bisa ikutin instruksi iteung yaa.\nCaranya kakak tinggal ketikking perintah _Iteung ganti password siap passwordbarunya_.\nCukup gitu aja sih kak, iteung saranin pake password yang gampang diinget yaa, biar ga nyusahin iteung wkwkwk. Makasih kakk"
	return msg
}
