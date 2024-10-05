- Kubernetes uses YAML files as inputs for creation of **objects** such as:
    - pods
    - replicas
    - deployment services

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
- `kubectl run <any-podname> --image=<name_from_repo>`: Creating a new pod.
- `kubectl create -f pod-definition.yaml`: Kubernetes creates the pod.
- `kubectl get pods`: Retrieves list of pods
- `kubectl describe pod myapp-pod`: Detailed description of a pod.
- `kubectl edit pod redis`: Allows you to edit the pod configuration details.<br>
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