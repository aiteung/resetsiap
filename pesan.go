package siappkg

import "github.com/aiteung/module/model"

func MessageBerhasilReset(mhs TblMhs, PasswordBaru string) string {
	msg := "*Reset Password*\n"
	msg = msg + "Hai kak _*" + mhs.NamaMhs + "*_,\ndengan nomor telepon *" + mhs.TlpMhs + "*,\nNPM *" + mhs.Nim + "*, \npassword kakak berhasil di reset.\nSilahkan kakak coba login lagi di https://siapmhs.ulbi.ac.id/login\n*Password baru kamu : " + PasswordBaru + "*"
	return msg
}

func MessageBerhasilResetDosen(dosen TblDosen, newPassword string) string {
	msg := "*Reset Password*\n"
	msg = msg + "Hai kak _*" + dosen.Nama + "*_,\ndengan nomor telepon *" + dosen.Phone + "*,\nNIDN *" + dosen.Nidn + "*, \npassword kakak berhasil di reset.\nSilahkan kakak coba login lagi di https://siapdosen.ulbi.ac.id/login\n*Password baru kamu : " + newPassword + "*"
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

func MessageSalahKeyword() string {
	msg := "*Gagal Reset Password*\n"
	msg = msg + "Aduhhh keyword kakak belum bener nihh buat ganti password. Yang bener itu 'Iteung ganti password siap mahasiswa \"password_barunya\"' atau 'Iteung reset password siap mahasiswa \"password_barunya\"'. Ganti password_barunya dengan password yang kamu pengen yaa. Kalo password_barunya ngga kamu inputin, iteung kasih default password sariasih54 lohh"
	return msg
}
