// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package policy

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	intstr "k8s.io/apimachinery/pkg/util/intstr"
	reflect "reflect"
)

// Deprecated: register deep-copy functions.
func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// Deprecated: RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Eviction).DeepCopyInto(out.(*Eviction))
			return nil
		}, InType: reflect.TypeOf(&Eviction{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PodDisruptionBudget).DeepCopyInto(out.(*PodDisruptionBudget))
			return nil
		}, InType: reflect.TypeOf(&PodDisruptionBudget{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PodDisruptionBudgetList).DeepCopyInto(out.(*PodDisruptionBudgetList))
			return nil
		}, InType: reflect.TypeOf(&PodDisruptionBudgetList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PodDisruptionBudgetSpec).DeepCopyInto(out.(*PodDisruptionBudgetSpec))
			return nil
		}, InType: reflect.TypeOf(&PodDisruptionBudgetSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*PodDisruptionBudgetStatus).DeepCopyInto(out.(*PodDisruptionBudgetStatus))
			return nil
		}, InType: reflect.TypeOf(&PodDisruptionBudgetStatus{})},
	)
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Eviction) DeepCopyInto(out *Eviction) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.DeleteOptions != nil {
		in, out := &in.DeleteOptions, &out.DeleteOptions
		if *in == nil {
			*out = nil
		} else {
			*out = new(v1.DeleteOptions)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new Eviction.
func (x *Eviction) DeepCopy() *Eviction {
	if x == nil {
		return nil
	}
	out := new(Eviction)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (x *Eviction) DeepCopyObject() runtime.Object {
	if c := x.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodDisruptionBudget) DeepCopyInto(out *PodDisruptionBudget) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new PodDisruptionBudget.
func (x *PodDisruptionBudget) DeepCopy() *PodDisruptionBudget {
	if x == nil {
		return nil
	}
	out := new(PodDisruptionBudget)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (x *PodDisruptionBudget) DeepCopyObject() runtime.Object {
	if c := x.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodDisruptionBudgetList) DeepCopyInto(out *PodDisruptionBudgetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PodDisruptionBudget, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new PodDisruptionBudgetList.
func (x *PodDisruptionBudgetList) DeepCopy() *PodDisruptionBudgetList {
	if x == nil {
		return nil
	}
	out := new(PodDisruptionBudgetList)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (x *PodDisruptionBudgetList) DeepCopyObject() runtime.Object {
	if c := x.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodDisruptionBudgetSpec) DeepCopyInto(out *PodDisruptionBudgetSpec) {
	*out = *in
	if in.MinAvailable != nil {
		in, out := &in.MinAvailable, &out.MinAvailable
		if *in == nil {
			*out = nil
		} else {
			*out = new(intstr.IntOrString)
			**out = **in
		}
	}
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		if *in == nil {
			*out = nil
		} else {
			*out = new(v1.LabelSelector)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.MaxUnavailable != nil {
		in, out := &in.MaxUnavailable, &out.MaxUnavailable
		if *in == nil {
			*out = nil
		} else {
			*out = new(intstr.IntOrString)
			**out = **in
		}
	}
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new PodDisruptionBudgetSpec.
func (x *PodDisruptionBudgetSpec) DeepCopy() *PodDisruptionBudgetSpec {
	if x == nil {
		return nil
	}
	out := new(PodDisruptionBudgetSpec)
	x.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodDisruptionBudgetStatus) DeepCopyInto(out *PodDisruptionBudgetStatus) {
	*out = *in
	if in.DisruptedPods != nil {
		in, out := &in.DisruptedPods, &out.DisruptedPods
		*out = make(map[string]v1.Time, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, creating a new PodDisruptionBudgetStatus.
func (x *PodDisruptionBudgetStatus) DeepCopy() *PodDisruptionBudgetStatus {
	if x == nil {
		return nil
	}
	out := new(PodDisruptionBudgetStatus)
	x.DeepCopyInto(out)
	return out
}
