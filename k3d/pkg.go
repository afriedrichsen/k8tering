package k3d

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// K3D executes a 'k3d xxx' command and returns the literal result.
// The first two parameters influence output behavior: withstderr adds stderr output,
// of the k3d invocation and verbose gives additional details. For example:
//
// K3D(false, false, "~/bin/kubectl", "get", "--namespace=foo", "pods", "--output=yaml")
func K3D(withstderr, verbose bool, kubectlbin, cmd string, args ...string) (result string, err error) {
	if kubectlbin == "" {
		bin, err := shellout(withstderr, false, "which", "k3d")
		if err != nil {
			if verbose {
				perr("Can't find k3d", err)
			}
			return "", err
		}
		kubectlbin = bin
	}
	all := append([]string{cmd}, args...)
	result, err = shellout(withstderr, verbose, kubectlbin, all...)
	if err != nil {
		if verbose {
			perr("Can't maintain cluster using k3d: ", err)
		}
		return "", err
	}
	return result, nil
}

// shellout shells out to execute a command with a variable number of arguments
// and returns the literal result, optionally including stderr output.
func shellout(withstderr, verbose bool, cmd string, args ...string) (result string, err error) {
	var out bytes.Buffer
	if verbose {
		pinfo(cmd + " " + strings.Join(args, " "))
	}
	c := exec.Command(cmd, args...)
	c.Env = os.Environ()
	if withstderr {
		c.Stderr = os.Stderr
	}
	c.Stdout = &out
	err = c.Run()
	if err != nil {
		return "", err
	}
	result = strings.TrimSpace(out.String())
	return result, nil
}

// pinfo writes msg in light blue to stderr
// see also https://misc.flogisoft.com/bash/tip_colors_and_formatting
func pinfo(msg string) {
	_, _ = fmt.Fprintf(os.Stderr, "\x1b[94m%v\x1b[0m\n", msg)
}

// perr writes message and error in light red to stderr
// see also https://misc.flogisoft.com/bash/tip_colors_and_formatting
func perr(msg string, err error) {
	_, _ = fmt.Fprintf(os.Stderr, "%v: \x1b[91m%v\x1b[0m\n", msg, err)
}
