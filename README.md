# Jorge Díaz Alcojor

# Sistema de Gestión de Taller Mecánico (Go)

## Descripción general

Este programa, desarrollado en **Go**, implementa un sistema de gestión para un **taller mecánico**.
Utiliza únicamente las estructuras de control, datos y punteros que hemos usado en los ejercicios vistos en clase.

El sistema se ejecuta íntegramente por consola, con un **menú principal** y **submenús** que reflejan los diferentes módulos de gestión: **clientes, vehículos, incidencias, mecánicos y plazas**.

---

## Estructura del programa

El código está contenido en un único archivo `Taller.go`, organizado en las siguientes secciones:

* **Estructuras de datos**: definición de `Taller`, `Plaza`, `Cliente`, `Vehiculo`, `Incidencia` y `Mecanico`.
* **Métodos asociados**: comportamiento propio de cada estructura (getters, setters y funciones de utilidad).
* **Funciones globales**: operaciones comunes como búsqueda, inicialización y validaciones.
* **Menú principal y submenús**: gestión independiente de cada módulo.
* **Funciones**: creación, lectura, actualización y eliminación de entidades.

---

## Menú principal y submenús

* **Clientes** → Crear, listar, modificar, eliminar (liberando plazas si corresponde).
* **Vehículos** → Crear, listar, modificar, eliminar, registrar o consultar incidencia.
* **Incidencias** → Crear (una por vehículo), listar, modificar, eliminar, cambiar estado.
* **Mecánicos** → Crear, listar, modificar, eliminar, dar de alta o baja (recalcula plazas).
* **Plazas / Taller** → Asignar vehículo a plaza, liberar implícitamente, visualizar estado actual y porcentaje de ocupación (usa `math.Round`).

---

## Funcionalidad implementada

* **Inicialización automática** de plazas (2 por cada mecánico activo).
* **Asignación controlada** de vehículos a plazas (solo si hay plazas libres).
* **Gestión de incidencias** asociadas a vehículos (una por vehículo).
* **Control de mecánicos activos**: solo los activos pueden asignarse a plazas o incidencias.
* **Cálculo de ocupación** del taller con porcentaje (`math`).

---

## Validaciones

* No se permite crear clientes o mecánicos con IDs duplicados.
* No se permite registrar vehículos con matrícula repetida.
* Un vehículo solo puede tener **una incidencia activa**.
* No se pueden asignar vehículos si **no hay plazas disponibles**.
* Las plazas se **liberan automáticamente** al eliminar un cliente o mecánico.
* Las plazas se **recalculan** al añadir o eliminar mecánicos (2 × nº mecánicos activos).

---

## Conceptos de Go aplicados

* **Estructuras (`struct`)** y **métodos** asociados.
* **Punteros (`*`)** para evitar duplicación de memoria y relacionar objetos.
* **Slices dinámicos (`append`)** para almacenar clientes, vehículos, incidencias, etc.
* **Control de flujo (`for`, `if`, `switch`)** para menús y decisiones.
* **Uso del paquete `math`** para redondear porcentajes en las estadísticas del taller.

---