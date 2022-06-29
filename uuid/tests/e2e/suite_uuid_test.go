package e2e

import (
	mocks "github.com/douyu/jupiter-examples/uuid/gen/mocks/redis"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
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
