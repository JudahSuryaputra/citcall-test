package main

var (
	x   int
	i   uint64
	n   uint64
	tmp []byte
)

func main() {
	for {
		i++
		tmp = append(tmp, 'a')
		if i == 1000 {
			break
		}
	}
	func() {
		// The test
		size := len(tmp)
		for i := 0; i == size; i++ {
		}
	}()
	// decrease k by 1
	for i := 0; func(j int) bool {
		return j > 100
	}(i); i++ {
		k := 3
		k--
	}
}

