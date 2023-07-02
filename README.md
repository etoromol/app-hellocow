# app-hellocow 🐮

This repository documents the process to containerize a [Go](https://go.dev/) Web Application named `hellowcow` using [Docker](https://www.docker.com/). 

            -----------------------
          /  Mooooo, World!        /
         /  From GITHUB.COM       /
        /  Version 1.0.0         /
         -----------------------
                     \
                      \
                        ^__^
                        (oo)\\_______ 
                        (__)\\       )\/\ 
                            ||----w | 
                            ||     ||    

`Hellocow` serves as a basic `Hello World` service and leverage a Dockerfile to build the application. By containerizing the application, it becomes easily deployable and scalable within [Kubernetes](https://kubernetes.io/) environments as a microservice test-case.

Cloud developers can utilize this repository as a reference and baseline project. The ultimate goal is to empower developers to seamlessly deploy the application on various Kubernetes cloud services, such as [GKE](https://cloud.google.com/kubernetes-engine/), [EKS](https://aws.amazon.com/eks/) and [AKS](https://learn.microsoft.com/en-us/azure/aks/).

Within this directory, you will find:

- `main.go` contains the HTTP server implementation. It prints `Mooooo, World!`, providing a fun response to client requests.
- `Dockerfile` facilitates the creation of the [Container Image](https://docs.docker.com/engine/reference/commandline/images/) of the application.

*"Container Image" and "Docker Image" refer to the same concept. Docker, being a popular containerization platform, introduced the concept of container images and uses the term "Docker images" to specifically refer to the images created and managed using Docker tools.*

This application is already built and available in [Docker Hub](https://hub.docker.com/) which is a cloud-based [Container Registry]() for container images:

- `https://hub.docker.com/r/etoromol/hellocow`

## ¿Why Docker?

Developing apps today requires so much more than writing code. Multiple languages, frameworks, architectures, and discontinuous interfaces between tools for each lifecycle stage creates enormous complexity. Docker simplifies and accelerates your workflow, while giving developers the freedom to innovate with their choice of tools, application stacks, and deployment environments for each project.

## Prerequisites

To successfully build and deploy this application, please ensure that you have the following prerequisites:

* [Docker Desktop](https://www.docker.com/products/docker-desktop/) (4.21.0+) installed.

*Alternatively, you may opt to use an alternative container engine such as Podman. [Podman](https://podman.io/).*

## Installation and deployment

1. Clone the hellocow application repository to your local device:
```bash
git clone https://github.com/etoromol/app-hellowcow.git && 
cd app-hellowcow
```

2. Build the container image for the application with the following command:
```bash
docker build -t app/hellocow:1.0.0 .
```

3. To create the Container image and start the Container, you can use the docker run command as a single step:
```bash
docker run --hostname app-hellocow --name hellocow-1.0.0 --publish 8080:8080 app/hellocow:1.0.0
```

4. (optional) If you prefer more granularity, you can divide the previous step into two individual commands. First, create the container using the container image we built in step (2):
```bash
docker create --hostname app-hellocow --name hellocow-1.0.0 --publish 8080 app/hellocow:1.0.0
```

5. (optional) To manually start the container, use the following command:
```bash
docker start hellocow-1.0.0
````

##  Running the Application

6. To verify that the application is running, you may perform a quick validation using the following curl command:
```bash
curl http://127.0.0.1:8080
```
*Alternatively, you can access the application by opening the URL in your web browser.*


## License

Copyright (c) 2023 Eduardo Toro.

Licensed under the [MIT](LICENSE) license.
