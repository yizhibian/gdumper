package model

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type PodInfo struct {
	Name         string            `json:"name"`
	Namespace    string            `json:"namespace"`
	Status       string            `json:"status"`
	IP           string            `json:"ip"`
	Node         string            `json:"node"`
	CreationTime metav1.Time       `json:"creationTime"`
	Labels       map[string]string `json:"labels"`
	Containers   []ContainerInfo   `json:"containers"`
	Events       []EventInfo       `json:"events,omitempty"`
}

type ContainerInfo struct {
	Name       string `json:"name"`
	Image      string `json:"image"`
	Ready      bool   `json:"ready"`
	Restarts   int32  `json:"restarts"`
	RequestCPU string `json:"requestCpu,omitempty"`
	LimitCPU   string `json:"limitCpu,omitempty"`
	RequestMem string `json:"requestMem,omitempty"`
	LimitMem   string `json:"limitMem,omitempty"`
}

type EventInfo struct {
	Type    string `json:"type"`
	Reason  string `json:"reason"`
	Message string `json:"message"`
	Age     string `json:"age"`
}
