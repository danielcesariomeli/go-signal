package counter

import (
	"fmt"
	"time"
)

type Service struct {
	quit chan bool
}

func NewService(quit chan bool) *Service {
	return &Service{
		quit: quit,
	}
}

func (s *Service) StartCounter() error {
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
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()
	return nil
}

func (s *Service) StopCounter() error {
	s.quit <- true
	return nil
}
