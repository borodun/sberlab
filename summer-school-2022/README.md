## K8s installation 
### K0sctl
#### Installation
[Docs](https://github.com/k0sproject/k0sctl)

Create **_default.yaml_** config and change it according to your needs, using [example config](https://docs.k0sproject.io/v1.23.8+k0s.0/configuration/) or [**k0s-config.yaml**](./k0s-config.yaml):
```shell
k0sctl init --k0s > default.yaml
```
Add _hosts_, change _k0s version_ to a stable one and change _network provider_ to _calico_. See [**k0s-config.yaml**](./k0s-config.yaml).
```yaml
k0s:
  version: 1.23.8+k0s.0
  config:
    network:
      provider: calico
      podCIDR: 10.244.0.0/16
      serviceCIDR: 10.96.0.0/12
```
Apply config:
```shell
k0sctl apply --config default.yaml
```
Get cluster config for _kubectl_:
```shell
k0sctl kubeconfig > kubeconfig
```
Move kubeconfig to **_.kube_** if you doesn't have installed other k8s clusters:
```shell
cp kubeconfig ~/.kube/config
```
Or add to an existing config:
```shell
cp ~/.kube/config ~/.kube/config.backup
KUBECONFIG=/path/to/new/kubeconfig:~/.kube/config.backup kubectl config view --flatten > ~/.kube/config
```
Delete k8s cluster:
```shell
k0sctl reset --config default.yaml
```

#### Testing
Try some _kubectl_ commands:
```shell
alias k=kubectl 
k get nodes
k get all --all-namespaces
k apply -f manifests/
```
Open _nat-ip_ in browser and check that default nginx page is displayed. Your NAT should allow connection external-ip:30080 -> internal-ip:30080 \
Check k8s network with [kube-detective](https://github.com/sapcc/kube-detective):
```shell
./kube-detective
```

Try to run [conformance tests](https://github.com/cncf/k8s-conformance) with [sonobuoy](https://github.com/vmware-tanzu/sonobuoy):
```shell
./sonobuoy run --mode=certified-conformance --wait
./sonobuoy retrieve
./sonobuoy results archive.tar.gz --mode=dump
./sonobuoy results archive.tar.gz --mode=detailed | jq 'select(.status=="passed")'
./sonobuoy delete --wait
```

### Minikube
[Docs](https://minikube.sigs.k8s.io/docs/start/)



