package reset

import (
	"database/sql"
	"strings"

	"github.com/aiteung/module/model"
)

func Handler(Pesan model.IteungMessage, db *sql.DB) (reply string) {
	mahasiswa, _ := GetMahasiswaByPhoneNumber(db, Pesan.Phone_number)
	if strings.Contains(Pesan.Message, "ganti") {
		// Split pesan menjadi kata-kata
		pesanSplit := strings.Fields(Pesan.Message)
		foundSiap := false
		for i, kata := range pesanSplit {
			if kata == "siap" && i+1 < len(pesanSplit) {
				// Password baru adalah kata setelah "siap"
				PasswordBaru := pesanSplit[i+1]
				reply = ResetPassword(db, PasswordBaru, Pesan)
				foundSiap = true
				break
			}
		}

		// Jika tidak ada kata "siap" atau password baru, berikan pesan error
		if !foundSiap {
			reply = "Keyword kakak belum benar nihh, kakak harus ganti password dengan cara 'Iteung ganti password siap [password_baru]'. Maaciww kakakkk"
		}
	} else if strings.Contains(Pesan.Message, "cara") {
		reply = CaraResetPassword(mahasiswa)
	} else {
		return "Terjadi Error"
	}
	return
}

func ResetPassword(db *sql.DB, PasswordBaru string, Pesan model.IteungMessage) (reply string) {
	mahasiswa, _ := GetMahasiswaByPhoneNumber(db, Pesan.Phone_number)
	// Lakukan perintah SQL untuk mengganti password
	_, err := db.Exec("UPDATE tblMHS SET Password = ? WHERE Tlp_Mhs = ?", PasswordBaru, Pesan.Phone_number)
	if err != nil {
		return MessageGagalReset(mahasiswa)
	}

	return MessageBerhasilReset(mahasiswa)
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

func GetMahasiswaByPhoneNumber(db *sql.DB, phoneNumber string) (TblMhs, error) {
	// Query untuk mengambil data dari tabel tblMHS dengan kondisi WHERE Nomor Telepon
	query := "SELECT Nim, Nama_Mhs, Tlp_Mhs FROM tblMHS WHERE Tlp_Mhs = ?"

	var result TblMhs

	// Eksekusi query dan ambil data
	err := db.QueryRow(query, phoneNumber).Scan(&result.Nim, &result.NamaMhs, &result.TlpMhs)
	if err != nil {
		return TblMhs{}, err
	}

	return result, nil
}
