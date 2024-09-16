package object

import (
	"errors"
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("HashObject", func() {
	When("Hash Object call", func() {
		var (
			tigDir string
			data   = []byte("hello !")
		)

		BeforeEach(func() {
			t := GinkgoT()
			tigDir = filepath.Join(t.TempDir(), ".tig")
		})

		Context("DryRun = false", func() {
			It("Return key, error is nil, creating object file actually", func() {
				k, err := HashObject(HashObjectParam{
					DryRun: false,
					TigDir: tigDir,
					Type:   Blob,
					Data:   data,
				})

				Expect(err).NotTo(HaveOccurred())

				expectedKey := newKey(Blob, data)
				Expect(k).To(Equal(string(expectedKey)))

				stat, err := os.Stat(Key(k).Path(tigDir))
				Expect(err).NotTo(HaveOccurred())
				Expect(stat.IsDir()).NotTo(BeTrue())

			})
		})

		Context("DryRun = true", func() {
			It("Return key, error is nil, no object file actually", func() {
				k, err := HashObject(HashObjectParam{
					DryRun: true,
					TigDir: tigDir,
					Type:   Blob,
					Data:   data,
				})

				Expect(err).NotTo(HaveOccurred())

				expectedKey := newKey(Blob, data)
				Expect(k).To(Equal(string(expectedKey)))

				_, err = os.Stat(Key(k).Path(tigDir))
				Expect(err).To(HaveOccurred())
				Expect(errors.Is(err, os.ErrNotExist)).To(BeTrue())

			})
		})

	})
})
