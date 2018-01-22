// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1beta1

import (
	unsafe "unsafe"

	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	custom_metrics "k8s.io/metrics/pkg/apis/custom_metrics"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1beta1_MetricValue_To_custom_metrics_MetricValue,
		Convert_custom_metrics_MetricValue_To_v1beta1_MetricValue,
		Convert_v1beta1_MetricValueList_To_custom_metrics_MetricValueList,
		Convert_custom_metrics_MetricValueList_To_v1beta1_MetricValueList,
	)
}

func autoConvert_v1beta1_MetricValue_To_custom_metrics_MetricValue(in *MetricValue, out *custom_metrics.MetricValue, s conversion.Scope) error {
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.DescribedObject, &out.DescribedObject, 0); err != nil {
		return err
	}
	out.MetricName = in.MetricName
	out.Timestamp = in.Timestamp
	out.WindowSeconds = (*int64)(unsafe.Pointer(in.WindowSeconds))
	out.Value = in.Value
	return nil
}

// Convert_v1beta1_MetricValue_To_custom_metrics_MetricValue is an autogenerated conversion function.
func Convert_v1beta1_MetricValue_To_custom_metrics_MetricValue(in *MetricValue, out *custom_metrics.MetricValue, s conversion.Scope) error {
	return autoConvert_v1beta1_MetricValue_To_custom_metrics_MetricValue(in, out, s)
}

func autoConvert_custom_metrics_MetricValue_To_v1beta1_MetricValue(in *custom_metrics.MetricValue, out *MetricValue, s conversion.Scope) error {
	// TODO: Inefficient conversion - can we improve it?
	if err := s.Convert(&in.DescribedObject, &out.DescribedObject, 0); err != nil {
		return err
	}
	out.MetricName = in.MetricName
	out.Timestamp = in.Timestamp
	out.WindowSeconds = (*int64)(unsafe.Pointer(in.WindowSeconds))
	out.Value = in.Value
	return nil
}

// Convert_custom_metrics_MetricValue_To_v1beta1_MetricValue is an autogenerated conversion function.
func Convert_custom_metrics_MetricValue_To_v1beta1_MetricValue(in *custom_metrics.MetricValue, out *MetricValue, s conversion.Scope) error {
	return autoConvert_custom_metrics_MetricValue_To_v1beta1_MetricValue(in, out, s)
}

func autoConvert_v1beta1_MetricValueList_To_custom_metrics_MetricValueList(in *MetricValueList, out *custom_metrics.MetricValueList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]custom_metrics.MetricValue)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1beta1_MetricValueList_To_custom_metrics_MetricValueList is an autogenerated conversion function.
func Convert_v1beta1_MetricValueList_To_custom_metrics_MetricValueList(in *MetricValueList, out *custom_metrics.MetricValueList, s conversion.Scope) error {
	return autoConvert_v1beta1_MetricValueList_To_custom_metrics_MetricValueList(in, out, s)
}

func autoConvert_custom_metrics_MetricValueList_To_v1beta1_MetricValueList(in *custom_metrics.MetricValueList, out *MetricValueList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]MetricValue)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_custom_metrics_MetricValueList_To_v1beta1_MetricValueList is an autogenerated conversion function.
func Convert_custom_metrics_MetricValueList_To_v1beta1_MetricValueList(in *custom_metrics.MetricValueList, out *MetricValueList, s conversion.Scope) error {
	return autoConvert_custom_metrics_MetricValueList_To_v1beta1_MetricValueList(in, out, s)
}
