package siappkg

import "github.com/aiteung/module/model"

func MessageBerhasilReset(mhs TblMhs, PasswordBaru string) string {
	msg := "*Reset Password*\n"
	msg = msg + "Hai kak _*" + mhs.NamaMhs + "*_,\ndengan nomor telepon *" + mhs.TlpMhs + "*,\nNPM *" + mhs.Nim + "*, \npassword kakak berhasil di reset.\nSilahkan kakak coba login lagi di https://siapmhs.ulbi.ac.id/login\n*Password baru kamu : " + PasswordBaru + "*"
	return msg
}

func MessageBerhasilResetDosen(dosen TblDosen) string {
	msg := "*Reset Password*\n"
	msg = msg + "Hai kak _*" + dosen.Nama + "*_,\ndengan nomor telepon *" + dosen.Phone + "*,\nNIDN *" + dosen.Nidn + "*, \npassword kakak berhasil di reset.\nSilahkan kakak coba login lagi di https://siapdosen.ulbi.ac.id/login"
	return msg
}

func MessageGagalReset(Pesan model.IteungMessage) string {
	msg := "*Gagal Reset Password*\n"
	msg = msg + "Data Kamu dengan Nomor Telepon " + Pesan.Phone_number + " Gaada nihh.\nCoba cek dulu nomor kamu udah sama kayak di SIAP Mahasiswa atau belum yaa..."
	return msg
}

func MessageGagalResetDosen(Pesan model.IteungMessage) string {
	msg := "*Gagal Reset Password*\n"
	msg = msg + "Data Kamu dengan Nomor Telepon " + Pesan.Phone_number + " Gaada nihh.\nCoba cek dulu nomor kamu udah sama kayak di SIAP Dosen atau belum yaa..."
	return msg
}

func CaraResetPassword(mhs TblMhs) string {
	msg := "*Reset Password*\n"
	msg = msg + "Hai kak _*" + mhs.NamaMhs + "*_,\ndengan nomor telepon *" + mhs.TlpMhs + "*,\nNPM *" + mhs.Nim + "*, \nKalo kakak mau ganti password SIAP kakak, kakak bisa ikutin instruksi iteung yaa.\nCaranya kakak tinggal ketikkin perintah _Iteung ganti password siap mahasiswa passwordbarunya_.\nCukup gitu aja sih kak, iteung saranin pake password yang gampang diinget yaa, biar ga nyusahin iteung wkwkwk. Makasih kakk"
	return msg
}

func MessageSudahApproval(mahasiswa TblMhs, prw Perwalian) string {
	msg := "*Cek Approval*\n"
	msg = msg + "Hai kak_*" + mahasiswa.NamaMhs + "*_\nDengan nomor telepon *" + mahasiswa.TlpMhs + "*,\nNPM *" + mahasiswa.Nim + "*, \nTahun Akademik *" + prw.TahunAkademik + "*\nKakak *SUDAH DI APPROVE* sama dosen wali kakak. Kuliah nya yg semangat yaa kakk kata gue tehh...."
	return msg
}

func MessageBelumApproval(mahasiswa TblMhs, prw Perwalian) string {
	msg := "*Cek Approval*\n"
	msg = msg + "Hai kak_*" + mahasiswa.NamaMhs + "*_\nDengan nomor telepon *" + mahasiswa.TlpMhs + "*,\nNPM *" + mahasiswa.Nim + "*, \nTahun Akademik *" + prw.TahunAkademik + "*\nKakak *BELUM DI APPROVE* sama dosen wali kakak nihh ahh.\nCoba kakak hubungi dosen wali kakak dehh biar di approve dan kuliah kakak ga keganggu. Cepet lakuin ya kak apa kata iteung :). Kalo ngga, iteung marah nihh..."
	return msg
}

func MessageUpdateNomorDiSiap() string {
	msg := "*KETERANGAN*\n"
	msg = msg + "Kakak belum update nomor WhatsApp kakak di SIAP nihh. Kakak update dulu yaww nomor nya di SIAP sebelum mau akses aq. Caranya gampang kok kak, kakak tinggal buka https://siapmhs.ulbi.ac.id/login terus ke bagian profil mahasiswa. Disitu ada tanda pensil kak yang artinya edit. Nah disitu kakak update deh nomor WhatsApp kakak. Maaciww kakk"
	return msg
}

func MessageLengkapiData() string {
	msg := "*SKMK*\n"
	msg = msg + "Data kakak di siap belum lengkap nihh. Kakak lengkapi dulu ya, kayak profil kakak nya, pekerjaan orang tua kakak nya. Kalau udah, kakak boleh minta SKMK lagi sama akuu, maaciwww...."
	return msg
}

func MessageBerhasilMintaSkmk(mhs TblSkmk) string {
	msg := "*SKMK*\n"
	msg = msg + "SKMK Lagi dikirim sama iTeung ke email " + mhs.Email + ", di cek aja ya kak..."
	return msg
}

func MessageGagalMintaSkmk() string {
	msg := "*SKMK*\n"
	msg = msg + "SKMK gagal dikirim sama iTeung di email, coba kakak hubungi TIK yaa..."
	return msg
}
