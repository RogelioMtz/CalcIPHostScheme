# IP Subnet Calculator

This Go program calculates subnet details from CIDR notation and generates IP schemes based on required hosts. It provides a simple command-line interface to perform these calculations.

## Features

- Calculate subnet details from CIDR notation
- Generate IP scheme based on required hosts
- Display possible networks based on CIDR

## Usage

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/ip-subnet-calculator.git
    cd ip-subnet-calculator
    ```

2. Build the program:
    ```sh
    go build -o ip-subnet-calculator calcHostSchemeIP.go
    ```

3. Run the program:
    ```sh
    ./ip-subnet-calculator
    ```

## Menu Options

1. **Calculate subnet details from CIDR**
    - Enter an IP address in CIDR notation (e.g., `192.168.1.1/24`).
    - The program will display the IP address, subnet mask, CIDR notation, first usable IP, last usable IP, and total hosts.

2. **Generate IP scheme based on required hosts**
    - Enter a base IP address (e.g., `192.168.1.0`).
    - Enter the number of required hosts.
    - The program will display the IP address, broadcast address, subnet mask, CIDR notation, first usable IP, last usable IP, and total hosts.

3. **Exit**
    - Exit the program.

## Example

```sh
Welcome to the IP Subnet Calculator!

-----Options:-----
1. Calculate subnet details from CIDR
2. Generate IP scheme based on required hosts
3. Exit
Choose an option (1-3): 1
Enter IP address in CIDR notation (e.g., 192.168.1.1/24): 192.168.1.1/24

----- Subnet Details -----
IP Address:         192.168.1.1
Subnet Mask:        255.255.255.0
CIDR Notation:      /24
First Usable IP:    192.168.1.1
Last Usable IP:     192.168.1.254
Total Hosts:        254

----- All Possible /24 Networks for 192.168.1.0 -----
Network Address      Usable Host Range                Broadcast Address
192.168.1.0          192.168.1.1 - 192.168.1.254      192.168.1.255
```

##License

This project is licensed under the Mozilla Public License 2.0.

##Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.
