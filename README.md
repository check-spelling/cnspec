# cnspec

```
  ___ _ __  ___ _ __   ___  ___
 / __| '_ \/ __| '_ \ / _ \/ __|
| (__| | | \__ \ |_) |  __/ (__
 \___|_| |_|___/ .__/ \___|\___|
   mondoo™     |_|
```

**Open source, cloud-native security and policy project**

`cnspec` is a cloud-native solution to assess the security and compliance of your business-critical infrastructure. `cnspec` finds vulnerabilities and misconfigurations on all systems in your infrastructure including: public and private cloud environments, Kubernetes clusters, containers, container registries, servers, and endpoints, SaaS products, infrastructure as code, APIs, and more.

`cnspec` is a powerful Policy as Code engine built on [`cnquery`](https://github.com/mondoohq/cnquery), and comes configured with default security policies that run right out of the box. It's both fast and simple to use!

![cnspec scan example](docs/gif/cnspec-scan.gif)

## Installation

Install `cnspec` with our installation script:

**Linux and macOS**

```bash
bash -c "$(curl -sSL https://install.mondoo.com/sh/cnspec)"
```

**Windows**

```powershell
Set-ExecutionPolicy Unrestricted -Scope Process -Force;
[System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072;
iex ((New-Object System.Net.WebClient).DownloadString('https://install.mondoo.com/ps1/cnquery'));
Install-Mondoo -Product cnspec;
```

If you prefer a package, find it on [GitHub releases](https://github.com/mondoohq/cnspec/releases).

## Run a scan with policies

Use the `cnspec scan` subcommand to check local and remote targets for misconfigurations and vulnerabilities.

### Local scan

This command evaluates the security of your local machine:

```bash
cnspec scan local
```

### Remote scan targets

You can also specify [remote targets](#supported-targets) to scan. For example:

```bash
# to scan a docker image:
cnspec scan docker image ubuntu:22.04

# scan public ECR registry
aws ecr-public get-login-password --region us-east-1 | docker login --username AWS --password-stdin public.ecr.aws/r6z5b8t4
cnspec scan docker image public.ecr.aws/r6z5b8t4

# to scan an AWS account using the local AWS config
cnspec scan aws

# scan an EC2 instance with EC2 Instance Connect
cnspec scan aws ec2 instance-connect root@i-1234567890abcdef0

# to scan a Kubernetes cluster via your local kubectl config or a local manifest file
cnspec scan k8s
cnspec scan k8s manifest.yaml

# to scan a GitHub repository
export GITHUB_TOKEN=<personal_access_token>
cnspec scan github repo <org/repo>
```
