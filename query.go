package siappkg

import (
	"context"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/aiteung/module/model"
)

func ResetPassword(db *sql.DB, PasswordBaru string, Pesan model.IteungMessage) (reply string) {
	mahasiswa, err := GetMahasiswaByPhoneNumber(db, Pesan.Phone_number)
	if err != nil {
		return MessageGagalReset(Pesan)
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

func CheckMahasiswaApproval(db *sql.DB, Pesan model.IteungMessage, TahunAkademik string) (TblMhs, Perwalian, error) {
	// Mendapatkan tahun akademik saat ini dalam format "YYYY-YYYY"
	TahunAkademik = GetCurrentAcademicYear()

	// Query SQL untuk menggabungkan data dari tabel tblMHS dan Perwalian
	query := `
        SELECT m.Nim, m.Nama_Mhs, m.Tlp_Mhs, p.AppDosenWali
        FROM tblMHS m
        LEFT JOIN Perwalian p ON m.Nim = p.Nim
        WHERE m.Tlp_Mhs = ? AND p.Thn_Akademik = ?
    `

	var mhs TblMhs
	var prw Perwalian
	row := db.QueryRowContext(context.Background(), query, Pesan.Phone_number, TahunAkademik)
	if err := row.Scan(&mhs.Nim, &mhs.NamaMhs, &mhs.TlpMhs, &prw.AppDosenWali); err != nil {
		return TblMhs{}, Perwalian{}, err
	}

	prw.TahunAkademik = TahunAkademik

	return mhs, prw, nil
}

func IsPhoneNumberExist(db *sql.DB, Pesan model.IteungMessage) bool {
	// Query database untuk memeriksa apakah nomor handphone sudah ada
	query := "SELECT COUNT(*) FROM tblMHS WHERE Tlp_Mhs = ?"
	var count int
	err := db.QueryRow(query, Pesan.Phone_number).Scan(&count)
	if err != nil {
		log.Println(err)
		return false // Terjadi kesalahan saat mengakses database
	}
	return count > 0
}

func GetTahunAkademik(db *sql.DB) (*AcademicYear, error) {
	// Query to get the active academic year
	query := "SELECT DISTINCT Thn_Akademik FROM Perwalian WHERE Tgl_Prw <= GETDATE() AND GETDATE() <= DATEADD(YEAR, 1, Tgl_Prw)"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Loop through the query results
	for rows.Next() {
		var tahunAkademik string
		if err := rows.Scan(&tahunAkademik); err != nil {
			return nil, err
		}

		// Return the AcademicYear struct with the retrieved value
		return &AcademicYear{ThnAkademik: tahunAkademik}, nil
	}

	// If no rows were returned, you may want to handle this case accordingly
	return nil, fmt.Errorf("no active academic year found")
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
