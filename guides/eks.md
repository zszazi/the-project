pre-req
- eksctl 
- aws cli 

To Auth with AWS CLI 

```bash
aws configure
```

Create EKS Cluster - Takes around 20 mins to create
```
eksctl create cluster --name the-project --region us-west-2
```

Or do it manually 

I. IAM Roles 

1. Search for Roles -> create role -> Use case : EKS - then EKS - Cluster -> give it a name -> click on Create 
2. Create new role -> Use Case EC2 - Just vanilla EC2 -> add below policies AmazonEKSWorkerNodePolicy, AmazonEC2ContainerRegistryReadOnly and AmazonEKS_CNI_Policy -> Next -> give it a name -> Create 

II. EKS Cluster 

3. search for EKS -> Add Cluter - Create -> give it a name, cluster service role select from step 1. -> Next Next ..Ne..Xt -> Create 

Takes up around 10 to 15 mins for cluster to be created. 

III. Create Node Group 

4. Compute tab -> Add Node group -> give a node group name and select Node IAM role as from Step 2 -> Next -> give correct instance t3.medium , t3.micro -> Create . Give it 5 mins to be fully up 


Configure kubectl to point to kubernetes cluster , first two not required if you used aws configure
```
export AWS_ACCESS_KEY_ID=${Access Key ID}
export AWS_SECRET_ACCESS_KEY=${Secret Access Key}
export AWS_DEFAULT_REGION=us-west-2
export KUBECONFIG=kubeconfig
```

Set k8s Context
```
aws eks update-kubeconfig --name the-project
```
Now you can access the cluster from kubectl

### Install Nginx Controller 

```
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.11.1/deploy/static/provider/aws/deploy.yaml
```

### Apply app K8s manifest 

Before you apply make sure to build the app using instruction in docker.md

Helm Chart - 

```helm install go-web-app helm/go-web-app-chart```

**OR**

Vanilla K8s Manifests -

```kubectl apply -f k8s/manifests/```

Give it some time, couple of mins. 

List the ingress 

```kubectl get ing```

Output
```
NAME         CLASS   HOSTS              ADDRESS                                                                         PORTS   AGE
go-web-app   nginx   go-web-app.local   acc61d95fd3bc416c80e9328f65c8784-2749dcfda409de0c.elb.us-west-2.amazonaws.com   80      2m47s
```
do a nslookup on the ing fqdn
eg 

```nslookup acc61d95fd3bc416c80e9328f65c8784-2749dcfda409de0c.elb.us-west-2.amazonaws.com```

Output
```
Server:         2001:420:200:1::a
Address:        2001:420:200:1::a#53

Non-authoritative answer:
Name:   acc61d95fd3bc416c80e9328f65c8784-2749dcfda409de0c.elb.us-west-2.amazonaws.com
Address: 52.41.153.228
Name:   acc61d95fd3bc416c80e9328f65c8784-2749dcfda409de0c.elb.us-west-2.amazonaws.com
Address: 52.43.201.236
Name:   acc61d95fd3bc416c80e9328f65c8784-2749dcfda409de0c.elb.us-west-2.amazonaws.com
Address: 34.223.29.148
Name:   acc61d95fd3bc416c80e9328f65c8784-2749dcfda409de0c.elb.us-west-2.amazonaws.com
Address: 52.13.230.90
```
add below line to /etc/hosts file 

```
52.13.230.90    go-web-app.local
```

Now open ```go-web-app.local/courses``` in the browser 