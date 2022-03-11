package main

type Kendaraan struct {
	totalRoda       int
	kecepatanPerJam float32
}

type Mobil struct {
	Kendaraan
}

func (mobil *Mobil) berjalan() {
	mobil.tambahKecepatan(10)
}

func (mobil *Mobil) tambahKecepatan(kecepatanBaru float32) {
	mobil.kecepatanPerJam += kecepatanBaru
}

func main() {
	mobilCepat := Mobil{}
	mobilCepat.berjalan()
	mobilCepat.berjalan()
	mobilCepat.berjalan()

	mobilLamban := Mobil{}
	mobilLamban.berjalan()
}
