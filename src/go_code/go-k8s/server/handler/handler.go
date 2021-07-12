package handler

import (
	"k8s-manage/server/auth"
	"k8s-manage/server/client"
	"k8s-manage/server/resources/configmap"
	deploy "k8s-manage/server/resources/deployment"
	"k8s-manage/server/resources/namespace"
	"k8s-manage/server/resources/node"
	"k8s-manage/server/resources/pod"
	"k8s-manage/server/resources/replicaset"
	"k8s-manage/server/resources/secret"
	"k8s-manage/server/resources/service"
	"k8s-manage/server/resources/statefulset"
	"k8s-manage/server/utils"
	"log"
	"net/http"

	"github.com/emicklei/go-restful/v3"
)

func CreateHttpAPIHandler() http.Handler {
	wsContainer := restful.NewContainer()

	cors := restful.CrossOriginResourceSharing{
		AllowedHeaders: []string{"Content-Type", "Accept"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedDomains: []string{"*"},
		CookiesAllowed: false,
		Container:      wsContainer}
	wsContainer.Filter(cors.Filter)
	wsContainer.Filter(wsContainer.OPTIONSFilter)
	wsContainer.Filter(headerFilter)

	wsContainer.EnableContentEncoding(true)

	authv1 := new(restful.WebService)
	authv1.Path("/api").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	wsContainer.Add(authv1)

	// token
	authv1.Route(
		authv1.POST("token").
			To(handleGetToken).
			Writes(auth.Token{}))
	authv1.Route(
		authv1.GET("").
			To(handleAPI).
			Returns(http.StatusOK, "server is running", nil))

	apiV1 := new(restful.WebService)
	apiV1.Filter(tokenVerifyFilter)

	apiV1.Path("/api/v1").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)
	wsContainer.Add(apiV1)

	// pod
	apiV1.Route(
		apiV1.GET("pod").
			To(handleGetPods).
			Writes(pod.PodList{}))
	apiV1.Route(
		apiV1.GET("pod/{namespace}").
			To(handleGetPods).
			Writes(pod.PodList{}))
	apiV1.Route(
		apiV1.GET("pod/{namespace}/{pod}").
			To(handleGetPodDetail).
			Writes(pod.PodDetail{}))
	apiV1.Route(
		apiV1.GET("pod/{namespace}/{pod}/container").
			To(handleGetPodContainers).
			Writes(pod.PodDetail{}))

	// namespace
	apiV1.Route(
		apiV1.GET("namespace").
			To(handlerGetNamespaces).
			Writes(namespace.NamespaceList{}))
	apiV1.Route(
		apiV1.GET("namespace/{name}").
			To(handlerGetNamespaceDetail).
			Writes(namespace.NamespaceDetail{}))

	// node
	apiV1.Route(
		apiV1.GET("node").
			To(handleGetNodes).
			Writes(node.NodeList{}))
	apiV1.Route(
		apiV1.GET("node/{name}").
			To(handleGetNodeDetail).
			Writes(node.NodeDetail{}))
	apiV1.Route(
		apiV1.GET("node/{name}/pods").
			To(handlerGetNodePods).
			Writes(pod.PodList{}))

	// deployment
	apiV1.Route(
		apiV1.GET("deployment").
			To(handleGetDeployments).
			Writes(deploy.DeploymentList{}))
	apiV1.Route(
		apiV1.GET("deployment/{namespace}").
			To(handleGetDeployments).
			Writes(deploy.DeploymentList{}))
	apiV1.Route(
		apiV1.GET("deployment/{namespace}/{deployment}").
			To(handleGetDeploymentDetail).
			Writes(deploy.DeploymentDetail{}))
	apiV1.Route(
		apiV1.GET("deployment/{namespace}/{deployment}/oldreplicaset").
			To(handleGetDeploymentOldReplicaset).
			Writes(replicaset.ReplicaSetList{}))
	apiV1.Route(
		apiV1.GET("deployment/{namespace}/{deployment}/newreplicaset").
			To(handleGetDeploymentNewReplicaset).
			Writes(replicaset.ReplicaSet{}))

	// statefulset
	apiV1.Route(
		apiV1.GET("statefulset").
			To(handleGetStatefulSetList).
			Writes(statefulset.StatefulSetList{}))
	apiV1.Route(
		apiV1.GET("statefulset/{namespace}").
			To(handleGetStatefulSetList).
			Writes(statefulset.StatefulSetList{}))
	apiV1.Route(
		apiV1.GET("statefulset/{namespace}/{statefulset}").
			To(handleGetStatefulsetDetail).
			Writes(statefulset.StatefulSetDetail{}))
	apiV1.Route(
		apiV1.GET("statefulset/{namespace}/{statefulset}/pod").
			To(handleGetStatefulSetPods).
			Writes(pod.PodList{}))

	// service
	apiV1.Route(
		apiV1.GET("service/").
			To(handleGetServices).
			Writes(service.ServiceList{}))
	apiV1.Route(
		apiV1.GET("service/{namespace}").
			To(handleGetServices).
			Writes(service.ServiceList{}))
	apiV1.Route(
		apiV1.GET("service/{namespace}/{service}").
			To(handleGetServiceDetail).
			Writes(service.ServiceDetail{}))
	apiV1.Route(
		apiV1.GET("service/{namespace}/{service}/pod").
			To(handleGetServicePod).
			Writes(pod.PodList{}))

	// secret
	apiV1.Route(
		apiV1.GET("secret").
			To(handleGetSecretList).
			Writes(secret.SecretList{}))
	apiV1.Route(
		apiV1.GET("secret/{namespace}").
			To(handleGetSecretList).
			Writes(secret.SecretList{}))
	apiV1.Route(
		apiV1.GET("secret/{namespace}/{secret}").
			To(handleGetSecretDetail).
			Writes(secret.SecretDetail{}))
	// configmap
	apiV1.Route(
		apiV1.GET("configmap").
			To(handleGetConfigMapList).
			Writes(configmap.ConfigMapList{}))
	apiV1.Route(
		apiV1.GET("configmap/{namespace}").
			To(handleGetConfigMapList).
			Writes(configmap.ConfigMapList{}))
	apiV1.Route(
		apiV1.GET("configmap/{namespace}/{configmap}").
			To(handleGetConfigMapDetail).
			Writes(configmap.ConfigMapDetail{}))

	log.Println("api server init success")

	return wsContainer

}

func headerFilter(request *restful.Request, response *restful.Response, chain *restful.FilterChain) {
	response.AddHeader("Access-Control-Allow-Origin", "*")
	chain.ProcessFilter(request, response)
}

func handleAPI(request *restful.Request, response *restful.Response) {
	resultStr, err := utils.GetJsonData("api server is running", nil)
	code, result := utils.GetJsonResponse(http.StatusOK, "success", resultStr, err)

	response.WriteHeaderAndJson(code, result, "application/json;charset=UTF-8")
}

func tokenVerifyFilter(request *restful.Request, response *restful.Response, chain *restful.FilterChain) {
	htoken := request.HeaderParameter("Authorization")
	// log.Printf("request token = [%s]\n", htoken)
	if htoken == "" {
		response.WriteErrorString(http.StatusBadRequest, "not found token in request header")
		return
	}
	if !auth.IsValidToken(htoken) {
		response.WriteErrorString(http.StatusBadRequest, "invalid token")
		return
	}
	chain.ProcessFilter(request, response)
}

func handleGetToken(request *restful.Request, response *restful.Response) {
	log.Println("get request,", request.Request.URL)
	user := auth.User{}
	err := request.ReadEntity(&user)
	// log.Println(user)
	if err != nil {
		response.WriteErrorString(http.StatusBadRequest, "get token failed")
		return
	}
	result, err := auth.GetToken(&user)
	if err != nil {
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
	// TODO
	response.WriteHeaderAndEntity(http.StatusOK, result)
}

func handleGetPods(request *restful.Request, response *restful.Response) {
	var k8sClient = client.GetClient()
	namespace := request.PathParameter("namespace")
	result, err := pod.GetPodList(k8sClient, namespace)

	if err != nil {
		response.WriteErrorString(http.StatusInternalServerError, "get pods list tailed, please try again.")
		return
	}
	response.WriteHeaderAndEntity(http.StatusOK, result)
}

func handleGetPodDetail(request *restful.Request, response *restful.Response) {
	var k8sClient = client.GetClient()
	namespace := request.PathParameter("namespace")
	podName := request.PathParameter("pod")
	result, err := pod.GetPodDetail(k8sClient, namespace, podName)

	if err != nil {
		response.WriteErrorString(http.StatusInternalServerError, "get pod detail failed, please try again.")
		return
	}
	response.WriteHeaderAndEntity(http.StatusOK, result)
}

func handleGetPodContainers(request *restful.Request, response *restful.Response) {
	var k8sClient = client.GetClient()
	namespace := request.PathParameter("namespace")
	podName := request.PathParameter("pod")
	result, err := pod.GetPodContainers(k8sClient, namespace, podName)
	if err != nil {
		response.WriteErrorString(http.StatusInternalServerError, "get containers list in pod tailed, please try again")
		return
	}
	response.WriteHeaderAndEntity(http.StatusOK, result)
}

func handlerGetNamespaces(request *restful.Request, response *restful.Response) {
	var k8sClient = client.GetClient()
	result, err := namespace.GetNamespacesList(k8sClient)

	if err != nil {
		response.WriteErrorString(http.StatusInternalServerError, "get namespae list failed, please try again.")
		return
	}
	response.WriteHeaderAndEntity(http.StatusOK, result)
}

func handlerGetNamespaceDetail(request *restful.Request, response *restful.Response) {
	var k8sClient = client.GetClient()
	nsName := request.PathParameter("name")
	result, err := namespace.GetNamespaceDetail(k8sClient, nsName)

	if err != nil {
		response.WriteErrorString(http.StatusInternalServerError, "get namespae detail failed, please try again.")
		return
	}
	response.WriteHeaderAndEntity(http.StatusOK, result)
}

func handleGetNodes(request *restful.Request, response *restful.Response) {
	var k8sClient = client.GetClient()
	result, err := node.GetNodesList(k8sClient)

	if err != nil {
		response.WriteErrorString(http.StatusInternalServerError, "get nodes failed, please try again.")
		return
	}
	response.WriteHeaderAndEntity(http.StatusOK, result)
}

func handleGetNodeDetail(request *restful.Request, response *restful.Response) {
	var k8sClient = client.GetClient()
	nodeName := request.PathParameter("name")
	result, err := node.GetNodeDetail(k8sClient, nodeName)

	if err != nil {
		response.WriteErrorString(http.StatusInternalServerError, "get node detail failed, please try again.")
		return
	}
	response.WriteHeaderAndEntity(http.StatusOK, result)
}

func handlerGetNodePods(request *restful.Request, response *restful.Response) {
	var k8sClient = client.GetClient()
	nodeName := request.PathParameter("name")
	result, err := node.GetNodePods(k8sClient, nodeName)

	if err != nil {
		response.WriteErrorString(http.StatusInternalServerError, "get pods list in node failed, please try again.")
		return
	}
	response.WriteHeaderAndEntity(http.StatusOK, result)
}

func handleGetDeployments(request *restful.Request, response *restful.Response) {
	var k8sClient = client.GetClient()
	namespace := request.PathParameter("namespace")
	result, err := deploy.GetDeploymentList(k8sClient, namespace)

	if err != nil {
		response.WriteErrorString(http.StatusInternalServerError, "get deployment list failed, please try again.")
		return
	}
	response.WriteHeaderAndEntity(http.StatusOK, result)
}

func handleGetDeploymentDetail(request *restful.Request, response *restful.Response) {
	var k8sClient = client.GetClient()
	namespace := request.PathParameter("namespace")
	deployName := request.PathParameter("deployment")
	result, err := deploy.GetDeploymentDetail(k8sClient, namespace, deployName)

	if err != nil {
		response.WriteErrorString(http.StatusInternalServerError, "get deployment detail failed, please try again.")
		return
	}
	response.WriteHeaderAndEntity(http.StatusOK, result)
}

func handleGetDeploymentOldReplicaset(request *restful.Request, response *restful.Response) {
	var k8sClient = client.GetClient()
	namespace := request.PathParameter("namespace")
	deployment := request.PathParameter("deployment")
	result, err := deploy.GetDeploymentOldReplicasetList(k8sClient, namespace, deployment)

	if err != nil {
		response.WriteErrorString(http.StatusInternalServerError, "get deployment oldreplicasets failed, please try again.")
		return
	}
	response.WriteHeaderAndEntity(http.StatusOK, result)
}

func handleGetDeploymentNewReplicaset(request *restful.Request, response *restful.Response) {
	var k8sClient = client.GetClient()
	namespace := request.PathParameter("namespace")
	deployment := request.PathParameter("deployment")
	result, err := deploy.GetDeploymentNewReplicaset(k8sClient, namespace, deployment)

	if err != nil || result == nil {
		response.WriteErrorString(http.StatusInternalServerError, "get deployment newreplicasets failed, please try again.")
		return
	}
	response.WriteHeaderAndEntity(http.StatusOK, result)
}

func handleGetStatefulSetList(request *restful.Request, response *restful.Response) {
	k8sClient := client.GetClient()
	namespace := request.PathParameter("namespace")
	result, err := statefulset.GetStatefulSetList(k8sClient, namespace)
	if err != nil {
		response.WriteErrorString(http.StatusInternalServerError, "get statefulSet list failed, please try again.")
		return
	}
	response.WriteHeaderAndEntity(http.StatusOK, result)
}

func handleGetStatefulsetDetail(request *restful.Request, response *restful.Response) {
	k8sClient := client.GetClient()
	namespace := request.PathParameter("namespace")
	name := request.PathParameter("statefulset")
	result, err := statefulset.GetStatefulsetDetail(k8sClient, namespace, name)
	if err != nil {
		response.WriteErrorString(http.StatusInternalServerError, "get statefulSet detail failed, please try again.")
		return
	}
	response.WriteHeaderAndEntity(http.StatusOK, result)
}

func handleGetStatefulSetPods(request *restful.Request, response *restful.Response) {
	k8sClient := client.GetClient()
	namespace := request.PathParameter("namespace")
	name := request.PathParameter("statefulset")
	result, err := statefulset.GetStatefulsetPods(k8sClient, namespace, name)
	if err != nil {
		response.WriteErrorString(http.StatusInternalServerError, "get statefulSet pod list failed, please try again.")
		return
	}
	response.WriteHeaderAndEntity(http.StatusOK, result)
}

func handleGetServices(request *restful.Request, response *restful.Response) {
	var k8sClient = client.GetClient()
	namespace := request.PathParameter("namespace")
	result, err := service.GetServiceList(k8sClient, namespace)
	if err != nil || result == nil {
		response.WriteErrorString(http.StatusInternalServerError, "get service list failed, please try again.")
		return
	}
	response.WriteHeaderAndEntity(http.StatusOK, result)
}

func handleGetServiceDetail(request *restful.Request, response *restful.Response) {
	var k8sClient = client.GetClient()
	namespace := request.PathParameter("namespace")
	svcName := request.PathParameter("service")
	result, err := service.GetServiceDetail(k8sClient, namespace, svcName)
	if err != nil {
		response.WriteErrorString(http.StatusInternalServerError, "get service detail failed, please try again.")
		return
	}
	response.WriteHeaderAndEntity(http.StatusOK, result)
}

func handleGetServicePod(request *restful.Request, response *restful.Response) {
	k8sClient := client.GetClient()
	namespace := request.PathParameter("namespace")
	svcName := request.PathParameter("service")
	result, err := service.GetServicePod(k8sClient, namespace, svcName)
	if err != nil {
		response.WriteErrorString(http.StatusInternalServerError, "get service pods failed, please try again.")
		return
	}
	response.WriteHeaderAndEntity(http.StatusOK, result)
}

func handleGetSecretList(request *restful.Request, response *restful.Response) {
	k8sClient := client.GetClient()
	namespace := request.PathParameter("namespace")
	result, err := secret.GetSecretList(k8sClient, namespace)
	if err != nil {
		response.WriteErrorString(http.StatusInternalServerError, "get secret list failed, please try again.")
		return
	}
	response.WriteHeaderAndEntity(http.StatusOK, result)
}

func handleGetSecretDetail(request *restful.Request, response *restful.Response) {
	k8sClient := client.GetClient()
	namespace := request.PathParameter("namespace")
	name := request.PathParameter("secret")
	result, err := secret.GetSecretDetail(k8sClient, namespace, name)
	if err != nil {
		response.WriteErrorString(http.StatusInternalServerError, "get secret detail failed, please try again.")
		return
	}
	response.WriteHeaderAndEntity(http.StatusOK, result)

}

func handleGetConfigMapList(request *restful.Request, response *restful.Response) {
	k8sClient := client.GetClient()
	namespace := request.PathParameter("namespace")
	result, err := configmap.GetConfigMapList(k8sClient, namespace)
	if err != nil {
		response.WriteErrorString(http.StatusInternalServerError, "get configmap list failed, please try again.")
		return
	}
	response.WriteHeaderAndEntity(http.StatusOK, result)
}

func handleGetConfigMapDetail(request *restful.Request, response *restful.Response) {
	k8sClient := client.GetClient()
	namespace := request.PathParameter("namespace")
	cmpName := request.PathParameter("configmap")
	result, err := configmap.GetConfigMapDetail(k8sClient, namespace, cmpName)
	if err != nil {
		response.WriteErrorString(http.StatusInternalServerError, "get configmap detail failed, please try again.")
		return
	}
	response.WriteHeaderAndEntity(http.StatusOK, result)
}
