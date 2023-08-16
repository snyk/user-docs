# Verifying CLI standalone binaries

You can verify both the shasum of downloaded binaries and their GPG signatures.

The download location on `static.snyk.io` contains a file called `sha256sums.txt.asc`. You can download it directly: `https://static.snyk.io/cli/latest/sha256sums.txt.asc` or for a specific version like `https://static.snyk.io/cli/v1.666.0/sha256sums.txt.asc`.

To check that a downloaded file matches the checksum, use a `sha256sum` command, for example:

```
grep snyk-macos sha256sums.txt.asc | sha256sum -c -
```

If you want to verify Snyk CLI standalone binaries against [Snyk CLI GPG key](https://github.com/snyk/cli/blob/master/help/\_about-this-project/snyk-code-signing-public.pgp), first import the CPG key:

```
# A22665FB96CAB0E0973604C83676C4B8289C296E is the key belonging to code-signing@snyk.io
# Copy of this public key is also in this repository /help/_about-this-project/snyk-code-signing-public.pgp
gpg --keyserver hkps://keys.openpgp.org --recv-keys A22665FB96CAB0E0973604C83676C4B8289C296E
```

Then verify the file is signed with:

```
gpg --verify sha256sums.txt.asc
```

The command output should look like the following:

```
gpg: Signature made So  8 Jan 14:11:44 2023 CET
gpg:                using EDDSA key A22665FB96CAB0E0973604C83676C4B8289C296E
gpg: Good signature from "Snyk Limited <code-signing@snyk.io>" [unknown]
gpg: WARNING: This key is not certified with a trusted signature!
gpg:          There is no indication that the signature belongs to the owner.
Primary key fingerprint: A226 65FB 96CA B0E0 9736  04C8 3676 C4B8 289C 296E
```
