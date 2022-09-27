package utils

import (
	"context"
	"fmt"
	"io/ioutil"
	"time"

	"go.uber.org/zap"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func GetKubeClient(configPath string, logger *zap.Logger) *kubernetes.Clientset {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		logger.Fatal("Could not read kube config", zap.Any("ERROR", err))
	}

	config, err := clientcmd.RESTConfigFromKubeConfig(data)
	if err != nil {
		logger.Fatal("Could not parse kube config", zap.Any("ERROR", err))
	}

	kubeClient, err := kubernetes.NewForConfig(config)
	if err != nil {
		logger.Fatal("Could not get the kube client", zap.Any("ERROR", err))
	}

	return kubeClient
}

func GetAllPodsOfNamespace(kubeClient *kubernetes.Clientset, namespace string, logger *zap.Logger) *v1.PodList {
	podInterface := kubeClient.CoreV1().Pods(namespace)
	podList, err := podInterface.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logger.Fatal("Could not get the pod list in the namespace: "+namespace, zap.Any("ERROR", err))
	}
	for _, pod := range podList.Items {
		fmt.Println(pod.GetName())
	}
	return podList
}

func WaitForPodToRun(kubeClient *kubernetes.Clientset, namespace string, podLabel string, logger *zap.Logger) {
	for true {
		// waiting for the resources to have pending state
		for true {
			if len(GetAllPodsOfNamespaceByLabel(kubeClient, namespace, podLabel, logger).Items) != 0 {
				break
			}
			time.Sleep(1000 * time.Millisecond)
		}

		podList := GetAllPodsOfNamespaceByLabel(kubeClient, namespace, podLabel, logger)
		pod := podList.Items[0]
		if pod.Status.Phase == "Running" {
			break
		}
		time.Sleep(1000 * time.Millisecond)
		logger.Info("Waiting for pod " + pod.GetName() + " in the namespace " + namespace + " to be in running state.")
	}
}

func GetAllPodsOfNamespaceByLabel(kubeClient *kubernetes.Clientset, namespace string, label string, logger *zap.Logger) *v1.PodList {
	podInterface := kubeClient.CoreV1().Pods(namespace)
	podList, err := podInterface.List(context.TODO(), metav1.ListOptions{LabelSelector: label})
	if err != nil {
		logger.Fatal("Could not get the pod list in the namespace: "+namespace, zap.Any("ERROR", err))
	}
	return podList
}

func AddLabelToARunningPod(kubeClient *kubernetes.Clientset, namespace string, existingLabel string, newLabelKey string, newLabelValue string, logger *zap.Logger) {
	// waiting for the resources to have pending state
	for true {
		if len(GetAllPodsOfNamespaceByLabel(kubeClient, namespace, existingLabel, logger).Items) != 0 {
			break
		}
		time.Sleep(1000 * time.Millisecond)
	}
	podList := GetAllPodsOfNamespaceByLabel(kubeClient, namespace, existingLabel, logger)
	pod := podList.Items[0]
	payload := `{"metadata": {"labels": {"` + newLabelKey + `": "` + newLabelValue + `"}}}`
	_, err := kubeClient.CoreV1().Pods(namespace).Patch(context.TODO(), pod.GetName(), types.MergePatchType, []byte(payload), metav1.PatchOptions{})
	if err != nil {
		logger.Fatal("Could not patch the pod: "+pod.GetName(), zap.Any("ERROR", err))
	}
}

func CheckK8sResource(resourceName string, resourceType string) {

}
