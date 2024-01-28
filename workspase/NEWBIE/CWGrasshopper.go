package main
import
"fmt"

func main() {
	var damage, health int
	fmt.Println("Enter a hit damage")
	fmt.Scanf("%d\n",&damage)
	fmt.Println("Enter a start health")
	fmt.Scanf("%d\n",&health)
	if health > 100 {
		health = 100
	}
	hitRegistration(damage, health)
	
}

func hitRegistration (damage int, health int)  {
	health -= damage
	if health <= 0 {
		fmt.Println("YOU DEAD! RESPAWNING")
		health = 100
	} 
	fmt.Printf("Your health is %d\n", health)
}