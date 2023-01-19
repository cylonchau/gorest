module github.io/cylonchau/kube-haproxy

go 1.18

require (
	github.com/haproxytech/models v1.2.5-0.20191122125615-30d0235b81ec
	k8s.io/api v0.19.10
	k8s.io/apimachinery v0.19.10
	k8s.io/apiserver v0.19.10
	k8s.io/client-go v0.19.10
	k8s.io/klog/v2 v2.80.1
	k8s.io/kubernetes v1.19.10
	k8s.io/utils v0.0.0-20221107191617-1a15be271d1d
)

require (
	github.com/PuerkitoBio/purell v1.1.1 // indirect
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect
	github.com/antelman107/net-wait-go v0.0.0-20220211074630-12d8a944b87d // indirect
	github.com/asaskevich/govalidator v0.0.0-20190424111038-f61b66f89f4a // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/blang/semver v3.5.0+incompatible // indirect
	github.com/cespare/xxhash/v2 v2.1.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-openapi/analysis v0.19.5 // indirect
	github.com/go-openapi/errors v0.19.2 // indirect
	github.com/go-openapi/jsonpointer v0.19.3 // indirect
	github.com/go-openapi/jsonreference v0.19.3 // indirect
	github.com/go-openapi/loads v0.19.4 // indirect
	github.com/go-openapi/runtime v0.19.4 // indirect
	github.com/go-openapi/spec v0.19.3 // indirect
	github.com/go-openapi/strfmt v0.19.3 // indirect
	github.com/go-openapi/swag v0.19.5 // indirect
	github.com/go-openapi/validate v0.19.5 // indirect
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/gofuzz v1.1.0 // indirect
	github.com/google/uuid v1.1.2 // indirect
	github.com/googleapis/gnostic v0.4.1 // indirect
	github.com/hashicorp/golang-lru v0.5.1 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/mailru/easyjson v0.7.0 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.2-0.20181231171920-c182affec369 // indirect
	github.com/mitchellh/mapstructure v1.1.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/prometheus/client_golang v1.7.1 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.10.0 // indirect
	github.com/prometheus/procfs v0.1.3 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/stretchr/testify v1.6.1 // indirect
	go.mongodb.org/mongo-driver v1.1.2 // indirect
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9 // indirect
	golang.org/x/net v0.3.1-0.20221206200815-1e63c2f08a10 // indirect
	golang.org/x/oauth2 v0.0.0-20220223155221-ee480838109b // indirect
	golang.org/x/sys v0.3.0 // indirect
	golang.org/x/text v0.5.0 // indirect
	golang.org/x/time v0.0.0-20220210224613-90d013bbcef8 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k8s.io/component-base v0.19.10 // indirect
	k8s.io/kube-openapi v0.0.0-20200805222855-6aeccd4b50c6 // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.2.3 // indirect
	sigs.k8s.io/yaml v1.3.0 // indirect
)

replace k8s.io/api => k8s.io/api v0.19.10

replace k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.19.10

replace k8s.io/apimachinery => k8s.io/apimachinery v0.19.11-rc.0

replace k8s.io/apiserver => k8s.io/apiserver v0.19.10

replace k8s.io/cli-runtime => k8s.io/cli-runtime v0.19.10

replace k8s.io/client-go => k8s.io/client-go v0.19.10

replace k8s.io/cloud-provider => k8s.io/cloud-provider v0.19.10

replace k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.19.10

replace k8s.io/code-generator => k8s.io/code-generator v0.19.13-rc.0

replace k8s.io/component-base => k8s.io/component-base v0.19.10

replace k8s.io/controller-manager => k8s.io/controller-manager v0.19.17-rc.0

replace k8s.io/cri-api => k8s.io/cri-api v0.19.13-rc.0

replace k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.19.10

replace k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.19.10

replace k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.19.10

replace k8s.io/kube-proxy => k8s.io/kube-proxy v0.19.10

replace k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.19.10

replace k8s.io/kubectl => k8s.io/kubectl v0.19.10

replace k8s.io/kubelet => k8s.io/kubelet v0.19.10

replace k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.19.10

replace k8s.io/metrics => k8s.io/metrics v0.19.10

replace k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.19.10

replace k8s.io/sample-cli-plugin => k8s.io/sample-cli-plugin v0.19.10

replace k8s.io/sample-controller => k8s.io/sample-controller v0.19.10
