package main

type user struct {
	id       int
	username int
	password int
}

type userservice struct { // penamaan tipe data tidak bagus, seharusnya userService
	t []user // penamaan variabel tidak jelas, karena artian t sangat luas
}

func (u userservice) getallusers() []user { // penamaan fungsi tidak bagus, seharusnya getAllUsers
	return u.t // penamaan variabel tidak jelas, karena artian u dan t sangat luas
}

func (u userservice) getuserbyid(id int) user { // penamaan fungsi tidak bagus, seharusnya getUserById
	for _, r := range u.t { // penamaan variabel tidak jelas, karena artian r, u, dan t sangat luas
		if id == r.id {
			return r
		}
	}

	return user{}
}

// Kurangnya komentar yang dimana seharusnya menjelaskan fungsi, tipe data, dan variabel yang digunakan

// Berapa banyak kekurangan dalam penulisan kode tersebut?
// Jawaban: 3

// Bagian mana saja terjadi kekurangan tersebut?
// Jawaban: penamaan variabel, penamaan fungsi, dan tidak ada komentar

// Tuliskan alasan dari setiap kekurangan tersebut.
// Alasan bisa diberikan dalam bentuk komentar pada kode yang disediakan.
