package kube

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/yizhibian/gdumper/api/v1/monitor"
	"github.com/yizhibian/gdumper/internal/app/monitor/model"
	"github.com/yizhibian/gdumper/internal/app/monitor/service"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func init() {
	service.RegisterKube(New())
}

// 创建 Kubernetes 客户端
func createK8sClient(ctx context.Context) (*kubernetes.Clientset, error) {
	// 1. 优先使用 In-Cluster 配置（当运行在 Kubernetes 集群中时）
	config, err := rest.InClusterConfig()
	if err != nil {
		g.Log().Info(ctx, "error when init first time", err.Error())
		// 2. 如果不在集群内，则使用 kubeconfig 文件
		kubeconfig := os.Getenv("KUBECONFIG")
		if kubeconfig == "" {
			kubeconfig = os.Getenv("HOME") + "/.kube/config"
		}
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			g.Log().Info(ctx, "error when init second time", err.Error())

			return nil, fmt.Errorf("failed to build kubeconfig: %v", err)
		}
	}

	// 创建客户端集合
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		g.Log().Info(context.Background(), "error when get clientset", err.Error())

		return nil, fmt.Errorf("failed to create kubernetes client: %v", err)
	}

	return clientset, nil
}

func New() *sKube {
	client, err := createK8sClient(context.Background())
	if err != nil {
		g.Log().Info(context.Background(), "k8s wrong", err.Error())
	}
	return &sKube{
		clientset: client,
	}
}

type sKube struct {
	clientset *kubernetes.Clientset
}

func (s *sKube) GetSomeShit(c context.Context, namespace string) (err error) {
	g.Log().Info(c, "print some shit here")
	return
}

func (s *sKube) GetPodsByNamespace(c context.Context, namespace string) (res *monitor.PodListRes, err error) {
	res = &monitor.PodListRes{}
	if err != nil {
		g.Log().Info(c, "error when get create", err.Error())
		return
	}
	// 获取 Pod 列表
	pods, err := s.clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		g.Log().Info(c, "error when get pods", err.Error())
		return
	}
	// 获取事件信息
	events, err := s.clientset.CoreV1().Events(namespace).List(context.TODO(), metav1.ListOptions{
		FieldSelector: "involvedObject.kind=Pod",
	})
	if err != nil {
		g.Log().Info(c, "error when get events", err.Error())
		return
	}
	//if len(pods.Items) == 0 {
	//	g.Log().Info(c, "null pod")
	//}
	//
	//if len(events.Items) == 0 {
	//	g.Log().Info(c, "null events")
	//}

	// 转换为展示格式
	var podInfos []*model.PodInfo
	for _, pod := range pods.Items {
		podInfo := s.convertPodToInfo(pod, events.Items)
		podInfos = append(podInfos, podInfo)
	}

	res.List = podInfos

	return
}

func (s *sKube) convertPodToInfo(pod corev1.Pod, events []corev1.Event) *model.PodInfo {
	// 转换容器信息
	var containers []model.ContainerInfo
	for _, container := range pod.Spec.Containers {
		ci := model.ContainerInfo{
			Name:  container.Name,
			Image: container.Image,
		}

		// 获取容器状态
		for _, cs := range pod.Status.ContainerStatuses {
			if cs.Name == container.Name {
				ci.Ready = cs.Ready
				ci.Restarts = cs.RestartCount
				break
			}
		}

		// 获取资源请求/限制
		if requests := container.Resources.Requests; requests != nil {
			ci.RequestCPU = requests.Cpu().String()
			ci.RequestMem = requests.Memory().String()
		}
		if limits := container.Resources.Limits; limits != nil {
			ci.LimitCPU = limits.Cpu().String()
			ci.LimitMem = limits.Memory().String()
		}

		containers = append(containers, ci)
	}

	// 获取关联事件
	var podEvents []model.EventInfo
	for _, event := range events {
		if event.InvolvedObject.UID == pod.UID {
			podEvents = append(podEvents, model.EventInfo{
				Type:    event.Type,
				Reason:  event.Reason,
				Message: event.Message,
				Age:     metav1.Now().Sub(event.LastTimestamp.Time).Round(time.Second).String(),
			})
		}
	}

	return &model.PodInfo{
		Name:         pod.Name,
		Namespace:    pod.Namespace,
		Status:       string(pod.Status.Phase),
		IP:           pod.Status.PodIP,
		Node:         pod.Spec.NodeName,
		CreationTime: pod.CreationTimestamp,
		Labels:       pod.Labels,
		Containers:   containers,
		Events:       podEvents,
	}
}
