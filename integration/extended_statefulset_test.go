package integration_test

import (
	"fmt"

	essv1 "code.cloudfoundry.org/cf-operator/pkg/kube/apis/extendedstatefulset/v1alpha1"
	"code.cloudfoundry.org/cf-operator/testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ExtendedStatefulSet", func() {
	var (
		extendedStatefulSet      essv1.ExtendedStatefulSet
		wrongExtendedStatefulSet essv1.ExtendedStatefulSet
	)

	Context("when correctly setup", func() {
		BeforeEach(func() {
			essName := fmt.Sprintf("testess-%s", testing.RandString(5))
			extendedStatefulSet = env.DefaultExtendedStatefulSet(essName)

			wrongEssName := fmt.Sprintf("wrong-testess-%s", testing.RandString(5))
			wrongExtendedStatefulSet = env.WrongExtendedStatefulSet(wrongEssName)
		})

		It("should create a statefulset and eventually a pod", func() {
			// Create an ExtendedStatefulSet
			var ess *essv1.ExtendedStatefulSet
			ess, tearDown, err := env.CreateExtendedStatefulSet(env.Namespace, extendedStatefulSet)
			Expect(err).NotTo(HaveOccurred())
			Expect(ess).NotTo(Equal(nil))
			defer tearDown()

			// check for pod
			err = env.WaitForPods(env.Namespace, "testpod=yes")
			Expect(err).NotTo(HaveOccurred())

			// check for extendedStatefulSet available
			err = env.WaitForExtendedStatefulSetAvailable(env.Namespace, ess.GetName())
			Expect(err).NotTo(HaveOccurred())

			// check for extendedStatefulSet versions
			ess, err = env.GetExtendedStatefulSet(env.Namespace, ess.GetName())
			Expect(err).NotTo(HaveOccurred())
			Expect(ess.Status.Versions).To(Equal(map[int]bool{
				1: true,
			}))
		})

		It("should update a statefulset", func() {
			// Create an ExtendedStatefulSet
			var ess *essv1.ExtendedStatefulSet
			ess, tearDown, err := env.CreateExtendedStatefulSet(env.Namespace, extendedStatefulSet)
			Expect(err).NotTo(HaveOccurred())
			Expect(ess).NotTo(Equal(nil))
			defer tearDown()

			// check for pod
			err = env.WaitForPods(env.Namespace, "testpod=yes")
			Expect(err).NotTo(HaveOccurred())

			// Check for extendedStatefulSet available
			err = env.WaitForExtendedStatefulSetAvailable(env.Namespace, ess.GetName())
			Expect(err).NotTo(HaveOccurred())

			// check for extendedStatefulSet versions
			ess, err = env.GetExtendedStatefulSet(env.Namespace, ess.GetName())
			Expect(err).NotTo(HaveOccurred())
			Expect(ess.Status.Versions).To(Equal(map[int]bool{
				1: true,
			}))

			// Update the ExtendedStatefulSet
			ess.Spec.Template.Spec.Template.ObjectMeta.Labels["testpodupdated"] = "yes"
			essUpdated, tearDown, err := env.UpdateExtendedStatefulSet(env.Namespace, *ess)
			Expect(err).NotTo(HaveOccurred())
			Expect(essUpdated).NotTo(Equal(nil))
			defer tearDown()

			// check for pod
			err = env.WaitForPods(env.Namespace, "testpodupdated=yes")
			Expect(err).NotTo(HaveOccurred())

			// Check for extendedStatefulSet available
			err = env.WaitForExtendedStatefulSetAvailable(env.Namespace, ess.GetName())
			Expect(err).NotTo(HaveOccurred())

			// check for extendedStatefulSet versions
			ess, err = env.GetExtendedStatefulSet(env.Namespace, ess.GetName())
			Expect(err).NotTo(HaveOccurred())
			Expect(ess.Status.Versions).To(Equal(map[int]bool{
				2: true,
			}))

			// check that old statefulset is deleted
			ess, err = env.GetExtendedStatefulSet(env.Namespace, ess.GetName())
			Expect(ess.Status.Versions).To(Equal(map[int]bool{
				2: true,
			}))
		})

		It("should do nothing if nothing has changed", func() {
			// Create an ExtendedStatefulSet
			var ess *essv1.ExtendedStatefulSet

			ess, tearDown, err := env.CreateExtendedStatefulSet(env.Namespace, extendedStatefulSet)
			Expect(err).NotTo(HaveOccurred())
			Expect(ess).NotTo(Equal(nil))
			defer tearDown()

			// check for pod
			err = env.WaitForPods(env.Namespace, "testpod=yes")
			Expect(err).NotTo(HaveOccurred())

			// Check for extendedStatefulSet available
			err = env.WaitForExtendedStatefulSetAvailable(env.Namespace, ess.GetName())
			Expect(err).NotTo(HaveOccurred())

			// check for extendedStatefulSet versions
			ess, err = env.GetExtendedStatefulSet(env.Namespace, ess.GetName())
			Expect(err).NotTo(HaveOccurred())
			Expect(ess.Status.Versions).To(Equal(map[int]bool{
				1: true,
			}))

			// Update the ExtendedStatefulSet
			ess.Labels = map[string]string{
				"essupdated": "yes",
			}
			essUpdated, tearDown, err := env.UpdateExtendedStatefulSet(env.Namespace, *ess)
			Expect(err).NotTo(HaveOccurred())
			Expect(essUpdated).NotTo(Equal(nil))
			defer tearDown()

			// check for pod
			err = env.WaitForExtendedStatefulSets(env.Namespace, "essupdated=yes")
			Expect(err).NotTo(HaveOccurred())

			expectedMsg := fmt.Sprintf("StatefulSet '%s-v1' for ExtendedStatefulSet '%s/%s' has not changed, checking if any other changes are necessary.", extendedStatefulSet.Name, env.Namespace, extendedStatefulSet.Name)
			msgs := env.ObservedLogs.FilterMessage(expectedMsg)
			Expect(msgs.Len()).NotTo(Equal(0))
		})

		It("should keeps two versions if all are not running", func() {
			// Create an ExtendedStatefulSet
			ess, tearDown, err := env.CreateExtendedStatefulSet(env.Namespace, wrongExtendedStatefulSet)
			Expect(err).NotTo(HaveOccurred())
			Expect(ess).NotTo(Equal(nil))
			defer tearDown()

			// check for pod
			err = env.WaitForPods(env.Namespace, "wrongpod=yes")
			Expect(err).To(HaveOccurred())

			ess, err = env.GetExtendedStatefulSet(env.Namespace, ess.GetName())
			Expect(err).NotTo(HaveOccurred())
			Expect(ess).NotTo(Equal(nil))

			// Update the ExtendedStatefulSet
			ess.Spec.Template.Spec.Template.ObjectMeta.Labels["testpodupdated"] = "yes"
			essUpdated, tearDown, err := env.UpdateExtendedStatefulSet(env.Namespace, *ess)
			Expect(err).NotTo(HaveOccurred())
			Expect(essUpdated).NotTo(Equal(nil))
			defer tearDown()

			// check for pod
			err = env.WaitForPods(env.Namespace, "wrongpod=yes")
			Expect(err).To(HaveOccurred())

			// check that old statefulset is deleted
			ess, err = env.GetExtendedStatefulSet(env.Namespace, ess.GetName())
			Expect(err).NotTo(HaveOccurred())
			Expect(ess.Status.Versions).To(Equal(map[int]bool{
				1: false,
				2: false,
			}))
		})
	})
})
