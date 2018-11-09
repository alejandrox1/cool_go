/*
This implements our image entrypoint. It:
- waits for SIGUSR1
- then execs (argv[1], argv[1:], env[:])

This allows us to perform other actions in the "node" container via
`docker exec` _before_ we actually "booted" the init and everything to go along
with it.

We can then send SIGUSR1 to this process to trigger starting the "actual"
entrypoint when we are done performing any provisions on the "node".
*/
package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

// yes this should be the c macro, but on linux in docker you're going to get this anyhow
// http://man7.org/linux/man-pages/man7/signal.7.html
// https://github.com/moby/moby/blob/562df8c2d6f48601c8d1df7256389569d25c0bf1/pkg/signal/signal_linux.go#L10
const sigrtmin = 34

func main() {
	// Prevent zombie processes since we will be PID1 for a while.
	// https://linux.die.net/man/2/waitpid
	signal.Ignore(syscall.SIGCHLD)

	// Grab the real entrypoint and command and args from our args.
	if len(os.Args) < 2 {
		log.Fatal("Not enoguh arguments to entrypoint!")
	}
	cmd, argv := os.Args[1], os.Args[1:]

	// Wait for SIGUSER1 (or exit on SIGRTMIN+3 to match systemd).
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGUSR1, syscall.Signal(sigrtmin+3))
	log.Println("Waiting for SIGUSR1...")

	sig := <-c
	if sig != syscall.SIGUSR1 {
		log.Printf("Exiting after signal %v != SIGUSR1\n", sig)
		return
	}

	// Now exec the real entrypoint, keeping the env.
	log.Printf("Received SIGUSR1, execing to %v %v\n", cmd, argv)
	syscall.Exec(cmd, argv, os.Environ())
}
