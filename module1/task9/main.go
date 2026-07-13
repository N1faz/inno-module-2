// main.go
package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("=" + repeat("=", 48))
	fmt.Println("   🧮 GO МАТЕМАТИЧЕСКИЙ КАЛЬКУЛЯТОР")
	fmt.Println("=" + repeat("=", 48))
	
	// Базовые операции
	a := 10
	b := 5
	
	fmt.Println("📊 Базовые операции:")
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
	fmt.Printf("%d - %d = %d\n", a, b, a-b)
	fmt.Printf("%d × %d = %d\n", a, b, a*b)
	fmt.Printf("%d ÷ %d = %d\n", a, b, a/b)
	fmt.Printf("%d %% %d = %d\n", a, b, a%b)
	
	// Степени
	fmt.Println("\n📈 Степени:")
	fmt.Printf("2^3 = %.0f\n", math.Pow(2, 3))
	fmt.Printf("5^2 = %.0f\n", math.Pow(5, 2))
	
	// Квадратный корень
	fmt.Println("\n🔢 Корни:")
	fmt.Printf("√16 = %.1f\n", math.Sqrt(16))
	fmt.Printf("√25 = %.1f\n", math.Sqrt(25))
	
	// Константы
	fmt.Println("\n📐 Константы:")
	fmt.Printf("PI = %.15f\n", math.Pi)
	fmt.Printf("E = %.15f\n", math.E)
	
	// Тригонометрия
	fmt.Println("\n📐 Тригонометрия:")
	fmt.Printf("sin(30°) = %.10f\n", math.Sin(30*math.Pi/180))
	fmt.Printf("cos(60°) = %.10f\n", math.Cos(60*math.Pi/180))
	fmt.Printf("tan(45°) = %.10f\n", math.Tan(45*math.Pi/180))
	
	// Случайные числа
	rand.Seed(time.Now().UnixNano())
	fmt.Println("\n🎲 Случайные числа:")
	fmt.Printf("Random 0-1: %.6f\n", rand.Float64())
	fmt.Printf("Random 1-100: %d\n", rand.Intn(100)+1)
	
	// Дополнительные математические функции
	fmt.Println("\n📊 Дополнительно:")
	fmt.Printf("Максимум (10, 20) = %d\n", max(10, 20))
	fmt.Printf("Минимум (10, 20) = %d\n", min(10, 20))
	fmt.Printf("Абсолютное значение -5 = %d\n", abs(-5))
	
	fmt.Println("\n" + "=" + repeat("=", 48))
	fmt.Println("✅ Go приложение работает в Distroless!")
}

// Вспомогательные функции
func repeat(s string, n int) string {
	result := ""
	for i := 0; i < n; i++ {
		result += s
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
