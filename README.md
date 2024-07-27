## ScanNet

ScanNet is a Go program that scans a network for active hosts and displays their IP addresses.

### Installation

1. Ensure you have Go installed on your system.
2. Clone the repository:
   ```
   git clone https://github.com/jac11/NetScan.git
   ```
3. Navigate to the project directory:
   ```
   cd NetScan
   ```
4. Build the program:
   ```
   go build -o ScanNet ScanNet.go
   ```

### Usage

1. Run the program:
   ```
   ./ScanNet
   ```
2. The program will scan the local network and display the IP addresses of active hosts.

### Configuration

The program can be configured by passing command-line arguments:

- `--Port`: The port to scan (default is 80)
- `--Domain`: The IP address or domain name to scan
- `--StartScan`: The starting port for the range scan
- `--EndScan`: The ending port for the range scan

Example usage:

```
./ScanNet --Port 22 --Domain example.com
./ScanNet --StartScan 1 --EndScan 1000 --Domain 192.168.1.1
```

### Features

- Supports scanning a single port or a range of ports
- Displays the connection status for each port (success or failure)
- Includes a stylized banner for the program

### Contributing

If you find any issues or have suggestions for improvements, please feel free to open an issue or submit a pull request on the [GitHub repository](https://github.com/jac11/NetScan).

