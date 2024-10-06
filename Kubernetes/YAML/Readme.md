- Kubernetes uses YAML files as inputs for creation of **objects** such as:
    - pods
    - replicas
    - deployment 
    - services

## Defining a YAML file<br>
`pod-definition.yaml`
```
apiVersion: v1
kind: Pod
metadata:
    name: myapp-pod
    labels: ------------------ (dictionary)
        app: myapp
        tier: front-end
        ...
        ...
        anything else:    
spec:
    containers: ------------------ (list/array)
        - name: nginx-container ------------------ (the dash indicates that this is the first item in the list)
          image: nginx ------------------ (name of the docker image in the docker repository)
```
- `apiVersion`: Version of the kubernetes API you are using to create the object
- `kind`: Type of object we are trying to create. Ex: `Pod`, `Service`, `ReplicaSet`, `Deployment`, etc.
- `metadata`: Data about the object
- `spec`: Additional information to kubernetes pertaining to that object.

### Commands
##### Pods:
- `kubectl run <any-podname> --image=<name_from_repo>`: Creating a new pod.
- `kubectl create -f pod-definition.yaml`: Kubernetes creates the pod.
- `kubectl get pods`: Retrieves list of pods.
- `kubectl get all`: To retrieve all objects created.
- `kubectl describe pod myapp-pod`: Detailed description of a pod.
- `kubectl edit pod redis`: Allows you to edit the pod configuration details.

##### Replicas:
- `kubectl delete replicaset <replicaset-name>`: to delete a replicaset.
    - OR use `kubectl delete -f <file-name>`.
- `kubectl edit replicaset <replicaset-name>`: to edit a replicaset.
- `kubectl api-resources | grep replicaset`: to check the apiVersion of replicaset by command .
- `kubectl explain replicaset | grep VERSION`

##### Deployment:
Create a new Deployment with the below attributes using your own deployment definition file.
<br>
- Name: `httpd-frontend`;
- Replicas: `3`;
- Image: `httpd:2.4-alpine`

```
apiVersion: apps/v1
kind: Deployment
metadata:
    name: httpd-frontend
spec:
    replicas: 3
    selector:
        matchLabels:
            name: httpd-frontend
    template:
        metadata:
            labels:
                name: httpd-frontend
        spec:
            containers:
            - name: httpd-frontend
              image: httpd:2.4-alpine
```

Then run: `kubectl create -f own-deployment-definition-1.yaml`

##### Rollout, Update and Rollback:
- `kubectl apply -f <name of the deployment>`: UPDATE
    - `kubectl set image <name of the deployment> nginx=nginx:1.9.1`: Sets image - but **not recommended**.
- `kubectl rollout status <name of the deployment>`: To check the status of the deployment.
- `kubectl rollout history <name of the deployment>`: To check the history of our deployment.
- `kubectl rollout undo <name of the deployment>`: UNDO Deployment.
    - Example for "name of the deployment": `deployment/app-deployment`

#### Example:

```
Q. Create a new pod with the name redis and the image redis123. Use a pod-definition YAML file. And yes the image name is wrong!

Ans:
- Use this command to create a pod definition YAML file: `kubectl run redit --image=redis123`
- Then: `kubectl create -f` command to create a resource from the manifest file
```
- We use kubectl run command with --dry-run=client -o yaml option to create a manifest file: `kubectl run redis --image=redis123 --dry-run=client -o yaml > redis-definition.yaml`
- After that, using kubectl create -f command to create a resource from the manifest file: `kubectl create -f redis-definition.yaml`
- Verify the work by running kubectl get command: `kubectl get pods`