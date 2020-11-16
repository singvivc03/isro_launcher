package launcher

import (
	"fmt"
	"github/singvivc03/isro_launcher/src/counter"
)

type launcher struct {
	name string
}

// Launcher ...
type Launcher interface {
	Launch(int, int) <-chan int
}

// New ...
func New(name string) Launcher {
	return &launcher{name: name}
}

func (l *launcher) Launch(count int, groupof int) <-chan int {
	signal := make(chan int)
	go func() {
		fmt.Printf("Launching satellites from launch pad %s in...\n", l.name)
		for sat := count; sat < count+groupof; sat++ {
			fmt.Printf("Launching SAT%d ", sat)
			counter.Count()
			fmt.Printf("SAT%d launched successfully\n", sat)
		}
		fmt.Println()
		signal <- count + groupof
	}()
	return signal
}
