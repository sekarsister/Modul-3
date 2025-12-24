package main
import (
    "fmt"
    "math"
)

func main() {
    var (
        jariJari,
        tKerucut,
        tTabung,
        vT,
        vK,
        total float64
    )
    
    const pi = 3.14
    
    fmt.Print("Input jari-jari : ")
    fmt.Scan(&jariJari)
    
    fmt.Print("Input tinggi kerucut : ")
    fmt.Scan(&tKerucut)
    
    fmt.Print("Input tinggi tabung : ")
    fmt.Scan(&tTabung)
    
    vT = pi * math.Pow(jariJari, 2) * tTabung
    vK = (pi * math.Pow(jariJari, 2) * tKerucut) / 3
    total = vT + vK
    
    fmt.Println("Volume tabung : ", math.Round(vT))
    fmt.Println("Volume kerucut : ", math.Round(vK))
    fmt.Println("Volume total : ", math.Round(total))
}
