package siappkg

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"github.com/aiteung/module/model"
)

func ResetPassword(db *sql.DB, PasswordBaru string, Pesan model.IteungMessage) (reply string) {
	mahasiswa, err := GetMahasiswaByPhoneNumber(db, Pesan.Phone_number)
	if err != nil {
		return MessageGagalReset(Pesan)
	}

	// Jika PasswordBaru tidak diberikan, gunakan nilai default
	if PasswordBaru == "" {
		PasswordBaru = "sariasih54"
	}

	// Lakukan perintah SQL untuk mengganti password
	_, err = db.Exec("UPDATE tblMHS SET Password = ? WHERE Tlp_Mhs = ?", PasswordBaru, Pesan.Phone_number)
	if err != nil {
		return "Gagal mereset password. Terjadi kesalahan pada proses reset."
	}

	return MessageBerhasilReset(mahasiswa, PasswordBaru)
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

func GetDosenByPhoneNumber(db *sql.DB, phoneNumber string) (TblDosen, error) {
	// Query untuk mengambil data dari tabel tblMHS dengan kondisi WHERE Nomor Telepon
	query := "SELECT NIDN, Nama, Phone FROM tblDosen WHERE Phone = ?"

	var result TblDosen

	// Eksekusi query dan ambil data
	err := db.QueryRow(query, phoneNumber).Scan(&result.Nidn, &result.Nama, &result.Phone)
	if err != nil {
		return TblDosen{}, err
	}

	return result, nil
}

func ResetPasswordDosen(db *sql.DB, Pesan model.IteungMessage, newPassword string) (reply string) {
	dosen, err := GetDosenByPhoneNumber(db, Pesan.Phone_number)
	if err != nil {
		return MessageGagalResetDosen(Pesan)
	}
	// Generate MD5 hash for the new password
	hashedPassword, err := GenerateMD5Hash(newPassword)
	if err != nil {
		return "Gagal generate password MD5"
	}

	// Update password in the database
	_, err = db.Exec("UPDATE tblDosen SET Password = ? WHERE Phone = ?", hashedPassword, Pesan.Phone_number)
	if err != nil {
		return MessageGagalResetDosen(Pesan)
	}
	return MessageBerhasilResetDosen(dosen, newPassword)
}

func GenerateMD5Hash(password string) (string, error) {
	passwordBytes := []byte(password)
	hasher := md5.New()
	_, err := hasher.Write(passwordBytes)
	if err != nil {
		return "", err
	}
	hashBytes := hasher.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	return hashString, nil
}
