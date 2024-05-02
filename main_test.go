package main

import "testing"

// Funcion para testear la funcion crearProyecto
func TestCrearProyecto(t *testing.T) {
	// Crear un mapa de proyectos vacío
	listaProyecto := make(map[string]*Proyecto)

	// Datos del proyecto
	nombre := "Proyecto X"
	integrantes := []string{"Integrante1", "Integrante2"}
	inicio := "2024-01-01"
	hitos := []string{"Hito1", "Hito2"}
	fechasHito := []string{"2024-02-01", "2024-03-01"}
	fin := "2024-12-31"

	// Llamar a la función para crear el proyecto
	crearProyecto(listaProyecto, nombre, integrantes, inicio, hitos, fechasHito, fin)

	// Verificar que el proyecto existe en el mapa
	proyectoCreado := listaProyecto[nombre]
	if proyectoCreado == nil {
		t.Errorf("El proyecto no se creó correctamente")
	}
}

// Funcion para testear la funcion buscarPorProyecto
func TestBuscarPorProyecto(t *testing.T) {

	// Crear un mapa de proyectos vacío
	listaProyecto := make(map[string]*Proyecto)

	// Datos del proyecto
	nombre := "Proyecto X"
	integrantes := []string{"Integrante1", "Integrante2"}
	inicio := "2024-01-01"
	hitos := []string{"Hito1", "Hito2"}
	fechasHito := []string{"2024-02-01", "2024-03-01"}
	fin := "2024-12-31"

	// Llamar a la función para crear el proyecto
	crearProyecto(listaProyecto, nombre, integrantes, inicio, hitos, fechasHito, fin)

	encontrado := buscarPorProyecto(listaProyecto, "Proyecto X")
	if encontrado == nil {
		t.Errorf("Proyecto no se encontrado")
	}
}
