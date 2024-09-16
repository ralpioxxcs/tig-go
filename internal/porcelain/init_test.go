package porcelain

import (
	"os"
	"path/filepath"
	"tig/internal/config"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const (
	baseDir        = ".git"
	headFileName   = "HEAD"
	objectsDirName = "objects"
	refsDirName    = "refs"
)

var _ = Describe("Init", func() {
	When("When call Init func", func() {
		var tempDir string
		BeforeEach(func() {
			tempDir = GinkgoT().TempDir()
		})

		It("config, HEAD, ojbect, refs file is created", func() {
			cfg := config.Config{
				User: config.User{
					Name:  "jaehong",
					Email: "jaehong@mail.cc",
				},
			}

			param := InitParam{
				WorkingCopy: tempDir,
				Config:      cfg,
			}

			Expect(Init(param)).Should(BeNil())

			createdCfg, err := config.ReadConfigFile(filepath.Join(tempDir, baseDir))
			Expect(err).To(BeNil())
			Expect(createdCfg).Should(Equal(cfg))

			f, err := os.Stat(filepath.Join(tempDir, baseDir, headFileName))
			Expect(err).To(BeNil())
			Expect(f.IsDir()).NotTo(BeTrue())

			d, err := os.Stat(filepath.Join(tempDir, baseDir, objectsDirName))
			Expect(err).To(BeNil())
			Expect(d.IsDir()).NotTo(BeTrue())

			d, err = os.Stat(filepath.Join(tempDir, baseDir, refsDirName))
			Expect(err).To(BeNil())
			Expect(d.IsDir()).NotTo(BeTrue())

		})

	})

})
