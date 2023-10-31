/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/JeronimoMendes/spendr/cmd"
	_ "github.com/JeronimoMendes/spendr/cmd/category"
	_ "github.com/JeronimoMendes/spendr/cmd/expense"
)

func main() {
	cmd.Execute()
}
