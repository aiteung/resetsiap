package reset

import (
	"context"
	"database/sql"
	"log"

	"github.com/aiteung/module/model"
)

func ResetPassword(db *sql.DB, PasswordBaru string, Pesan model.IteungMessage) (reply string) {
	mahasiswa, _ := GetMahasiswaByPhoneNumber(db, Pesan.Phone_number)
	// Lakukan perintah SQL untuk mengganti password
	_, err := db.Exec("UPDATE tblMHS SET Password = ? WHERE Tlp_Mhs = ?", PasswordBaru, Pesan.Phone_number)
	if err != nil {
		return MessageGagalReset(mahasiswa)
	}

	return MessageBerhasilReset(mahasiswa)
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
