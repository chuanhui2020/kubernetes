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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v2

// PodsMetricStatusApplyConfiguration represents a declarative configuration of the PodsMetricStatus type for use
// with apply.
type PodsMetricStatusApplyConfiguration struct {
	Metric  *MetricIdentifierApplyConfiguration  `json:"metric,omitempty"`
	Current *MetricValueStatusApplyConfiguration `json:"current,omitempty"`
}

// PodsMetricStatusApplyConfiguration constructs a declarative configuration of the PodsMetricStatus type for use with
// apply.
func PodsMetricStatus() *PodsMetricStatusApplyConfiguration {
	return &PodsMetricStatusApplyConfiguration{}
}
func (b PodsMetricStatusApplyConfiguration) IsApplyConfiguration() {}

// WithMetric sets the Metric field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Metric field is set to the value of the last call.
func (b *PodsMetricStatusApplyConfiguration) WithMetric(value *MetricIdentifierApplyConfiguration) *PodsMetricStatusApplyConfiguration {
	b.Metric = value
	return b
}

// WithCurrent sets the Current field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Current field is set to the value of the last call.
func (b *PodsMetricStatusApplyConfiguration) WithCurrent(value *MetricValueStatusApplyConfiguration) *PodsMetricStatusApplyConfiguration {
	b.Current = value
	return b
}
