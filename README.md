# kubectl-plugins

Some of my plugins for `kubectl`.

## Installation 

To install the plugin, clone the repo and simply move this executable file to anywhere on your `PATH`.

## Check Installation

`kubectl` provides a command `kubectl plugin list` that searches your `PATH` for valid plugin executables. 

Usage:
```shell script
$ kubectl plugin list
The following compatible plugins are available:

/home/chris/.local/bin/kubectl-bash
```

## Plugins

### kubectl bash

When you want to access kubernetes resources inside a cluster, you either:
a) need to expose them via e.g. an Ingress
b) run kube-proxy
c) forward ports between your local machine and services running inside the cluster.

Via kubectl bash you can run plain curl command without exposing any services or forwarding ports.

It works by running a pod with radial/busyboxplus:curl image inside the cluster.

Usage: 
```shell script
$ kubectl bash
```

