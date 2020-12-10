package main

import (
	"fmt"
	"time"
)

var opcion uint
var procesos = make([]Proceso, 5)
var PM = make(map[int]bool)
var PMC = make(map[int]bool)
var numero int = 0
var cstart = make(chan string)
var cmostrar = make(chan bool)

type Proceso struct {
	numero int
	tiempo int
}

func main() {
	for {
		time.Sleep(time.Second*1)
		fmt.Println("1) Agregar proceso")
		fmt.Println("2) Mostrar proceso")
		fmt.Println("3) Eliminar proceso")
		fmt.Println("4) Salír")
		fmt.Scan(&opcion)
		switch opcion {
		case 1:
			if numero < 5 {
				procesos[numero] = Proceso{
					numero: numero,
					tiempo: 0,
				}
				go procesos[numero].Start()
				PM[numero] = true
				numero = numero + 1
			} else {
				fmt.Println("No se pueden agregar mas de 5 procesos")
			}
		case 2:
			var input string
			for i := range procesos {
				if PM[i] == true && PMC[i] != true{
					go procesos[i].Mostrar()
					PMC[i] = true
					go func() {
						for input != "0" {
							cmostrar <- true
						}
					}()
				}
				if PM[i] == true{
					go func() {
						for input != "0" {
							cmostrar <- true
						}
					}()
				}
			}
			fmt.Scan(&input)
		case 3:
			Terminar()
		case 0:
			return
		}
	}
}

//Start comienza el proceso
func (p *Proceso) Start() {
	i := int(0)
	for {
		p.tiempo = i
		i = i + 1
		time.Sleep(time.Millisecond * 500)
	}
}

func (p *Proceso) Mostrar() {
	for {
		select {
		case canal := <-cmostrar:
			if canal == true {
				fmt.Printf("%d : %d\n", p.numero, p.tiempo)
				time.Sleep(time.Millisecond * 500)
			}
		}
	}
}

func Terminar() {
	var elproc int
	fmt.Println("Ingresa el proceso a eliminar")
	fmt.Scan(&elproc)
	PM[elproc] = false
	PMC[elproc] = false
	numero = 0
	procesos[elproc].numero = 0
	procesos[elproc].tiempo = 0
}
