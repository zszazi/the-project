## Multi Stage docker file 

Since we are using golang, its better practice to use multi stage docker file


>Create docker image 

```bash 
docker build -t zszazi/the-project:v0.0.1 .
```


>Run the docker image 
```bash 
docker run -p 8080:8080 -it zszazi/the-project:v0.0.1 
```

>Push the docker image 
```docker push zszazi/the-project:v0.0.1```