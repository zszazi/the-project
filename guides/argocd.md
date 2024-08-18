### Argo CD

In this project Argocd is on the same cluster as other app, but in a different namespace

We have to set this up in the EKS cluster 


>Create new NS and install argo cd for CD
```
kubectl create namespace argocd
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
```

>make the argo CD UI Accessible 

```
kubectl patch svc argocd-server -n argocd -p '{"spec": {"type": "LoadBalancer"}}'
```

> Get LB IP 

```
kubectl get svc argocd-server -n argocd
```

Output will be something like this 
```
NAME            TYPE           CLUSTER-IP     EXTERNAL-IP                                                               PORT(S)                      AGE
argocd-server   LoadBalancer   10.100.220.1   a8ae587b500314889b9df0ff224f1778-1354122885.us-west-2.elb.amazonaws.com   80:31190/TCP,443:30828/TCP   46s
```
> Get the External IP (eg a8ae587b500314889b9df0ff224f1778-1354122885.us-west-2.elb.amazonaws.com) and paste it in Browser, it will take some time for ELB to register and expose on internet

> Login to Argo UI 
username - admin

Password - (Initial Password is stored as secret)

```
kubectl get secret -n argocd argocd-initial-admin-secret -o yaml
```

data.password field is the password but is base64 encoded 

```
echo <some-Random-string> | base64 -D
```
copy the decoded value without the string at the end 


>Login to UI and Setup CD for the-project

I. Connect the private repo 

Settings -> Repositories -> Connect Repo -> Fill the username and password (github PAT), and test for successful connection

II. Setup CD 

New App -> Give a name -> Sync Policy make it Automatic -> Check the Self Heal -> Select the Github Repo from dropdown which we added in step I, and for PATH , Argocd will auto detect helm chart just accept it 

Give Cluster URL as https://kubernetes.default.svc (which means we are deploying on same EKS cluster) and namespace as default, also under HELM , select the values.yaml from the dropdown , then finally click CREATE