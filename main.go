package goarea

import "math"

//Pi e uma proporção numérica definida pela relação entre

const Pi = 3.1416

//Circ é responsavel por calculoar a area da circunferencia
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// rECT É RESPONSAVEL POR CALCULAR A AREA DE UM Retangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

// Não é visivel!
func _Triangulo(base, altura float64) float64 {
	return (base * altura) / 2
}
