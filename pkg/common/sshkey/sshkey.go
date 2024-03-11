package sshkey

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/PatrickLaabs/frigg/pkg/common/vars"
	"github.com/fatih/color"
	"golang.org/x/crypto/ssh"
	"os"
	"path/filepath"
)

func KeypairGen() {
	println(color.GreenString("Generating the ssh key pair"))

	homedir, err := os.UserHomeDir()
	if err != nil {
		println(color.RedString("Error on accessing the working directory: %v\n", err))
		return
	}
	friggDir := filepath.Join(homedir, vars.FriggDirName)

	savePublicFileTo := filepath.Join(friggDir, vars.PublickeyName)
	savePrivateFileTo := filepath.Join(friggDir, vars.PrivatekeyName)
	bitSize := 4096

	privateKey, err := generatePrivateKey(bitSize)
	if err != nil {
		println(color.RedString("error on private key generating: %v\n", err))
	}

	publicKeyBytes, err := generatePublicKey(&privateKey.PublicKey)
	if err != nil {
		println(color.RedString("error on public key generating: %v\n", err))
	}

	privateKeyBytes := encodePrivateKeyToPEM(privateKey)

	err = writeKeyToFile(privateKeyBytes, savePrivateFileTo)
	if err != nil {
		println(color.RedString("error on writing private key: %v\n", err))
	}

	err = writeKeyToFile(publicKeyBytes, savePublicFileTo)
	if err != nil {
		println(color.RedString("error on writing public key: %v\n", err))
	}
}

// generatePrivateKey creates a RSA Private Key of specified byte size
func generatePrivateKey(bitSize int) (*rsa.PrivateKey, error) {
	// Private Key generation
	privateKey, err := rsa.GenerateKey(rand.Reader, bitSize)
	if err != nil {
		return nil, err
	}

	// Validate Private Key
	err = privateKey.Validate()
	if err != nil {
		return nil, err
	}

	println(color.YellowString("Private Key generated"))
	return privateKey, nil
}

// encodePrivateKeyToPEM encodes Private Key from RSA to PEM format
func encodePrivateKeyToPEM(privateKey *rsa.PrivateKey) []byte {
	// Get ASN.1 DER format
	privDER := x509.MarshalPKCS1PrivateKey(privateKey)

	// pem.Block
	privBlock := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   privDER,
	}

	// Private key in PEM format
	privatePEM := pem.EncodeToMemory(&privBlock)

	return privatePEM
}

// generatePublicKey take a rsa.PublicKey and return bytes suitable for writing to .pub file
// returns in the format "ssh-rsa ..."
func generatePublicKey(privatekey *rsa.PublicKey) ([]byte, error) {
	publicRsaKey, err := ssh.NewPublicKey(privatekey)
	if err != nil {
		return nil, err
	}

	pubKeyBytes := ssh.MarshalAuthorizedKey(publicRsaKey)

	println(color.YellowString("Public key generated"))
	return pubKeyBytes, nil
}

// writePemToFile writes keys to a file
func writeKeyToFile(keyBytes []byte, saveFileTo string) error {
	err := os.WriteFile(saveFileTo, keyBytes, 0600)
	if err != nil {
		return err
	}

	println(color.YellowString("Key saved to: %v\n", saveFileTo))
	return nil
}
