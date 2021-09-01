package yagu

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/kevinburke/ssh_config"
)

type SSHMachine struct {
	User string
	Keys *ssh.PublicKeys
}

func SSHCredentials(machine string) (SSHMachine, error) {
	fmt.Println("ssh.CredsLookup")
	// try to ssh config file
	pk := ""
	pka, err := ssh_config.GetAllStrict(machine, "IdentityFile")
	fmt.Println(pka, err)
	if err != nil {
		// try to load id_rsa.pub
		hdir, err := os.UserHomeDir()
		if err != nil {
			// no home dir?
			return SSHMachine{}, err
		}

		// set pk file name to git's expected default, often the one uploaded per GitHub's docs
		pk = filepath.Join(hdir, ".ssh", "id_rsa.pub")
	}


	if strings.HasPrefix(pk, "~") {
		if hdir, err := os.UserHomeDir(); err == nil {
			pk = strings.Replace(pk, "~", hdir, 1)
		}
	}
	usr := ssh_config.Get(machine, "User")
	if usr == "" {
		usr = "git"
	}

	fmt.Println("  ", usr, pk)

	pks, err := ssh.NewPublicKeysFromFile(usr, pk, "")
	if err != nil {
		return SSHMachine{}, err
	}

	return SSHMachine{usr, pks}, nil
}
