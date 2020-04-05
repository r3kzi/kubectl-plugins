# kubectl-plugins

Some of my plugins for `kubectl`.

## Installation 

To install the plugin, clone the repo and simply move this executable files to anywhere on your `PATH`.

## Check Installation

`kubectl` provides a command `kubectl plugin list` that searches your `PATH` for valid plugin executables. 

Usage:
```shell script
$ kubectl plugin list
The following compatible plugins are available:

/home/chris/.local/bin/kubectl-bash
/home/chris/.local/bin/kubectl-irsa
```

## Plugins

- [kubectl-bash](./kubectl-bash/README.md)
- [kubectl-irsa](./kubectl-irsa/README.md)