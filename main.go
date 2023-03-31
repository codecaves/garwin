package main

/*
 * garwin — go version of the 'arwin' utility:
 * (https://vividmachines.com/shellcode/arwin.c)
 *
 * Finds the absolute address of a function in a given DLL.
 * Happy shellcoding :)
 */

import (
	"fmt"
	"os"
	"syscall"
)

func isErr(err error) bool {
	return err != nil
}

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "usage: garwin <dll> <function>\n")
		os.Exit(-1)
	}

	if tgtDll, err := syscall.LoadLibrary(os.Args[1]); isErr(err) {
		fmt.Fprintf(os.Stderr, "The module '%s' was not found!\n%s", os.Args[1], err)
		os.Exit(-1)
	} else {
		defer syscall.FreeLibrary(tgtDll)
		tgtProcAddr, err := syscall.GetProcAddress(tgtDll, os.Args[2])

		if err != nil {
			fmt.Fprintf(os.Stderr, "The function '%s' was not found in %s!\n%s", os.Args[2], os.Args[1], err)
			os.Exit(-1)
		}

		fmt.Println("garwin - win32 address resolver - mortal - v.01")
		fmt.Printf("'%s' is located at 0x%08x in %s", os.Args[2], tgtProcAddr, os.Args[1])
	}
}