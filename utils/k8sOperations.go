package utils

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/remotecommand"
)

// Utility function to get the kube client, required by other k8s functions to perform
func GetKubeClient(configPath string, clusterContext string, logger *zap.Logger) *kubernetes.Clientset {
	configLoadingRules := &clientcmd.ClientConfigLoadingRules{ExplicitPath: configPath}
	configOverrides := &clientcmd.ConfigOverrides{CurrentContext: clusterContext}
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(configLoadingRules, configOverrides)

	restClientConfig, err := kubeConfig.ClientConfig()
	if err != nil {
		logger.Fatal("Could not get the kube client", zap.Any("ERROR", err))
	}
	kubeClient, err := kubernetes.NewForConfig(restClientConfig)
	if err != nil {
		logger.Fatal("Could not get the kube client", zap.Any("ERROR", err))
	}
	return kubeClient
}

// Utility function to get the k8s rest config file, required for certain functions
func GetK8sRestConfig(configPath string, clusterContext string, logger *zap.Logger) *rest.Config {
	configLoadingRules := &clientcmd.ClientConfigLoadingRules{ExplicitPath: configPath}
	configOverrides := &clientcmd.ConfigOverrides{CurrentContext: clusterContext}
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(configLoadingRules, configOverrides)

	k8sRestConfig, err := kubeConfig.ClientConfig()
	if err != nil {
		logger.Fatal("Could not get the kube client", zap.Any("ERROR", err))
	}

	return k8sRestConfig
}

// Utility function to get the list of kubernetes node IP(s)
func GetK8sNodeIP(kubeClient *kubernetes.Clientset, logger *zap.Logger) []string {
	nodes, err := kubeClient.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logger.Fatal("Could not fetch the kubernetes node list", zap.Any("ERROR", err))
	}

	nodeIPs := []string{}
	for _, node := range nodes.Items {

		nodeIPs = append(nodeIPs, node.Status.Addresses[0].Address)
	}
	return nodeIPs
}

// Utility function to the get the NodePort of the given service within a namespace
func GetK8sServicePort(kubeClient *kubernetes.Clientset, namespace string, serviceName string, logger *zap.Logger) []v1.ServicePort {
	services, err := kubeClient.CoreV1().Services(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logger.Fatal("Could not fetch the service list in the namespace: "+namespace, zap.Any("ERROR", err))
	}
	for _, service := range services.Items {
		if service.GetName() == serviceName {
			return service.Spec.Ports
		}
	}
	return nil
}

// Utility function to get all the pods as a list within a given namespace
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

// Utility function to wait for a pod to run based on its pod label
func WaitForPodToRun(kubeClient *kubernetes.Clientset, namespace string, podLabel string, logger *zap.Logger) {
	for true {
		// waiting for the resources to have pending state
		for true {
			if len(GetAllPodsOfNamespaceByLabel(kubeClient, namespace, podLabel, logger).Items) != 0 {
				break
			}
			time.Sleep(2500 * time.Millisecond)
		}

		podList := GetAllPodsOfNamespaceByLabel(kubeClient, namespace, podLabel, logger)
		pod := podList.Items[0]
		if pod.Status.Phase == "Running" {
			break
		}
		time.Sleep(2500 * time.Millisecond)
		logger.Info("Waiting for pod " + pod.GetName() + " in the namespace " + namespace + " to be in running state.")
	}
}

// Utility function to get all the pods of the mentioned namespace based on a particular name
func GetAllPodsOfNamespaceByLabel(kubeClient *kubernetes.Clientset, namespace string, label string, logger *zap.Logger) *v1.PodList {
	podInterface := kubeClient.CoreV1().Pods(namespace)
	podList, err := podInterface.List(context.TODO(), metav1.ListOptions{LabelSelector: label})
	if err != nil {
		logger.Fatal("Could not get the pod list in the namespace: "+namespace, zap.Any("ERROR", err))
	}
	return podList
}

// Utility function to add a label to a pod in a given namespace with an existing label
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

// Utility function to check if a pod is running in a given cluster within the mentioned namespace
func CheckIfPodRunning(kubeClient *kubernetes.Clientset, podName string, namespace string) bool {
	pod, _ := kubeClient.CoreV1().Pods(namespace).Get(context.TODO(), podName, metav1.GetOptions{})
	if pod.Status.Phase == "Running" {
		return true
	} else {
		return false
	}
}

// Utility function to check if the kubernetes resource exists or not in a given namespace
func CheckK8sResource(kubeClient *kubernetes.Clientset, resourceName string, resourceType string, namespace string) bool {
	if resourceType == "pod" {
		return CheckIfPodRunning(kubeClient, resourceName, namespace)
	} else {
		return false
	}
}

// Utility function to execute a command within the mentioned container of a given namespace
func KubectlExecCmd(k8sRestConfig *rest.Config, podName string, containerName string, namespace string, cmdString string, logger *zap.Logger) string {
	buf := &bytes.Buffer{}
	errBuf := &bytes.Buffer{}

	kubeClient, err := kubernetes.NewForConfig(k8sRestConfig)
	if err != nil {
		logger.Fatal("Could not get the kube client", zap.Any("ERROR", err))
	}
	if CheckK8sResource(kubeClient, podName, "pod", namespace) {
		execRequest := kubeClient.CoreV1().RESTClient().Post().Resource("pods").Name(podName).Namespace(namespace).SubResource("exec").VersionedParams(&v1.PodExecOptions{
			Command: []string{"/bin/sh", "-c", cmdString},
			Stdin:   false,
			Stdout:  true,
			Stderr:  true,
			TTY:     false,
		}, scheme.ParameterCodec)
		exec, err := remotecommand.NewSPDYExecutor(k8sRestConfig, "POST", execRequest.URL())
		if err != nil {
			logger.Fatal("Could not get the bidirectional multiplexed stream from the corresponding k8s rest config", zap.Any("ERROR", err))
		}
		err = exec.Stream(remotecommand.StreamOptions{
			Stdout: buf,
			Stderr: errBuf,
		})
		if err != nil {
			logger.Fatal("Failed executing command: "+cmdString+" on "+containerName+":"+podName+" in the namespace "+namespace, zap.Any("ERROR", err))
		}
		return buf.String()
	} else {
		logger.Fatal("The requested pod "+podName+" doesn't exist in the namespace "+namespace, zap.Any("ERROR", err))
		return ""
	}
}
