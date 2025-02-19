/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/mohdcode/kubeAnalyzer/cmd"
	_ "github.com/mohdcode/kubeAnalyzer/cmd/kubeAnalyzer/kubercmd"
	_ "github.com/mohdcode/kubeAnalyzer/cmd/kubeAnalyzer/kubercmd/apis"
	_ "github.com/mohdcode/kubeAnalyzer/cmd/kubeAnalyzer/kubercmd/comparison"
	_ "github.com/mohdcode/kubeAnalyzer/cmd/kubeAnalyzer/kubercmd/configmap"
	_ "github.com/mohdcode/kubeAnalyzer/cmd/kubeAnalyzer/kubercmd/currentAPIversion"
	_ "github.com/mohdcode/kubeAnalyzer/cmd/kubeAnalyzer/kubercmd/daemonsets"
	_ "github.com/mohdcode/kubeAnalyzer/cmd/kubeAnalyzer/kubercmd/deployments"
	_ "github.com/mohdcode/kubeAnalyzer/cmd/kubeAnalyzer/kubercmd/ingress"
	_ "github.com/mohdcode/kubeAnalyzer/cmd/kubeAnalyzer/kubercmd/pods"
	_ "github.com/mohdcode/kubeAnalyzer/cmd/kubeAnalyzer/kubercmd/prefferedAPIversion"
	_ "github.com/mohdcode/kubeAnalyzer/cmd/kubeAnalyzer/kubercmd/services"
	_ "github.com/mohdcode/kubeAnalyzer/cmd/kubeAnalyzer/kubercmd/statefulsets"
)

func main() {
	cmd.Execute()
}
