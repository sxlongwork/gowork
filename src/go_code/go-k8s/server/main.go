package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"k8s-manage/server/args"
	"k8s-manage/server/handler"
	resinformer "k8s-manage/server/handler/informer"
	"log"
	"net/http"
)

var (
	kubeconfig   = flag.String("kubeconfig", "/root/.kube/config", "absolute path to the kubeconfig file")
	insecureHost = flag.String("insecureHost", "0.0.0.0", "The IP address on which to serve the --insecurePort")
	insecurePort = flag.String("insecurePort", "8088", "The port to listen to for incoming HTTP requests.")
	secureHost   = flag.String("secureHost", "0.0.0.0", "The IP address on which to serve the --securePort")
	securePort   = flag.String("securePort", "8443", "The port to listen to for incoming HTTPS requests.")
	certFile     = flag.String("tls-cert-file", "", "File containing the default x509 Certificate for HTTPS.")
	keyFile      = flag.String("tls-key-file", "", "File containing the default x509 private key matching --tls-cert-file.")
)

func main() {

	flag.Parse()
	// utils.ConfParse()
	initHolder()
	go resinformer.ResourseInformer()
	// fmt.Println(args.Holder.GetKubeconfig())
	apiHandler := handler.CreateHttpAPIHandler()

	http.Handle("/api/", apiHandler)
	if args.Holder.GetCertFile() != "" && args.Holder.GetKeyFile() != "" {
		log.Printf("Serving securely on HTTPS port: %s", args.Holder.GetSecurePort())
		secureAddr := fmt.Sprintf("%s:%s", args.Holder.GetSecureHost(), args.Holder.GetSecurePort())
		serverCert, err := tls.LoadX509KeyPair(args.Holder.GetCertFile(), args.Holder.GetKeyFile())
		if err != nil {
			log.Fatal("get serverCerts failed")
		}
		server := &http.Server{
			Addr:    secureAddr,
			Handler: http.DefaultServeMux,
			TLSConfig: &tls.Config{
				Certificates: []tls.Certificate{serverCert},
				MinVersion:   tls.VersionTLS12,
			},
		}
		go func() { log.Fatal(server.ListenAndServeTLS("", "")) }()
	} else {
		insecureAddr := fmt.Sprintf("%s:%s", args.Holder.GetInsecureHost(), args.Holder.GetInsecurePort())
		log.Printf("Serving insecurely on HTTP port %s", args.Holder.GetInsecurePort())
		server := &http.Server{
			Addr:    insecureAddr,
			Handler: http.DefaultServeMux,
		}
		go func() { log.Fatal(server.ListenAndServe()) }()
	}
	select {}

}

func initHolder() {
	builder := args.GetBuilder()
	builder.SetKubeconfig(*kubeconfig)
	builder.SetInsecureHost(*insecureHost)
	builder.SetInsecurePort(*insecurePort)
	builder.SetSecureHost(*secureHost)
	builder.SetSecurePort(*securePort)
	builder.SetCertFile(*certFile)
	builder.SetKeyFile(*keyFile)
}
