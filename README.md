
# ⚠️The aim of this project was solely to make progress in Golang: it is no longer active and requires a complete rewrite.⚠️


<div align="center">
<p align="center">
  <img src="https://raw.githubusercontent.com/PKief/vscode-material-icon-theme/ec559a9f6bfd399b82bb44393651661b08aaf7ba/icons/folder-markdown-open.svg" width="100" />
</p>
<p align="center">
    <h1 align="center">BURNTCOFFEE</h1>
</p>
<p align="center">
    <em>Fuel your code with BurntCoffee's tech prowess!</em>
</p>
<p align="center">
	<img src="https://img.shields.io/github/license/qjoly/burntcoffee?style=default" alt="license">
	<img src="https://img.shields.io/github/last-commit/qjoly/burntcoffee?style=default" alt="last-commit">
	<img src="https://img.shields.io/github/languages/top/qjoly/burntcoffee?style=default" alt="repo-top-language">
	<img src="https://img.shields.io/github/languages/count/qjoly/burntcoffee?style=default" alt="repo-language-count">
<p>
<p align="center">
	<!-- default option, no dependency badges. -->
</p>

 
</div>
<hr>

##  Quick Links
- [ Quick Links](#-quick-links)
- [ Overview](#-overview)
- [ Features](#-features)
- [ Repository Structure](#-repository-structure)
- [ Modules](#modules)
- [ Getting Started](#-getting-started)
    - [ Installation](#-installation)
    - [ Running burntcoffee](#-running-burntcoffee)
    - [ Tests](#-tests)
- [ Contributing](#-contributing)
- [ License](#-license)
- [ Acknowledgments](#-acknowledgments)

---

##  Overview

Burntcoffee is a project that aims to provide a simple and efficient solution for managing and configuring firecracker instances. With its core functionalities, users can easily set up and control firecracker virtual machines, allowing them to create and manage lightweight, secure, and isolated environments. Burntcoffee's value proposition lies in simplifying the process of managing firecracker instances, making it an ideal choice for developers and system administrators who require efficient and secure virtualization for their applications.

---

##  Features

|    | Feature           | Description                                                                                                       |
|----|--------------------|--------------------------------------------------------------------------------------------------------------------|
| 📄 | **Start a job**  | Start a job on an available firecracker socket.|
| 🧩 | **Stop all jobs**     | Stop all running VMs in all firecracker sockets|
| 🧩 | **Stop a specific jobs**     | Stop a single job |
| 🧪 | **Config**        | Generate a configuration file that contains sockets and IP available for jobs |

---

##  Repository Structure

```sh
└── burntcoffee/
    ├── config/
    │   ├── config.go
    │   ├── go.mod
    │   └── go.sum
    ├── firecracker/
    │   ├── firec.go
    │   └── go.mod
    ├── go.mod
    ├── go.sum
    ├── go.work
    └── main.go

```

---

##  Modules

<details closed><summary>.</summary>

| File                                                              | Summary                                                                                                                                                                                                                                                                                                                                                                        |
| ---                                                               | ---                                                                                                                                                                                                                                                                                                                                                                            |
| [go.work](https://github.com/qjoly/burntcoffee/blob/main/go.work) | The code snippet in the main.go file integrates the config and firecracker packages. It leverages dependencies to provide key functionalities for the burntcoffee repository's architecture.                                                                                                                                                                                   |
| [go.sum](https://github.com/qjoly/burntcoffee/blob/main/go.sum)   | The code snippet in the burntcoffee repository is responsible for configuring and interacting with the Firecracker virtualization technology. It provides essential functionalities for managing and controlling Firecracker instances.                                                                                                                                        |
| [main.go](https://github.com/qjoly/burntcoffee/blob/main/main.go) | This code snippet is the main entry point of a command-line application called burntcoffee. It utilizes the Cobra framework to define and execute various commands related to managing firecracker virtual machines. The code includes commands to start, stop, and show the status of jobs running on VMs, as well as commands to generate and display configuration details. |
| [go.mod](https://github.com/qjoly/burntcoffee/blob/main/go.mod)   | The code snippet in the `firecracker` directory is a critical feature of the `burntcoffee` repository. It is responsible for handling firecracker operations and dependencies. It allows for the configuration and management of firecracker instances within the larger software architecture.                                                                                |

</details>

<details closed><summary>config</summary>

| File                                                                         | Summary                                                                                                                                                                                                                                                                                                                                                                   |
| ---                                                                          | ---                                                                                                                                                                                                                                                                                                                                                                       |
| [go.sum](https://github.com/qjoly/burntcoffee/blob/main/config/go.sum)       | This code snippet, located in the config directory, is responsible for managing the configuration settings for the parent repository. It utilizes the dependencies gopkg.in/check.v1 and gopkg.in/yaml.v2 to load and parse YAML files.                                                                                                                                   |
| [config.go](https://github.com/qjoly/burntcoffee/blob/main/config/config.go) | The `config/config.go` file in the `burntcoffee` repository is responsible for generating and retrieving configuration files for the application. It includes functions to generate a YAML config file with predefined instances and retrieve the config from a file. This file plays a critical role in managing and accessing the application's configuration settings. |
| [go.mod](https://github.com/qjoly/burntcoffee/blob/main/config/go.mod)       | The code snippet in the config directory provides configuration management for the burntcoffee repository. It uses the gopkg.in/yaml.v2 library to handle YAML configuration files.                                                                                                                                                                                       |

</details>

<details closed><summary>firecracker</summary>

| File                                                                            | Summary                                                                                                                                                                                                                                                                                              |
| ---                                                                             | ---                                                                                                                                                                                                                                                                                                  |
| [firec.go](https://github.com/qjoly/burntcoffee/blob/main/firecracker/firec.go) | The `firecracker/firec.go` file in the `burntcoffee` repository contains functions for managing jobs in a distributed system. It allows starting, stopping, and showing the status of jobs on multiple instances. The functions make use of the HTTP protocol to send requests and handle responses. |
| [go.mod](https://github.com/qjoly/burntcoffee/blob/main/firecracker/go.mod)     | This code snippet, located in the `firecracker` module, is a key file in the `burntcoffee` repository. It utilizes the `firecracker/go.mod` dependency and plays a critical role in the overall architecture of the system.                                                                          |

</details>

---

###  Installation

1. Clone the burntcoffee repository:
```sh
git clone https://github.com/qjoly/burntcoffee
```

2. Change to the project directory:
```sh
cd burntcoffee
```

3. Install the dependencies:
```sh
go build -o myapp
```

###  Running burntcoffee
Use the following command to run burntcoffee:
```sh
./burntcoffee
```

###  Tests
To execute tests, run:
```sh
go test
```

---

##  Contributing

Contributions are welcome! Here are several ways you can contribute:

- **[Submit Pull Requests](https://github.com/qjoly/burntcoffee/blob/main/CONTRIBUTING.md)**: Review open PRs, and submit your own PRs.
- **[Join the Discussions](https://github.com/qjoly/burntcoffee/discussions)**: Share your insights, provide feedback, or ask questions.
- **[Report Issues](https://github.com/qjoly/burntcoffee/issues)**: Submit bugs found or log feature requests for burntcoffee.

<details closed>
<summary>Contributing Guidelines</summary>

1. **Fork the Repository**: Start by forking the project repository to your GitHub account.
2. **Clone Locally**: Clone the forked repository to your local machine using a Git client.
   ```sh
   git clone <your-forked-repo-url>
   ```
3. **Create a New Branch**: Always work on a new branch, giving it a descriptive name.
   ```sh
   git checkout -b new-feature-x
   ```
4. **Make Your Changes**: Develop and test your changes locally.
5. **Commit Your Changes**: Commit with a clear and concise message describing your updates.
   ```sh
   git commit -m 'Implemented new feature x.'
   ```
6. **Push to GitHub**: Push the changes to your forked repository.
   ```sh
   git push origin new-feature-x
   ```
7. **Submit a Pull Request**: Create a PR against the original project repository. Clearly describe the changes and their motivations.

Once your PR is reviewed and approved, it will be merged into the main branch.

</details>

---

##  License

This project is protected under the [MIT](https://github.com/QJoly/BurntCoffee/blob/main/LICENSE) License. For more details, refer to the [LICENSE](https://github.com/QJoly/BurntCoffee/blob/main/LICENSE) file.

---

##  Acknowledgments

- List any resources, contributors, inspiration, etc. here.

[**Return**](#-quick-links)

---
