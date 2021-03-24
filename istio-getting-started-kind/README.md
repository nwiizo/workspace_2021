```
 kind create cluster --config kind-config.yaml --name istio-getting-started
 kind get kubeconfig --name istio-getting-started > kubeconfig.yaml
 curl -L https://istio.io/downloadIstio | sh -
 cd istio-1.9.1
 cp bin/istioctl /usr/bin/
 istioctl install --set profile=demo -y
 kubectl label namespace default istio-injection=enabled
 kubectl get services,pods
 kubectl exec ratings-v1-b6994bb9-swqz9 -c ratings -- curl -sS productpage:9080/productpage | grep -o "<title>.*</title>"
 kubectl apply -f istio-1.9.1/samples/bookinfo/networking/bookinfo-gateway.yaml
```
