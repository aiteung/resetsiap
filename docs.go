package siappkg

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/base64"
	"fmt"
	"io"
	"log"

	"github.com/JPratama7/gwrap"
	docs2 "github.com/JPratama7/gwrap/docs"
	drive2 "github.com/JPratama7/gwrap/drive"
	"github.com/aiteung/module/model"
	"google.golang.org/api/docs/v1"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

// Retrieves a token, saves the token, then returns the generated client.
func ReplaceSkmk(data TblSkmk, akademik AcademicYear, db *sql.DB) (val []byte) {
	//Get Current Month
	bulan := convertRomanMonth()
	tahun := GetCurrentYear()
	thnakademik, _ := GetTahunAkademik(db)
	tglbulantahun := GetCurrentDate()
	ctx := context.Background()
	filepath := "credentials.json"
	cfg, err := gwrap.NewGoogleConfig(filepath, drive.DriveScope, drive.DriveReadonlyScope, docs.DocumentsScope, docs.DocumentsReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v\n", err)
		return
	}
	client := gwrap.GetClient(cfg, "token.json")

	srvDocs, err := docs.NewService(ctx, option.WithHTTPClient(client))
	srvDrive, err := drive.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Docs client: %v", err)
		return
	}

	// Prints the title of the requested doc:
	// https://docs.google.com/document/d/195j9eDD3ccgjQRttHhJPymLJUCOUjs-jmwTrekvdjFE/edit
	docId := "1ZfM-AoyDeLrt4G5RhbCQsHDCYj6QK9Vwq6jv_bkSNyI"

	drv := drive2.NewGoogleDrive(srvDrive)
	doc := docs2.NewGoogleDocs(srvDocs)

	namafile := data.Nim

	docDup, err := drv.CreateDuplicate(docId, fmt.Sprintf("SKMK-%s", namafile), "", nil)
	if err != nil {
		log.Fatalf("Unable to create duplicate: %v\n", err)
		return
	}
	// fmt.Printf("Duplicate ID : %s\n", docDup)

	listReplace := make([]*docs.Request, 0, 11)
	req1 := docs2.ReplaceTextDocs("{{Nama_Mhs}}", data.NamaMhs)
	req2 := docs2.ReplaceTextDocs("{{TempatTglLahir}}", data.TempatTglLahir)
	listReplace = append(listReplace, req1, req2)
	listReplace = append(listReplace, docs2.ReplaceTextDocs("{{Bulan}}", bulan))
	listReplace = append(listReplace, docs2.ReplaceTextDocs("{{Tahun}}", tahun))
	listReplace = append(listReplace, docs2.ReplaceTextDocs("{{NamaAgama}}", data.NamaAgama))
	listReplace = append(listReplace, docs2.ReplaceTextDocs("{{AlamatMhs}}", data.AlamatMhs))
	listReplace = append(listReplace, docs2.ReplaceTextDocs("{{Prodi}}", data.Prodi))
	listReplace = append(listReplace, docs2.ReplaceTextDocs("{{Nim}}", data.Nim))
	listReplace = append(listReplace, docs2.ReplaceTextDocs("{{NamaAyah}}", data.NamaAyah))
	listReplace = append(listReplace, docs2.ReplaceTextDocs("{{NamaPekerjaan}}", data.NamaPekerjaan))
	listReplace = append(listReplace, docs2.ReplaceTextDocs("{{AlamatOrangTua}}", data.AlamatOrangTua))
	listReplace = append(listReplace, docs2.ReplaceTextDocs("{{KotaKodePos}}", data.KotaKodePos))
	listReplace = append(listReplace, docs2.ReplaceTextDocs("{{ThnAkademik}}", thnakademik.ThnAkademik))
	listReplace = append(listReplace, docs2.ReplaceTextDocs("{{TglBulanTahun}}", tglbulantahun))

	err = doc.FindAndReplace(docDup, listReplace...)
	if err != nil {
		log.Fatalf("Unable to find and replace: %v", err)
		return
	}

	res, err := drv.DownloadFile(docDup, "application/pdf")
	if err != nil {
		log.Fatalf("Unable to retrieve data from document: %v", err)
	}

	defer res.Body.Close()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, res.Body)

	val = buf.Bytes()

	//Hapus file di google docs
	_, err = drv.DeleteFiles(docDup)

	return
}

func SendSkmk(urlEmail string, db *sql.DB, Pesan model.IteungMessage) (reply string) {
	dataMhs, err := GetMhsByPhoneNumber(db, Pesan.Phone_number)
	if err != nil {
		reply = MessageLengkapiData()
		return reply
	}

	tahunakademik, err := GetTahunAkademik(db)
	if err != nil {
		reply = "Tidak ada tahun akademik aktif"
		return reply
	}
	// Handle logika untuk "kirim skmk"
	SkmkFile := ReplaceSkmk(dataMhs, *tahunakademik, db)
	baseString := base64.StdEncoding.EncodeToString(SkmkFile)
	filename := fmt.Sprintf("SKMK - %s - %s.pdf", dataMhs.NamaMhs, dataMhs.Nim)

	attachEmail := Files{
		Mimetype: "application/pdf",
		Name:     filename,
		Base64:   baseString,
	}

	// Send email with attachment
	StatusSend := SendEmailTo(
		urlEmail,
		dataMhs.Email,
		"Surat Keterangan Masih Kuliah ULBI",
		fmt.Sprintf("Hai hai haiii Kakak %s, inii iteung bawain SKMK nih buat kakak. Gunakan seperlunya ya kakk.....", dataMhs.NamaMhs),
		attachEmail)

	if !StatusSend {
		return MessageGagalMintaSkmk()
	}

	return MessageBerhasilMintaSkmk(dataMhs)
}
