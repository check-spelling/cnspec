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

### Policies

`cnspec` policies are built on the concept of [policy as code](https://mondoo.com/policy-as-code/). `cnspec` comes with default security policies configured for all supported targets. The default policies are available via the [cnspec-policies](https://github.com/mondoohq/cnspec-policies) GitHub repo.

## Vulnerability Scan

`cnspec` supports vulnerability scanning for a wide-range of platforms. The vulnerability scanning is not restricted to container images, it works for build and runtime.

![cnspec vuln example](docs/gif/cnspec-vuln.gif)

NOTE: The current version requires you to be logged in to Mondoo Platform. Future versions will be able to scan the platforms without the requirement to be logged in.

### Examples

```bash
# scan container image
cnspec vuln docker debian:10

# scan aws instance via EC@ instance connect
cnspec vuln aws ec2 instance-connect root@i-1234567890abcdef0

# scan instance via SSH
cnspec vuln ssh user@host

# scan windows via SSH or Winrm
cnspec vuln ssh user@host --ask-pass
cnspec vuln winrm user@host --ask-pass

# scan VMware vSphere ESXi hosts
cnspec vuln vsphere user@host --ask-pass

# scan Linux, Windows
cnspec vuln local
```

| Platform                 | Versions                 |
| ------------------------ | ------------------------ |
| Alpine                   | 3.10 - 3.16              |
| AlmaLinux                | 8, 9                     |
| Amazon Linux             | 1, 2, 2022               |
| Arch Linux               | Rolling                  |
| CentOS                   | 6, 7                     |
| Debian                   | 8, 9, 10, 11             |
| Fedora                   | 30 - 36                  |
| openSUSE                 | Leap 15.4                |
| Oracle Linux             | 6, 7, 8                  |
| Photon Linux             | 2, 3, 4                  |
| Red Hat Enterprise Linux | 6, 7, 8                  |
| Rocky Linux              | 8                        |
| SUSE Linux Enterprise    | 12, 15                   |
| Ubuntu                   | 18.04, 20.04, 22.04      |
| VMware vSphere ESXi      | 6, 7                     |
| Windows                  | 10, 11, 2016, 2019, 2022 |

## cnspec interactive shell

`cnspec` also provides an interactive shell to explore assertions. It helps you understand the assertions that policies use, and write your own as well. It's also a great way to interact with both local and remote targets on the fly.

### Local system shell

```bash
cnspec shell local
```

The shell provides a `help` command to get help on the resources that power `cnspec`. Running `help` without any arguments lists all of the available resources and their fields. You can also run `help <resource>` to get more information on a specific resource. For example:

```bash
cnquery> help ports
ports:              TCP/IP ports on the system
  list []port:      TCP/IP ports on the system
  listening []port: TCP/IP ports on the system
```

The shell uses auto-complete, which makes it easy to explore. Once inside the shell, you can enter MQL assertions like this:

```coffeescript
> ports.listening.none( port == 23 )
```

To clear the terminal, type `clear`.

To exit, either hit CTRL + D or type `exit`.
