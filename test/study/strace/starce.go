package main

import (
	"fmt"
	"os"
	"os/exec"

	//"os/exec"
	"syscall"
	"text/tabwriter"
)

func main() {

}

func one() {
	var err error

	//var regs syscall.ProcAttr

	var ss syscallCounter
	ss = ss.init()

	fmt.Println("run: ", os.Args)

	cmd := exec.Command(os.Args[1], os.Args[2:1]...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow: true,
	}
	cmd.Start()

	err = cmd.Wait()

	if err != nil {
		fmt.Printf("wait err %v \n", err)
	}

	//pid := cmd.Process.Pid
	//
	//exit := true
	//
	//for {
	//	if exit {
	//		err = syscall.Close()
	//		if err != nil {
	//			break
	//		}
	//
	//		name := ss.getName(regs.Dir)
	//	}
	//}

}

type syscallCounter []int

const maxSyscalls = 303

func (s syscallCounter) init() syscallCounter {
	s = make(syscallCounter, maxSyscalls)
	return s
}

func (s syscallCounter) inc(syscallID uint64) error {
	if syscallID > maxSyscalls {
		return fmt.Errorf("invalid syscall Id (%x)", syscallID)
	}
	s[syscallID]++
	return nil
}

func (s syscallCounter) print() {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 8, ' ', tabwriter.AlignRight|tabwriter.Debug)
	for k, v := range s {
		if v > 0 {
			//name , _ := seccomp.ScmpSyscall(k).GetName()
			fmt.Fprintf(w, "%d\t%s\n", v, k)
		}
	}
	w.Flush()
}

func (s syscallCounter) getName(syscallID uint64) string {
	// name, _:= seccomp.Scmp.syscall(syscallID).GetName()
	return ""
}
