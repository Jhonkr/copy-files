package main

/*
#include "logger.h"
*/
import "C"

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {

	oldDir := os.Args[1]
	newDir := os.Args[2]

	initialText := []string{"[LOG] ", fmt.Sprintf("Copying from %s to %s", oldDir, newDir)}
	loggerC(initialText)

	logs := make(chan []string)
	go copyDir(oldDir, newDir, logs)

	loggerC(<-logs)

}

func copyDir(src string, dest string, logs chan<- []string) {

	// Command CP
	cmd := exec.Command("cp", "-rip", src, dest)
	_, err := cmd.CombinedOutput()

	if err != nil {
		logs <- []string{"ERROR", err.Error()}
	}
	logs <- []string{"LOG: ", "SUCCESS"}
}

func loggerC(text []string){
	C.writeLogger(C.CString(text[0]), C.CString(text[1]))
}
