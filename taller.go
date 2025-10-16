package main

import "fmt"

type Taller struct {
	Plazas   int
	Vehiculo Vehiculo
	Mecanico Mecanico
}

type Vehiculo struct {
	Matricula  string
	Modelo     string
	FechaIn    string
	FechaOut   string
	Problema   string
	Coste      int
	MecanicoID int
	Estado     string // "En espera", "En reparación", "Terminado"
}

type Mecanico struct {
	ID              int
	Nombre          string
	Especialidad    string
	AñosExperiencia int
	Disponibilidad  bool // true si esta activo, false si esta de baja
	Vehiculos       Vehiculo
}

func registrarV() Vehiculo {
	var v Vehiculo

	fmt.Println("=== REGISTRO DE VEHÍCULO ===")

	fmt.Print("Matrícula: ")
	fmt.Scanln(&v.Matricula)
	fmt.Print("Modelo: ")
	fmt.Scanln(&v.Modelo)
	fmt.Print("Fecha de entrada: ")
	fmt.Scanln(&v.FechaIn)
	fmt.Print("Fecha estimada de salida: ")
	fmt.Scanln(&v.FechaOut)
	fmt.Print("Problema detectado: ")
	fmt.Scanln(&v.Problema)
	fmt.Print("Coste estimado: ")
	fmt.Scanln(&v.Coste)

	v.MecanicoID = -1
	v.Estado = "En espera"

	fmt.Println("Vehículo registrado con éxito.")
	return v
}

func mostrarVehiculo(v Vehiculo) {
	fmt.Println("=== INFORMACIÓN DEL VEHÍCULO ===")
	fmt.Println("Matrícula: ", v.Matricula)
	fmt.Println("Modelo: ", v.Modelo)
	fmt.Println("Fecha de entrada: ", v.FechaIn)
	fmt.Println("Fecha de salida: ", v.FechaOut)
	fmt.Println("Problema detectado: ", v.Problema)
	fmt.Println("Coste de la reparación: ", v.Coste)
	fmt.Println("Mecanico asignado: ", v.MecanicoID)
	fmt.Println("Estado: ", v.Estado)
}

func showMenu() int {
	var option int
	fmt.Println("Opciones de la calculadora geométrica:")
	fmt.Println("1.- Registrar enrada de un vehículo")
	fmt.Println("2.- Asignar mecánico a un vehículo")
	fmt.Println("3.- Actualizar estado de la reparación")
	fmt.Println("4.- Registrar salida de un vehículo")
	fmt.Println("5.- Visualizar estado actual del taller (plazas ocupadas/libres)")
	fmt.Println("6.- Consultar información de un vehículo específico")
	fmt.Println("7.- Listar todos los vehículos asignados a un mecánico")
	fmt.Print("Elija su opción: ")
	fmt.Scanln(&option)
	return option
}

func main() {
	v1 := registrarV()
	mostrarVehiculo(v1)
}
