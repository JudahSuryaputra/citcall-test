package main

import "fmt"

type Bebek struct {
	energi       int
	hidup        bool
	bisaTerbang  bool
	suaraTerbang string
}

func main() {
	donald := Bebek{
		hidup:        true,
		bisaTerbang:  true,
		suaraTerbang: "plek keplek keplek terbang ah",
	}

	donald.hidup = true
	donald.bisaTerbang = true

	fmt.Println("donald hidup?", donald.hidup)

	donald.Makan()
	donald.Makan()
	donald.Makan()

	fmt.Println("energi donald berapa?", donald.energi)

	donald.Terbang()
	donald.Terbang()

	fmt.Println("masih hidup?", donald.hidup)
}

func (b *Bebek) Mati() {
	b.hidup = false
}

func (b *Bebek) Terbang() {
	if b.hidup && b.bisaTerbang {
		fmt.Println(b.suaraTerbang)
		b.energi--
		if b.energi == 0 {
			b.Mati()
		}
	}
}

func (b *Bebek) Makan() {
	if b.hidup == true {
		b.energi++
	}
}
