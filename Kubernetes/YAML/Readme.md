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
- `kubectl create -f pod-definition.yaml`: Kubernetes creates the pod.
- `kubectl get pods`: Retrieves list of pods
- `kubectl describe pod myapp-pod`: Detailed description of a pod.
