# Action Documentation

## Release

GitHub does **not** maintain an action to interact with releases. Both related
action are archived:
+ [actions/upload-release-asset](https://github.com/actions/upload-release-asset)
+ [actions/create-release](https://github.com/actions/create-release)

GitHub links to
[softprops/action-gh-release](https://github.com/softprops/action-gh-release)
for a maintained action to interact with releases.


## Sigstore

Verify using rekor:
```
uuid=$(rekor-cli search --artifact calculator | tail -n 1)
sig=$(rekor-cli get --uuid=$uuid --format=json | jq -r .Body.HashedRekordObj.signature.content)
cosign verify-blob --key cosign.pub --signature <(echo $sig) calculator
```
