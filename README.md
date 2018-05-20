# brutus
Brute force your iOS Restrictions Passcode from the key &amp; salt

From my blog post: https://medium.com/@rastogikirin/how-i-brute-forced-my-iphones-restrictions-password-492cb97ce5b7

## Find your iOS Restrictions Password in a backup
`cat $(fd com.apple.restrictionspassword.plist /var)`

RestrictionsPasswordKey is the -k

RestrictionsPasswordSalt is the -s

## Example

`./brutus -k 8+az5Jgc5v3IRMFny2T1iEe+Lxw= -s jvPLSA==`

`pass is 6161`
