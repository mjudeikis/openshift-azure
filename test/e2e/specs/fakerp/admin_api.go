package fakerp

import (
	"context"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/Azure/go-autorest/autorest/to"

	"github.com/openshift/openshift-azure/test/clients/azure"
)

var _ = Describe("Validate AdimAPI field readability [AdminAPI][Fake]", func() {
	var (
		cli *azure.Client
	)

	BeforeEach(func() {
		var err error
		cli, err = azure.NewClientFromEnvironment(false)
		Expect(err).NotTo(HaveOccurred())
	})

	It("should not be possible to update pluginVersion using AdminAPI", func() {
		By("Reading the cluster state")
		before, err := cli.OpenShiftManagedClustersAdmin.Get(context.Background(), os.Getenv("RESOURCEGROUP"), os.Getenv("RESOURCEGROUP"))
		Expect(err).NotTo(HaveOccurred())
		Expect(before).NotTo(BeNil())

		By("Updating the cluster state")
		before.Config.PluginVersion = to.StringPtr("v0.1")
		update, err := cli.OpenShiftManagedClustersAdmin.CreateOrUpdate(context.Background(), os.Getenv("RESOURCEGROUP"), os.Getenv("RESOURCEGROUP"), before)
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(HavePrefix("400: 400 Bad Request"))
		Expect(update).To(BeNil())

		By("Reading the cluster state")
		after, err := cli.OpenShiftManagedClustersAdmin.Get(context.Background(), os.Getenv("RESOURCEGROUP"), os.Getenv("RESOURCEGROUP"))
		Expect(err).NotTo(HaveOccurred())
		Expect(after).NotTo(BeNil())
		Expect(*after.Config.PluginVersion).To(Equal(*before.Config.PluginVersion))
	})
})
