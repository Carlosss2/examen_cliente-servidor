package main

import (
	servidorprincipal "examenC1/servidor-principal"
	servidorreplica "examenC1/servidor-replica"
)

func main() {
	
	go servidorprincipal.Run()
	go servidorreplica.Run()
	select {}
}