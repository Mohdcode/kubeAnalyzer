# kubeAnalyzer

kubeAnalyzer is a command-line tool for analyzing the health of Kubernetes resources ,API endpoints and API versions in a given namespace. It provides detailed health reports for different Kubernetes resources and checks the status of API endpoints.

## Installation

To install kubeAnalyzer, follow these steps:

1. Clone the repository:

```bash
git clone https://github.com/mohdcode/kubeAnalyzer.git
```
Change to the project directory:
```bash
cd kubeAnalyzer
```
Build the project:
```bash
go build
```
Move the binary to a directory in your system's PATH:
```bash
go install
```
## Usage
To use kubeAnalyzer, you need to have a valid kubeconfig file with access to your Kubernetes cluster.

```bash
kubeAnalyzer check [command]
```
### Commands

#### Preferred APIversion command
Check preferred api-versions for all resources.

```bash
kubeAnalyzer check recomapiv
```

![preferred api demo](./demo/recom.png)

NOTE: Providing kubeconfig path or namespace is optional if you dont provide config path it will try to get path from ~/.kube/config, and if you dont provide any namespace it will autometically serlect the default namespace.

#### Compare APIversion command
Check preferred api-version and current api version are same or not for all resources.

```bash
kubeAnalyzer check compareapiv
```
![compare demo](./demo/compare.png)

#### Current APIversion command
Check current api version for all resources.

```bash
kubeAnalyzer check currentapiv
```

![current api demo](./demo/current.png)

#### API Command
Check the preferred Kubernetes API versions for resources in a given namespace.

```bash
kubeAnalyzer check api --config=<path/to/kubeconfig> --ns=<namespace>
```
### Pods Command
Analyze the health of Pods in a given namespace.

```bash
kubeAnalyzer check pods --config=<path/to/kubeconfig> --ns=<namespace>
```
![pods demo](./demo/pods.png)

#### Services Command
Analyze the health of Services in a given namespace.

```bash
kubeAnalyzer check services --config=<path/to/kubeconfig> --ns=<namespace>
```
![service demo](./demo/service.png)

#### Deployments Command
Analyze the health of Deployments in a given namespace.

```bash
kubeAnalyzer check deployments --config=<path/to/kubeconfig> --ns=<namespace>
```
#### Daemonsets Command
Analyze the health of DaemonSets in a given namespace.

```bash
kubeAnalyzer check daemonsets --config=<path/to/kubeconfig> --ns=<namespace>
```
#### Statefulsets command
Analyze the health of Statefulsets in a given namespace.

```bash
kubeAnalyzer check statefulsets --config=<path/to/kubeconfig> --ns=<namespace>
```

#### Configmaps Command
Analyze the health of ConfigMaps in a given namespace.

```bash
kubeAnalyzer check configmaps --config=<path/to/kubeconfig> --ns=<namespace>
```
#### Ingress Command
Analyze the health of Ingress resources in a given namespace.

```bash
kubeAnalyzer check ingress --config=<path/to/kubeconfig> --ns=<namespace>
```

We will be add more helath analysis features.
Feel free to contribute..