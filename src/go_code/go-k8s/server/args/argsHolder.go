package args

var Holder = &argsHolder{}

type argsHolder struct {
	kubeconfig   string
	insecureHost string
	insecurePort string
	secureHost   string
	securePort   string
	certFile     string
	keyFile      string
}

func (argshd *argsHolder) GetKubeconfig() string {
	return argshd.kubeconfig
}

func (argshd *argsHolder) GetInsecureHost() string {
	return argshd.insecureHost
}

func (argshd *argsHolder) GetInsecurePort() string {
	return argshd.insecurePort
}

func (argshd *argsHolder) GetSecureHost() string {
	return argshd.secureHost
}

func (argshd *argsHolder) GetSecurePort() string {
	return argshd.securePort
}

func (argshd *argsHolder) GetCertFile() string {
	return argshd.certFile
}

func (argshd *argsHolder) GetKeyFile() string {
	return argshd.keyFile
}
