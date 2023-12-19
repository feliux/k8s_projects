# Kubernetes

**Certs**

```sh
$ openssl x509 -inform der -in smartdiagnosis.txt -out smartdiagnosis-dev.crt

$ kubectl create configmap-cacerts -n jira --from-file=cacerts=./cacerts-nueva --dry-run=client -o yaml > configmap-cacerts.yaml
$ kubectl create secret tls smart-diagnosis-cert --key=smartdiagnosis.key --cert=smartdiagnosis.crt -n humsapps-qa-release-1 --dry-run=client -o yaml > secret-smart-cert.yaml
```

```sg
kubeadm init phase certs all -h
```
