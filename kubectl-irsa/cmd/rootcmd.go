/*
Copyright © 2020 Christian Niehoff <mail@christian-niehoff.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"os"
	"path/filepath"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const irsaAnnotation = "eks.amazonaws.com/role-arn"

var irsa = &cobra.Command{
	Use:   "irsa",
	Short: "Shows all ServiceAccounts that use IAM Roles for Service Accounts",
	Run:   run,
}

func run(cmd *cobra.Command, args []string) {
	kubeconfig, _ := cmd.Flags().GetString("kubeconfig")

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	serviceAccounts := clientset.CoreV1().ServiceAccounts(apiv1.NamespaceAll)
	serviceAccountList, _ := serviceAccounts.List(metav1.ListOptions{})

	for _, serviceAccount := range serviceAccountList.Items {
		if _, ok := serviceAccount.Annotations[irsaAnnotation]; ok {
			fmt.Println(serviceAccount.Name, serviceAccount.Namespace)
		}
	}
}

func Execute() {
	if home := homedir.HomeDir(); home != "" {
		irsa.Flags().String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		irsa.Flags().String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	if err := irsa.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
