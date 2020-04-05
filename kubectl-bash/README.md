# kubectl-bash

When you want to access kubernetes resources inside a cluster, you either:

- need to expose them via e.g. an Ingress
- run kube-proxy
- forward ports between your local machine and services running inside the cluster.

Via kubectl bash you can run plain `curl` command without exposing any services or forwarding ports.

It works by running a pod with `radial/busyboxplus:curl` image inside the cluster.

Usage: 
```shell script
$ kubectl bash
```
