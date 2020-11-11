package main

import "fmt"
import "time"


type Process struct {
	id        int
	value     uint64
	mustPrint bool
}

func (p *Process) start(c chan Process) {
	go printer(c, p.mustPrint)
	for {
		p.value++
		c <- *p
		time.Sleep(time.Millisecond * 500)
	}
}

func (p *Process) show(mustShow bool) {
	p.mustPrint = mustShow
}

func printer(c chan Process, show bool) {
	for {
		msg := <- c
		if msg.mustPrint == true {
			fmt.Printf("\nid %d: %d", msg.id, msg.value)
		}
	}
}

func create(id int, c chan Process) *Process {
	p := Process{id:id, value:0, mustPrint: false} 
	go p.start(c)
	fmt.Printf("Proceso %d creado\n", id)
	return &p
}

func show(processes []*Process, mustShow bool) {
	for i := 0; i < len(processes); i++ {
		processes[i].show(mustShow)
	}
}

func delete(processes []*Process) []*Process {
	var input int
	fmt.Print("Id del proceso a eliminar: ")
	fmt.Scan(&input)
	for i := 0; i < len(processes); i++  {
		if input == processes[i].id {
			copy(processes[i:], processes[i+1:])
			processes[len(processes)-1] = nil
			processes = processes[:len(processes)-1]
		}
	}
	return processes
}


func main() {
	op := 1
	var id = 1
	var processes []*Process
	c := make(chan Process)
	mustShow := false
	for op != 0 {
		fmt.Println("1) Agregar proceso")
		fmt.Println("2) Mostrar procesos")
		fmt.Println("3) Terminar proceso")
		fmt.Println("0) Salir")
		fmt.Scan(&op)
		switch op {
		case 0:
			return
		case 1:
			p := create(id, c)
			processes = append(processes, p)
			id++
		case 2:
			mustShow = !mustShow
			show(processes, mustShow)
		case 3:
			processes = delete(processes)
		}
	}
}