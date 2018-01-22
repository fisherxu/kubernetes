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

// This file was autogenerated by defaulter-gen. Do not edit it manually!

package v1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// RegisterDefaults adds defaulters functions to the given scheme.
// Public to allow building arbitrary schemes.
// All generated defaulters are covering - they call all nested defaulters.
func RegisterDefaults(scheme *runtime.Scheme) error {
	scheme.AddTypeDefaultingFunc(&v1.ConfigMap{}, func(obj interface{}) { SetObjectDefaults_ConfigMap(obj.(*v1.ConfigMap)) })
	scheme.AddTypeDefaultingFunc(&v1.ConfigMapList{}, func(obj interface{}) { SetObjectDefaults_ConfigMapList(obj.(*v1.ConfigMapList)) })
	scheme.AddTypeDefaultingFunc(&v1.Endpoints{}, func(obj interface{}) { SetObjectDefaults_Endpoints(obj.(*v1.Endpoints)) })
	scheme.AddTypeDefaultingFunc(&v1.EndpointsList{}, func(obj interface{}) { SetObjectDefaults_EndpointsList(obj.(*v1.EndpointsList)) })
	scheme.AddTypeDefaultingFunc(&v1.LimitRange{}, func(obj interface{}) { SetObjectDefaults_LimitRange(obj.(*v1.LimitRange)) })
	scheme.AddTypeDefaultingFunc(&v1.LimitRangeList{}, func(obj interface{}) { SetObjectDefaults_LimitRangeList(obj.(*v1.LimitRangeList)) })
	scheme.AddTypeDefaultingFunc(&v1.Namespace{}, func(obj interface{}) { SetObjectDefaults_Namespace(obj.(*v1.Namespace)) })
	scheme.AddTypeDefaultingFunc(&v1.NamespaceList{}, func(obj interface{}) { SetObjectDefaults_NamespaceList(obj.(*v1.NamespaceList)) })
	scheme.AddTypeDefaultingFunc(&v1.Node{}, func(obj interface{}) { SetObjectDefaults_Node(obj.(*v1.Node)) })
	scheme.AddTypeDefaultingFunc(&v1.NodeList{}, func(obj interface{}) { SetObjectDefaults_NodeList(obj.(*v1.NodeList)) })
	scheme.AddTypeDefaultingFunc(&v1.PersistentVolume{}, func(obj interface{}) { SetObjectDefaults_PersistentVolume(obj.(*v1.PersistentVolume)) })
	scheme.AddTypeDefaultingFunc(&v1.PersistentVolumeClaim{}, func(obj interface{}) { SetObjectDefaults_PersistentVolumeClaim(obj.(*v1.PersistentVolumeClaim)) })
	scheme.AddTypeDefaultingFunc(&v1.PersistentVolumeClaimList{}, func(obj interface{}) {
		SetObjectDefaults_PersistentVolumeClaimList(obj.(*v1.PersistentVolumeClaimList))
	})
	scheme.AddTypeDefaultingFunc(&v1.PersistentVolumeList{}, func(obj interface{}) { SetObjectDefaults_PersistentVolumeList(obj.(*v1.PersistentVolumeList)) })
	scheme.AddTypeDefaultingFunc(&v1.Pod{}, func(obj interface{}) { SetObjectDefaults_Pod(obj.(*v1.Pod)) })
	scheme.AddTypeDefaultingFunc(&v1.PodList{}, func(obj interface{}) { SetObjectDefaults_PodList(obj.(*v1.PodList)) })
	scheme.AddTypeDefaultingFunc(&v1.PodTemplate{}, func(obj interface{}) { SetObjectDefaults_PodTemplate(obj.(*v1.PodTemplate)) })
	scheme.AddTypeDefaultingFunc(&v1.PodTemplateList{}, func(obj interface{}) { SetObjectDefaults_PodTemplateList(obj.(*v1.PodTemplateList)) })
	scheme.AddTypeDefaultingFunc(&v1.ReplicationController{}, func(obj interface{}) { SetObjectDefaults_ReplicationController(obj.(*v1.ReplicationController)) })
	scheme.AddTypeDefaultingFunc(&v1.ReplicationControllerList{}, func(obj interface{}) {
		SetObjectDefaults_ReplicationControllerList(obj.(*v1.ReplicationControllerList))
	})
	scheme.AddTypeDefaultingFunc(&v1.ResourceQuota{}, func(obj interface{}) { SetObjectDefaults_ResourceQuota(obj.(*v1.ResourceQuota)) })
	scheme.AddTypeDefaultingFunc(&v1.ResourceQuotaList{}, func(obj interface{}) { SetObjectDefaults_ResourceQuotaList(obj.(*v1.ResourceQuotaList)) })
	scheme.AddTypeDefaultingFunc(&v1.Secret{}, func(obj interface{}) { SetObjectDefaults_Secret(obj.(*v1.Secret)) })
	scheme.AddTypeDefaultingFunc(&v1.SecretList{}, func(obj interface{}) { SetObjectDefaults_SecretList(obj.(*v1.SecretList)) })
	scheme.AddTypeDefaultingFunc(&v1.Service{}, func(obj interface{}) { SetObjectDefaults_Service(obj.(*v1.Service)) })
	scheme.AddTypeDefaultingFunc(&v1.ServiceList{}, func(obj interface{}) { SetObjectDefaults_ServiceList(obj.(*v1.ServiceList)) })
	return nil
}

func SetObjectDefaults_ConfigMap(in *v1.ConfigMap) {
	SetDefaults_ConfigMap(in)
}

func SetObjectDefaults_ConfigMapList(in *v1.ConfigMapList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_ConfigMap(a)
	}
}

func SetObjectDefaults_Endpoints(in *v1.Endpoints) {
	SetDefaults_Endpoints(in)
}

func SetObjectDefaults_EndpointsList(in *v1.EndpointsList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_Endpoints(a)
	}
}

func SetObjectDefaults_LimitRange(in *v1.LimitRange) {
	for i := range in.Spec.Limits {
		a := &in.Spec.Limits[i]
		SetDefaults_LimitRangeItem(a)
		SetDefaults_ResourceList(&a.Max)
		SetDefaults_ResourceList(&a.Min)
		SetDefaults_ResourceList(&a.Default)
		SetDefaults_ResourceList(&a.DefaultRequest)
		SetDefaults_ResourceList(&a.MaxLimitRequestRatio)
	}
}

func SetObjectDefaults_LimitRangeList(in *v1.LimitRangeList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_LimitRange(a)
	}
}

func SetObjectDefaults_Namespace(in *v1.Namespace) {
	SetDefaults_NamespaceStatus(&in.Status)
}

func SetObjectDefaults_NamespaceList(in *v1.NamespaceList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_Namespace(a)
	}
}

func SetObjectDefaults_Node(in *v1.Node) {
	SetDefaults_Node(in)
	SetDefaults_NodeStatus(&in.Status)
	SetDefaults_ResourceList(&in.Status.Capacity)
	SetDefaults_ResourceList(&in.Status.Allocatable)
}

func SetObjectDefaults_NodeList(in *v1.NodeList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_Node(a)
	}
}

func SetObjectDefaults_PersistentVolume(in *v1.PersistentVolume) {
	SetDefaults_PersistentVolume(in)
	SetDefaults_ResourceList(&in.Spec.Capacity)
	if in.Spec.PersistentVolumeSource.HostPath != nil {
		SetDefaults_HostPathVolumeSource(in.Spec.PersistentVolumeSource.HostPath)
	}
	if in.Spec.PersistentVolumeSource.RBD != nil {
		SetDefaults_RBDPersistentVolumeSource(in.Spec.PersistentVolumeSource.RBD)
	}
	if in.Spec.PersistentVolumeSource.ISCSI != nil {
		SetDefaults_ISCSIPersistentVolumeSource(in.Spec.PersistentVolumeSource.ISCSI)
	}
	if in.Spec.PersistentVolumeSource.AzureDisk != nil {
		SetDefaults_AzureDiskVolumeSource(in.Spec.PersistentVolumeSource.AzureDisk)
	}
	if in.Spec.PersistentVolumeSource.ScaleIO != nil {
		SetDefaults_ScaleIOPersistentVolumeSource(in.Spec.PersistentVolumeSource.ScaleIO)
	}
}

func SetObjectDefaults_PersistentVolumeClaim(in *v1.PersistentVolumeClaim) {
	SetDefaults_PersistentVolumeClaim(in)
	SetDefaults_ResourceList(&in.Spec.Resources.Limits)
	SetDefaults_ResourceList(&in.Spec.Resources.Requests)
	SetDefaults_ResourceList(&in.Status.Capacity)
}

func SetObjectDefaults_PersistentVolumeClaimList(in *v1.PersistentVolumeClaimList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_PersistentVolumeClaim(a)
	}
}

func SetObjectDefaults_PersistentVolumeList(in *v1.PersistentVolumeList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_PersistentVolume(a)
	}
}

func SetObjectDefaults_Pod(in *v1.Pod) {
	SetDefaults_Pod(in)
	SetDefaults_PodSpec(&in.Spec)
	for i := range in.Spec.Volumes {
		a := &in.Spec.Volumes[i]
		SetDefaults_Volume(a)
		if a.VolumeSource.HostPath != nil {
			SetDefaults_HostPathVolumeSource(a.VolumeSource.HostPath)
		}
		if a.VolumeSource.Secret != nil {
			SetDefaults_SecretVolumeSource(a.VolumeSource.Secret)
		}
		if a.VolumeSource.ISCSI != nil {
			SetDefaults_ISCSIVolumeSource(a.VolumeSource.ISCSI)
		}
		if a.VolumeSource.RBD != nil {
			SetDefaults_RBDVolumeSource(a.VolumeSource.RBD)
		}
		if a.VolumeSource.DownwardAPI != nil {
			SetDefaults_DownwardAPIVolumeSource(a.VolumeSource.DownwardAPI)
			for j := range a.VolumeSource.DownwardAPI.Items {
				b := &a.VolumeSource.DownwardAPI.Items[j]
				if b.FieldRef != nil {
					SetDefaults_ObjectFieldSelector(b.FieldRef)
				}
			}
		}
		if a.VolumeSource.ConfigMap != nil {
			SetDefaults_ConfigMapVolumeSource(a.VolumeSource.ConfigMap)
		}
		if a.VolumeSource.AzureDisk != nil {
			SetDefaults_AzureDiskVolumeSource(a.VolumeSource.AzureDisk)
		}
		if a.VolumeSource.Projected != nil {
			SetDefaults_ProjectedVolumeSource(a.VolumeSource.Projected)
			for j := range a.VolumeSource.Projected.Sources {
				b := &a.VolumeSource.Projected.Sources[j]
				if b.DownwardAPI != nil {
					for k := range b.DownwardAPI.Items {
						c := &b.DownwardAPI.Items[k]
						if c.FieldRef != nil {
							SetDefaults_ObjectFieldSelector(c.FieldRef)
						}
					}
				}
			}
		}
		if a.VolumeSource.ScaleIO != nil {
			SetDefaults_ScaleIOVolumeSource(a.VolumeSource.ScaleIO)
		}
	}
	for i := range in.Spec.InitContainers {
		a := &in.Spec.InitContainers[i]
		SetDefaults_Container(a)
		for j := range a.Ports {
			b := &a.Ports[j]
			SetDefaults_ContainerPort(b)
		}
		for j := range a.Env {
			b := &a.Env[j]
			if b.ValueFrom != nil {
				if b.ValueFrom.FieldRef != nil {
					SetDefaults_ObjectFieldSelector(b.ValueFrom.FieldRef)
				}
			}
		}
		SetDefaults_ResourceList(&a.Resources.Limits)
		SetDefaults_ResourceList(&a.Resources.Requests)
		if a.LivenessProbe != nil {
			SetDefaults_Probe(a.LivenessProbe)
			if a.LivenessProbe.Handler.HTTPGet != nil {
				SetDefaults_HTTPGetAction(a.LivenessProbe.Handler.HTTPGet)
			}
		}
		if a.ReadinessProbe != nil {
			SetDefaults_Probe(a.ReadinessProbe)
			if a.ReadinessProbe.Handler.HTTPGet != nil {
				SetDefaults_HTTPGetAction(a.ReadinessProbe.Handler.HTTPGet)
			}
		}
		if a.Lifecycle != nil {
			if a.Lifecycle.PostStart != nil {
				if a.Lifecycle.PostStart.HTTPGet != nil {
					SetDefaults_HTTPGetAction(a.Lifecycle.PostStart.HTTPGet)
				}
			}
			if a.Lifecycle.PreStop != nil {
				if a.Lifecycle.PreStop.HTTPGet != nil {
					SetDefaults_HTTPGetAction(a.Lifecycle.PreStop.HTTPGet)
				}
			}
		}
	}
	for i := range in.Spec.Containers {
		a := &in.Spec.Containers[i]
		SetDefaults_Container(a)
		for j := range a.Ports {
			b := &a.Ports[j]
			SetDefaults_ContainerPort(b)
		}
		for j := range a.Env {
			b := &a.Env[j]
			if b.ValueFrom != nil {
				if b.ValueFrom.FieldRef != nil {
					SetDefaults_ObjectFieldSelector(b.ValueFrom.FieldRef)
				}
			}
		}
		SetDefaults_ResourceList(&a.Resources.Limits)
		SetDefaults_ResourceList(&a.Resources.Requests)
		if a.LivenessProbe != nil {
			SetDefaults_Probe(a.LivenessProbe)
			if a.LivenessProbe.Handler.HTTPGet != nil {
				SetDefaults_HTTPGetAction(a.LivenessProbe.Handler.HTTPGet)
			}
		}
		if a.ReadinessProbe != nil {
			SetDefaults_Probe(a.ReadinessProbe)
			if a.ReadinessProbe.Handler.HTTPGet != nil {
				SetDefaults_HTTPGetAction(a.ReadinessProbe.Handler.HTTPGet)
			}
		}
		if a.Lifecycle != nil {
			if a.Lifecycle.PostStart != nil {
				if a.Lifecycle.PostStart.HTTPGet != nil {
					SetDefaults_HTTPGetAction(a.Lifecycle.PostStart.HTTPGet)
				}
			}
			if a.Lifecycle.PreStop != nil {
				if a.Lifecycle.PreStop.HTTPGet != nil {
					SetDefaults_HTTPGetAction(a.Lifecycle.PreStop.HTTPGet)
				}
			}
		}
	}
}

func SetObjectDefaults_PodList(in *v1.PodList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_Pod(a)
	}
}

func SetObjectDefaults_PodTemplate(in *v1.PodTemplate) {
	SetDefaults_PodSpec(&in.Template.Spec)
	for i := range in.Template.Spec.Volumes {
		a := &in.Template.Spec.Volumes[i]
		SetDefaults_Volume(a)
		if a.VolumeSource.HostPath != nil {
			SetDefaults_HostPathVolumeSource(a.VolumeSource.HostPath)
		}
		if a.VolumeSource.Secret != nil {
			SetDefaults_SecretVolumeSource(a.VolumeSource.Secret)
		}
		if a.VolumeSource.ISCSI != nil {
			SetDefaults_ISCSIVolumeSource(a.VolumeSource.ISCSI)
		}
		if a.VolumeSource.RBD != nil {
			SetDefaults_RBDVolumeSource(a.VolumeSource.RBD)
		}
		if a.VolumeSource.DownwardAPI != nil {
			SetDefaults_DownwardAPIVolumeSource(a.VolumeSource.DownwardAPI)
			for j := range a.VolumeSource.DownwardAPI.Items {
				b := &a.VolumeSource.DownwardAPI.Items[j]
				if b.FieldRef != nil {
					SetDefaults_ObjectFieldSelector(b.FieldRef)
				}
			}
		}
		if a.VolumeSource.ConfigMap != nil {
			SetDefaults_ConfigMapVolumeSource(a.VolumeSource.ConfigMap)
		}
		if a.VolumeSource.AzureDisk != nil {
			SetDefaults_AzureDiskVolumeSource(a.VolumeSource.AzureDisk)
		}
		if a.VolumeSource.Projected != nil {
			SetDefaults_ProjectedVolumeSource(a.VolumeSource.Projected)
			for j := range a.VolumeSource.Projected.Sources {
				b := &a.VolumeSource.Projected.Sources[j]
				if b.DownwardAPI != nil {
					for k := range b.DownwardAPI.Items {
						c := &b.DownwardAPI.Items[k]
						if c.FieldRef != nil {
							SetDefaults_ObjectFieldSelector(c.FieldRef)
						}
					}
				}
			}
		}
		if a.VolumeSource.ScaleIO != nil {
			SetDefaults_ScaleIOVolumeSource(a.VolumeSource.ScaleIO)
		}
	}
	for i := range in.Template.Spec.InitContainers {
		a := &in.Template.Spec.InitContainers[i]
		SetDefaults_Container(a)
		for j := range a.Ports {
			b := &a.Ports[j]
			SetDefaults_ContainerPort(b)
		}
		for j := range a.Env {
			b := &a.Env[j]
			if b.ValueFrom != nil {
				if b.ValueFrom.FieldRef != nil {
					SetDefaults_ObjectFieldSelector(b.ValueFrom.FieldRef)
				}
			}
		}
		SetDefaults_ResourceList(&a.Resources.Limits)
		SetDefaults_ResourceList(&a.Resources.Requests)
		if a.LivenessProbe != nil {
			SetDefaults_Probe(a.LivenessProbe)
			if a.LivenessProbe.Handler.HTTPGet != nil {
				SetDefaults_HTTPGetAction(a.LivenessProbe.Handler.HTTPGet)
			}
		}
		if a.ReadinessProbe != nil {
			SetDefaults_Probe(a.ReadinessProbe)
			if a.ReadinessProbe.Handler.HTTPGet != nil {
				SetDefaults_HTTPGetAction(a.ReadinessProbe.Handler.HTTPGet)
			}
		}
		if a.Lifecycle != nil {
			if a.Lifecycle.PostStart != nil {
				if a.Lifecycle.PostStart.HTTPGet != nil {
					SetDefaults_HTTPGetAction(a.Lifecycle.PostStart.HTTPGet)
				}
			}
			if a.Lifecycle.PreStop != nil {
				if a.Lifecycle.PreStop.HTTPGet != nil {
					SetDefaults_HTTPGetAction(a.Lifecycle.PreStop.HTTPGet)
				}
			}
		}
	}
	for i := range in.Template.Spec.Containers {
		a := &in.Template.Spec.Containers[i]
		SetDefaults_Container(a)
		for j := range a.Ports {
			b := &a.Ports[j]
			SetDefaults_ContainerPort(b)
		}
		for j := range a.Env {
			b := &a.Env[j]
			if b.ValueFrom != nil {
				if b.ValueFrom.FieldRef != nil {
					SetDefaults_ObjectFieldSelector(b.ValueFrom.FieldRef)
				}
			}
		}
		SetDefaults_ResourceList(&a.Resources.Limits)
		SetDefaults_ResourceList(&a.Resources.Requests)
		if a.LivenessProbe != nil {
			SetDefaults_Probe(a.LivenessProbe)
			if a.LivenessProbe.Handler.HTTPGet != nil {
				SetDefaults_HTTPGetAction(a.LivenessProbe.Handler.HTTPGet)
			}
		}
		if a.ReadinessProbe != nil {
			SetDefaults_Probe(a.ReadinessProbe)
			if a.ReadinessProbe.Handler.HTTPGet != nil {
				SetDefaults_HTTPGetAction(a.ReadinessProbe.Handler.HTTPGet)
			}
		}
		if a.Lifecycle != nil {
			if a.Lifecycle.PostStart != nil {
				if a.Lifecycle.PostStart.HTTPGet != nil {
					SetDefaults_HTTPGetAction(a.Lifecycle.PostStart.HTTPGet)
				}
			}
			if a.Lifecycle.PreStop != nil {
				if a.Lifecycle.PreStop.HTTPGet != nil {
					SetDefaults_HTTPGetAction(a.Lifecycle.PreStop.HTTPGet)
				}
			}
		}
	}
}

func SetObjectDefaults_PodTemplateList(in *v1.PodTemplateList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_PodTemplate(a)
	}
}

func SetObjectDefaults_ReplicationController(in *v1.ReplicationController) {
	SetDefaults_ReplicationController(in)
	if in.Spec.Template != nil {
		SetDefaults_PodSpec(&in.Spec.Template.Spec)
		for i := range in.Spec.Template.Spec.Volumes {
			a := &in.Spec.Template.Spec.Volumes[i]
			SetDefaults_Volume(a)
			if a.VolumeSource.HostPath != nil {
				SetDefaults_HostPathVolumeSource(a.VolumeSource.HostPath)
			}
			if a.VolumeSource.Secret != nil {
				SetDefaults_SecretVolumeSource(a.VolumeSource.Secret)
			}
			if a.VolumeSource.ISCSI != nil {
				SetDefaults_ISCSIVolumeSource(a.VolumeSource.ISCSI)
			}
			if a.VolumeSource.RBD != nil {
				SetDefaults_RBDVolumeSource(a.VolumeSource.RBD)
			}
			if a.VolumeSource.DownwardAPI != nil {
				SetDefaults_DownwardAPIVolumeSource(a.VolumeSource.DownwardAPI)
				for j := range a.VolumeSource.DownwardAPI.Items {
					b := &a.VolumeSource.DownwardAPI.Items[j]
					if b.FieldRef != nil {
						SetDefaults_ObjectFieldSelector(b.FieldRef)
					}
				}
			}
			if a.VolumeSource.ConfigMap != nil {
				SetDefaults_ConfigMapVolumeSource(a.VolumeSource.ConfigMap)
			}
			if a.VolumeSource.AzureDisk != nil {
				SetDefaults_AzureDiskVolumeSource(a.VolumeSource.AzureDisk)
			}
			if a.VolumeSource.Projected != nil {
				SetDefaults_ProjectedVolumeSource(a.VolumeSource.Projected)
				for j := range a.VolumeSource.Projected.Sources {
					b := &a.VolumeSource.Projected.Sources[j]
					if b.DownwardAPI != nil {
						for k := range b.DownwardAPI.Items {
							c := &b.DownwardAPI.Items[k]
							if c.FieldRef != nil {
								SetDefaults_ObjectFieldSelector(c.FieldRef)
							}
						}
					}
				}
			}
			if a.VolumeSource.ScaleIO != nil {
				SetDefaults_ScaleIOVolumeSource(a.VolumeSource.ScaleIO)
			}
		}
		for i := range in.Spec.Template.Spec.InitContainers {
			a := &in.Spec.Template.Spec.InitContainers[i]
			SetDefaults_Container(a)
			for j := range a.Ports {
				b := &a.Ports[j]
				SetDefaults_ContainerPort(b)
			}
			for j := range a.Env {
				b := &a.Env[j]
				if b.ValueFrom != nil {
					if b.ValueFrom.FieldRef != nil {
						SetDefaults_ObjectFieldSelector(b.ValueFrom.FieldRef)
					}
				}
			}
			SetDefaults_ResourceList(&a.Resources.Limits)
			SetDefaults_ResourceList(&a.Resources.Requests)
			if a.LivenessProbe != nil {
				SetDefaults_Probe(a.LivenessProbe)
				if a.LivenessProbe.Handler.HTTPGet != nil {
					SetDefaults_HTTPGetAction(a.LivenessProbe.Handler.HTTPGet)
				}
			}
			if a.ReadinessProbe != nil {
				SetDefaults_Probe(a.ReadinessProbe)
				if a.ReadinessProbe.Handler.HTTPGet != nil {
					SetDefaults_HTTPGetAction(a.ReadinessProbe.Handler.HTTPGet)
				}
			}
			if a.Lifecycle != nil {
				if a.Lifecycle.PostStart != nil {
					if a.Lifecycle.PostStart.HTTPGet != nil {
						SetDefaults_HTTPGetAction(a.Lifecycle.PostStart.HTTPGet)
					}
				}
				if a.Lifecycle.PreStop != nil {
					if a.Lifecycle.PreStop.HTTPGet != nil {
						SetDefaults_HTTPGetAction(a.Lifecycle.PreStop.HTTPGet)
					}
				}
			}
		}
		for i := range in.Spec.Template.Spec.Containers {
			a := &in.Spec.Template.Spec.Containers[i]
			SetDefaults_Container(a)
			for j := range a.Ports {
				b := &a.Ports[j]
				SetDefaults_ContainerPort(b)
			}
			for j := range a.Env {
				b := &a.Env[j]
				if b.ValueFrom != nil {
					if b.ValueFrom.FieldRef != nil {
						SetDefaults_ObjectFieldSelector(b.ValueFrom.FieldRef)
					}
				}
			}
			SetDefaults_ResourceList(&a.Resources.Limits)
			SetDefaults_ResourceList(&a.Resources.Requests)
			if a.LivenessProbe != nil {
				SetDefaults_Probe(a.LivenessProbe)
				if a.LivenessProbe.Handler.HTTPGet != nil {
					SetDefaults_HTTPGetAction(a.LivenessProbe.Handler.HTTPGet)
				}
			}
			if a.ReadinessProbe != nil {
				SetDefaults_Probe(a.ReadinessProbe)
				if a.ReadinessProbe.Handler.HTTPGet != nil {
					SetDefaults_HTTPGetAction(a.ReadinessProbe.Handler.HTTPGet)
				}
			}
			if a.Lifecycle != nil {
				if a.Lifecycle.PostStart != nil {
					if a.Lifecycle.PostStart.HTTPGet != nil {
						SetDefaults_HTTPGetAction(a.Lifecycle.PostStart.HTTPGet)
					}
				}
				if a.Lifecycle.PreStop != nil {
					if a.Lifecycle.PreStop.HTTPGet != nil {
						SetDefaults_HTTPGetAction(a.Lifecycle.PreStop.HTTPGet)
					}
				}
			}
		}
	}
}

func SetObjectDefaults_ReplicationControllerList(in *v1.ReplicationControllerList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_ReplicationController(a)
	}
}

func SetObjectDefaults_ResourceQuota(in *v1.ResourceQuota) {
	SetDefaults_ResourceList(&in.Spec.Hard)
	SetDefaults_ResourceList(&in.Status.Hard)
	SetDefaults_ResourceList(&in.Status.Used)
}

func SetObjectDefaults_ResourceQuotaList(in *v1.ResourceQuotaList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_ResourceQuota(a)
	}
}

func SetObjectDefaults_Secret(in *v1.Secret) {
	SetDefaults_Secret(in)
}

func SetObjectDefaults_SecretList(in *v1.SecretList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_Secret(a)
	}
}

func SetObjectDefaults_Service(in *v1.Service) {
	SetDefaults_Service(in)
}

func SetObjectDefaults_ServiceList(in *v1.ServiceList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_Service(a)
	}
}
