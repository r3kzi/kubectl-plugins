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
	"context"
	"fmt"
	"github.com/spf13/cobra"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"os"
	"path/filepath"
	"text/tabwriter"
)

const annotation = "eks.amazonaws.com/role-arn"

var irsa = &cobra.Command{
	Use:   "irsa",
	Short: "Shows all ServiceAccounts that use IAM Roles for Service Accounts",
	RunE:  run,
}

func run(cmd *cobra.Command, _ []string) error {
	kubeconfig, err := cmd.Flags().GetString("kubeconfig")
	if err != nil {
		return err
	}

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}

	serviceAccountList, err := clientset.CoreV1().ServiceAccounts(apiv1.NamespaceAll).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return err
	}

	printServiceAccounts(serviceAccountList)

	return nil
}

func init() {
	if home := homedir.HomeDir(); home != "" {
		irsa.Flags().String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		irsa.Flags().String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
}

func Execute() {
	if err := irsa.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func printServiceAccounts(serviceAccounts *apiv1.ServiceAccountList) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.TabIndent)
	fmt.Fprintln(w, "NAME\tNAMESPACE\tIAM ROLE\t")
	for _, serviceAccount := range serviceAccounts.Items {
		if iamRole, ok := serviceAccount.Annotations[annotation]; ok {
			fmt.Fprintln(w, fmt.Sprintf("%s\t%s\t%s\t", serviceAccount.Name, serviceAccount.Namespace, iamRole))
		}
	}
	w.Flush()
}
