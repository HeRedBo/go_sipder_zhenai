package rpc

import "fmt"

type DemoService struct {}

type Args struct {
	 A, B int
}

func (DemoService) Div(args Args, result *float64) error {
	if args.B == 0 {
		return fmt.Errorf("division divisor is zero not allowed")
	}
	*result = float64(args.A ) / float64(args.B)
	return nil
}

// telnet
// {"method":"DemoService.Div","params":[{"A":10,"B":2}],"id":1}
// {"method":"DemoService.Div","params":[{"A":10,"B":0}],"id":1}