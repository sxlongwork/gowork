package args

var builder = &argsBuilder{holder: Holder}

type argsBuilder struct {
	holder *argsHolder
}

func (builder *argsBuilder) SetKubeconfig(kubeconfig string) {
	builder.holder.kubeconfig = kubeconfig
}

func (builder *argsBuilder) SetInsecureHost(insecureHost string) {
	builder.holder.insecureHost = insecureHost
}

func (builder *argsBuilder) SetInsecurePort(insecurePort string) {
	builder.holder.insecurePort = insecurePort
}

func (builder *argsBuilder) SetSecureHost(secureHost string) {
	builder.holder.secureHost = secureHost
}

func (builder *argsBuilder) SetSecurePort(securePort string) {
	builder.holder.securePort = securePort
}

func (builder *argsBuilder) SetCertFile(certFile string) {
	builder.holder.certFile = certFile
}

func (builder *argsBuilder) SetKeyFile(keyFile string) {
	builder.holder.keyFile = keyFile
}

func GetBuilder() *argsBuilder {
	return builder
}
