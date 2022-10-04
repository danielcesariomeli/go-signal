package counter

import (
	"fmt"
	"time"
)

type Service struct {
	quit      chan bool
	isRunning bool
}

func NewService(quit chan bool, isRunnig bool) *Service {
	return &Service{
		quit:      quit,
		isRunning: isRunnig,
	}
}

func (s *Service) StartCounter() error {
	if s.isRunning {
		fmt.Println("Counter already is running")
		return nil
	}
	s.isRunning = true

	go func() {
		i := 1
		for {
			select {
			case <-s.quit:
				fmt.Println("The last number before stop is", i)
				return
			default:
				fmt.Println(i)
				i++
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()
	return nil
}

func (s *Service) StopCounter() error {
	s.isRunning = false
	go func() {
		s.quit <- false
	}()
	return nil
}
