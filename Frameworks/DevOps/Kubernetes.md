[toc]

# Kubernetes

## Overview

**Prompt: Tell me about the basics of Kubernetes.**

Kubernetes, often abbreviated as K8s, is an open-source container orchestration platform designed to automate the deployment, scaling, and management of containerized applications. It was originally developed by Google and is now maintained by the Cloud Native Computing Foundation (CNCF). Kubernetes has become the de facto standard for container orchestration and is widely used in both on-premises and cloud environments.

Here are the key basics of Kubernetes:

1. Containers:
   - Kubernetes is designed to work with containers, which are lightweight, portable, and consistent environments for running applications. Docker is one of the most popular containerization technologies, and Kubernetes can manage Docker containers along with other container runtimes like containerd and CRI-O.

2. Cluster:
   - In Kubernetes, you operate a cluster of machines, known as nodes, which can be physical or virtual servers. Each node runs a container runtime (e.g., Docker), and collectively, they form the infrastructure for running your applications.

3. Master and Nodes:
   - A Kubernetes cluster consists of two main components:
     - Master Node: The master node is responsible for managing the overall state of the cluster. It coordinates activities, schedules workloads, and monitors the health of nodes and applications.
     - Worker Nodes (Minions): Worker nodes are the machines where containers are actually deployed and run. These nodes are also referred to as minions.

4. Pod:
   - The smallest deployable unit in Kubernetes is a "Pod." A Pod can contain one or more containers that share the same network namespace and storage volume. Containers within the same Pod can communicate with each other via localhost.

5. Replication and Scaling:
   - Kubernetes allows you to define desired states for your application, and it ensures that the specified number of replicas (instances) are running at all times. You can scale your application up or down easily by adjusting the number of replicas.

6. Services:
   - Kubernetes abstracts the network to provide stable IP addresses and DNS names for applications. Services enable load balancing across Pods and provide a way for one set of Pods to discover and connect to another set.

7. Deployment:
   - Deployments are a higher-level abstraction that allows you to declaratively manage the desired state of your application, including its replicas, rolling updates, and rollbacks.

8. Configurations:
   - Kubernetes provides ConfigMaps and Secrets to manage configuration data and sensitive information like API keys, database credentials, etc., separately from application code.

9. Health Checks:
   - Kubernetes performs health checks (readiness and liveness probes) to ensure that containers are running and healthy. It can automatically restart containers that fail these checks.

10. Storage:
    - Kubernetes provides abstractions for managing persistent storage with features like Persistent Volumes (PVs) and Persistent Volume Claims (PVCs).

11. Namespace:
    - Kubernetes allows you to create multiple virtual clusters within a physical cluster using namespaces. This helps in isolating and organizing resources for different teams or applications.

12. Extensibility:
    - Kubernetes is highly extensible and can be customized using Custom Resource Definitions (CRDs) and Operators, enabling the management of complex applications beyond what is available out of the box.

Kubernetes is a powerful tool for managing containerized workloads, and its ecosystem continues to evolve with a rich set of features and tools to support various use cases and deployment scenarios. It can be complex to set up and manage but offers great flexibility and scalability for modern container-based applications.

## Architecture

**Prompt: Tell me about the Kubernetes architecture and modules.**

![Kubernetes Components | Kubernetes](../../Images/components-of-kubernetes.svg)

Kubernetes has a modular and highly scalable architecture that allows it to manage containerized applications effectively. Understanding the architecture and the key modules within Kubernetes is crucial for effectively deploying and managing applications in a Kubernetes cluster. Here's an overview of the Kubernetes architecture and its primary components/modules:

1. **Control Plane (Master Node):**
   - The control plane is the brain of the Kubernetes cluster, responsible for managing the overall state of the system.

   - Key components of the control plane include:
     - **API Server:** Acts as the front-end for the Kubernetes control plane. It exposes the Kubernetes API, which is used by both users and Kubernetes components to interact with the cluster.

     - **etcd:** A distributed key-value store that stores all cluster data, including configuration settings, state information, and metadata. It serves as the cluster's single source of truth.

     - **Controller Manager:** Watches the state of the cluster through the API Server and ensures that the desired state is maintained. Various controllers (e.g., Replication Controller, ReplicaSet Controller, Deployment Controller) are responsible for managing different types of resources.

     - **Scheduler:** Assigns work (Pods) to worker nodes based on resource requirements, constraints, and availability. It makes decisions about where to place Pods in the cluster.

2. **Node(s) (Minion):**
   - Nodes are the worker machines in the Kubernetes cluster, where containerized applications run. They are responsible for executing and managing Pods.

   - Key components on each node include:
     - **Kubelet:** Acts as the primary node agent that communicates with the control plane. It ensures that containers in the Pod are running in a healthy state.

     - **Container Runtime:** The container runtime (e.g., Docker, containerd) is responsible for pulling and running container images as specified in the Pod specification.

     - **Kube Proxy:** Maintains network rules on nodes, allowing network communication to and from Pods. It implements the Kubernetes Service abstraction.

3. **Add-Ons:**
   - These are optional components that enhance the functionality of a Kubernetes cluster. Add-ons are not part of the core Kubernetes architecture but are often included for various purposes, such as networking, monitoring, and logging.

   - Common add-ons include:
     - **Ingress Controllers:** Manage external access to services within the cluster.
     - **Container Network Interface (CNI):** Provides networking and connectivity between Pods.
     - **Cluster DNS:** Provides DNS-based service discovery within the cluster.
     - **Dashboard:** A web-based user interface for managing and monitoring the cluster.
     - **Monitoring and Logging Tools:** Such as Prometheus for monitoring and Fluentd/ELK Stack for logging.

4. **Custom Resource Definitions (CRDs) and Operators:**
   - Kubernetes allows you to extend its functionality by defining custom resources and controllers through CRDs and Operators. This allows you to manage custom application resources and automate their operations.

5. **API Objects:**
   - Kubernetes uses various API objects to define and manage resources in the cluster. Common objects include Pods, Services, Deployments, ConfigMaps, and Secrets. Users interact with the cluster by creating, updating, and deleting these objects using `kubectl` or other Kubernetes client tools.

In summary, Kubernetes has a multi-tier architecture with a control plane (master node) and worker nodes, all communicating through the Kubernetes API. Key components like etcd, the API server, controllers, and schedulers work together to maintain the desired state of the cluster. Nodes run containerized workloads, and various add-ons enhance cluster functionality. Kubernetes' extensibility allows you to define custom resources and controllers to manage complex applications efficiently. Understanding this architecture is essential for effectively deploying and managing applications in a Kubernetes environment.

## Workflow

**Prompt: Show me how to deploy an app to Kubernetes.**

Deploying an application to Kubernetes involves several steps, including creating Kubernetes resources like Pods, Services, and Deployments, as well as managing container images. Here's a high-level guide on how to deploy a simple web application to a Kubernetes cluster:

**Prerequisites:**
1. A running Kubernetes cluster. You can use a local cluster like Minikube or a cloud-managed Kubernetes service like Google Kubernetes Engine (GKE), Amazon EKS, or Azure Kubernetes Service (AKS).

2. `kubectl` command-line tool installed and configured to communicate with your Kubernetes cluster.

**Steps to Deploy an App to Kubernetes:**

1. **Containerize Your Application:**
   - Before deploying to Kubernetes, you need to containerize your application using a containerization technology like Docker. Create a `Dockerfile` that defines how your application should be packaged into a container image.

2. **Build and Push the Container Image:**
   - Build the container image from your application code and Dockerfile, and then push it to a container registry (e.g., Docker Hub, Google Container Registry, Amazon ECR).

   ```bash
   # Build the Docker image
   docker build -t your-image-name:tag .
   
   # Push the Docker image to a container registry
   docker push your-registry/your-image-name:tag
   ```

3. **Create Kubernetes Deployment YAML:**
   - Create a Kubernetes Deployment YAML file that defines how your application should run. Below is an example YAML file:

   ```yaml
   apiVersion: apps/v1
   kind: Deployment
   metadata:
     name: your-deployment-name
   spec:
     replicas: 3
     selector:
       matchLabels:
         app: your-app-name
     template:
       metadata:
         labels:
           app: your-app-name
       spec:
         containers:
         - name: your-container-name
           image: your-registry/your-image-name:tag
           ports:
           - containerPort: 80
   ```

   Save this YAML file, replacing `your-deployment-name`, `your-app-name`, `your-container-name`, and `your-registry/your-image-name:tag` with your own values.

4. **Apply the Deployment to Kubernetes:**
   - Use `kubectl` to apply the Deployment to your Kubernetes cluster:

   ```bash
   kubectl apply -f your-deployment.yaml
   ```

   This command will create the Deployment, which will manage the desired number of Pods running your application.

5. **Expose Your Application with a Service:**
   - To make your application accessible, you'll typically create a Kubernetes Service to expose it. Here's an example YAML for a simple LoadBalancer Service:

   ```yaml
   apiVersion: v1
   kind: Service
   metadata:
     name: your-service-name
   spec:
     selector:
       app: your-app-name
     ports:
       - protocol: TCP
         port: 80
         targetPort: 80
     type: LoadBalancer
   ```

   Apply this YAML file using `kubectl apply -f your-service.yaml`.

6. **Access Your Application:**
   - Depending on your Kubernetes setup, it might take some time for the LoadBalancer Service to provision an external IP address. You can check the status using:

   ```bash
   kubectl get svc your-service-name
   ```

   Once an external IP is assigned, you can access your application using that IP.

7. **Scale Your Application (Optional):**
   - You can scale your application up or down by adjusting the `replicas` field in your Deployment YAML file and reapplying it using `kubectl apply`.

That's a basic overview of deploying an application to Kubernetes. Depending on your application's complexity and requirements, you may need to set up additional resources like ConfigMaps, Secrets, and Ingress controllers. Monitoring, logging, and continuous deployment practices can also be integrated as needed for production-ready applications.

## Monitoring

**Prompt: Show me some kubectl commands to interact with Kubernetes cluster to monitor its state.**

Certainly! Here are some `kubectl` commands that you can use to interact with a Kubernetes cluster and monitor its state:

1. **View Cluster Information:**

   - View the Kubernetes cluster information:
     ```bash
     kubectl cluster-info
     ```

   - Display the current context, which includes the cluster, user, and namespace:
     ```bash
     kubectl config current-context
     ```

   - List all available contexts:
     ```bash
     kubectl config get-contexts
     ```

2. **View Cluster Nodes and Resources:**

   - List all nodes in the cluster:
     ```bash
     kubectl get nodes
     ```

   - List all pods in the current namespace:
     ```bash
     kubectl get pods
     ```

   - List pods in all namespaces:
     ```bash
     kubectl get pods --all-namespaces
     ```

   - List all services in the current namespace:
     ```bash
     kubectl get services
     ```

   - List all namespaces in the cluster:
     ```bash
     kubectl get namespaces
     ```

3. **View Resource Details:**

   - Describe a specific pod to see its details, events, and configurations:
     ```bash
     kubectl describe pod <pod-name>
     ```

   - Get detailed information about a service:
     ```bash
     kubectl describe service <service-name>
     ```

4. **Logs and Exec into Pods:**

   - View the logs of a pod:
     ```bash
     kubectl logs <pod-name>
     ```

   - Stream logs from a pod in real-time:
     ```bash
     kubectl logs -f <pod-name>
     ```

   - Execute a command inside a running container in a pod:
     ```bash
     kubectl exec -it <pod-name> -- /bin/bash
     ```

5. **Resource Health and Events:**

   - Check the status of a deployment:
     ```bash
     kubectl rollout status deployment/<deployment-name>
     ```

   - View events related to a resource:
     ```bash
     kubectl get events
     ```

   - Get events for a specific resource:
     ```bash
     kubectl describe <resource-type> <resource-name>
     ```

6. **Scaling Resources:**

   - Scale the number of replicas in a deployment:
     ```bash
     kubectl scale deployment/<deployment-name> --replicas=<desired-replicas>
     ```

7. **Resource Deletion:**

   - Delete a resource:
     ```bash
     kubectl delete <resource-type> <resource-name>
     ```

   - Delete all resources in a specific namespace:
     ```bash
     kubectl delete all --all -n <namespace>
     ```

8. **Resource Creation and Updates:**

   - Create or update resources from YAML files:
     ```bash
     kubectl apply -f <file.yaml>
     ```

   - Apply configurations from a directory (e.g., for multiple resources):
     ```bash
     kubectl apply -f <directory>
     ```

These `kubectl` commands are essential for monitoring the state of a Kubernetes cluster, inspecting resources, and troubleshooting issues. By using these commands, you can gain insights into the health and status of your Kubernetes applications and infrastructure.