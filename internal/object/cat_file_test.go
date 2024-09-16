package object

import (
	"path/filepath"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("CatFile", func() {
	When("CatFile is called with pretty-print", func() {

		var (
			t      GinkgoTInterface
			tigDir string
			data   = []byte("hello !")
		)

		BeforeEach(func() {
			t = GinkgoT()
			tigDir = filepath.Join(t.TempDir(), ".tig")
		})

		It("", func() {
			hash, err := HashObject(HashObjectParam{
				DryRun: false,
				TigDir: tigDir,
				Type:   Blob,
				Data:   data,
			})
			if err != nil {
				t.Fatal(err)
			}

			c, err := CatFile(CatFileParam{
				TigDir:        tigDir,
				OperationType: CatFileOperationTypePrettyPrint,
				ObjectHash:    hash,
			})

			Expect(err).ShouldNot(HaveOccurred())
			Expect(c).To(Equal(string(data)))
		})

	})
})
