package main

// Kode setelah perubahan

// pembuatan tipe data baru bernama mobil

type mobil struct {
	roda      int
	kecepatan int
}

// pembuatan fungsi baru bernama start, yang dimana fungsi ini akan memanggil fungsi tambahKecepatan

func (car *mobil) start() {

	car.tambahKecepatan(10)

}

// pembuatan fungsi baru bernama tambahKecepatan, yang dimana fungsi ini akan menambahkan kecepatan mobil

func (car *mobil) tambahKecepatan(kecepatan int) {

	car.kecepatan = car.kecepatan + kecepatan
	car.roda = 4

}

// pembuatan fungsi baru bernama main, yang dimana fungsi ini akan memanggil fungsi start

func main() {

	firstCar := mobil{}
	firstCar.start()
	secondCar := mobil{}
	secondCar.start()

}
