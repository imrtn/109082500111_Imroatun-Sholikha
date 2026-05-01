package main
import "fmt"

const NMAX = 127
type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var ch string
	*n = 0

	fmt.Scan(&ch)
	for ch != "." && *n < NMAX {
		t[*n] = rune(ch[0])
		*n++

		fmt.Scan(&ch)
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(string(t[i]), " ")
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
}

func palindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var n int

	fmt.Print("Teks: ")
	isiArray(&tab, &n)

	fmt.Print("Reverse teks: ")
	balikanArray(&tab, n)
	cetakArray(tab, n)

	if palindrom(tab, n) {
		fmt.Println("Palindrom")
	} else {
		fmt.Println("Bukan palindrom")
	}
}