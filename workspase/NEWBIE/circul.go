package main

import (
	"errors"
	"fmt"
)

const pi = 3.1456
var choice int
var radius, input1, input2 float32

func main() {
	fmt.Println("Выберите необходимую фигуру: \n 1.Круг \n 2.Прямоугольник \n 3.Треугольник")
	fmt.Scanf("%d\n",&choice)
	switch choice {
		case 1: 
			fmt.Println("Вы выбрали круг, введите радиус окружности")
			fmt.Scanf("%f\n", &radius)
			printCircleArea(radius)
		case 2: 
			fmt.Println("Вы выбрали прямоугольник, введите длины сторон")
			fmt.Scanf("%f\n", &input1)
			fmt.Scanf("%f\n", &input2)
			printRectangleArea(input1,input2)
		case 3: 
			fmt.Println("Вы выбрали треугольник, введите длину основания и высоты")
			fmt.Scanf("%f\n", &input1)
			fmt.Scanf("%f\n", &input2)
			printTriangleArea(input1,input2)
		
		default: fmt.Println("Ошибка")
	}
	
}

func printCircleArea(radius float32) {
	circleArea, err := calculateArea(radius,input1,input2)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Радиус круга: %f см.\n", radius)
	fmt.Println("Формула расчета площади круга: A=пr2\n")
	fmt.Printf("Площадь круга: %f см. кв.\n\n", circleArea)
}
func printRectangleArea(input1 float32, input2 float32) {
	circleArea, err := calculateArea(radius,input1,input2)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Стороны прямоугльника равны: %f и %f см.\n", input1, input2)
	fmt.Println("Формула расчета площади круга: S= a x b")
	fmt.Printf("Площадь прямоугольника равна: %f см. кв.\n\n", circleArea)
}
func printTriangleArea(input1 float32, input2 float32) {
	circleArea, err := calculateArea(radius,input1,input2)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Стороны треугольника равны: %f и %f см.\n", input1, input2)
	fmt.Println("Формула расчета площади круга: S= a x b")
	fmt.Printf("Площадь треугольника равна: %f см. кв.\n\n", circleArea)
}
func calculateArea(radius,input1,input2 float32) (float32, error) {
	switch choice{
	case 1: 
		if radius <= 0 {
			return float32(0), errors.New("Такой фигуры не существует")
		}
		return radius * radius * pi, nil
	case 2: 
		if (input1 <= 0)||(input2 <=0) {
			return float32(0), errors.New("Такой фигуры не существует")
		}
		return input1 * input2, nil
	case 3: 
		if (input1 <= 0)||(input2 <=0) {
			return float32(0), errors.New("Такой фигуры не существует")
		}
		return input1 * input2/2, nil
	default: fmt.Println("Ошибка")
}
	return input1, nil
}
