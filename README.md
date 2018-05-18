# brutus
Brute force your iOS Restrictions Passcode from the key &amp; salt

## Find your iOS Restrictions Password in a backup
`cat $(fd com.apple.restrictionspassword.plist /var)`

RestrictionsPasswordKey is the -k

RestrictionsPasswordSalt is the -s

## Example

`./brutus -k 8+az5Jgc5v3IRMFny2T1iEe+Lxw= -s jvPLSA==`

`pass is 6161`
