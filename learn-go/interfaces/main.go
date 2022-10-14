package main
import (
	"fmt"
)

var print = fmt.Println

type Animal interface{
	AngrySound()
	HappySound()
}

type Cat string

func (c Cat) Attack() {
	print("attack mouses")
}

func (c Cat) Name() string{
	return string(c)
}
func (c Cat) AngrySound() {
	print("hiss")
}

func (c Cat) HappySound() {
	print("meow purr")
}
func main() {
	var kitty Animal
	kitty = Cat("kitty")
	kitty.AngrySound()
	var kitty2 Cat = kitty.(Cat)
	kitty2.Attack()
}
