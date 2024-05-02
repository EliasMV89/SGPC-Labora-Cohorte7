package main

import (
	"fmt"
	"os"
	"strings"
)

// 5. SGPC - Sistema de gestión de proyectos creativos

type Proyecto struct {
	Nombre      string
	Integrantes []string
	Inicio      string
	Hitos       []string
	FechasHito  []string
	Fin         string
}

// Función para crear un Proyecto
func crearProyecto(listaProyecto map[string]*Proyecto, nombre string, integrantes []string, inicio string, hitos []string, fechasHito []string, fin string) {
	nuevoProyecto := &Proyecto{
		Nombre:      nombre,
		Integrantes: integrantes,
		Inicio:      inicio,
		Hitos:       hitos,
		FechasHito:  fechasHito,
		Fin:         fin,
	}
	listaProyecto[nombre] = nuevoProyecto
}

// Función para mostrar todos los Proyectos
func mostrarProyecto(listaProyecto map[string]*Proyecto) {
	fmt.Println("Todos los proyectos:")
	for _, proyecto := range listaProyecto {
		fmt.Printf("Nombre: %s\nIntegrantes: %v\nInicio: %s\nHitos: %v\nFechas de hito: %v\nFin: %s\n\n", proyecto.Nombre, proyecto.Integrantes, proyecto.Inicio, proyecto.Hitos, proyecto.FechasHito, proyecto.Fin)
	}
}

// Función para buscar proyecto por nombre
func buscarPorProyecto(listaProyecto map[string]*Proyecto, nombre string) *Proyecto {
	if proyecto, existe := listaProyecto[nombre]; existe {
		return proyecto
	}
	return nil
}

// Función para buscar proyecto por integrante
func buscarPorIntegrante(listaProyecto map[string]*Proyecto, integranteBuscado string) []*Proyecto {
	var proyectosEncontrados []*Proyecto
	for _, proyecto := range listaProyecto {
		for _, integrante := range proyecto.Integrantes {
			if integrante == integranteBuscado {
				proyectosEncontrados = append(proyectosEncontrados, proyecto)
				break // No es necesario seguir buscando en este proyecto
			}
		}
	}
	return proyectosEncontrados
}

// Función para agregar un nuevo hito a un proyecto existente
func agregarHito(listaProyecto map[string]*Proyecto, nombre string, nuevoHito string, nuevaFechaHito string) {
	if proyecto, encontrado := listaProyecto[nombre]; encontrado {
		proyecto.Hitos = append(proyecto.Hitos, nuevoHito)
		proyecto.FechasHito = append(proyecto.FechasHito, nuevaFechaHito)
	}
}

// Función para eliminar un proyecto
func eliminarProyecto(listaProyecto map[string]*Proyecto, nombre string) {
	if _, encontrado := listaProyecto[nombre]; encontrado {
		delete(listaProyecto, nombre)
		fmt.Printf("Proyecto '%s' eliminado.\n", nombre)
	} else {
		fmt.Printf("Proyecto '%s' no encontrado.\n", nombre)
	}
}

func main() {
	listaProyecto := make(map[string]*Proyecto)

	// Opciones del sistema
	for {
		fmt.Printf("Bienvenido!")
		fmt.Printf("******************************************************\n")
		fmt.Printf("Elige una opcion.\n")

		fmt.Printf("******************************************************\n")
		fmt.Println("1.Crear un nuevo proyecto")
		fmt.Println("2.Ver todos los proyectos")
		fmt.Println("3.Buscar proyectos por nombre")
		fmt.Println("4.Buscar proyectos por integrante")
		fmt.Println("5.Agregar un nuevo hito a un proyecto existente")
		fmt.Println("6.Eliminar proyecto")
		fmt.Println("7.Salir del sistema")
		fmt.Printf("******************************************************\n")

		var choice int
		fmt.Print("Ingrese su opción: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Print("Ingrese el nombre del proyecto: ")
			var nombre string
			fmt.Scanln(&nombre)
			var integrante []string
			var nombreIntegrante string
			for {
				fmt.Print("Ingrese el nombre del integrante o 'x' para continuar: ")
				fmt.Scanln(&nombreIntegrante)
				if nombreIntegrante == "x" {
					break
				}
				integrante = append(integrante, nombreIntegrante)
			}
			fmt.Print("Ingrese la fecha de inicio del proyecto: ")
			var inicio string
			fmt.Scanln(&inicio)
			// Ingreso por consola para los hitos del proyecto con validacion para evitar que se ingrese vacios
			var hito []string
			var nombreHito string
			for {
				fmt.Print("Ingrese un hito del proyecto: ")
				fmt.Scanln(&nombreHito)
				if strings.TrimSpace(nombreHito) != "" {
					hito = append(hito, nombreHito)
					break
				} else {
					fmt.Printf("El hito no puede estar vacío.\n")
				}
			}
			// Ingreso por consola para la fecha de los hitos con validacion para evitar que se ingrese vacios
			var fecha string
			var fechaHito []string
			for {
				fmt.Print("Ingrese la fecha del hito del proyecto: ")
				fmt.Scanln(&fecha)
				if strings.TrimSpace(fecha) != "" {
					fechaHito = append(fechaHito, fecha)
					break
				} else {
					fmt.Printf("La fecha del hito no puede estar vacío.\n")
				}
			}
			fmt.Print("Ingrese la fecha de finalizacion del proyecto: ")
			var fin string
			fmt.Scanln(&fin)
			crearProyecto(listaProyecto, nombre, integrante, inicio, hito, fechaHito, fin)
		case 2:
			mostrarProyecto(listaProyecto)
		case 3:
			fmt.Print("Ingrese el nombre del proyecto que desea buscar: ")
			var nombreProyecto string
			fmt.Scanln(&nombreProyecto)
			proyecto := buscarPorProyecto(listaProyecto, nombreProyecto)
			if proyecto != nil {
				fmt.Printf("Proyecto encontrado: %+v\n", proyecto)
			} else {
				fmt.Println("No Proyecto encontrado")
			}
		case 4:
			fmt.Print("Ingrese el nombre del integrante que desea buscar: ")
			var integranteBuscado string
			fmt.Scanln(&integranteBuscado)
			proyectos := buscarPorIntegrante(listaProyecto, integranteBuscado)
			if len(proyectos) > 0 {
				fmt.Printf("Proyectos encontrados para el integrante %s:\n", integranteBuscado)
				for _, proyecto := range proyectos {
					fmt.Printf("- %s\n", proyecto.Nombre)
				}
			} else {
				fmt.Printf("No se encontraron proyectos para el integrante %s\n", integranteBuscado)
			}
		case 5:
			fmt.Print("Ingrese el nombre del proyecto que desea agregar un nuevo hito: ")
			var nombre string
			fmt.Scanln(&nombre)
			proyecto := buscarPorProyecto(listaProyecto, nombre)
			if proyecto != nil {
				// Ingreso por consola para los hitos del proyecto
				var hito string
				for {
					fmt.Print("Ingrese un hito del proyecto: ")
					fmt.Scanln(&hito)
					if strings.TrimSpace(hito) != "" {
						break
					} else {
						fmt.Printf("El hito no puede estar vacío.\n")
					}
				}
				// Ingreso por consola para la fecha de los hitos
				var fecha string
				for {
					fmt.Print("Ingrese la fecha del hito del proyecto: ")
					fmt.Scanln(&fecha)
					if strings.TrimSpace(fecha) != "" {
						break
					} else {
						fmt.Printf("La fecha del hito no puede estar vacío.\n")
					}
				}
				agregarHito(listaProyecto, nombre, hito, fecha)
			} else {
				fmt.Println("No Proyecto encontrado")
			}
		case 6:
			fmt.Print("Ingrese el nombre del proyecto que desea eliminar: ")
			var nombre string
			fmt.Scanln(&nombre)
			proyecto := buscarPorProyecto(listaProyecto, nombre)
			if proyecto != nil {
				eliminarProyecto(listaProyecto, nombre)
			} else {
				fmt.Println("No Proyecto encontrado")
			}
		case 7:
			fmt.Println("salir")
			os.Exit(0)
		default:
			fmt.Println("Opción inválida. Por favor, intente de nuevo.")
		}
	}
}
