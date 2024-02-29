package sshkey

import (
	"github.com/fatih/color"
	"os"
	"os/exec"
)

func KeypairGen() {
	println(color.GreenString("Generating the ssh key pair"))

	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("Error on accessing the working directory: %v\n", err))
		return
	}
	friggDir := homedir + "/" + friggDirName
	sshKeypairName := "frigg-sshkeypair_gen"
	keypairSavePath := friggDir + "/" + sshKeypairName

	// ssh-keygen -t rsa -C "frigg ssh keypar" -N "" -f frigg-sshkeypair
	cmd := exec.Command("ssh-keygen", "-t", "rsa",
		"-C", `frigg ssh keypar`, "-N", `""`, "-f", keypairSavePath,
	)

	//Capture the output of the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		println(color.RedString("error creating ssh keypair: %v\n", err))
		println(color.YellowString(string(output)))
		return
	}

	keyvalue, err := os.ReadFile(keypairSavePath)
	if err != nil {
		println(color.RedString("error reading ssh key file: %v\n", err))
	}

	err = os.WriteFile(keypairSavePath, keyvalue, 0775)
	if err != nil {
		println(color.RedString("Error on writing ssh key pairs: %v\n", err))
		return
	}
}
