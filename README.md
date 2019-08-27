# brutus
Brute force your iOS Restrictions Passcode from the key &amp; salt

# Use

## Installation

```bash
go install github.com/dipshit/brutus
```

## Find your iOS Restrictions Password in a backup

`cat $(fd com.apple.restrictionspassword.plist /var)`

RestrictionsPasswordKey is the -k

RestrictionsPasswordSalt is the -s

## Example

```bash
brutus -k 8+az5Jgc5v3IRMFny2T1iEe+Lxw= -s jvPLSA==
pass is 6161
```
