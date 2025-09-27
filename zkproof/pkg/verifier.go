package zkproof

import "os/exec"

func Verify(vkey, proof, public string) error {
	cmd := exec.Command("snarkjs", "groth16", "verify", vkey, public, proof)
	return cmd.Run()
}
