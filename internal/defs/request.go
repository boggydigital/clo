package defs

import "fmt"

type Request struct {
	Flags     []string
	Command   string
	Arguments map[string][]string
}

func (req *Request) Print() {
	if len(req.Flags) > 0 {
		fmt.Printf("Flags:%v\n", req.Flags)
	}
	fmt.Printf("Command:%v\n", req.Command)
	if len(req.Arguments) > 0 {
		fmt.Println("Arguments:")
		for k, v := range req.Arguments {
			fmt.Printf("- %v=%v\n", k, v)
		}
	}
}
