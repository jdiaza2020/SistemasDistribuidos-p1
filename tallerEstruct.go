package main

// Taller representa el sistema general del taller
type Taller struct {
	MaxPlazas       int
	ClientesTaller  []*Cliente
	MecanicosTaller []*Mecanico
	PlazasTaller    []*Plaza
}

// Plaza representa a una plaza del taller
type Plaza struct {
	IDPlaza  int
	Ocupada  bool
	Cliente  *Cliente
	Mecanico *Mecanico
}

// Cliente representa a un cliente del taller
type Cliente struct {
	IDCliente int
	Nombre    string
	Telefono  string
	Email     string
	Vehiculos []*Vehiculo
}

// Vehiculo representa un coche registrado en el taller
type Vehiculo struct {
	Matricula    string
	Marca        string
	Modelo       string
	FechaEntrada string
	FechaSalida  string
	Incidencias  []*Incidencia
}

// Incidencia representa un trabajo o avería a reparar
type Incidencia struct {
	IDIncidencia int
	Mecanicos    []*Mecanico
	Tipo         string
	Prioridad    string
	Descripcion  string
	Estado       string
}

// Mecanico representa a un trabajador del taller
type Mecanico struct {
	IDMecanico       int
	Nombre           string
	Especialidad     string
	AniosExperiencia int
	Activo           bool
}
