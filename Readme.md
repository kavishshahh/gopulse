# My Hot-Reloading Tool gopulse

My Hot-Reloading Tool is a command-line utility that allows you to automate the process of monitoring changes in your Go code and automatically reloading your application.

## Installation

To install and use the hot-reloading tool, follow these steps:

1. Make sure you have Go installed on your system. You can download and install it from the [official website](https://golang.org/dl/).

2. Install the tool using the `go install` command:

    ```bash
   go install github.com/kavishshahh/gopulse
   ```
Verify that the installation was successful by running the following command:

   ```bash
   gopulse --version
   ```
## Usage
To use the hot-reloading tool, follow these steps:

Navigate to the directory of your Go project that you want to monitor for changes.

Run the following command to start the hot-reloading watcher:

```bash
gopulse
```

The tool will begin monitoring your project directory for changes.

Make changes to your Go code files. When you save a file, the tool will automatically detect the change, rebuild your application, and restart it.

Observe the console output to see the status of your application and any build or runtime errors.
