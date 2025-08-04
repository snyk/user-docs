# Verifying Broker image signatures

Beginning with version 4.169.1, all Broker container images are signed using Cosign.

Cosign is the tool for signing and verifying containers. Cosign Includes storage in an Open Container Initiative (OCI) registry.

Cosign is designed to enhance the security of container images by providing a simple and efficient way to sign and verify them. It leverages the concept of digital signatures, with which you sign the container image with your private key, and the recipient can verify your signature using the corresponding public key.

## Prerequisite for verifying Broker image signatures

You must [install Cosign](https://docs.sigstore.dev/system_config/installation/) version 2.2.0 or higher.

## Verify signed container images

To verify the signed image, you must use the built-in `cosign verify` command.

It is not necessary to pull the Broker container image to perform the verification step

```
$ cosign verify --key cosign.pub snyk/broker:4.169.1-github-com

Verification for index.docker.io/snyk/broker:4.169.1-github-com --
The following checks were performed on each of these signatures:
  - The cosign claims were validated
  - Existence of the claims in the transparency log was verified offline
  - The signatures were verified against the specified public key

[{"critical":{"identity":{"docker-reference":"index.docker.io/snyk/broker"},"image":{"docker-manifest-digest":"sha256:a2ff856b180a532c3e31a90b9788cad567fa05d78c84bccc637de54c6f46ebf2"},"type":"cosign container image signature"},"optional":{"Bundle":{"SignedEntryTimestamp":"MEUCIEOmElvKK0eC/hvpM9SE66RAekcV6DpF6NSO4Gz5aftrAiEAlQY8lKe1RUqYtCK1WRwWhWT/f/PvyBTJC8qBjgU20kU=","Payload":{"body":"eyJhcGlWZXJzaW9uIjoiMC4wLjEiLCJraW5kIjoiaGFzaGVkcmVrb3JkIiwic3BlYyI6eyJkYXRhIjp7Imhhc2giOnsiYWxnb3JpdGhtIjoic2hhMjU2IiwidmFsdWUiOiJjOTE2NDQ1N2YxMDA0NTQxNWNlMjBlN2I3YjNmYjg4YjZmMmNhNzI4MDNkODY4NTk0ZDhlY2UzMGJkYTFiZjQ4In19LCJzaWduYXR1cmUiOnsiY29udGVudCI6Ik1FWUNJUUNRQVp6VVdqbkNFai9GZkpxTGU4YVdoYXhacWJzZnZTc21JNXRiRzZuRmdBSWhBTFgyMWVLRHl6OXNqWHQrVStVZUZNTUFyN1oyV09Gd0k1b1oxclc0dlJBUCIsInB1YmxpY0tleSI6eyJjb250ZW50IjoiTFMwdExTMUNSVWRKVGlCUVZVSk1TVU1nUzBWWkxTMHRMUzBLVFVacmQwVjNXVWhMYjFwSmVtb3dRMEZSV1VsTGIxcEplbW93UkVGUlkwUlJaMEZGWkNzeWJVVlhlVVJyT0VOdGJUQkRSREZhT0dwamMxaEhZVkV5YVFwelREaHdlRWh5ZDI5SlNEUkVlRzFrZVVveWJucDNWMkY0V1daelpscE5OazV2UTFKV2MyZFpRVlpsTlVkQ2FFWmljalpvZW1OcU5XZDNQVDBLTFMwdExTMUZUa1FnVUZWQ1RFbERJRXRGV1MwdExTMHRDZz09In19fX0=","integratedTime":1698788303,"logIndex":46732614,"logID":"c0d23d6ad406973f9559f3ba2d1ca01f84147d8ffc5b8445c224f98b9591801d"}},"tag":"4.169.1-github-com"}}]
```

## Broker Cosign public key

```
-----BEGIN PUBLIC KEY-----
MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEd+2mEWyDk8Cmm0CD1Z8jcsXGaQ2i
sL8pxHrwoIH4DxmdyJ2nzwWaxYfsfZM6NoCRVsgYAVe5GBhFbr6hzcj5gw==
-----END PUBLIC KEY-----
```
