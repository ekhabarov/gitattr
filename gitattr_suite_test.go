package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGitattr(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gitattr Suite")
}
