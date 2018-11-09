// Package cluster implements kind local cluster management.
package cluster

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

// ClusterLabelKey is applied to each "node" docker container for
// identification.
const ClusterLabelKey = "io.k8s.kind-cluster"

// Config contains cluster configurations.
type Config struct {
	clusterID string
}

// NewConfig returns a new Config with default settings. If clusterID is "",
// then the clusterID will be set to "1".
func NewConfig(clusterID string) Config {
	if clusterID == "" {
		clusterID = "1"
	}
	return Config{clusterID: clusterID}
}

// Context contains a Config and is used to create/manage the
// kubernetes-in-docker cluster.
type Context struct {
	Config
}

func (c *Config) clusterLabel() string {
	return fmt.Sprintf("%s=%s", ClusterLabelKey, c.clusterID)
}

// NewContext returns a new cluster management context with a config.
func NewContext(config Config) *Context {
	return &Context{Config: config}
}

// Create provisions and starts a kubernetes-in-docker cluster.
func (c *Context) Create() error {
	// TODO: provision multiple nodes.
	return c.provisionNode()
}

func (c *Context) provisionNode() error {
	nodeName := "kind-" + c.Config.clusterID + "-master"

	// Create the "node" container.
	if err := c.createNode(nodeName); err != nil {
		return err
	}

	// systemd in a container should have a read-only /sys
	// https://www.freedesktop.org/wiki/Software/systemd/ContainerInterface/
	// however, we need other things from `docker run --privileged`. This flag
	// also happens to make /sys rw, among other things.
	if err := c.runOnNode(nodeName, []string{
		"mount", "-o", "remount,ro", "/sys",
	}); err != nil {
		// TODO: logging here.
		c.deleteNodes(nodeName)
		return err
	}

	// TODO: Insert other provisioning here (e.g., enabling/disabling units,
	// installing kube).

	// Signal the node to boot into systemd.
	if err := c.actuallyStartNode(nodeName); err != nil {
		// TODO: logging here.
		c.deleteNodes(nodeName)
		return err
	}

	return nil
}

// createNode `docker run`'s the node image. Note that due to
// images/node/entrypoint being the entrypoint, this container will effectively
// be paused until we call actuallyStartNode().
func (c *Context) createNode(name string) error {
	// TODO: use logging and derive the run flags from the config.
	cmd := exec.Command("docker", "run")
	cmd.Args = append(cmd.Args,
		"-d", // Run the container detached.
		"-t", // Need a Pseudo-TTY for systemd logs.
		// Running containers in a container requires privileged.
		// Note: we could try to replicate this with --cap-add, and use less
		// privileges, but this flag also changes some mounts that are
		// necessary, including some that Docker would otherwise do by default.
		"--privileged",
		"--security-opt", "seccomp=unconfined", // Ignore seccomp.
		"--tmpfs", "/tmp", // Various things depend on a working /tmp.
		"--tmpfs", "/run", // systemd wants a writable /run.
		// Docker in docker needs this, so as not to stack overlays.
		"--tmpfs", "/var/lib/docker:exec",
		//"-v", "/sys/fs/cgroup:/sys/fs/cgroup:ro",
		// Some k8s things want /lib/modules.
		"-v", "/lib/modules:/lib/modules:ro",
		"--hostname", name, // Make hostname match container name.
		"--name", name, // Set the container name.
		// Label the node with the clster ID.
		"--label", c.Config.clusterLabel(),
		"kind-node",
	)
	// TODO: collect output.
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

// runOnNode execs command on the named node.
func (c *Context) runOnNode(nameOrID string, command []string) error {
	// TODO: use config and logging.
	cmd := exec.Command("docker", "exec")
	cmd.Args = append(cmd.Args,
		"-t",           // Use a TTY so we can get output.
		"--privileged", // Run with privileges so we can remount.
		nameOrID,
	)
	cmd.Args = append(cmd.Args, command...)
	// TODO: collect output instead of connecting.
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

// actuallyStartNode signals our entrypoint (images/node/entrypoint) to boot.
func (c *Context) actuallyStartNode(name string) error {
	// TODO: use config and logging.
	cmd := exec.Command("docker", "kill")
	cmd.Args = append(cmd.Args,
		"-s", "SIGUSR1",
		name,
	)
	// TODO: collect output instead of conecting these.
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func (c *Context) deleteNodes(names ...string) error {
	cmd := exec.Command("docker", "rm")
	cmd.Args = append(cmd.Args, "-f")
	cmd.Args = append(cmd.Args, names...)
	return cmd.Run()
}

// Delete tears down a kubernetes-in-docker cluster.
func (c *Context) Delete() error {
	// TODO: find and delete nodes.
	nodes, err := c.ListNodes(true)
	if err != nil {
		return fmt.Errorf("Error listing nodes: %v", err)
	}
	return c.deleteNodes(nodes...)
}

// ListNodes returns a list of container IDs for the "nodes" in the cluster.
func (c *Context) ListNodes(alsoStopped bool) ([]string, error) {
	cmd := exec.Command("docker", "ps")
	cmd.Args = append(cmd.Args,
		"-q", // Quiet output for parsing.
		"--filter", "label="+c.Config.clusterLabel(),
	)
	// Optionally show nodes that are stopped.
	if alsoStopped {
		cmd.Args = append(cmd.Args, "-a")
	}

	return cmdLines(cmd)
}

func cmdLines(cmd *exec.Cmd) ([]string, error) {
	var buf bytes.Buffer
	cmd.Stderr = &buf
	cmd.Stdout = &buf

	if err := cmd.Run(); err != nil {
		return nil, err
	}

	var lines []string
	scanner := bufio.NewScanner(&buf)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}
