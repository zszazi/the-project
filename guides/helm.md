### Helm chart for any project 

To create helm chart with the structure, use below command

```helm create go-web-app-chart```

This will create a folder structure as below 

```
.
└── go-web-app-chart
    ├── Chart.yaml
    ├── charts
    ├── templates
    │   ├── NOTES.txt
    │   ├── _helpers.tpl
    │   ├── deployment.yaml
    │   ├── hpa.yaml
    │   ├── ingress.yaml
    │   ├── service.yaml
    │   ├── serviceaccount.yaml
    │   └── tests
    │       └── test-connection.yaml
    └── values.yaml
```

But we don't need all the files 
refer helm/go-web-app-chart for current structure 

Helm Install 

```helm install go-web-app helm/go-web-app-chart```

Helm Template 

```helm template helm/go-web-app-chart```

Helm Uninstall 

```helm uninstall go-web-app```

