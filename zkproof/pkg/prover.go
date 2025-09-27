package zkproof

import "os/exec"

func Prove(circuit, witness string) error {
	cmd := exec.Command("snarkjs", "groth16", "prove", circuit, witness)
	return cmd.Run()
}
