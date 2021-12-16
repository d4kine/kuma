// Generated by tools/resource-gen
// Run "make generate" to update this file.

// nolint:whitespace
package v1alpha1

import (
	"github.com/golang/protobuf/proto"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	mesh_proto "github.com/kumahq/kuma/api/mesh/v1alpha1"
	"github.com/kumahq/kuma/pkg/plugins/resources/k8s/native/pkg/model"
	"github.com/kumahq/kuma/pkg/plugins/resources/k8s/native/pkg/registry"
)

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster
type CircuitBreaker struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Mesh string                     `json:"mesh,omitempty"`
	Spec *mesh_proto.CircuitBreaker `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced
type CircuitBreakerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CircuitBreaker `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CircuitBreaker{}, &CircuitBreakerList{})
}

func (cb *CircuitBreaker) GetObjectMeta() *metav1.ObjectMeta {
	return &cb.ObjectMeta
}

func (cb *CircuitBreaker) SetObjectMeta(m *metav1.ObjectMeta) {
	cb.ObjectMeta = *m
}

func (cb *CircuitBreaker) GetMesh() string {
	return cb.Mesh
}

func (cb *CircuitBreaker) SetMesh(mesh string) {
	cb.Mesh = mesh
}

func (cb *CircuitBreaker) GetSpec() proto.Message {
	return cb.Spec
}

func (cb *CircuitBreaker) SetSpec(spec proto.Message) {
	cb.Spec = proto.Clone(spec).(*mesh_proto.CircuitBreaker)
}

func (cb *CircuitBreaker) Scope() model.Scope {
	return model.ScopeCluster
}

func (l *CircuitBreakerList) GetItems() []model.KubernetesObject {
	result := make([]model.KubernetesObject, len(l.Items))
	for i := range l.Items {
		result[i] = &l.Items[i]
	}
	return result
}

func init() {
	registry.RegisterObjectType(&mesh_proto.CircuitBreaker{}, &CircuitBreaker{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "CircuitBreaker",
		},
	})
	registry.RegisterListType(&mesh_proto.CircuitBreaker{}, &CircuitBreakerList{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "CircuitBreakerList",
		},
	})
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced
type Dataplane struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Mesh string                `json:"mesh,omitempty"`
	Spec *mesh_proto.Dataplane `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster
type DataplaneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Dataplane `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Dataplane{}, &DataplaneList{})
}

func (cb *Dataplane) GetObjectMeta() *metav1.ObjectMeta {
	return &cb.ObjectMeta
}

func (cb *Dataplane) SetObjectMeta(m *metav1.ObjectMeta) {
	cb.ObjectMeta = *m
}

func (cb *Dataplane) GetMesh() string {
	return cb.Mesh
}

func (cb *Dataplane) SetMesh(mesh string) {
	cb.Mesh = mesh
}

func (cb *Dataplane) GetSpec() proto.Message {
	return cb.Spec
}

func (cb *Dataplane) SetSpec(spec proto.Message) {
	cb.Spec = proto.Clone(spec).(*mesh_proto.Dataplane)
}

func (cb *Dataplane) Scope() model.Scope {
	return model.ScopeNamespace
}

func (l *DataplaneList) GetItems() []model.KubernetesObject {
	result := make([]model.KubernetesObject, len(l.Items))
	for i := range l.Items {
		result[i] = &l.Items[i]
	}
	return result
}

func init() {
	registry.RegisterObjectType(&mesh_proto.Dataplane{}, &Dataplane{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "Dataplane",
		},
	})
	registry.RegisterListType(&mesh_proto.Dataplane{}, &DataplaneList{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "DataplaneList",
		},
	})
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced
type DataplaneInsight struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Mesh   string                       `json:"mesh,omitempty"`
	Status *mesh_proto.DataplaneInsight `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster
type DataplaneInsightList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataplaneInsight `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataplaneInsight{}, &DataplaneInsightList{})
}

func (cb *DataplaneInsight) GetObjectMeta() *metav1.ObjectMeta {
	return &cb.ObjectMeta
}

func (cb *DataplaneInsight) SetObjectMeta(m *metav1.ObjectMeta) {
	cb.ObjectMeta = *m
}

func (cb *DataplaneInsight) GetMesh() string {
	return cb.Mesh
}

func (cb *DataplaneInsight) SetMesh(mesh string) {
	cb.Mesh = mesh
}

func (cb *DataplaneInsight) GetSpec() proto.Message {
	return cb.Status
}

func (cb *DataplaneInsight) SetSpec(spec proto.Message) {
	cb.Status = proto.Clone(spec).(*mesh_proto.DataplaneInsight)
}

func (cb *DataplaneInsight) Scope() model.Scope {
	return model.ScopeNamespace
}

func (l *DataplaneInsightList) GetItems() []model.KubernetesObject {
	result := make([]model.KubernetesObject, len(l.Items))
	for i := range l.Items {
		result[i] = &l.Items[i]
	}
	return result
}

func init() {
	registry.RegisterObjectType(&mesh_proto.DataplaneInsight{}, &DataplaneInsight{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "DataplaneInsight",
		},
	})
	registry.RegisterListType(&mesh_proto.DataplaneInsight{}, &DataplaneInsightList{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "DataplaneInsightList",
		},
	})
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster
type ExternalService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Mesh string                      `json:"mesh,omitempty"`
	Spec *mesh_proto.ExternalService `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced
type ExternalServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ExternalService `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ExternalService{}, &ExternalServiceList{})
}

func (cb *ExternalService) GetObjectMeta() *metav1.ObjectMeta {
	return &cb.ObjectMeta
}

func (cb *ExternalService) SetObjectMeta(m *metav1.ObjectMeta) {
	cb.ObjectMeta = *m
}

func (cb *ExternalService) GetMesh() string {
	return cb.Mesh
}

func (cb *ExternalService) SetMesh(mesh string) {
	cb.Mesh = mesh
}

func (cb *ExternalService) GetSpec() proto.Message {
	return cb.Spec
}

func (cb *ExternalService) SetSpec(spec proto.Message) {
	cb.Spec = proto.Clone(spec).(*mesh_proto.ExternalService)
}

func (cb *ExternalService) Scope() model.Scope {
	return model.ScopeCluster
}

func (l *ExternalServiceList) GetItems() []model.KubernetesObject {
	result := make([]model.KubernetesObject, len(l.Items))
	for i := range l.Items {
		result[i] = &l.Items[i]
	}
	return result
}

func init() {
	registry.RegisterObjectType(&mesh_proto.ExternalService{}, &ExternalService{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "ExternalService",
		},
	})
	registry.RegisterListType(&mesh_proto.ExternalService{}, &ExternalServiceList{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "ExternalServiceList",
		},
	})
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster
type FaultInjection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Mesh string                     `json:"mesh,omitempty"`
	Spec *mesh_proto.FaultInjection `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced
type FaultInjectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FaultInjection `json:"items"`
}

func init() {
	SchemeBuilder.Register(&FaultInjection{}, &FaultInjectionList{})
}

func (cb *FaultInjection) GetObjectMeta() *metav1.ObjectMeta {
	return &cb.ObjectMeta
}

func (cb *FaultInjection) SetObjectMeta(m *metav1.ObjectMeta) {
	cb.ObjectMeta = *m
}

func (cb *FaultInjection) GetMesh() string {
	return cb.Mesh
}

func (cb *FaultInjection) SetMesh(mesh string) {
	cb.Mesh = mesh
}

func (cb *FaultInjection) GetSpec() proto.Message {
	return cb.Spec
}

func (cb *FaultInjection) SetSpec(spec proto.Message) {
	cb.Spec = proto.Clone(spec).(*mesh_proto.FaultInjection)
}

func (cb *FaultInjection) Scope() model.Scope {
	return model.ScopeCluster
}

func (l *FaultInjectionList) GetItems() []model.KubernetesObject {
	result := make([]model.KubernetesObject, len(l.Items))
	for i := range l.Items {
		result[i] = &l.Items[i]
	}
	return result
}

func init() {
	registry.RegisterObjectType(&mesh_proto.FaultInjection{}, &FaultInjection{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "FaultInjection",
		},
	})
	registry.RegisterListType(&mesh_proto.FaultInjection{}, &FaultInjectionList{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "FaultInjectionList",
		},
	})
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster
type Gateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Mesh string              `json:"mesh,omitempty"`
	Spec *mesh_proto.Gateway `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced
type GatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Gateway `json:"items"`
}

func (cb *Gateway) GetObjectMeta() *metav1.ObjectMeta {
	return &cb.ObjectMeta
}

func (cb *Gateway) SetObjectMeta(m *metav1.ObjectMeta) {
	cb.ObjectMeta = *m
}

func (cb *Gateway) GetMesh() string {
	return cb.Mesh
}

func (cb *Gateway) SetMesh(mesh string) {
	cb.Mesh = mesh
}

func (cb *Gateway) GetSpec() proto.Message {
	return cb.Spec
}

func (cb *Gateway) SetSpec(spec proto.Message) {
	cb.Spec = proto.Clone(spec).(*mesh_proto.Gateway)
}

func (cb *Gateway) Scope() model.Scope {
	return model.ScopeCluster
}

func (l *GatewayList) GetItems() []model.KubernetesObject {
	result := make([]model.KubernetesObject, len(l.Items))
	for i := range l.Items {
		result[i] = &l.Items[i]
	}
	return result
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster
type GatewayRoute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Mesh string                   `json:"mesh,omitempty"`
	Spec *mesh_proto.GatewayRoute `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced
type GatewayRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GatewayRoute `json:"items"`
}

func (cb *GatewayRoute) GetObjectMeta() *metav1.ObjectMeta {
	return &cb.ObjectMeta
}

func (cb *GatewayRoute) SetObjectMeta(m *metav1.ObjectMeta) {
	cb.ObjectMeta = *m
}

func (cb *GatewayRoute) GetMesh() string {
	return cb.Mesh
}

func (cb *GatewayRoute) SetMesh(mesh string) {
	cb.Mesh = mesh
}

func (cb *GatewayRoute) GetSpec() proto.Message {
	return cb.Spec
}

func (cb *GatewayRoute) SetSpec(spec proto.Message) {
	cb.Spec = proto.Clone(spec).(*mesh_proto.GatewayRoute)
}

func (cb *GatewayRoute) Scope() model.Scope {
	return model.ScopeCluster
}

func (l *GatewayRouteList) GetItems() []model.KubernetesObject {
	result := make([]model.KubernetesObject, len(l.Items))
	for i := range l.Items {
		result[i] = &l.Items[i]
	}
	return result
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster
type HealthCheck struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Mesh string                  `json:"mesh,omitempty"`
	Spec *mesh_proto.HealthCheck `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced
type HealthCheckList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HealthCheck `json:"items"`
}

func init() {
	SchemeBuilder.Register(&HealthCheck{}, &HealthCheckList{})
}

func (cb *HealthCheck) GetObjectMeta() *metav1.ObjectMeta {
	return &cb.ObjectMeta
}

func (cb *HealthCheck) SetObjectMeta(m *metav1.ObjectMeta) {
	cb.ObjectMeta = *m
}

func (cb *HealthCheck) GetMesh() string {
	return cb.Mesh
}

func (cb *HealthCheck) SetMesh(mesh string) {
	cb.Mesh = mesh
}

func (cb *HealthCheck) GetSpec() proto.Message {
	return cb.Spec
}

func (cb *HealthCheck) SetSpec(spec proto.Message) {
	cb.Spec = proto.Clone(spec).(*mesh_proto.HealthCheck)
}

func (cb *HealthCheck) Scope() model.Scope {
	return model.ScopeCluster
}

func (l *HealthCheckList) GetItems() []model.KubernetesObject {
	result := make([]model.KubernetesObject, len(l.Items))
	for i := range l.Items {
		result[i] = &l.Items[i]
	}
	return result
}

func init() {
	registry.RegisterObjectType(&mesh_proto.HealthCheck{}, &HealthCheck{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "HealthCheck",
		},
	})
	registry.RegisterListType(&mesh_proto.HealthCheck{}, &HealthCheckList{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "HealthCheckList",
		},
	})
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster
type Mesh struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Mesh string           `json:"mesh,omitempty"`
	Spec *mesh_proto.Mesh `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced
type MeshList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Mesh `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Mesh{}, &MeshList{})
}

func (cb *Mesh) GetObjectMeta() *metav1.ObjectMeta {
	return &cb.ObjectMeta
}

func (cb *Mesh) SetObjectMeta(m *metav1.ObjectMeta) {
	cb.ObjectMeta = *m
}

func (cb *Mesh) GetMesh() string {
	return cb.Mesh
}

func (cb *Mesh) SetMesh(mesh string) {
	cb.Mesh = mesh
}

func (cb *Mesh) GetSpec() proto.Message {
	return cb.Spec
}

func (cb *Mesh) SetSpec(spec proto.Message) {
	cb.Spec = proto.Clone(spec).(*mesh_proto.Mesh)
}

func (cb *Mesh) Scope() model.Scope {
	return model.ScopeCluster
}

func (l *MeshList) GetItems() []model.KubernetesObject {
	result := make([]model.KubernetesObject, len(l.Items))
	for i := range l.Items {
		result[i] = &l.Items[i]
	}
	return result
}

func init() {
	registry.RegisterObjectType(&mesh_proto.Mesh{}, &Mesh{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "Mesh",
		},
	})
	registry.RegisterListType(&mesh_proto.Mesh{}, &MeshList{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "MeshList",
		},
	})
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster
type MeshInsight struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Mesh string                  `json:"mesh,omitempty"`
	Spec *mesh_proto.MeshInsight `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced
type MeshInsightList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MeshInsight `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MeshInsight{}, &MeshInsightList{})
}

func (cb *MeshInsight) GetObjectMeta() *metav1.ObjectMeta {
	return &cb.ObjectMeta
}

func (cb *MeshInsight) SetObjectMeta(m *metav1.ObjectMeta) {
	cb.ObjectMeta = *m
}

func (cb *MeshInsight) GetMesh() string {
	return cb.Mesh
}

func (cb *MeshInsight) SetMesh(mesh string) {
	cb.Mesh = mesh
}

func (cb *MeshInsight) GetSpec() proto.Message {
	return cb.Spec
}

func (cb *MeshInsight) SetSpec(spec proto.Message) {
	cb.Spec = proto.Clone(spec).(*mesh_proto.MeshInsight)
}

func (cb *MeshInsight) Scope() model.Scope {
	return model.ScopeCluster
}

func (l *MeshInsightList) GetItems() []model.KubernetesObject {
	result := make([]model.KubernetesObject, len(l.Items))
	for i := range l.Items {
		result[i] = &l.Items[i]
	}
	return result
}

func init() {
	registry.RegisterObjectType(&mesh_proto.MeshInsight{}, &MeshInsight{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "MeshInsight",
		},
	})
	registry.RegisterListType(&mesh_proto.MeshInsight{}, &MeshInsightList{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "MeshInsightList",
		},
	})
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster
type ProxyTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Mesh string                    `json:"mesh,omitempty"`
	Spec *mesh_proto.ProxyTemplate `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced
type ProxyTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProxyTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ProxyTemplate{}, &ProxyTemplateList{})
}

func (cb *ProxyTemplate) GetObjectMeta() *metav1.ObjectMeta {
	return &cb.ObjectMeta
}

func (cb *ProxyTemplate) SetObjectMeta(m *metav1.ObjectMeta) {
	cb.ObjectMeta = *m
}

func (cb *ProxyTemplate) GetMesh() string {
	return cb.Mesh
}

func (cb *ProxyTemplate) SetMesh(mesh string) {
	cb.Mesh = mesh
}

func (cb *ProxyTemplate) GetSpec() proto.Message {
	return cb.Spec
}

func (cb *ProxyTemplate) SetSpec(spec proto.Message) {
	cb.Spec = proto.Clone(spec).(*mesh_proto.ProxyTemplate)
}

func (cb *ProxyTemplate) Scope() model.Scope {
	return model.ScopeCluster
}

func (l *ProxyTemplateList) GetItems() []model.KubernetesObject {
	result := make([]model.KubernetesObject, len(l.Items))
	for i := range l.Items {
		result[i] = &l.Items[i]
	}
	return result
}

func init() {
	registry.RegisterObjectType(&mesh_proto.ProxyTemplate{}, &ProxyTemplate{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "ProxyTemplate",
		},
	})
	registry.RegisterListType(&mesh_proto.ProxyTemplate{}, &ProxyTemplateList{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "ProxyTemplateList",
		},
	})
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster
type RateLimit struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Mesh string                `json:"mesh,omitempty"`
	Spec *mesh_proto.RateLimit `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced
type RateLimitList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RateLimit `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RateLimit{}, &RateLimitList{})
}

func (cb *RateLimit) GetObjectMeta() *metav1.ObjectMeta {
	return &cb.ObjectMeta
}

func (cb *RateLimit) SetObjectMeta(m *metav1.ObjectMeta) {
	cb.ObjectMeta = *m
}

func (cb *RateLimit) GetMesh() string {
	return cb.Mesh
}

func (cb *RateLimit) SetMesh(mesh string) {
	cb.Mesh = mesh
}

func (cb *RateLimit) GetSpec() proto.Message {
	return cb.Spec
}

func (cb *RateLimit) SetSpec(spec proto.Message) {
	cb.Spec = proto.Clone(spec).(*mesh_proto.RateLimit)
}

func (cb *RateLimit) Scope() model.Scope {
	return model.ScopeCluster
}

func (l *RateLimitList) GetItems() []model.KubernetesObject {
	result := make([]model.KubernetesObject, len(l.Items))
	for i := range l.Items {
		result[i] = &l.Items[i]
	}
	return result
}

func init() {
	registry.RegisterObjectType(&mesh_proto.RateLimit{}, &RateLimit{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "RateLimit",
		},
	})
	registry.RegisterListType(&mesh_proto.RateLimit{}, &RateLimitList{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "RateLimitList",
		},
	})
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster
type Retry struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Mesh string            `json:"mesh,omitempty"`
	Spec *mesh_proto.Retry `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced
type RetryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Retry `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Retry{}, &RetryList{})
}

func (cb *Retry) GetObjectMeta() *metav1.ObjectMeta {
	return &cb.ObjectMeta
}

func (cb *Retry) SetObjectMeta(m *metav1.ObjectMeta) {
	cb.ObjectMeta = *m
}

func (cb *Retry) GetMesh() string {
	return cb.Mesh
}

func (cb *Retry) SetMesh(mesh string) {
	cb.Mesh = mesh
}

func (cb *Retry) GetSpec() proto.Message {
	return cb.Spec
}

func (cb *Retry) SetSpec(spec proto.Message) {
	cb.Spec = proto.Clone(spec).(*mesh_proto.Retry)
}

func (cb *Retry) Scope() model.Scope {
	return model.ScopeCluster
}

func (l *RetryList) GetItems() []model.KubernetesObject {
	result := make([]model.KubernetesObject, len(l.Items))
	for i := range l.Items {
		result[i] = &l.Items[i]
	}
	return result
}

func init() {
	registry.RegisterObjectType(&mesh_proto.Retry{}, &Retry{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "Retry",
		},
	})
	registry.RegisterListType(&mesh_proto.Retry{}, &RetryList{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "RetryList",
		},
	})
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster
type ServiceInsight struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Mesh string                     `json:"mesh,omitempty"`
	Spec *mesh_proto.ServiceInsight `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced
type ServiceInsightList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceInsight `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ServiceInsight{}, &ServiceInsightList{})
}

func (cb *ServiceInsight) GetObjectMeta() *metav1.ObjectMeta {
	return &cb.ObjectMeta
}

func (cb *ServiceInsight) SetObjectMeta(m *metav1.ObjectMeta) {
	cb.ObjectMeta = *m
}

func (cb *ServiceInsight) GetMesh() string {
	return cb.Mesh
}

func (cb *ServiceInsight) SetMesh(mesh string) {
	cb.Mesh = mesh
}

func (cb *ServiceInsight) GetSpec() proto.Message {
	return cb.Spec
}

func (cb *ServiceInsight) SetSpec(spec proto.Message) {
	cb.Spec = proto.Clone(spec).(*mesh_proto.ServiceInsight)
}

func (cb *ServiceInsight) Scope() model.Scope {
	return model.ScopeCluster
}

func (l *ServiceInsightList) GetItems() []model.KubernetesObject {
	result := make([]model.KubernetesObject, len(l.Items))
	for i := range l.Items {
		result[i] = &l.Items[i]
	}
	return result
}

func init() {
	registry.RegisterObjectType(&mesh_proto.ServiceInsight{}, &ServiceInsight{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "ServiceInsight",
		},
	})
	registry.RegisterListType(&mesh_proto.ServiceInsight{}, &ServiceInsightList{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "ServiceInsightList",
		},
	})
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster
type Timeout struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Mesh string              `json:"mesh,omitempty"`
	Spec *mesh_proto.Timeout `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced
type TimeoutList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Timeout `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Timeout{}, &TimeoutList{})
}

func (cb *Timeout) GetObjectMeta() *metav1.ObjectMeta {
	return &cb.ObjectMeta
}

func (cb *Timeout) SetObjectMeta(m *metav1.ObjectMeta) {
	cb.ObjectMeta = *m
}

func (cb *Timeout) GetMesh() string {
	return cb.Mesh
}

func (cb *Timeout) SetMesh(mesh string) {
	cb.Mesh = mesh
}

func (cb *Timeout) GetSpec() proto.Message {
	return cb.Spec
}

func (cb *Timeout) SetSpec(spec proto.Message) {
	cb.Spec = proto.Clone(spec).(*mesh_proto.Timeout)
}

func (cb *Timeout) Scope() model.Scope {
	return model.ScopeCluster
}

func (l *TimeoutList) GetItems() []model.KubernetesObject {
	result := make([]model.KubernetesObject, len(l.Items))
	for i := range l.Items {
		result[i] = &l.Items[i]
	}
	return result
}

func init() {
	registry.RegisterObjectType(&mesh_proto.Timeout{}, &Timeout{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "Timeout",
		},
	})
	registry.RegisterListType(&mesh_proto.Timeout{}, &TimeoutList{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "TimeoutList",
		},
	})
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster
type TrafficLog struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Mesh string                 `json:"mesh,omitempty"`
	Spec *mesh_proto.TrafficLog `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced
type TrafficLogList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TrafficLog `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TrafficLog{}, &TrafficLogList{})
}

func (cb *TrafficLog) GetObjectMeta() *metav1.ObjectMeta {
	return &cb.ObjectMeta
}

func (cb *TrafficLog) SetObjectMeta(m *metav1.ObjectMeta) {
	cb.ObjectMeta = *m
}

func (cb *TrafficLog) GetMesh() string {
	return cb.Mesh
}

func (cb *TrafficLog) SetMesh(mesh string) {
	cb.Mesh = mesh
}

func (cb *TrafficLog) GetSpec() proto.Message {
	return cb.Spec
}

func (cb *TrafficLog) SetSpec(spec proto.Message) {
	cb.Spec = proto.Clone(spec).(*mesh_proto.TrafficLog)
}

func (cb *TrafficLog) Scope() model.Scope {
	return model.ScopeCluster
}

func (l *TrafficLogList) GetItems() []model.KubernetesObject {
	result := make([]model.KubernetesObject, len(l.Items))
	for i := range l.Items {
		result[i] = &l.Items[i]
	}
	return result
}

func init() {
	registry.RegisterObjectType(&mesh_proto.TrafficLog{}, &TrafficLog{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "TrafficLog",
		},
	})
	registry.RegisterListType(&mesh_proto.TrafficLog{}, &TrafficLogList{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "TrafficLogList",
		},
	})
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster
type TrafficPermission struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Mesh string                        `json:"mesh,omitempty"`
	Spec *mesh_proto.TrafficPermission `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced
type TrafficPermissionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TrafficPermission `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TrafficPermission{}, &TrafficPermissionList{})
}

func (cb *TrafficPermission) GetObjectMeta() *metav1.ObjectMeta {
	return &cb.ObjectMeta
}

func (cb *TrafficPermission) SetObjectMeta(m *metav1.ObjectMeta) {
	cb.ObjectMeta = *m
}

func (cb *TrafficPermission) GetMesh() string {
	return cb.Mesh
}

func (cb *TrafficPermission) SetMesh(mesh string) {
	cb.Mesh = mesh
}

func (cb *TrafficPermission) GetSpec() proto.Message {
	return cb.Spec
}

func (cb *TrafficPermission) SetSpec(spec proto.Message) {
	cb.Spec = proto.Clone(spec).(*mesh_proto.TrafficPermission)
}

func (cb *TrafficPermission) Scope() model.Scope {
	return model.ScopeCluster
}

func (l *TrafficPermissionList) GetItems() []model.KubernetesObject {
	result := make([]model.KubernetesObject, len(l.Items))
	for i := range l.Items {
		result[i] = &l.Items[i]
	}
	return result
}

func init() {
	registry.RegisterObjectType(&mesh_proto.TrafficPermission{}, &TrafficPermission{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "TrafficPermission",
		},
	})
	registry.RegisterListType(&mesh_proto.TrafficPermission{}, &TrafficPermissionList{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "TrafficPermissionList",
		},
	})
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster
type TrafficRoute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Mesh string                   `json:"mesh,omitempty"`
	Spec *mesh_proto.TrafficRoute `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced
type TrafficRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TrafficRoute `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TrafficRoute{}, &TrafficRouteList{})
}

func (cb *TrafficRoute) GetObjectMeta() *metav1.ObjectMeta {
	return &cb.ObjectMeta
}

func (cb *TrafficRoute) SetObjectMeta(m *metav1.ObjectMeta) {
	cb.ObjectMeta = *m
}

func (cb *TrafficRoute) GetMesh() string {
	return cb.Mesh
}

func (cb *TrafficRoute) SetMesh(mesh string) {
	cb.Mesh = mesh
}

func (cb *TrafficRoute) GetSpec() proto.Message {
	return cb.Spec
}

func (cb *TrafficRoute) SetSpec(spec proto.Message) {
	cb.Spec = proto.Clone(spec).(*mesh_proto.TrafficRoute)
}

func (cb *TrafficRoute) Scope() model.Scope {
	return model.ScopeCluster
}

func (l *TrafficRouteList) GetItems() []model.KubernetesObject {
	result := make([]model.KubernetesObject, len(l.Items))
	for i := range l.Items {
		result[i] = &l.Items[i]
	}
	return result
}

func init() {
	registry.RegisterObjectType(&mesh_proto.TrafficRoute{}, &TrafficRoute{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "TrafficRoute",
		},
	})
	registry.RegisterListType(&mesh_proto.TrafficRoute{}, &TrafficRouteList{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "TrafficRouteList",
		},
	})
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster
type TrafficTrace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Mesh string                   `json:"mesh,omitempty"`
	Spec *mesh_proto.TrafficTrace `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced
type TrafficTraceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TrafficTrace `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TrafficTrace{}, &TrafficTraceList{})
}

func (cb *TrafficTrace) GetObjectMeta() *metav1.ObjectMeta {
	return &cb.ObjectMeta
}

func (cb *TrafficTrace) SetObjectMeta(m *metav1.ObjectMeta) {
	cb.ObjectMeta = *m
}

func (cb *TrafficTrace) GetMesh() string {
	return cb.Mesh
}

func (cb *TrafficTrace) SetMesh(mesh string) {
	cb.Mesh = mesh
}

func (cb *TrafficTrace) GetSpec() proto.Message {
	return cb.Spec
}

func (cb *TrafficTrace) SetSpec(spec proto.Message) {
	cb.Spec = proto.Clone(spec).(*mesh_proto.TrafficTrace)
}

func (cb *TrafficTrace) Scope() model.Scope {
	return model.ScopeCluster
}

func (l *TrafficTraceList) GetItems() []model.KubernetesObject {
	result := make([]model.KubernetesObject, len(l.Items))
	for i := range l.Items {
		result[i] = &l.Items[i]
	}
	return result
}

func init() {
	registry.RegisterObjectType(&mesh_proto.TrafficTrace{}, &TrafficTrace{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "TrafficTrace",
		},
	})
	registry.RegisterListType(&mesh_proto.TrafficTrace{}, &TrafficTraceList{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "TrafficTraceList",
		},
	})
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster
type VirtualOutbound struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Mesh string                      `json:"mesh,omitempty"`
	Spec *mesh_proto.VirtualOutbound `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced
type VirtualOutboundList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualOutbound `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VirtualOutbound{}, &VirtualOutboundList{})
}

func (cb *VirtualOutbound) GetObjectMeta() *metav1.ObjectMeta {
	return &cb.ObjectMeta
}

func (cb *VirtualOutbound) SetObjectMeta(m *metav1.ObjectMeta) {
	cb.ObjectMeta = *m
}

func (cb *VirtualOutbound) GetMesh() string {
	return cb.Mesh
}

func (cb *VirtualOutbound) SetMesh(mesh string) {
	cb.Mesh = mesh
}

func (cb *VirtualOutbound) GetSpec() proto.Message {
	return cb.Spec
}

func (cb *VirtualOutbound) SetSpec(spec proto.Message) {
	cb.Spec = proto.Clone(spec).(*mesh_proto.VirtualOutbound)
}

func (cb *VirtualOutbound) Scope() model.Scope {
	return model.ScopeCluster
}

func (l *VirtualOutboundList) GetItems() []model.KubernetesObject {
	result := make([]model.KubernetesObject, len(l.Items))
	for i := range l.Items {
		result[i] = &l.Items[i]
	}
	return result
}

func init() {
	registry.RegisterObjectType(&mesh_proto.VirtualOutbound{}, &VirtualOutbound{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "VirtualOutbound",
		},
	})
	registry.RegisterListType(&mesh_proto.VirtualOutbound{}, &VirtualOutboundList{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "VirtualOutboundList",
		},
	})
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced
type ZoneIngress struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Mesh string                  `json:"mesh,omitempty"`
	Spec *mesh_proto.ZoneIngress `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster
type ZoneIngressList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ZoneIngress `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ZoneIngress{}, &ZoneIngressList{})
}

func (cb *ZoneIngress) GetObjectMeta() *metav1.ObjectMeta {
	return &cb.ObjectMeta
}

func (cb *ZoneIngress) SetObjectMeta(m *metav1.ObjectMeta) {
	cb.ObjectMeta = *m
}

func (cb *ZoneIngress) GetMesh() string {
	return cb.Mesh
}

func (cb *ZoneIngress) SetMesh(mesh string) {
	cb.Mesh = mesh
}

func (cb *ZoneIngress) GetSpec() proto.Message {
	return cb.Spec
}

func (cb *ZoneIngress) SetSpec(spec proto.Message) {
	cb.Spec = proto.Clone(spec).(*mesh_proto.ZoneIngress)
}

func (cb *ZoneIngress) Scope() model.Scope {
	return model.ScopeNamespace
}

func (l *ZoneIngressList) GetItems() []model.KubernetesObject {
	result := make([]model.KubernetesObject, len(l.Items))
	for i := range l.Items {
		result[i] = &l.Items[i]
	}
	return result
}

func init() {
	registry.RegisterObjectType(&mesh_proto.ZoneIngress{}, &ZoneIngress{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "ZoneIngress",
		},
	})
	registry.RegisterListType(&mesh_proto.ZoneIngress{}, &ZoneIngressList{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "ZoneIngressList",
		},
	})
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced
type ZoneIngressInsight struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Mesh string                         `json:"mesh,omitempty"`
	Spec *mesh_proto.ZoneIngressInsight `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster
type ZoneIngressInsightList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ZoneIngressInsight `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ZoneIngressInsight{}, &ZoneIngressInsightList{})
}

func (cb *ZoneIngressInsight) GetObjectMeta() *metav1.ObjectMeta {
	return &cb.ObjectMeta
}

func (cb *ZoneIngressInsight) SetObjectMeta(m *metav1.ObjectMeta) {
	cb.ObjectMeta = *m
}

func (cb *ZoneIngressInsight) GetMesh() string {
	return cb.Mesh
}

func (cb *ZoneIngressInsight) SetMesh(mesh string) {
	cb.Mesh = mesh
}

func (cb *ZoneIngressInsight) GetSpec() proto.Message {
	return cb.Spec
}

func (cb *ZoneIngressInsight) SetSpec(spec proto.Message) {
	cb.Spec = proto.Clone(spec).(*mesh_proto.ZoneIngressInsight)
}

func (cb *ZoneIngressInsight) Scope() model.Scope {
	return model.ScopeNamespace
}

func (l *ZoneIngressInsightList) GetItems() []model.KubernetesObject {
	result := make([]model.KubernetesObject, len(l.Items))
	for i := range l.Items {
		result[i] = &l.Items[i]
	}
	return result
}

func init() {
	registry.RegisterObjectType(&mesh_proto.ZoneIngressInsight{}, &ZoneIngressInsight{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "ZoneIngressInsight",
		},
	})
	registry.RegisterListType(&mesh_proto.ZoneIngressInsight{}, &ZoneIngressInsightList{
		TypeMeta: metav1.TypeMeta{
			APIVersion: GroupVersion.String(),
			Kind:       "ZoneIngressInsightList",
		},
	})
}
