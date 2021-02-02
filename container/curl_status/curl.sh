#!/bin/sh
for i in `kubectl get ingress -A -o go-template="{{range .items}}{{range .spec.rules}}{{.host}} {{end}}{{end}}"`;do
    curl -LI ${i} -o /dev/null -w '%{http_code}\n' -s
    echo ${i}
done
