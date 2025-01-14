### Authentication System

#### Introduction

In a distributed network, the following issues need to be addressed to establish associations among nodes (or objects):

- Identify a user or object
- Prove the identity of a user or object
- Describe user permissions
- Describe relationships between users
- Describe user credibility
- Describe users' ability to serve each other

Common concepts used in the system are as follows:

**Object Identifier**

For a specific file or content, we generate a digest and convert it into a visible 28-byte string, which can be used to identify a specific content object.

**Identity Identifier**

A pair of keys represents an identity. Keys are binary and too long to be used as identifiers. The public key can be digested and converted into a visible 28-byte string. This ID uniquely identifies the key pair and can be used to represent the identity of a user or node. This approach is inspired by Bitcoin wallets.

### 1. Keys

The keys discussed here are asymmetric encryption public keys, using the ed25519 key algorithm. The node's key is placed in the root directory, named `hostkey.cfg`. When Leither starts for the first time, the system checks the root directory and automatically generates this file if it does not exist. Execute the `lpki` command in the root directory, requiring the `key` parameter. If not specified, the default node key is used. Keys can be used for signing and verification.

- [Generate Key Pair](../api/LPki.md#genkey)
- [Export Public Key](../api/LPki.md#genpkkey)

### 2. Certificates

Certificates usually consist of the following:

- Content information, typically a series of key-value pairs.
- Signature information
- Signer information

**Self-Signed Certificate**

If the content of the certificate is to prove the user's own identity, it is a self-signed certificate.

**CA Certificate**

A certificate used to prove someone else's identity is a CA certificate.

- [Generate Certificate with Private Key](../api/LPki.md#gencert)
- [Generate Public Key Certificate](../api/LPki.md#genpkcert)

### 3. Passes

A pass is a certificate for logging into a node with certain permissions. The content of a pass usually includes the visitor's identity, validity period, and service content.

- [Generate Pass via Command Line](../api/LPki.md#signppt)

**Mac Pass**

In an internal network, devices like phones and computers can directly obtain a Mac-based pass issued by the node when accessing it. This can be bound to a user and enable automatic login. Mac passes can be obtained via RPC or HTTP API.

**WeChat Node**

In the Leither architecture, a service node bound to a public account is called a WeChat node. WeChat nodes do not store user information and data; their main function is to provide an interface for user devices and WeChat integration and identity verification.

**WeChat Pass**

When a WeChat user accesses a public account, Tencent's server sends unique information representing the user to the WeChat node. If this unique information is bound to the user's information, it allows WeChat users to automatically log into the user node. A pass issued by the WeChat node representing the WeChat user's identity is called a WeChat pass.

### 4. Services

Requesting permissions from a node

After logging in, before using node resources, users must first request service access permissions.

- [Request Service](../api/LPki.md#RequestService)