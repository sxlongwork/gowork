package resinformer

import (
	"fmt"
	"k8s-manage/server/client"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/tools/cache"
)

var podslist []*corev1.Pod

func ResourseInformer() {

	// 初始化 client
	clientset := client.GetClient()
	stopper := make(chan struct{})
	defer close(stopper)

	// 初始化 informe
	factory := informers.NewSharedInformerFactory(clientset, 0)
	podInformer := factory.Core().V1().Pods()
	informer := podInformer.Informer()
	defer runtime.HandleCrash()

	// 启动 informer，list & watch
	go factory.Start(stopper)

	// 从 apiserver 同步资源，即 list
	if !cache.WaitForCacheSync(stopper, informer.HasSynced) {
		runtime.HandleError(fmt.Errorf("Timed out waiting for caches to sync"))
		return
	}

	// 使用自定义 handle
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    onAdd,
		UpdateFunc: onUpdate,
		DeleteFunc: onDelete,
	})

	// 创建 lister
	podLister := podInformer.Lister().Pods("default")
	// 从 lister 中获取所有 items
	var set = make(map[string]string)
	set["app"] = "wechat"
	lebelSelector := labels.SelectorFromSet(set)

	podList, err := podLister.List(lebelSelector)
	if err != nil {
		fmt.Println(err)
	}
	podslist = podList
	fmt.Println("podlist:", podList)
	<-stopper
}

func onAdd(obj interface{}) {
	pod := obj.(*corev1.Pod)
	fmt.Println("add a pod:", pod.Name, pod.Status.Phase)

}

func onUpdate(old, new interface{}) {
	oldPod := old.(*corev1.Pod)
	newPod := new.(*corev1.Pod)
	fmt.Println("onUpdate:", oldPod.Name, oldPod.Status.Phase, newPod.Name, newPod.Status.Phase)
}

func onDelete(obj interface{}) {
	pod := obj.(*corev1.Pod)
	fmt.Println("onDelete:", pod.Name, pod.Status.Phase)
	fmt.Println("after updated:", podslist)
}
