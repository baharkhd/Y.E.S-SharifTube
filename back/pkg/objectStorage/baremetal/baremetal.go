package baremetal

import (
	scp "github.com/bramvdbogaerde/go-scp"
	"golang.org/x/crypto/ssh"
)

var publicKeyPath = "/home/kycilius/.ssh/id_rsa"

type BaremetalObjectStorageDriver struct {
	baseSir    string
	host       string
	sshSession *ssh.Session
	scpSession *scp.Client
}

func New(host, username, baseDir string) (*BaremetalObjectStorageDriver, error) {
	client, err := newSSHClient(host, username, publicKeyPath)
	if err != nil {
		return nil, err
	}
	scpSession, err3 := scp.NewClientBySSH(client)
	if err3 != nil {
		return nil, err3
	}

	sshSession, err2 := client.NewSession()
	if err2 != nil {
		return nil, err2
	}

	return &BaremetalObjectStorageDriver{
		host:       host,
		baseSir:    baseDir,
		sshSession: sshSession,
		scpSession: &scpSession,
	}, nil
}
