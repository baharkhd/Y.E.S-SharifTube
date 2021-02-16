package baremetal

import (
	"golang.org/x/crypto/ssh"
	"io"
	"io/ioutil"
	"os"
)

func newSSHClient(host, username, publicKeyPath string) (*ssh.Client, error) {
	config := getSSHConfig(username, publicKeyPath)
	return ssh.Dial("tcp", host, config)

}

func getSSHConfig(username string, publicKeyPath string) *ssh.ClientConfig {
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			publicKey(publicKeyPath),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	return config
}

func publicKey(path string) ssh.AuthMethod {
	key, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		panic(err)
	}
	return ssh.PublicKeys(signer)
}

func (b *BaremetalObjectStorageDriver)Run(command string) error {

	session, err2 := b.client.NewSession()
	if err2 != nil {
		return err2
	}
	defer session.Close()

	sessStdOut, err := session.StdoutPipe()
	if err != nil {
		return err
	}
	go io.Copy(os.Stdout, sessStdOut)

	sessStderr, err := session.StderrPipe()
	if err != nil {
		panic(err)
	}
	go io.Copy(os.Stderr, sessStderr)

	err = session.Run(command)
	if err != nil {
		return err
	}
	return nil
}
