package main

// Taller representa el sistema general del taller
type Taller struct {
	MaxPlazas       int         // número máximo de plazas del taller (2 por mecánico)
	ClientesTaller  []*Cliente  // lista de clientes registrados
	MecanicosTaller []*Mecanico // lista de mecánicos disponibles
	PlazasTaller    []*Plaza    // lista de plazas del taller
}

// Plaza representa una plaza física dentro del taller
type Plaza struct {
	IDPlaza  int       // identificador único de la plaza
	Ocupada  bool      // true si la plaza está ocupada
	Cliente  *Cliente  // cliente asociado a la plaza (si hay vehículo)
	Mecanico *Mecanico // mecánico asignado a esa plaza
}

// Cliente representa a un cliente del taller
type Cliente struct {
	IDCliente int         // identificador único del cliente
	Nombre    string      // nombre del cliente
	Telefono  string      // teléfono de contacto
	Email     string      // correo electrónico del cliente
	Vehiculos []*Vehiculo // lista de vehículos que pertenecen al cliente
}

// Vehiculo representa un coche registrado en el taller
type Vehiculo struct {
	Matricula    string      // matrícula del vehículo (identificador único)
	Marca        string      // marca del vehículo
	Modelo       string      // modelo del vehículo
	FechaEntrada string      // fecha de entrada al taller
	FechaSalida  string      // fecha estimada o real de salida
	Incidencia   *Incidencia // incidencia actual asociada al vehículo
}

// Incidencia representa un trabajo o avería a reparar
type Incidencia struct {
	IDIncidencia int         // identificador único de la incidencia
	Mecanicos    []*Mecanico // lista de mecánicos asignados a la incidencia
	Tipo         string      // tipo de incidencia: "mecánica", "eléctrica" o "carrocería"
	Prioridad    string      // nivel de prioridad: "baja", "media" o "alta"
	Descripcion  string      // descripción breve del problema
	Estado       string      // estado actual: "abierta", "en proceso" o "cerrada"
}

// Mecanico representa a un trabajador del taller
type Mecanico struct {
	IDMecanico       int    // identificador único del mecánico
	Nombre           string // nombre del mecánico
	Especialidad     string // área de especialidad: "mecánica", "eléctrica" o "carrocería"
	AniosExperiencia int    // años de experiencia en el taller
	Activo           bool   // true = activo, false = de baja
}
