// A generated module for ConfluenceUpdater functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"
	"dagger/confluence-updater/internal/dagger"
	"fmt"
)

type ConfluenceUpdater struct {
	// +private
	SrcDirectory *dagger.Directory
	// +private
	ApiUser string
	// +private
	ApiToken *dagger.Secret
	// +private
	ApiFqdn string
}

func New(
	// +optional
	// +defaultPath="/"
	srcDirectory *dagger.Directory,
	// +optional
	// +default=""
	user string,
	// +optional
	token *dagger.Secret,
	// +optional
	// +default=""
	fqdn string,
) *ConfluenceUpdater {
	c := &ConfluenceUpdater{
		SrcDirectory: srcDirectory,
		ApiUser:      user,
		ApiToken:     token,
		ApiFqdn:      fqdn,
	}
	return c
}

func (m *ConfluenceUpdater) Run(ctx context.Context) (string, error) {

	token, err := m.ApiToken.Plaintext(ctx)
	if err != nil {
		return "", err
	}

	if m.ApiUser == "" {
		return "", fmt.Errorf("Confluence user not set.")
	}

	if token == "" {
		return "", fmt.Errorf("Confluence secret not set.")
	}

	if m.ApiFqdn == "" {
		return "", fmt.Errorf("Confluence FQDN not set.")
	}

	args := []string{"-u", m.ApiUser, "-s", "dummy", "--fqdn", m.ApiFqdn, "-c", "/workdir/confluence-updater.yaml"}

	output, nil := dag.Container().
		From("ghcr.io/kerwood/confluence-updater:2.3.1").
		WithMountedDirectory("/workdir", m.SrcDirectory).
		WithExec(args, dagger.ContainerWithExecOpts{UseEntrypoint: true}).
		WithSecretVariable("CU_SECRET", m.ApiToken).
		Stdout(ctx)

	return output, nil
}

func (m *ConfluenceUpdater) User(user string) *ConfluenceUpdater {
	m.ApiUser = user
	return m
}

func (m *ConfluenceUpdater) Token(token *dagger.Secret) *ConfluenceUpdater {
	m.ApiToken = token
	return m
}

func (m *ConfluenceUpdater) Fqdn(fqdn string) *ConfluenceUpdater {
	m.ApiFqdn = fqdn
	return m
}

func (m *ConfluenceUpdater) SourceDirectory(srcDirectory *dagger.Directory) *ConfluenceUpdater {
	m.SrcDirectory = srcDirectory
	return m
}
