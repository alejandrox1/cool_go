// Package build implements functionality to build the kind images.
package build

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/golang/glog"
)

// NodeImageBuildContext is used to build the kind node image, and contains
// buid configuration.
type NodeImageBuildContext struct {
	SourceDir string
	ImageTag  string
	GoCmd     string
	Arch      string
}

// NewNodeImageBuildContext creates a NodeImageBuildContext with default
// configuration.
func NewNodeImageBuildContext() *NodeImageBuildContext {
	return &NodeImageBuildContext{
		ImageTag: "kind-node",
		GoCmd:    "go",
		Arch:     "amd64",
	}
}

func runCmd(cmd *exec.Cmd) error {
	glog.Infof("Running: %v %v", cmd.Path, cmd.Args)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func copyDir(src, dst string) error {
	src = filepath.Clean(src) + string(filepath.Separator) + "."
	dst = filepath.Clean(dst)
	cmd := exec.Command("cp", "-r", src, dst)
	return runCmd(cmd)
}

func (c *NodeImageBuildContext) Build() error {
	// Create a tmp dir to build in.
	dir, err := ioutil.TempDir("", "kind-build")
	if err != nil {
		return err
	}
	defer os.RemoveAll(dir)

	// Populate with image sources.
	err = copyDir(c.SourceDir, dir)
	if err != nil {
		glog.Errorf("failed to copy sources to build directory: %v", dir)
		return err
	}

	glog.Infof("Building node image in: %s", dir)

	// Build entrypoint binary.
	glog.Info("Building entrypoint binary...")
	entrypointSrc := filepath.Join(dir, "entrypoint", "main.go")
	entrypointDest := filepath.Join(dir, "entrypoint", "entrypoint")
	cmd := exec.Command(c.GoCmd, "build", "-o", entrypointDest, entrypointSrc)
	cmd.Env = []string{"GOOS=linux", "GOARCH=" + c.Arch}
	if err = cmd.Run(); err != nil {
		glog.Errorf("Entrypoint build failed! %v", err)
		return err
	}
	glog.Info("Entrypoint build completed.")

	glog.Info("Starting Docker build...")
	err = runCmd(exec.Command("docker", "build", "-t", c.ImageTag, dir))
	if err != nil {
		glog.Errorf("Docker build failed! %v", err)
		return err
	}
	glog.Info("Docker build completed.")

	return nil
}
