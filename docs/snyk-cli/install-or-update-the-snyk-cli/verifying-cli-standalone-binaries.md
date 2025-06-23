# Verifying CLI standalone binaries

You can verify both the shasum of downloaded binaries and their GPG signatures.

The download location on `downloads.snyk.io` contains a file called `sha256sums.txt.asc`. You can download it directly: `https://downloads.snyk.io/cli/stable/sha256sums.txt.asc` or for a specific version like `https://downloads.snyk.io/cli/v1.666.0/sha256sums.txt.asc`.

To check that a downloaded file matches the checksum, use a `sha256sum` command, for example:

```
grep snyk-macos sha256sums.txt.asc | sha256sum -c -
```

If you want to verify Snyk CLI standalone binaries against the [Snyk CLI GPG key](https://github.com/snyk/cli/blob/master/help/_about-this-project/snyk-code-signing-public.pgp), first import the GPG key:

```
# 467717A30B2B4658415975629691DA64D0025194 is the key belonging to code-signing@snyk.io
# Copy of this public key is also in this repository /help/_about-this-project/snyk-code-signing-public.pgp
gpg --keyserver hkps://keys.openpgp.org --recv-keys 467717A30B2B4658415975629691DA64D0025194
```

Then verify the file is signed with:

```
gpg --verify sha256sums.txt.asc
```

The command output should look like the following:

```
gpg: Signature made So  8 Jan 14:11:44 2025 CET
gpg:                using EDDSA key 467717A30B2B4658415975629691DA64D0025194
gpg: Good signature from "Snyk Limited <code-signing@snyk.io>" [unknown]
gpg: WARNING: This key is not certified with a trusted signature!
gpg:          There is no indication that the signature belongs to the owner.
Primary key fingerprint: 4677 17A3 0B2B 4658 4159  7562 9691 DA64 D002 5194
```
