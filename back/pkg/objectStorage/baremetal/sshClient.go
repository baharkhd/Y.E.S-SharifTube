package baremetal

import (
	"golang.org/x/crypto/ssh"
	"io"
	"io/ioutil"
	"os"
)

func newSSHClient(host, username, publicKeyPath string) (*ssh.Client, error) {
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			publicKey(publicKeyPath),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	return ssh.Dial("tcp", host, config)

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
	sessStdOut, err := b.sshSession.StdoutPipe()
	if err != nil {
		return err
	}
	go io.Copy(os.Stdout, sessStdOut)

	sessStderr, err := b.sshSession.StderrPipe()
	if err != nil {
		panic(err)
	}
	go io.Copy(os.Stderr, sessStderr)

	err = b.sshSession.Run(command)
	if err != nil {
		return err
	}
	return nil
}
