package object

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = ginkgo.Describe("UpdateIndex", func() {
	var (
		t          ginkgo.GinkgoTInterface
		woringCopy string
		tigDir     string
	)

	ginkgo.BeforeEach(func() {
		t = ginkgo.GinkgoT()
		woringCopy = t.TempDir()
		tigDir = filepath.Join(woringCopy, ".tig")
	})

	//
	// Caches
	//
	ginkgo.When("Using Caches field", func() {
		ginkgo.Context("If not found object hash in objects directory", func() {
			var (
				fileName   = "not-exists.txt"
				objectHash = "not_exists_hash_value"
			)

			ginkgo.It("return error", func() {
				param := UpdateIndexParam{
					Caches: []*Entry{
						{
							Mode:       os.ModePerm,
							File:       fileName,
							ObejctHash: objectHash,
						},
					},
					TigDir: tigDir,
				}

				err := UpdateIndex(param)
				Expect(err).To(HaveOccurred())
				Expect(errors.Is(err, os.ErrNotExist)).To(BeTrue())
			})
		})

		ginkgo.Context("Add = false", func() {
			ginkgo.Context("If object is exists in index file", func() {
				var (
					fileName     string
					originalHash string
				)
				ginkgo.BeforeEach(func() {})

				ginkgo.It("New object is updated into new index file", func() {})

			})

			ginkgo.Context("If object is not exists in index file", func() {
				ginkgo.It("return error", func() {

				})
			})

		})

		ginkgo.Context("Add = true", func() {
			ginkgo.Context("If object is not exists in index file", func() {
				var (
					fileName     string
					originalHash string
				)
				ginkgo.BeforeEach(func() {})

				ginkgo.It("Add new objects in new index file", func() {

				})
			})

		})
	})

	//
	// Files
	//
	ginkgo.When("Using Files field", func() {
		ginkgo.Context("If not exists file", func() {
			var (
				fileName = "not-exists.txt"
			)

			ginkgo.It("return error", func() {
				param := UpdateIndexParam{
					Files:  []string{fileName},
					TigDir: tigDir,
				}

				err := UpdateIndex(param)
				Expect(err).To(HaveOccurred())
				Expect(errors.Is(err, os.ErrNotExist)).To(BeTrue())
			})
		})

		ginkgo.Context("Add = false", func() {
			var (
				fileName     string
				originalHash string
			)
			ginkgo.BeforeEach(func() {

			})

			ginkgo.Context("If object is exists in index file", func() {
				ginkgo.It("New object is updated into new index file")
			})

			ginkgo.Context("If object is not exists in index file", func() {
				ginkgo.It("return error")

			})

		})

		ginkgo.Context("Add = true", func() {
			ginkgo.Context("If object is not exists in index file", func() {
				ginkgo.It("Add new objects in new index file")
			})

		})

	})
})
