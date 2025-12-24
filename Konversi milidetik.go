package main
import "fmt"

func main() {
    var ms, detik float64
    fmt.Print("Input Millisecond : ")
    fmt.Scan(&ms)
    detik = ms / 1000
    fmt.Println("Hasil : ", detik)
}
