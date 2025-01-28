package cmd

import "github.com/kodeyeen/omp"

var Commands = newCommandManager()

func init() {
	omp.ListenFunc(omp.EventTypePlayerCommandText, onPlayerCommandText)
}
