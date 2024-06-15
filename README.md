# Cloud-Judge

## Overview

Cloud-Judge 是一个基于 Kubernetes (k8s) 的云原生在线判题平台，专为编程竞赛和算法挑战设计。它利用云计算的强大能力，为用户提供了一个可伸缩、高可用的在线编程环境。Cloud-Judge 支持多种编程语言，包括但不限于 Java 和 Go，允许开发者在云上提交代码，即时获得反馈和评分。

## Key Features

- **多语言支持**：支持 Java、Go 等多种编程语言，满足不同开发者的需求。
- **云原生架构**：基于 Kubernetes 构建，确保平台的高可用性和弹性伸缩。
- **实时反馈**：提供即时的代码评测结果，帮助开发者快速定位问题。
- **安全性**：采用云安全最佳实践，保护用户数据和代码的安全。
- **易于集成**：提供丰富的 API，方便与其他系统集成。

## Description

Cloud-Judge 是为编程爱好者和专业开发者设计的在线判题平台。它不仅提供了一个展示编程技能的舞台，还为教育和企业培训提供了一个强大的工具。利用 Kubernetes 的强大功能，Cloud-Judge 能够轻松处理高并发的代码提交和评测任务，同时保持系统的稳定性和响应速度。

### 技术亮点

- **容器化部署**：所有服务都运行在 Docker 容器中，通过 Kubernetes 进行管理，确保环境一致性和快速部署。
- **自动扩展**：根据负载自动调整资源，优化成本和性能。
- **持续集成/持续部署 (CI/CD)**：集成了 CI/CD 流程，支持代码的自动化构建和部署。
- **微服务架构**：采用微服务架构设计，提高了系统的可维护性和可扩展性。

## Getting Started

### Prerequisites
- go version v1.22.0+
- docker version 17.03+.
- kubectl version v1.11.3+.
- Access to a Kubernetes v1.11.3+ cluster.

### To Deploy on the cluster
**Build and push your image to the location specified by `IMG`:**

```sh
make docker-build docker-push IMG=<some-registry>/cloud-judge:tag
```

**NOTE:** This image ought to be published in the personal registry you specified.
And it is required to have access to pull the image from the working environment.
Make sure you have the proper permission to the registry if the above commands don’t work.

**Install the CRDs into the cluster:**

```sh
make install
```

**Deploy the Manager to the cluster with the image specified by `IMG`:**

```sh
make deploy IMG=<some-registry>/cloud-judge:tag
```

> **NOTE**: If you encounter RBAC errors, you may need to grant yourself cluster-admin
privileges or be logged in as admin.

**Create instances of your solution**
You can apply the samples (examples) from the config/sample:

```sh
kubectl apply -k config/samples/
```

>**NOTE**: Ensure that the samples has default values to test it out.

### To Uninstall
**Delete the instances (CRs) from the cluster:**

```sh
kubectl delete -k config/samples/
```

**Delete the APIs(CRDs) from the cluster:**

```sh
make uninstall
```

**UnDeploy the controller from the cluster:**

```sh
make undeploy
```

## Project Distribution

Following are the steps to build the installer and distribute this project to users.

1. Build the installer for the image built and published in the registry:

```sh
make build-installer IMG=<some-registry>/cloud-judge:tag
```

NOTE: The makefile target mentioned above generates an 'install.yaml'
file in the dist directory. This file contains all the resources built
with Kustomize, which are necessary to install this project without
its dependencies.

2. Using the installer

Users can just run kubectl apply -f <URL for YAML BUNDLE> to install the project, i.e.:

```sh
kubectl apply -f https://raw.githubusercontent.com/<org>/cloud-judge/<tag or branch>/dist/install.yaml
```

## Contributing
// TODO(user): Add detailed information on how you would like others to contribute to this project

**NOTE:** Run `make help` for more information on all potential `make` targets

More information can be found via the [Kubebuilder Documentation](https://book.kubebuilder.io/introduction.html)

## License

Copyright 2024 coding-hui.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

