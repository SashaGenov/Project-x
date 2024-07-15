package main

import (
	"fmt"

	"github.com/Yandex-Practicum/go-first-floor-sprint-four/internal/ftracker"
)

func main() {
	fmt.Println(ftracker.ShowTrainingInfo(4000, "Бег", 0.15, 85, 185, 50, 5))
	fmt.Println(ftracker.ShowTrainingInfo(4000, "Ходьба", 1, 85, 185, 50, 5))
	fmt.Println(ftracker.ShowTrainingInfo(1000, "Плавание", 0.25, 85, 185, 100, 4))
	fmt.Println(ftracker.ShowTrainingInfo(1000, "Керлинг", 5, 85, 185, 50, 2))

}
