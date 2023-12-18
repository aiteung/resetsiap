package siappkg

import (
	"context"
	"database/sql"
	"fmt"
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

func GetMhsByPhoneNumber(db *sql.DB, PhoneNumber string) (TblSkmk, error) {
	// Query untuk mengambil data dari tabel tblMHS dengan kondisi WHERE Nomor Telepon
	query := "SELECT a.Nama_Mhs, CONCAT(a.Tmp_Lahir, ' / ', FORMAT(a.Tgl_Lahir, 'dd MMMM yyyy')) AS ttl, b.id_agama, b.nama_agama, CONCAT(a.Alamat_Mhs, ' Rt. ', a.rt, '/Rw. ', a.rw) AS alamat_mhs, c.Kode_Jp, CONCAT(c.Program, ' ', c.Jurusan) AS prodi, a.Nim, a.Nama_Ayah, a.id_pekerjaan_ayah, d.nama_pekerjaan, a.AlamatOrangTua, CONCAT(a.Kota_Mhs, ', ', a.Kodepos_Mhs) AS kota_kodepos, a.Tlp_Mhs, a.Email FROM tblMHS AS a JOIN feed_agama AS b ON a.id_agama = b.id_agama JOIN TblJurusan AS c ON a.Kode_Jp = c.Kode_Jp JOIN feed_pekerjaan AS d ON a.id_pekerjaan_ayah = d.id_pekerjaan WHERE Tlp_Mhs = ?"

	var result TblSkmk

	// Eksekusi query dan ambil data
	err := db.QueryRow(query, PhoneNumber).Scan(&result.NamaMhs, &result.TempatTglLahir, &result.IDAgama, &result.NamaAgama, &result.AlamatMhs, &result.KodeJp, &result.Prodi, &result.Nim, &result.NamaAyah, &result.IDPekerjaanAyah, &result.NamaPekerjaan, &result.AlamatOrangTua, &result.KotaKodePos, &result.TlpMhs, &result.Email)
	if err != nil {
		return TblSkmk{}, err
	}

	return result, nil
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
