# kubectl-irsa

Shows all `ServiceAccounts` that use [IAM Roles for Service Accounts](https://docs.aws.amazon.com/eks/latest/userguide/iam-roles-for-service-accounts.html).

Usage:
```shell script
$ kubectl irsa
NAME            NAMESPACE       IAM ROLE                        
irsa            default         arn:aws:iam::12345789:role/test 
irsa-2          default         arn:aws:iam::12345789:role/test 
```