package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
)

func main() {
	circuit := flag.String("circuit", "circuit/proof.circom", "Circuit file")
	witness := flag.String("witness", "circuit/witness.wtns", "Witness file")
	zkey := flag.String("zkey", "circuit/proof.zkey", "ZKey file")
	output := flag.String("output", "proof.json", "Proof output file")
	flag.Parse()

	// Generate proof using snarkjs and the provided circuit, witness, and zkey
	cmd := exec.Command(
		"snarkjs", "groth16", "prove",
		*zkey,
		*witness,
		*output,
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("proof generation failed: %v", err)
	}
	log.Printf("Proof generated as %s for circuit %s\n", *output, *circuit)
}
