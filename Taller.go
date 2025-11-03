package main

import (
	"fmt"
	"math"
)

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
	ocupada  bool      // true si la plaza está ocupada
	cliente  *Cliente  // cliente asociado a la plaza (si hay vehículo)
	mecanico *Mecanico // mecánico asignado a esa plaza
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
	incidencia   *Incidencia // incidencia actual asociada al vehículo
}

// Incidencia representa un trabajo o avería a reparar
type Incidencia struct {
	IDIncidencia int         // identificador único de la incidencia
	mecanicos    []*Mecanico // lista de mecánicos asignados a la incidencia
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

// MÉTODOS

func (t *Taller) InicializarPlazas() {
	t.MaxPlazas = 2 * len(t.MecanicosTaller)
	t.PlazasTaller = make([]*Plaza, t.MaxPlazas)
	for i := 0; i < t.MaxPlazas; i++ {
		t.PlazasTaller[i] = &Plaza{IDPlaza: i + 1}
	}
}

func (t *Taller) EstadoTaller() (ocupadas, libres int) {
	for _, p := range t.PlazasTaller {
		if p.ocupada {
			ocupadas++
		}
	}
	libres = len(t.PlazasTaller) - ocupadas
	return
}

func (t *Taller) BuscarVehiculo(matricula string) (*Cliente, *Vehiculo) {
	for _, c := range t.ClientesTaller {
		for _, v := range c.Vehiculos {
			if v.Matricula == matricula {
				return c, v
			}
		}
	}
	return nil, nil
}

func (t *Taller) ListarMecanicosDisponibles() []*Mecanico {
	var out []*Mecanico
	for _, m := range t.MecanicosTaller {
		if m.Activo {
			out = append(out, m)
		}
	}
	return out
}

// --- Plaza
func (p *Plaza) Ocupar(c *Cliente, m *Mecanico) {
	p.ocupada = true
	p.cliente = c
	p.mecanico = m
}
func (p *Plaza) Liberar() {
	p.ocupada = false
	p.cliente = nil
	p.mecanico = nil
}
func (p *Plaza) EstaLibre() bool { return !p.ocupada }
func (p *Plaza) GetCliente() *Cliente {
	return p.cliente
}
func (p *Plaza) GetMecanico() *Mecanico {
	return p.mecanico
}

// --- Vehiculo
func (v *Vehiculo) SetIncidencia(i *Incidencia) { v.incidencia = i }
func (v *Vehiculo) GetIncidencia() *Incidencia  { return v.incidencia }

// --- Incidencia
func (i *Incidencia) AsignarMecanico(m *Mecanico) {
	i.mecanicos = append(i.mecanicos, m)
}
func (i *Incidencia) GetMecanicos() []*Mecanico { return i.mecanicos }
func (i *Incidencia) SetEstado(estado string)   { i.Estado = estado }
func (i *Incidencia) GetEstado() string         { return i.Estado }
func (i *Incidencia) EsAltaPrioridad() bool     { return i.Prioridad == "alta" }

// --- Mecanico
func (m *Mecanico) CambiarEstado(activo bool) { m.Activo = activo }
func (m *Mecanico) Disponible() bool          { return m.Activo }

// VARIABLES GLOBALES
var app Taller
var nextIncID int = 1

// HELPERS

func findClienteByID(id int) (*Cliente, int) {
	for idx, c := range app.ClientesTaller {
		if c.IDCliente == id {
			return c, idx
		}
	}
	return nil, -1
}

func findMecanicoByID(id int) (*Mecanico, int) {
	for idx, m := range app.MecanicosTaller {
		if m.IDMecanico == id {
			return m, idx
		}
	}
	return nil, -1
}

func liberarPlazasDeCliente(c *Cliente) {
	for _, p := range app.PlazasTaller {
		if p.ocupada && p.cliente == c {
			p.Liberar()
		}
	}
}

func liberarPlazasDeMecanico(m *Mecanico) {
	for _, p := range app.PlazasTaller {
		if p.ocupada && p.mecanico == m {
			p.Liberar()
		}
	}
}

// MENÚS

// Menú: Clientes
func menuClientes() {
	var op int
	for {
		fmt.Println("\n===== GESTIÓN DE CLIENTES =====")
		fmt.Println("1. Crear cliente")
		fmt.Println("2. Visualizar clientes")
		fmt.Println("3. Modificar cliente")
		fmt.Println("4. Eliminar cliente")
		fmt.Println("0. Volver")
		fmt.Print("Opción: ")
		fmt.Scanln(&op)

		switch op {
		case 1:
			crearCliente()
		case 2:
			listarClientes()
		case 3:
			modificarCliente()
		case 4:
			eliminarCliente()
		case 0:
			return
		default:
			fmt.Println("Opción no válida.")
		}
	}
}

// Menú: Vehículos
func menuVehiculos() {
	var op int
	for {
		fmt.Println("\n===== GESTIÓN DE VEHÍCULOS =====")
		fmt.Println("1. Crear vehículo")
		fmt.Println("2. Visualizar vehículos")
		fmt.Println("3. Modificar vehículo")
		fmt.Println("4. Eliminar vehículo")
		fmt.Println("5. Registrar incidencia a un vehículo")
		fmt.Println("6. Consultar incidencia de un vehículo")
		fmt.Println("0. Volver")
		fmt.Print("Opción: ")
		fmt.Scanln(&op)

		switch op {
		case 1:
			crearVehiculo()
		case 2:
			listarVehiculos()
		case 3:
			modificarVehiculo()
		case 4:
			eliminarVehiculo()
		case 5:
			registrarIncidenciaVehiculo()
		case 6:
			consultarIncidenciaVehiculo()
		case 0:
			return
		default:
			fmt.Println("Opción no válida.")
		}
	}
}

// Menú: Incidencias
func menuIncidencias() {
	var op int
	for {
		fmt.Println("\n===== GESTIÓN DE INCIDENCIAS =====")
		fmt.Println("1. Crear incidencia (vehículo)")
		fmt.Println("2. Visualizar incidencias")
		fmt.Println("3. Modificar incidencia")
		fmt.Println("4. Eliminar incidencia")
		fmt.Println("5. Cambiar estado de incidencia")
		fmt.Println("0. Volver")
		fmt.Print("Opción: ")
		fmt.Scanln(&op)

		switch op {
		case 1:
			registrarIncidenciaVehiculo()
		case 2:
			listarIncidencias()
		case 3:
			modificarIncidencia()
		case 4:
			eliminarIncidencia()
		case 5:
			cambiarEstadoIncidencia()
		case 0:
			return
		default:
			fmt.Println("Opción no válida.")
		}
	}
}

// Menú: Mecánicos
func menuMecanicos() {
	var op int
	for {
		fmt.Println("\n===== GESTIÓN DE MECÁNICOS =====")
		fmt.Println("1. Crear mecánico")
		fmt.Println("2. Visualizar mecánicos")
		fmt.Println("3. Modificar mecánico")
		fmt.Println("4. Eliminar mecánico")
		fmt.Println("5. Dar de alta/baja a un mecánico")
		fmt.Println("0. Volver")
		fmt.Print("Opción: ")
		fmt.Scanln(&op)

		switch op {
		case 1:
			crearMecanico()
		case 2:
			listarMecanicos()
		case 3:
			modificarMecanico()
		case 4:
			eliminarMecanico()
		case 5:
			cambiarEstadoMecanico()
		case 0:
			return
		default:
			fmt.Println("Opción no válida.")
		}
	}
}

// CLIENTES
func crearCliente() {
	var id int
	var nombre, telefono, email string
	fmt.Print("ID cliente: ")
	fmt.Scanln(&id)
	if _, idx := findClienteByID(id); idx != -1 {
		fmt.Println("Ya existe un cliente con ese ID.")
		return
	}
	fmt.Print("Nombre: ")
	fmt.Scanln(&nombre)
	fmt.Print("Teléfono: ")
	fmt.Scanln(&telefono)
	fmt.Print("Email: ")
	fmt.Scanln(&email)

	c := &Cliente{IDCliente: id, Nombre: nombre, Telefono: telefono, Email: email}
	app.ClientesTaller = append(app.ClientesTaller, c)
	fmt.Println("Cliente creado.")
}

func listarClientes() {
	if len(app.ClientesTaller) == 0 {
		fmt.Println("No hay clientes.")
		return
	}
	fmt.Println("Listado de clientes:")
	for _, c := range app.ClientesTaller {
		fmt.Printf("- ID:%d | %s | Tel:%s | Email:%s | Vehículos:%d\n",
			c.IDCliente, c.Nombre, c.Telefono, c.Email, len(c.Vehiculos))
	}
}

func modificarCliente() {
	var id int
	fmt.Print("ID cliente a modificar: ")
	fmt.Scanln(&id)
	c, _ := findClienteByID(id)
	if c == nil {
		fmt.Println("Cliente no encontrado.")
		return
	}
	var nombre, tel, email string
	fmt.Print("Nuevo nombre: ")
	fmt.Scanln(&nombre)
	fmt.Print("Nuevo teléfono: ")
	fmt.Scanln(&tel)
	fmt.Print("Nuevo email: ")
	fmt.Scanln(&email)
	c.Nombre, c.Telefono, c.Email = nombre, tel, email
	fmt.Println("Cliente modificado.")
}

func eliminarCliente() {
	var id int
	fmt.Print("ID cliente a eliminar: ")
	fmt.Scanln(&id)
	c, idx := findClienteByID(id)
	if c == nil {
		fmt.Println("Cliente no encontrado.")
		return
	}
	// Liberar plazas ocupadas por este cliente (si las hubiera)
	liberarPlazasDeCliente(c)
	// Eliminar del slice
	app.ClientesTaller = append(app.ClientesTaller[:idx], app.ClientesTaller[idx+1:]...)
	fmt.Println("Cliente eliminado (y plazas liberadas si correspondía).")
}

// VEHÍCULOS
func crearVehiculo() {
	var idCliente int
	fmt.Print("ID del cliente propietario: ")
	fmt.Scanln(&idCliente)
	c, _ := findClienteByID(idCliente)
	if c == nil {
		fmt.Println("Cliente no encontrado.")
		return
	}
	var mat, marca, modelo, fIn, fOut string
	fmt.Print("Matrícula: ")
	fmt.Scanln(&mat)
	if _, v := app.BuscarVehiculo(mat); v != nil {
		fmt.Println("Ya existe un vehículo con esa matrícula.")
		return
	}
	fmt.Print("Marca: ")
	fmt.Scanln(&marca)
	fmt.Print("Modelo: ")
	fmt.Scanln(&modelo)
	fmt.Print("Fecha de entrada: ")
	fmt.Scanln(&fIn)
	fmt.Print("Fecha de salida: ")
	fmt.Scanln(&fOut)

	v := &Vehiculo{Matricula: mat, Marca: marca, Modelo: modelo, FechaEntrada: fIn, FechaSalida: fOut}
	c.Vehiculos = append(c.Vehiculos, v)
	fmt.Println("Vehículo creado y asignado al cliente.")
}

func listarVehiculos() {
	encontrados := 0
	for _, c := range app.ClientesTaller {
		for _, v := range c.Vehiculos {
			encontrados++
			estadoInc := "sin incidencia"
			if v.GetIncidencia() != nil {
				estadoInc = "incidencia " + v.GetIncidencia().Estado
			}
			fmt.Printf("- [%s] %s %s | Cliente:%s | %s\n",
				v.Matricula, v.Marca, v.Modelo, c.Nombre, estadoInc)
		}
	}
	if encontrados == 0 {
		fmt.Println("No hay vehículos registrados.")
	}
}

func modificarVehiculo() {
	var mat string
	fmt.Print("Matrícula del vehículo a modificar: ")
	fmt.Scanln(&mat)
	c, v := app.BuscarVehiculo(mat)
	if v == nil {
		fmt.Println("Vehículo no encontrado.")
		return
	}
	var marca, modelo, fIn, fOut string
	fmt.Print("Nueva marca: ")
	fmt.Scanln(&marca)
	fmt.Print("Nuevo modelo: ")
	fmt.Scanln(&modelo)
	fmt.Print("Nueva fecha de entrada: ")
	fmt.Scanln(&fIn)
	fmt.Print("Nueva fecha de salida: ")
	fmt.Scanln(&fOut)
	v.Marca, v.Modelo, v.FechaEntrada, v.FechaSalida = marca, modelo, fIn, fOut
	fmt.Printf("Vehículo %s del cliente %s modificado.\n", v.Matricula, c.Nombre)
}

func eliminarVehiculo() {
	var mat string
	fmt.Print("Matrícula del vehículo a eliminar: ")
	fmt.Scanln(&mat)
	c, v := app.BuscarVehiculo(mat)
	if v == nil {
		fmt.Println("Vehículo no encontrado.")
		return
	}
	// Si tuviera incidencia, la "eliminamos" (nil)
	v.SetIncidencia(nil)
	// Eliminar del slice del cliente
	pos := -1
	for i, vv := range c.Vehiculos {
		if vv == v {
			pos = i
			break
		}
	}
	if pos != -1 {
		c.Vehiculos = append(c.Vehiculos[:pos], c.Vehiculos[pos+1:]...)
	}
	fmt.Println("Vehículo eliminado.")
}

// INCIDENCIAS
func registrarIncidenciaVehiculo() {
	var mat string
	fmt.Print("Matrícula del vehículo: ")
	fmt.Scanln(&mat)
	c, v := app.BuscarVehiculo(mat)
	if v == nil {
		fmt.Println("Vehículo no encontrado.")
		return
	}
	if v.GetIncidencia() != nil {
		fmt.Println("Este vehículo ya tiene una incidencia (solo se permite una).")
		return
	}

	var tipo, prio, desc string
	fmt.Print("Tipo (mecánica/eléctrica/carrocería): ")
	fmt.Scanln(&tipo)
	fmt.Print("Prioridad (baja/media/alta): ")
	fmt.Scanln(&prio)
	fmt.Print("Descripción (una palabra o sin espacios): ")
	fmt.Scanln(&desc)

	inc := &Incidencia{
		IDIncidencia: nextIncID,
		Tipo:         tipo,
		Prioridad:    prio,
		Descripcion:  desc,
		Estado:       "abierta",
	}
	nextIncID++
	v.SetIncidencia(inc)

	fmt.Printf("Incidencia registrada al vehículo %s del cliente %s (ID=%d).\n",
		v.Matricula, c.Nombre, inc.IDIncidencia)
}

func consultarIncidenciaVehiculo() {
	var mat string
	fmt.Print("Matrícula del vehículo: ")
	fmt.Scanln(&mat)
	_, v := app.BuscarVehiculo(mat)
	if v == nil {
		fmt.Println("Vehículo no encontrado.")
		return
	}
	inc := v.GetIncidencia()
	if inc == nil {
		fmt.Println("El vehículo no tiene incidencia.")
		return
	}
	fmt.Printf("Incidencia ID:%d | Tipo:%s | Prioridad:%s | Estado:%s | Desc:%s | Mecánicos:%d\n",
		inc.IDIncidencia, inc.Tipo, inc.Prioridad, inc.Estado, inc.Descripcion, len(inc.GetMecanicos()))
}

func listarIncidencias() {
	total := 0
	for _, c := range app.ClientesTaller {
		for _, v := range c.Vehiculos {
			if inc := v.GetIncidencia(); inc != nil {
				total++
				fmt.Printf("- Vehículo [%s] de %s | IncID:%d | Tipo:%s | Prio:%s | Estado:%s\n",
					v.Matricula, c.Nombre, inc.IDIncidencia, inc.Tipo, inc.Prioridad, inc.Estado)
			}
		}
	}
	if total == 0 {
		fmt.Println("No hay incidencias registradas.")
	}
}

func modificarIncidencia() {
	var mat string
	fmt.Print("Matrícula del vehículo con incidencia: ")
	fmt.Scanln(&mat)
	_, v := app.BuscarVehiculo(mat)
	if v == nil || v.GetIncidencia() == nil {
		fmt.Println("Vehículo no encontrado o sin incidencia.")
		return
	}
	inc := v.GetIncidencia()
	var tipo, prio, desc string
	fmt.Print("Nuevo tipo (mecánica/eléctrica/carrocería): ")
	fmt.Scanln(&tipo)
	fmt.Print("Nueva prioridad (baja/media/alta): ")
	fmt.Scanln(&prio)
	fmt.Print("Nueva descripción (una palabra): ")
	fmt.Scanln(&desc)
	inc.Tipo, inc.Prioridad, inc.Descripcion = tipo, prio, desc
	fmt.Println("Incidencia modificada.")
}

func eliminarIncidencia() {
	var mat string
	fmt.Print("Matrícula del vehículo con incidencia a eliminar: ")
	fmt.Scanln(&mat)
	_, v := app.BuscarVehiculo(mat)
	if v == nil || v.GetIncidencia() == nil {
		fmt.Println("Vehículo no encontrado o sin incidencia.")
		return
	}
	v.SetIncidencia(nil)
	fmt.Println("Incidencia eliminada del vehículo.")
}

func cambiarEstadoIncidencia() {
	var mat, nuevo string
	fmt.Print("Matrícula del vehículo: ")
	fmt.Scanln(&mat)
	_, v := app.BuscarVehiculo(mat)
	if v == nil || v.GetIncidencia() == nil {
		fmt.Println("Vehículo no encontrado o sin incidencia.")
		return
	}
	fmt.Print("Nuevo estado (abierta/en proceso/cerrada): ")
	fmt.Scanln(&nuevo)
	v.GetIncidencia().SetEstado(nuevo)
	fmt.Println("Estado actualizado.")
}

// MECÁNICOS
func crearMecanico() {
	var id int
	var nombre, esp string
	var anios int
	fmt.Print("ID mecánico: ")
	fmt.Scanln(&id)
	if _, idx := findMecanicoByID(id); idx != -1 {
		fmt.Println("Ya existe un mecánico con ese ID.")
		return
	}
	fmt.Print("Nombre: ")
	fmt.Scanln(&nombre)
	fmt.Print("Especialidad (mecánica/eléctrica/carrocería): ")
	fmt.Scanln(&esp)
	fmt.Print("Años de experiencia: ")
	fmt.Scanln(&anios)

	m := &Mecanico{IDMecanico: id, Nombre: nombre, Especialidad: esp, AniosExperiencia: anios, Activo: true}
	app.MecanicosTaller = append(app.MecanicosTaller, m)
	app.InicializarPlazas()
	fmt.Println("Mecánico creado y plazas recalculadas.")
}

func listarMecanicos() {
	if len(app.MecanicosTaller) == 0 {
		fmt.Println("No hay mecánicos.")
		return
	}
	for _, m := range app.MecanicosTaller {
		status := "baja"
		if m.Activo {
			status = "activo"
		}
		fmt.Printf("- ID:%d | %s | %s | %d años | %s\n",
			m.IDMecanico, m.Nombre, m.Especialidad, m.AniosExperiencia, status)
	}
}

func modificarMecanico() {
	var id int
	fmt.Print("ID del mecánico a modificar: ")
	fmt.Scanln(&id)
	m, _ := findMecanicoByID(id)
	if m == nil {
		fmt.Println("No existe ese mecánico.")
		return
	}
	var nombre, esp string
	var anios int
	fmt.Print("Nuevo nombre: ")
	fmt.Scanln(&nombre)
	fmt.Print("Nueva especialidad (mecánica/eléctrica/carrocería): ")
	fmt.Scanln(&esp)
	fmt.Print("Nuevos años de experiencia: ")
	fmt.Scanln(&anios)
	m.Nombre, m.Especialidad, m.AniosExperiencia = nombre, esp, anios
	fmt.Println("Mecánico modificado.")
}

func eliminarMecanico() {
	var id int
	fmt.Print("ID del mecánico a eliminar: ")
	fmt.Scanln(&id)
	m, idx := findMecanicoByID(id)
	if m == nil {
		fmt.Println("No existe ese mecánico.")
		return
	}
	// Liberar plazas atendidas por este mecánico
	liberarPlazasDeMecanico(m)
	// Eliminar del slice
	app.MecanicosTaller = append(app.MecanicosTaller[:idx], app.MecanicosTaller[idx+1:]...)
	// Recalcular plazas por política 2 por mecánico
	app.InicializarPlazas()
	fmt.Println("Mecánico eliminado, plazas liberadas y recalculadas.")
}

func cambiarEstadoMecanico() {
	var id int
	var op int
	fmt.Print("ID del mecánico: ")
	fmt.Scanln(&id)
	m, _ := findMecanicoByID(id)
	if m == nil {
		fmt.Println("No existe ese mecánico.")
		return
	}
	fmt.Print("1=Activar, 2=Dar de baja: ")
	fmt.Scanln(&op)
	if op == 1 {
		m.CambiarEstado(true)
	} else if op == 2 {
		m.CambiarEstado(false)
	} else {
		fmt.Println("Opción inválida.")
		return
	}
	app.InicializarPlazas()
	fmt.Println("Estado del mecánico actualizado y plazas recalculadas.")
}

// PLAZAS / ESTADO TALLER
func asignarVehiculoAPlaza() {
	ocupadas, libres := app.EstadoTaller()
	if libres == 0 {
		fmt.Println("No hay plazas libres: taller lleno.")
		return
	}
	var mat string
	fmt.Print("Matrícula del vehículo a asignar: ")
	fmt.Scanln(&mat)
	cli, veh := app.BuscarVehiculo(mat)
	if veh == nil {
		fmt.Println("Vehículo no encontrado.")
		return
	}
	var idm int
	fmt.Print("ID del mecánico para asignar: ")
	fmt.Scanln(&idm)
	mec, _ := findMecanicoByID(idm)
	if mec == nil || !mec.Activo {
		fmt.Println("Mecánico inexistente o no activo.")
		return
	}
	for _, p := range app.PlazasTaller {
		if p.EstaLibre() {
			p.Ocupar(cli, mec)
			fmt.Printf("Vehículo %s asignado a plaza #%d con mecánico %s. (Ocupadas:%d→%d)\n",
				veh.Matricula, p.IDPlaza, mec.Nombre, ocupadas, ocupadas+1)
			return
		}
	}
	fmt.Println("No se encontró plaza libre (estado desactualizado).")
}

func consultarEstadoTaller() {
	ocupadas, libres := app.EstadoTaller()
	total := len(app.PlazasTaller)
	var pct float64 = 0
	if total > 0 {
		pct = math.Round((float64(ocupadas)/float64(total))*100.0 + 0.00001)
	}
	fmt.Printf("Plazas ocupadas: %d | libres: %d | total: %d | ocupación: %.0f%%\n", ocupadas, libres, total, pct)
	for _, p := range app.PlazasTaller {
		if p.ocupada {
			fmt.Printf(" - Plaza #%d: OCUPADA | Cliente:%s | Mecánico:%s\n",
				p.IDPlaza, p.GetCliente().Nombre, p.GetMecanico().Nombre)
		} else {
			fmt.Printf(" - Plaza #%d: libre\n", p.IDPlaza)
		}
	}
}

// MAIN

func main() {
	// Semilla de prueba
	app.MecanicosTaller = []*Mecanico{
		{IDMecanico: 1, Nombre: "Laura", Especialidad: "mecánica", AniosExperiencia: 3, Activo: true},
		{IDMecanico: 2, Nombre: "Pedro", Especialidad: "eléctrica", AniosExperiencia: 5, Activo: true},
	}
	app.ClientesTaller = []*Cliente{}
	app.InicializarPlazas()

	var opcion int
	for {
		fmt.Println("\n===== MENU PRINCIPAL =====")
		fmt.Println("1. Gestionar clientes")
		fmt.Println("2. Gestionar vehículos")
		fmt.Println("3. Gestionar incidencias")
		fmt.Println("4. Gestionar mecánicos")
		fmt.Println("5. Asignar vehículo a plaza")
		fmt.Println("6. Consultar estado del taller")
		fmt.Println("0. Salir")
		fmt.Print("Seleccione una opción: ")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			menuClientes()
		case 2:
			menuVehiculos()
		case 3:
			menuIncidencias()
		case 4:
			menuMecanicos()
		case 5:
			asignarVehiculoAPlaza()
		case 6:
			consultarEstadoTaller()
		case 0:
			fmt.Println("Saliendo del programa...")
			return
		default:
			fmt.Println("Opción no válida.")
		}
	}
}
