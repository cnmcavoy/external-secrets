// +build !ignore_autogenerated

/*
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSSMAuth) DeepCopyInto(out *AWSSMAuth) {
	*out = *in
	in.SecretRef.DeepCopyInto(&out.SecretRef)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSSMAuth.
func (in *AWSSMAuth) DeepCopy() *AWSSMAuth {
	if in == nil {
		return nil
	}
	out := new(AWSSMAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSSMAuthSecretRef) DeepCopyInto(out *AWSSMAuthSecretRef) {
	*out = *in
	in.AccessKeyID.DeepCopyInto(&out.AccessKeyID)
	in.SecretAccessKey.DeepCopyInto(&out.SecretAccessKey)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSSMAuthSecretRef.
func (in *AWSSMAuthSecretRef) DeepCopy() *AWSSMAuthSecretRef {
	if in == nil {
		return nil
	}
	out := new(AWSSMAuthSecretRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSSMProvider) DeepCopyInto(out *AWSSMProvider) {
	*out = *in
	in.Auth.DeepCopyInto(&out.Auth)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSSMProvider.
func (in *AWSSMProvider) DeepCopy() *AWSSMProvider {
	if in == nil {
		return nil
	}
	out := new(AWSSMProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSecretStore) DeepCopyInto(out *ClusterSecretStore) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSecretStore.
func (in *ClusterSecretStore) DeepCopy() *ClusterSecretStore {
	if in == nil {
		return nil
	}
	out := new(ClusterSecretStore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterSecretStore) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSecretStoreList) DeepCopyInto(out *ClusterSecretStoreList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterSecretStore, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSecretStoreList.
func (in *ClusterSecretStoreList) DeepCopy() *ClusterSecretStoreList {
	if in == nil {
		return nil
	}
	out := new(ClusterSecretStoreList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterSecretStoreList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalSecret) DeepCopyInto(out *ExternalSecret) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalSecret.
func (in *ExternalSecret) DeepCopy() *ExternalSecret {
	if in == nil {
		return nil
	}
	out := new(ExternalSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExternalSecret) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalSecretData) DeepCopyInto(out *ExternalSecretData) {
	*out = *in
	out.RemoteRef = in.RemoteRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalSecretData.
func (in *ExternalSecretData) DeepCopy() *ExternalSecretData {
	if in == nil {
		return nil
	}
	out := new(ExternalSecretData)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalSecretDataRemoteRef) DeepCopyInto(out *ExternalSecretDataRemoteRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalSecretDataRemoteRef.
func (in *ExternalSecretDataRemoteRef) DeepCopy() *ExternalSecretDataRemoteRef {
	if in == nil {
		return nil
	}
	out := new(ExternalSecretDataRemoteRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalSecretList) DeepCopyInto(out *ExternalSecretList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ExternalSecret, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalSecretList.
func (in *ExternalSecretList) DeepCopy() *ExternalSecretList {
	if in == nil {
		return nil
	}
	out := new(ExternalSecretList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExternalSecretList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalSecretSpec) DeepCopyInto(out *ExternalSecretSpec) {
	*out = *in
	out.SecretStoreRef = in.SecretStoreRef
	out.Target = in.Target
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = make([]ExternalSecretData, len(*in))
		copy(*out, *in)
	}
	if in.DataFrom != nil {
		in, out := &in.DataFrom, &out.DataFrom
		*out = make([]ExternalSecretDataRemoteRef, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalSecretSpec.
func (in *ExternalSecretSpec) DeepCopy() *ExternalSecretSpec {
	if in == nil {
		return nil
	}
	out := new(ExternalSecretSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalSecretStatus) DeepCopyInto(out *ExternalSecretStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ExternalSecretStatusCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalSecretStatus.
func (in *ExternalSecretStatus) DeepCopy() *ExternalSecretStatus {
	if in == nil {
		return nil
	}
	out := new(ExternalSecretStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalSecretStatusCondition) DeepCopyInto(out *ExternalSecretStatusCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	in.LastSyncTime.DeepCopyInto(&out.LastSyncTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalSecretStatusCondition.
func (in *ExternalSecretStatusCondition) DeepCopy() *ExternalSecretStatusCondition {
	if in == nil {
		return nil
	}
	out := new(ExternalSecretStatusCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalSecretTarget) DeepCopyInto(out *ExternalSecretTarget) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalSecretTarget.
func (in *ExternalSecretTarget) DeepCopy() *ExternalSecretTarget {
	if in == nil {
		return nil
	}
	out := new(ExternalSecretTarget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalSecretTemplate) DeepCopyInto(out *ExternalSecretTemplate) {
	*out = *in
	in.Metadata.DeepCopyInto(&out.Metadata)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalSecretTemplate.
func (in *ExternalSecretTemplate) DeepCopy() *ExternalSecretTemplate {
	if in == nil {
		return nil
	}
	out := new(ExternalSecretTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalSecretTemplateMetadata) DeepCopyInto(out *ExternalSecretTemplateMetadata) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalSecretTemplateMetadata.
func (in *ExternalSecretTemplateMetadata) DeepCopy() *ExternalSecretTemplateMetadata {
	if in == nil {
		return nil
	}
	out := new(ExternalSecretTemplateMetadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretStore) DeepCopyInto(out *SecretStore) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretStore.
func (in *SecretStore) DeepCopy() *SecretStore {
	if in == nil {
		return nil
	}
	out := new(SecretStore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecretStore) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretStoreList) DeepCopyInto(out *SecretStoreList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SecretStore, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretStoreList.
func (in *SecretStoreList) DeepCopy() *SecretStoreList {
	if in == nil {
		return nil
	}
	out := new(SecretStoreList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecretStoreList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretStoreProvider) DeepCopyInto(out *SecretStoreProvider) {
	*out = *in
	if in.AWSSM != nil {
		in, out := &in.AWSSM, &out.AWSSM
		*out = new(AWSSMProvider)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretStoreProvider.
func (in *SecretStoreProvider) DeepCopy() *SecretStoreProvider {
	if in == nil {
		return nil
	}
	out := new(SecretStoreProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretStoreRef) DeepCopyInto(out *SecretStoreRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretStoreRef.
func (in *SecretStoreRef) DeepCopy() *SecretStoreRef {
	if in == nil {
		return nil
	}
	out := new(SecretStoreRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretStoreSpec) DeepCopyInto(out *SecretStoreSpec) {
	*out = *in
	if in.Provider != nil {
		in, out := &in.Provider, &out.Provider
		*out = new(SecretStoreProvider)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretStoreSpec.
func (in *SecretStoreSpec) DeepCopy() *SecretStoreSpec {
	if in == nil {
		return nil
	}
	out := new(SecretStoreSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretStoreStatus) DeepCopyInto(out *SecretStoreStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]SecretStoreStatusCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretStoreStatus.
func (in *SecretStoreStatus) DeepCopy() *SecretStoreStatus {
	if in == nil {
		return nil
	}
	out := new(SecretStoreStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretStoreStatusCondition) DeepCopyInto(out *SecretStoreStatusCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretStoreStatusCondition.
func (in *SecretStoreStatusCondition) DeepCopy() *SecretStoreStatusCondition {
	if in == nil {
		return nil
	}
	out := new(SecretStoreStatusCondition)
	in.DeepCopyInto(out)
	return out
}
