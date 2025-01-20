// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/openshift/api/operator/v1"
	operatorv1 "github.com/openshift/client-go/operator/applyconfigurations/operator/v1"
	typedoperatorv1 "github.com/openshift/client-go/operator/clientset/versioned/typed/operator/v1"
	gentype "k8s.io/client-go/gentype"
)

// fakeKubeControllerManagers implements KubeControllerManagerInterface
type fakeKubeControllerManagers struct {
	*gentype.FakeClientWithListAndApply[*v1.KubeControllerManager, *v1.KubeControllerManagerList, *operatorv1.KubeControllerManagerApplyConfiguration]
	Fake *FakeOperatorV1
}

func newFakeKubeControllerManagers(fake *FakeOperatorV1) typedoperatorv1.KubeControllerManagerInterface {
	return &fakeKubeControllerManagers{
		gentype.NewFakeClientWithListAndApply[*v1.KubeControllerManager, *v1.KubeControllerManagerList, *operatorv1.KubeControllerManagerApplyConfiguration](
			fake.Fake,
			"",
			v1.SchemeGroupVersion.WithResource("kubecontrollermanagers"),
			v1.SchemeGroupVersion.WithKind("KubeControllerManager"),
			func() *v1.KubeControllerManager { return &v1.KubeControllerManager{} },
			func() *v1.KubeControllerManagerList { return &v1.KubeControllerManagerList{} },
			func(dst, src *v1.KubeControllerManagerList) { dst.ListMeta = src.ListMeta },
			func(list *v1.KubeControllerManagerList) []*v1.KubeControllerManager {
				return gentype.ToPointerSlice(list.Items)
			},
			func(list *v1.KubeControllerManagerList, items []*v1.KubeControllerManager) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
