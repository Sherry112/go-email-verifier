# Email Domain Checker

This Go program checks the DNS records of a given domain to determine if it has MX, SPF, and DMARC records. It reads domain names from the standard input and outputs the findings for each domain.

## Features

- Checks for the presence of MX records
- Checks for the presence and value of SPF records
- Checks for the presence and value of DMARC records

## Requirements

- Go 1.16 or higher

## Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/your-username/go-email-checker.git
   cd go-email-checker

2. Build the program:
    ```sh
    go build -o email-checker

## Usage

1. Run the program:
    ```sh
    ./go-email-checker

2. Enter domain names one per line. For example:
    ```sh
    example.com
    google.com

3. The program will output the results in the following format:
    ```sh
    ./email-checker
    example.com
    Domain=example.com
    hasMx=true
    hasSPF=true
    spfRecord=v=spf1 include:_spf.example.com ~all
    hasDMARC=true
    dmarcRecord=v=DMARC1; p=none; rua=mailto:dmarc-reports@example.com
    google.com
    Domain=google.com
    hasMx=true
    hasSPF=true
    spfRecord=v=spf1 include:_spf.google.com ~all
    hasDMARC=true
    dmarcRecord=v=DMARC1; p=reject; rua=mailto:dmarc-reports@google.com
    ```

## Logging

If there are any errors during the DNS lookup process, they will be logged using the utils.ErrorLogger function from the utils package.