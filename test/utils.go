package test

import (
	//"bytes"

	"log"

	"golang.org/x/crypto/ssh"
)

func isConnect(ip string, port string, user string, pass string) bool {
	config := &ssh.ClientConfig{
		User:            user,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth: []ssh.AuthMethod{
			ssh.Password(pass),
		},
	}
	conn, err := ssh.Dial("tcp", ip+":"+port, config)
	if err != nil {
		log.Println(err)
		return false
	}
	defer conn.Close()

	session, err := conn.NewSession()
	if err != nil {
		log.Println(err)
		return false
	}
	defer session.Close()
	return true
}
