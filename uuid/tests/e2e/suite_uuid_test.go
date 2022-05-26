package e2e

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	mocks "uuid/gen/mocks/redis"
)

var _ = Describe("uuidService", func() {

	mockRedis := &mocks.RedisInterface{}

	uuidService := CreateUuidService(mockRedis)

	Context("List", func() {
		It("normal case", func() {
			Expect(uuidService).ShouldNot(BeNil())
		})
	})
})
