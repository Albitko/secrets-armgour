# Secrets-armGour

A serivce for storing secrets by developers.
It pulls up the storage of data such as:
- passwords from services
- binary data (for .pfx e.g.)
- credit cards data
- text data (for ssh keys e.g.)

Encryption keys are stored only on the client side.
Only you own your secrets.

The client for the servie is distributed in assemblies for macOS, Windows and Linux.
You will be able to deploy the server part of the service next to other services e.g. GitLab.
A single PostgreSQL dependency is required.

### Usage cli examples

First of all you could register
```
armgour-cli register -l john -p SecretPa$$
```
Or login if you are already registered
```
armgour-cli login -l john -p SecretPa$$
```
Then you could create secrets for storing in service
List of supported secrets can be recieved via `-h` flag
```
armgour-cli create -h

Secrets armGour version 0.1.0
Build at 05.08.2023
===================================
Create secret

Usage:
  armgour-cli create [command]

Available Commands:
  binary      Create user binary secret
  cards       Create user cards secrets
  credentials Save user credentials
  text        Save user text secrets

Flags:
  -h, --help   help for create
```

You can store credentials for your working accounts e.g.
```
armgour-cli create credentials -s gmail -l john@google.com -p SecretGmailPa$$
```
And then list already stored
```
armgour-cli list --card

Secrets armGour version 0.1.0
Build at 05.08.2023
===================================
ID: 3 Card number: 5105105105105100 Description: Master Card Citi
ID: 4 Card number: 4012888888881881 Description: Visa Bank of America

```
For getting full info about secret you could get it by index from previous output.
We will use short flag for card `-c`
```
armgour-cli get -c -i 3

Secrets armGour version 0.1.0
Build at 05.08.2023
===================================
Card holder:John Wayne number:5105105105105100 Validity period:05/2027 CVC:123 Description: Master Card Citi
```
And don`t forget to log out
```
armgour-cli logout

Secrets armGour version 0.1.0
Build at 05.08.2023
===================================
Logout successfully.
```
