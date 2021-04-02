kubebuilder init --domain xbc.io
kubebuilder create api --group opswagon --version v1 --kind JenkinsPipeline
kubebuilder create api --group opswagon --version v1 --kind JenkinsPlugin
# todo : edit  crd.go, and
make && make manifests
