# Leither and MiMei Technical Foundation

**Maintained by:** `cfa532/Leither`, in the fork-owned `cfa532-docs/` collection.
**Moved:** 2026-10-16 from LifeAlbum; the original remains a dated archive.

**Status:** Application-independent technical reference
**Leither CLI reviewed:** V0.24.02, API contract version 1; commit, pin and node-Files behavior rechecked on V0.24.23 (2026-09-16, section 5.6) and on V0.24.27 (2026-09-19, section 5.7)
**Online documentation reviewed:** `vzhan.cn/#doc.html`, August 2026
**Source repository reviewed:** [3and4/Leither](https://github.com/3and4/Leither), originally August 2026; documentation index, README, historical changelog and key-authentication guide refreshed on 2026-10-16 at commit `90362aef401c4d7ff53eb3573443f031f0be880f`
**CLI inspection refreshed:** gen8 `/home/pi/demo/Leither`, V0.24.63, amd64, API 1, 2026-10-16; all 159 help entries inspected recursively (section 12.1). Earlier runtime observations retain their own version and date.
**Reference implementation reviewed:** TweetBackendApp JavaScript and Go MApps, plus TweetWeb's browser RPC client
**Purpose:** Explain Leither, MiMei, storage types, versioning, distribution, identity, and the boundary between applications and a Leither node.

## 1. Executive Summary

Leither is a lightweight data-and-application container that combines stable data identities, versioned storage, application hosting, identity and permissions, and IPFS-compatible distribution.

The central storage unit is a **MiMei**. A MiMei has a stable external identifier, while each committed version points to immutable content. This gives applications a permanent object address without losing the integrity and deduplication benefits of content addressing.

MiMei supports two principal data types:

- **File type:** A versioned file or directory tree backed by content-addressed data.
- **Database type:** A versioned key-value database with richer data operations.

File type is the production-ready choice for applications that prioritize durable, portable, file-oriented storage. The database type provides richer operations but introduces substantially greater implementation and recovery complexity. Each application should choose its storage model according to its requirements and the maturity of the relevant Leither subsystem.

## 2. Leither System Model

A Leither installation can play several roles at once:

- **Container:** Holds MiMei data and Leither applications.
- **Storage node:** Persists immutable blocks, mutable File-type MiMei trees, versions, and local replicas.
- **Serving node:** Makes selected MiMei data available to authorized clients.
- **Provider:** Announces that it can supply the data for a MiMei or IPFS object.
- **Network participant:** Discovers peers and providers through DHT and IPFS-compatible protocols.
- **Identity endpoint:** Exposes the developing node/user key, certificate, passport, session, and rights model.
- **Application endpoint:** Hosts front-end assets and optional application logic.

Leither is therefore not simply an IPFS daemon. IPFS supplies content-addressed blocks and distributed retrieval, while MiMei supplies stable identity, mutable working state, versions, references, rights, applications, and publication semantics.

Leither exposes two related filesystem surfaces that must not be confused:

- The node-level **Files filesystem** is an IPFS-backed mutable filesystem addressed by ordinary root paths and operated through `Files*` calls.
- A **File-type MiMei filesystem** belongs to one stable `mid`, is opened at a selected version, and is operated through the returned MiMei/file handles and `MF*` calls.

The node-level Files filesystem can be used to assemble and flush a directory DAG before assigning its root CID to a File MiMei. It is a construction surface, not the identity or version boundary of the File MiMei itself.

## 3. MiMei Identity

### 3.1 Stable MiMei ID

Every MiMei has a stable identifier called a `mid`. The `mid` remains unchanged when the data changes. Applications use it to open, reference, synchronize, publish, and authorize the MiMei.

This differs from an IPFS CID:

| Identifier | Meaning | Changes when content changes? |
|---|---|---:|
| MiMei ID (`mid`) | Stable identity of a logical data object | No |
| MiMei version | Named immutable state of that object | New version is created |
| CID or Mac ID | Address of immutable content or a content tree | Yes |

The stable `mid` is the application-level identity. CIDs and Mac IDs are implementation-level content identities beneath it.

### 3.2 Unique and Deterministic Creation

`mimei create` defaults to `--mark "{{auto}}"`, which creates a new identity each time. A non-auto mark, together with the other creation parameters, can derive the same MiMei identity again.

Applications can use both modes deliberately:

- A deterministic mark identifies a singleton resource whose identity must be reproducible.
- An automatic mark creates an independent object whose identity must be unique.

Marks must be namespaced by application and purpose. They are identifiers, not user-visible names.

## 4. MiMei Versions

A MiMei exposes several version views:

- **`cur`:** Mutable working state. Writers make changes here.
- **`last`:** Most recently committed backup. Normal readers use this state.
- **`release`:** A selected stable version, useful for released applications or deliberately frozen data.
- **Numbered versions:** Immutable historical backups.

The normal write lifecycle is:

```text
open cur -> apply all related changes -> backup -> last advances
```

Backing up is the commit boundary. A successful file write does not by itself mean that normal readers or other nodes can see a complete new state.

Applications should batch one logical operation—such as content plus its metadata—into one `cur` mutation and one backup. File commands that automatically flush must be used carefully so that a multi-file operation does not accidentally create partially committed history.

`restore` copies a selected historical version back into `cur`; it does not erase history. `release` selects a stable version. `delver` removes selected retained versions when policy permits.

## 5. File-Type MiMei

### 5.1 Characteristics of File Type

File-type MiMei is suitable for production file-oriented storage because it offers:

- Ordinary files and directories with understandable recovery behavior.
- Immutable content-addressed backing.
- Efficient version reuse through Merkle-DAG structure.
- Direct compatibility with large media files.
- JSON as a portable, inspectable application-object format.
- Independence from a particular database implementation.
- Recovery and export using standard file-oriented tools.

Applications that choose File-type MiMei can still maintain disposable local databases or search indexes as acceleration layers, provided those caches can be rebuilt completely from authoritative files.

### 5.2 Two Filesystem Surfaces

The Leither repository makes the separation explicit:

| Surface | Identity and scope | Representative operations | Typical purpose |
|---|---|---|---|
| File-type MiMei filesystem | One `mid`, opened at `cur`, `last`, `release`, or a numbered version | `MMOpen`, `MFOpenByPath`, `MFOpenMacFile`, `MFReaddir`, `MFSetData`, `MFGetData`, `MFStat` | Authoritative versioned application data |
| Node Files filesystem | Node-local IPFS-backed mutable tree rooted at `/` | `FilesMkdir`, `FilesLs`, `FilesCopy`, `FilesMv`, `FilesRm`, `FilesStat`, `FilesFlush` | Staging, assembling, moving, and flushing an IPFS directory tree |
| Immutable IPFS/Mac content | Content ID rather than mutable path | `MFTemp2MacFile`, `MFTemp2Ipfs`, IPFS add/get operations | Immutable payload storage and reuse |

`MMOpen(sid, mid, ver)` returns an operational MiMei session (`mmsid`), not another MiMei ID. For File type, that session can represent the selected filesystem root. The repository's File-MiMei example opens a version and calls `MFReaddir(mmsid, -1)` directly to enumerate its root. A file below that root may be opened by path with `MFOpenByPath(sid, mmsid, path, flag)`, after which the returned file session is used for byte, object, size, type, and stat operations.

This produces the direct conceptual workflow:

```text
login -> MMCreate or resolve mid -> MMOpen(mid, "cur")
      -> open/mutate files and directories through the MiMei filesystem handle
      -> close resource handles -> MMBackup -> last advances
```

Older repository documentation also describes `FSFind`, `FSMkDir`, `FSStat`, `FSCopy`, and `FSRename` against a File-MiMei filesystem. Those calls are valuable architectural evidence, but they are absent from the repository's current Go `LApiStub`. Their exact availability, flags, signatures, and JavaScript names must therefore be confirmed against the installed runtime before an application depends on them.

The separate node-Files assembly workflow is:

```text
temporary file -> MFTemp2Files(node path)
               -> mutate with Files* -> FilesFlush(path) returns root CID
               -> MFSetCid(mid, root CID) assigns the completed tree to the File MiMei
```

This workflow is useful for large uploads and whole-tree assembly. It does not turn `/appDrive` or any other node path into the MiMei identity: the stable identity remains `mid`, and committed readers open a MiMei version.

#### `/.mm/<mid>`: the writable tree of a File MiMei in the Files namespace (experimental)

**Status: experimental.** Observed on gen8 (Leither V0.24.62, 2026-10-14) with the
CLI as the node's system user, then with a non-system session. It is not in the
Leither repository documentation reviewed here. Reads through it are not
access-controlled (see below); do not build application code on it.

The node Files root contains a `.mm` directory. It holds one subdirectory per
File-type MiMei on the node, named by `mid`. `/.mm/<mid>` is the MiMei's `cur`
tree, reachable through ordinary `Files*` calls:

- Its hash from `files stat` equals `mm://<mid>:cur` at every step of the test
  below, and differs from `mm://<mid>:last` whenever `cur` has uncommitted
  changes. When `cur` equals the latest commit, all three are the same.
- The directory does not exist until the first write: a newly created MiMei
  answers `file does not exist` for `/.mm/<mid>`, `mm://<mid>:cur` and
  `mm://<mid>:last`, and the first `files mkdir /.mm/<mid>/photos` creates it.
- Writes go to `cur` only. `last` does not appear, and does not move once it
  exists, until a commit.

Test on a throwaway MiMei (`mimei create -t file -e probe.test`):

| Step | `/.mm/<mid>` | `mm://<mid>:cur` | `mm://<mid>:last` |
|---|---|---|---|
| after create | does not exist | `file does not exist` | `ver[last] not find!` |
| `files mkdir /.mm/<mid>/photos` | new root | same | `ver[last] not find!` |
| `files cp ipfs/<cid> /.mm/<mid>/photos/a.txt` | new root | same | `ver[last] not find!` |
| `files flush /.mm/<mid>` (returns the root CID) | same | same | `ver[last] not find!` |
| `mimei setcid --mid <mid> <root CID>` | same | same | same (version 1) |
| `files mkdir /.mm/<mid>/docs`, flush | new root | same | unchanged |

After the second write the unchanged `photos` subtree kept its CID and only the
root changed, as expected for a Merkle tree. `FilesFlush` alone does not commit;
`mimei setcid` (`MFSetCid`) does, consistent with section 5.6.

Tested with a non-system session (2026-10-14, same node). The identity was a
fresh key, self-signed certificate and `CertFor=Self` passport, logged in with
`LoginWithPPT`, the way LifeDrive's identity service logs in a household user.
It had no grants unless stated. A control, a plain Files directory with
`files setright -u <uid> -d <dir> -r write`, accepted `mkdir`, `cp` and `rm` from
this identity, so the probe itself was sound.

| Case | Result |
|---|---|
| `FilesLs`/`FilesStat` on `/.mm/<mid>`, no grant | **Allowed**, for a MiMei with rights `72f6f04` and for a private MiMei created with rights `0`. A second private-MiMei probe listed `secret/a.txt` with its CID and size. |
| `FilesLs("/.mm")`, no grant | **Allowed**; returned every `mid` on the node (12). |
| `FilesLs`/`FilesStat` on `mm://<private mid>:last/…` and `:cur/…`, non-owner session | **Allowed**, same as `/.mm/<mid>`. (An earlier version of this row recorded `invalid mimei id`; that was a defect in the probe, which truncated the URL at its first colon.) |
| the same calls with **no login at all** (guest), also from the internet to gen8's public port | **Allowed** |
| `MMOpenUrl`/`MFGetData` on `mm://<private mid>:last/…` or on `/.mm/<mid>/…`, non-owner or guest | `1101: Permission denied` (contents are protected on this route) |
| `MFOpenTempFile` for a non-owner session (the first step of a MApp-style `MFTemp2Files` write) | `1101: Permission denied`; the maintainers' unchecked-`MFTemp2Files` claim could not be exercised |
| `FilesMkdir` under `/.mm/<mid>`, no grant | `5006: Insufficient permissions Run as administrator` |
| after `files setright -u <uid> -d /.mm/<mid> -r write` (also on `/.mm/<mid>/photos`) | still `5006`; `listrights` shows the grant but the write path ignores it |
| after `mimei setright --mid <mid> <uid> 0x07276704` | still `5006`; `getright` reads that uid back as `Read` only |
| after `mimei setright --mid <mid> <uid> 0x07` (read+write+execute) | **Allowed**: `mkdir`, `cp ipfs/<cid>`, `mv`, `rm`, `flush` all succeeded |
| after `mimei setright --mid <mid> <uid> 0` | `5006` again |

Consequences:

- Writes under `/.mm/<mid>` are governed by the MiMei's per-uid right, which needs
  the write bit. Directory-level `files setright` grants do not apply there. This
  matches "follow the MiMei author's permissions"; it is not the same mechanism
  LifeDrive uses for `/lifedrive-work/<user>`.
- A session's writes land in `cur`. After the non-system `flush`, `last` was
  unchanged (still `QmSTU…`); only the commit moves it. The test did not call
  `MFSetCid` from that session, so who may commit through it is untested.
- **Directory and metadata reads are not gated by the MiMei's rights, on any
  path form.** `FilesLs` and `FilesStat` answer for `/.mm/<mid>`, `mm://<mid>:cur/`
  and `mm://<mid>:last/` to a non-owner and to a caller with no login, including
  for a MiMei with rights `0`. gen8's public port answered an anonymous call from
  another machine, and an anonymous `FilesLs("/.mm")` returned every File MiMei id
  on the node (13). Names, sizes and CIDs of a private tree are therefore visible
  to anyone who can reach the node, with or without the `mid` (a `mid` is enough
  without `/.mm`). The byte route `MMOpenUrl` does enforce rights, but a CID
  obtained from the listing can be fetched with no session from
  `GET /ipfs/<cid>`: gen8's public port returned `200` for a CID with no login,
  and minipc's loopback port did the same for a household identity record
  (status and byte count only; the content was not read).
- **What that means for LifeDrive on minipc.** The identity MiMei is a File-type
  MiMei, and an anonymous `FilesStat("mm://<identity mid>:last/")` returned its
  hash and size (5,271 bytes); `/ipfs/<hash>` returned the same size. Anyone
  holding the identity `mid` could therefore read that user's identity record
  without a session. The `mid` is not secret in the current design (it is in the
  pairing payload and the profile reply). The legacy drive and backups MiMeis
  there are empty directories after the catalog migration, so they expose
  nothing. The Database-type catalog and user-root MiMeis answered `type is
  MiMei` and `not ipfs file id`, so they are **not** reachable through the Files
  RPC.
- This is Leither behavior, to be reported to its maintainers; the application
  cannot fix it. Until it is fixed, treat any File-type MiMei's names, sizes and
  CIDs, and the bytes behind its CIDs, as readable by anyone.

Leither maintainers' memo (2026-10-14, "files-mm-read-permission", measured on a
temporary V0.24.62 node) confirms and explains these results:

- Reads under `/.mm/<mid>` are not checked: `FilesLs`, `FilesStat` and `MFLs`
  take no rights decision, and a MApp reading `/.mm/<mid>/secret.txt` as an
  unauthorized user or a guest received the contents of a MiMei created with no
  Guest right. So content, not only names and CIDs, was readable.
- The data plane (`mm://<mid>:<ver>` read through `MMOpenUrl`/`MFGetData`) does
  check `DataRight_Read`: `1101: Permission denied` for a MiMei created with no
  Guest right. The leak is the Files-namespace path, not the MiMei.
- Writes under `/.mm/<mid>` check the MiMei's `DataRight_Write`; `files setright`
  grants are not consulted there. This matches the table above.
- A per-user `setright` value contributes only its low byte (`0x07276704` becomes
  `0x04`, Read) plus the group slot defaults, which is why it read back as Read.
  The full table in `Right :072f6f04` is for `mimei create -r`: the low byte is the
  Guest slot, so `0x072f6f04` (the `mimei create` default) is Guest-readable and
  `0x072f6f00` is not.
- They list it as gap A and propose a read gate with `off`, `observe` and `enforce`
  modes (the `FilesGateMode` family), covering `FilesLs`, `FilesStat`, `MFLs` and
  `MMOpenUrl("files://…")`, judging `DataRight_Read` and failing closed when the
  subject is unknown. Not in V0.24.62.
- They also state that `MFTemp2Files` to a `/.mm/<mid>` path is unchecked when
  `FilesGateMode=off`, the default. This document has not tested that.

What LifeDrive needs from any such gate:

- Share verification, `FilesStat("", mm://<share>:last/<name>)` in
  `apps/lifedrive-identity/shares.go`, uses an empty session on a share MiMei
  created with `0x07276704` (Guest Read). "Fail closed on an unknown subject"
  must mean checking the Guest slot, not rejecting.
- The identity service reads drives and catalogs through `mm://<mid>:last` with
  the user's authority session, which authored the MiMei and always passes.
- The memo gates `/.mm/<mid>` paths, while LifeDrive uses `mm://` URLs with
  `FilesLs` and `FilesStat` everywhere. The tests above show those forms are
  equally ungated today, so a fix limited to `/.mm/<mid>` would leave the leak
  open. The gate has to cover `mm://<mid>:<ver>/…` as well, and must not reject
  the empty-session share check on a Guest-readable MiMei.
- The memo says `MMOpenUrl("files://…")` is not judged. In this test
  `MMOpenUrl` on `files:///.mm/<mid>/…` returned `1101` for a non-owner, so the
  byte route already refuses there; the unprotected route is the listing and
  `/ipfs/<cid>`, which a read gate on `FilesLs`/`FilesStat` alone does not
  close. Whether `/ipfs/<cid>` should require a session is a question for the
  maintainers.

Not yet established:

- Whether a non-system session may commit (`MFSetCid`) a File MiMei it can write
  to through `/.mm/<mid>`.
- Whether this V0.24.62 read behavior was present on earlier releases.

Gen8 also has a root directory named `mm:` whose children are `<mid>:cur`
entries with hashes that differ from `/.mm/<mid>`. It exists for only the three
MiMeis that had a live `cur` tree. Its purpose is unknown; it appears to be an
internal cache of opened `mm://…:cur` trees.

The probe MiMeis `R8F75nZSJitgtGwPn952zk2bLye` (mark `claude-mm-probe-20261014`,
rights `72f6f04`), `xLAtmWH4DgvO4rBKFXoYEzPgQI7` (mark
`claude-mm-probe-private-20261014`, rights `0`) and `JSj2uUxGAe0e8BcVVqLn3Q9cmLg`
(mark `claude-mm-probe-share-20261014`, rights `0x07276704`, authored by a
throwaway identity whose key was deleted) remain on gen8; the CLI has no
command to delete a MiMei. The test identities, grants and probe binaries were
removed from gen8 and minipc.

#### Retaining a received IPFS file

The Leither repository's `lapi/lapi.go` pin interface and `Install.md` entry for
V0.18.36 document `IpfsPinAdd(sid string, r bool, ps string) error`. The path is
one string, not a variadic argument list. LifeDrive calls
`IpfsPinAdd(recipientSID, true, "/ipfs/" + cid)` on the recipient's paired node.
Unlike the Files calls, the pin path keeps its leading slash: on V0.24.23
`ipfs pin add ipfs/<cid>` fails with `invalid path "/ipfs/ipfs/<cid>"`.
Recursive pinning retains the complete DAG. The phone does not download and
re-upload the original to save it on that node.

A CID already present in a received link needs no sender lookup. For a share
token, resolve the sender's password/expiry gate first; then use `content_id`
or `FilesStat("", "mm://<share-mid>:last/<name>").Hash`. Do not substitute the
private drive MID or send the recipient's session to the sender.

To make a pinned item available to LifeDrive's existing Files context menu,
attach `ipfs/<cid>` (no leading slash, see the CLI table below) to the
recipient's Files tree with `FilesCopy`, flush the
tree, and commit with `MFSetCid`. This adds a reference to existing content, not
a second upload. Verify the committed `last` reference before reporting that
the item was saved in Files. An attachment failure can leave a valid pin;
report that partial result and allow a retry. Do not unpin it during rollback,
because other users or files on the same node may retain that CID.

This pin adapter was compiled against the documented RPC signature; live-node
pin permissions and availability still require deployment verification. See
[LifeDrive shared links](https://github.com/cfa532/LifeAlbum/blob/0876ac5a20e0b4f9225e63febe9426f8694795d9/docs/LIFEDRIVE_SHARED_LINKS.md) for the link and client contract.

### 5.3 File Operations

The installed CLI exposes the main file lifecycle through the `mimei` and `files` command families:

| Purpose | Representative CLI | Corresponding API responsibility |
|---|---|---|
| Create container | `mimei create --data-type file` | Create a File-type MiMei |
| Add local content | `mimei add --mid <mid> <path>` | Import a file or directory |
| Create directory | `mimei mkdir <path>` | Create a directory in the MiMei tree |
| Copy/link content | `mimei cp <source> <dest>` | Copy within the tree or attach an IPFS path |
| Replace root data | `mimei setdata --mid <mid> <filename>` | Set File-type content and commit metadata |
| Point to content | `mimei setcid --mid <mid> <cidpath>` | Set MiMei content from an existing CID path |
| Read/list | `mimei get`, `mimei ls` | Read files and enumerate directories |
| Commit | `mimei backup --mid <mid>` | Advance `last` from `cur` |
| Inspect | `mimei show --mid <mid>` | Read versions, rights, and state |

The node-level `files` family supplies mutable-filesystem operations such as `cp`, `mkdir`, `mv`, `rm`, `flush`, `ls`, and `stat`. Large or resumable uploads can use temporary-file APIs before attaching the completed immutable object to a File MiMei, either directly as a Mac/IPFS file or indirectly through a flushed Files tree.

The corresponding JavaScript MApp pattern verified against the installed API is:

| Purpose | JavaScript `lapi` operation | Contract |
|---|---|---|
| Create a File MiMei | `MMCreate(sid, appId, ext, mark, 1, right)` | Data type `1` is File; a non-auto mark deterministically resolves the same MiMei |
| Open writable state | `MMOpen(sid, mid, "cur")` | Returns the writable MiMei session used by operations on that MiMei |
| Open committed state | `MMOpen(sid, mid, "last")` | Opens the latest committed File MiMei version |
| Open a File-MiMei path | `MMOpenUrl(sid, "mm://<mid>:<ver>/<path>")` | Opens a file or directory from an ordinary File MiMei using its explicit versioned URL |
| Open a path from a MiMei handle | `MFOpenByPath(sid, mmsid, path, flag)` | Returns a resource session below an already opened File-MiMei root; flags require runtime confirmation |
| Enumerate an opened directory | `MFReaddir(fileSid, count)` | Lists directory entries; the repository example uses `-1` to request all entries |
| Write JSON-compatible data | `MFSetObject(fileSid, value)` | Stores structured application objects without Database calls |
| Read JSON-compatible data | `MFGetObject(fileSid)` | Reads a structured object from the selected MiMei version |
| Write bytes | `MFSetData(fileSid, bytes, offset)` | Supports bounded or resumable binary writes |
| Read bytes | `MFGetData(fileSid, start, count)` | Supports bounded range reads; `-1` may be used where the installed API permits read-to-end |
| Inspect size | `MFGetSize(fileSid)` | Completion should compare actual and declared byte sizes |
| Create staging file | `MFOpenTempFile(sid)` | Creates an upload staging handle outside the committed visible tree |
| Attach staged content | `MFTemp2Files(tempSid, "/appDrive/<path>")` | Places a completed temporary upload in the writable Files directory tree |
| Mutate directories | `FilesMkdir`, `FilesMv`, `FilesRm`, `FilesCopy` | Operate on ordinary paths under the writable `/appDrive` tree |
| Flush the tree | `FilesFlush(sid, "/appDrive")` | Returns the CID of the completed directory tree |
| Assign a flushed root | `MFSetCid(sid, mid, rootCid)` | Installed-runtime path for assigning a completed Files-tree CID to the File MiMei; returned version semantics must follow that runtime |
| Close a handle | `MMClose(handle)` or the installed file-close equivalent | Every successful or failed path must release opened handles |
| Snapshot/commit | `MMBackup(session, mid, memo, ...options)` | Creates a new immutable version and advances `last`; verify whether the installed binding expects an authenticated session or opened MiMei session |

The session contract deserves explicit verification. The repository API document names the first `MMBackup` argument `sid`, while its runnable database-backup example passes the `mmsid` returned by `MMOpen`. Existing JavaScript applications may use an author session obtained from `BELoginAsAuthor()`. This is API-version drift, not a semantic difference in what backup means. Adapters must probe or pin the installed signature rather than infer the required session type from the parameter name.

`MMBackup` and `MFSetCid` must likewise not be called blindly as two mandatory commits for one mutation. `MMBackup` is the general MiMei snapshot operation. `MFSetCid` is the installed File-MiMei path used to assign a completed root CID and may itself return the resulting version. The application adapter must follow the installed runtime's verified commit contract and expose one logical commit to callers.

### 5.4 Handle Ownership and Cleanup

Leither uses several handle types with different scopes:

- `sid`: authenticated API session.
- `mid`: stable MiMei identity.
- `mmsid`: opened MiMei/version session returned by `MMOpen`.
- `fsid`: opened file, directory, or temporary-file resource session.
- CID/Mac ID: immutable content identity, not a session.

A new MiMei has no `last` until the first commit. Inspect `GetVar(sid, "mmversions", mid)` before resolving `last` during initial creation; use `cur` while initializing. Check private data access through the MiMei Hprose APIs and rights, not by interpreting an application-asset `/mm/` HTTP response as a permission result.

Code must not substitute one kind for another merely because each is represented as a string. Opened MiMei and resource sessions must be closed on every success and failure path. Repository documentation names `MFClose` for file resources, while the current Go stub and examples expose `MMClose` for opened handles; the installed binding decides which close operation is callable.

### 5.5 File Paths as Relationships

File-oriented applications may represent parent-child relationships through directory nesting. For a cloud drive, the directory tree is the authoritative relationship graph and human-visible file names may be the stored path names, just as in a regular filesystem. Server entries must normalize paths, reject traversal and path separators in individual names, and define explicit conflict behavior.

### 5.6 Commit Cost and Node-Files Behavior Verified on V0.24.23

Observed on minipc (Leither V0.24.23, linux/amd64) on 2026-09-16, plus points
confirmed by the Leither maintainers the same day. Applications that commit large
File-type trees should read this before choosing a partition or commit policy.
The LifeDrive design that follows from it is in
[LifeDrive Backup Storage Plan](https://github.com/cfa532/LifeAlbum/blob/0876ac5a20e0b4f9225e63febe9426f8694795d9/docs/LIFEDRIVE_BACKUP_STORAGE_PLAN.md).

#### Commit pins the whole tree

`MFSetCid(sid, mid, rootCid)` records the version through
`backupWithFileId -> backupModifyInfo -> addRef`, and `addRef` runs
`IpfsPinAdd(recursive=true)` on the new root. The pin walks every block under the
root, not only the new ones:

| Tree committed | Blocks pinned | Pin duration |
|---|---|---|
| Small scratch tree | 13 | 31 ms |
| LifeDrive backup drive | 12,210 | 51.037 s |

Commit cost is therefore proportional to the size of the whole MiMei and grows
with every file added. Restructuring a tree into deeper directories inside one
MiMei does not help: the recursive pin still visits every block. Only a separate
MiMei is a separate pin root. A caller whose RPC deadline is shorter than the pin
sees a timeout while the commit still lands, so the result is ambiguous; retrying
pins the whole tree again.

#### References, pins and garbage collection (maintainer-confirmed)

- `MMAddRef` pins the referenced IPFS content; `MMDelRef` unpins it. `MMAddRef`
  is cheap and synchronous for content already on the node.
- Content referenced from a remote node is fetched over time; Leither handles
  that synchronization, and local path access waits for it.
- Garbage collection, which runs when disk space is low, removes only IPFS blocks
  with no remaining reference. It never removes ordinary files on disk.
- A reference to another MiMei is its MiMei ID, which does not change when that
  MiMei commits. Committing a child costs its parent nothing, and a
  references-only commit of the parent is cheap.
- Synchronization transfers only the committed `last` version, never `cur`.

#### Choosing a commit path for File-type MiMei

| Operation on a type-1 MiMei | Result on V0.24.23 |
|---|---|
| `FilesFlush(tree)` then `MFSetCid(mid, rootCid)` | Commits the directory tree; pins the whole tree (above) |
| `MMBackup(mid)` after writing `cur` | Snapshots `cur` as bytes. `last` becomes a bytes node, not a directory (`FilesLs mm://<mid>:last/photos` fails with `LookupBySegment called on a bytes node`) |
| `MMOpen(mid, "cur")` then `MFReaddir` | `not a directory` |
| `MFOpenByPath(sid, mid, "/photos/a.txt", O_RDWR\|O_CREATE)` | Server panic in `mmfs.(*MTreeFS).OpenFileForFind` (`MTreeNode2WebDav.go:281`), returned as an RPC error |
| `MFOpenByPath(…, flag 0)` then `MFSetData` | Returns a handle; the write fails (`fullName=/photos/a.txt left=`) |

TweetBackendApp reached the same conclusion: commit File MiMeis with
`FilesFlush` + `MFSetCid` and never call `MMBackup` after `MFSetCid`. The
Leither file-operations guide (2026-09) calls `MFSetCid` the correct way to point
a MiMei version at an existing CID ("path C"), and uses `MMBackup` for database
and Mac-file MiMei.

`MFOpenTempFile` is refused for an ordinary user session (`1101: Permission
denied`); it needs the node's `mimei` service grant
(`lpki grantservice --target-nodes=<uid> --service=mimei -k <hostkey>`).

#### CLI `mimei` file commands

| Command | Observed |
|---|---|
| `mimei mkdir mm/<mid>/photos` | Creates the directory and flushes it to `last` (default `--flush=true`) |
| `mimei mkdir --flush=false` on a new MiMei | Succeeds but creates no `last`; later `mimei cp` fails with `ver[last] not find!` |
| `mimei backup --mid <mid>` after a flushed mkdir | `ver[cur] is not exist` |
| `mimei add --mid <mid> <file>` | Replaces the MiMei root with that file; later copies fail with `root can't be a file` |
| `mimei cp /ipfs/<cid> mm/<mid>/…` or `files cp /ipfs/<cid> …` after CLI `ipfs add` | `cannot get node from path /ipfs/<cid>: file does not exist`, because a leading slash names a path in the node's Files tree |
| `files cp ipfs/<cid> /path` (also `ipfs://<cid>` or a bare CID) | `filesCp ok` immediately after the add, including an 80 MB file |
| `mimei cp mm://<other-mid>:last mm/<mid>/…` | Hung 219 s without output; the server-side `MFCopy` continued after the client was killed |

`mimei cp` calls `MFCopy(src, dest, flush, overwrite)`, which opens the target
MiMei's `last` tree and writes the destination path into it. Leither's own help
for `files cp` documents adding local content as `ipfs add --pin=false <file>`
followed by `files cp /ipfs/<CID> <mfs path>`, a lazy copy that references only
the root; on V0.24.23 the source must be written `ipfs/<CID>`.

#### `IpfsAdd` and the session message queue

`IpfsAdd(sid, localPath)` (and CLI `ipfs add`) returns after the content is
stored; the log shows `AsyncIpfsAdd` and then `node2Ipfs2 cid=…` before the
call returns. It sends progress messages to the calling session. The CLI pulls
them. A service session that never calls `PullMsg(sid, timeout)` fills its
10-message queue, after which each send waits 3 s (`[DIAG] sendmsg-timeout`),
so a 104 KB add took 9 s. Draining with `PullMsg` restored sub-second adds.
Over hprose, declare the result as a map: an interface result that comes back
empty panics the Go client (`reflect: … returned zero Value`).

Adding files this way leaves their blocks unpinned, so the next `MFSetCid`
pins them as part of the whole tree: 19,616 blocks took 3 m 31 s.

#### Node Files tree performance

- `FilesStat` on a directory walks its subtree to total the size. On the
  12,210-block drive, `FilesStat("mm://<mid>:last/")` did not return within 45 s.
  Existence checks should list the parent with `FilesLs` instead.
- `FilesLs("mm://<mid>:last/")` returned in about 350 ms.
- Later the same day every node-Files listing hung: `files ls /`,
  `/lifedrive-work` and `/tweet-file-work` exceeded 60–120 s, while `mm://`
  listings still answered in 353 ms. Anything built on `FilesCopy`, `FilesFlush`
  or work-tree paths stalls node-wide in that state.
- `Leither.log` rotated a 270 MB file every 2–3 minutes during this period.

### 5.7 Improvements and Open Costs Rechecked on V0.24.27

Rechecked on minipc on 2026-09-19, after the node moved from V0.24.23
(2026-09-16) through V0.24.26 to V0.24.27 (installed 2026-09-19 15:06 CST).
Two behaviors improved; the commit cost that motivated LifeDrive's commit
policy did not. The LifeDrive consequences are recorded in
[LifeDrive Backup Storage Plan](https://github.com/cfa532/LifeAlbum/blob/0876ac5a20e0b4f9225e63febe9426f8694795d9/docs/LIFEDRIVE_BACKUP_STORAGE_PLAN.md).

#### `FilesRm` no longer requires a node-administrator session

V0.24.23 refused `FilesRm` for every session but the node administrator's,
even for a tree the calling household session had created under
`/lifedrive-work`. LifeDrive's identity service had to borrow a short-lived
local administrator passport solely to remove its own unpublished staging
trees. On V0.24.27 a household session creates a tree under `/lifedrive-work`
and recursively removes it with the same session: verified with a probe that
logged in through the household authority certificate, created
`/lifedrive-work/<probe>/nested`, and removed the tree in 1 ms. The
administrator passport workaround was removed from the identity service in
LifeDrive 0.7.37.

#### An identical, already-pinned root CID skips the pin

Since the 2026-09-18 upgrade, committing a root CID that is already
recursively pinned no longer re-walks the tree. The log records it directly:

```text
[p2p][I] IpfsPinAdd skipped 1/1 CIDs already recursive-pinned
[p2p][I] IpfsPinAdd done: 1/1 CIDs pinned OK dur=0s
```

A repeated or retried commit of an unchanged tree is therefore free, where
V0.24.23 re-pinned every block on every commit. Retries that previously
doubled the pin load now cost nothing.

#### A changed root still pins the whole tree

The skip compares root CIDs; it is not incremental pinning. A probe copied
the Mate 30 backup container's `last` (19,756 blocks) into a fresh MiMei and
committed:

| Commit | Root CID | Measured | Server log |
|---|---|---|---|
| First, identical to the already-pinned source root | unchanged | 3 ms | skip path |
| After adding one empty directory | changed | 3m44.5s | `pin add OK … blocks=19756 dur=3m44.534s` |
| After adding a second directory | changed | 2m26.4s | `pin add OK … blocks=19756 dur=2m26.404s` |

Small containers now show repeat commits of ~1,000-block trees at 128–164 ms
(first commit after idle ~7 s), consistent with a per-block walk rather than
a diff against the previous version. Commit cost therefore still grows with
the whole MiMei, LifeDrive's quiet-hours commit policy, per-device containers
and 3-minute commit deadline remain necessary, and incremental pinning (pin
only the new blocks of a changed root, or pin asynchronously) remains an
open request to the Leither team.

### 5.8 Database MiMei Behaviour Verified on V0.24.27

Probed on minipc on 2026-09-19 with a throwaway identity and scratch type-`2`
MiMeis (`apps/lifedrive-identity/measure_catalog_test.go`, build tag `measure`).
These results decide how the LifeDrive household catalog commits
([Hybrid Storage](https://github.com/cfa532/LifeAlbum/blob/0876ac5a20e0b4f9225e63febe9426f8694795d9/docs/LIFEDRIVE_HYBRID_STORAGE.md)).

#### `MMBackup` cost follows the number of rows

One changed row, then `MMBackup`:

| Rows in the database | Backup time |
|---:|---:|
| 1,000 | 0.14 s |
| 10,000 | 1.6 s |
| 20,000 (60-byte values) | 1.4–1.6 s |
| 20,000 (600-byte values) | 1.5–1.7 s |
| 50,000 | 7.8 s |
| 100,000 | 28.4 s |

The cost tracks the row count, not the value bytes. `MMBackup` with nothing
changed returns the same version immediately. Commit many transactions and
back up once; never back up once per request.

#### One numbered version, a consistent snapshot

- Each `MMBackup` replaces the previous numbered version: `mmversions` read
  `["7","last"]`, then `["13","last"]`, then `["19","last"]`. There is no
  history of earlier database versions to retain or prune.
- A backup taken while other transactions commit is a point-in-time snapshot:
  over 9,000 concurrent transactions, `last` never held a row newer than the
  revision it recorded, and the commits were not blocked.
- `MMBackup` of a database that has never had a row written fails with
  `not exist`.

#### Transactions and reads

- `Begin`/`Hset`/`Commit` work per handle, at about 0.15–0.19 ms per `Hset`.
  `Rollback` discards the transaction's writes.
- `Hscan(handle, key, after, "", limit, false)` returns fields strictly after
  `after`, in order, at most `limit`; an absent hash returns an empty page.
- **`Hget` or `Get` of a missing field, and `MMOpen` of a missing numbered
  version, make the Go hprose binding panic** (`reflect: function created by
  MakeFunc using closure returned zero Value`) instead of returning an error.
  Read a database with `Hscan` only, and list versions with `mmversions`
  before opening one.
- Writes committed to `cur` and not yet backed up are visible to a new
  session.

#### References by CID

- `MMAddRef(sid, mid, cid)` accepts the node's own `Qm…` (CIDv0) identifiers,
  as `IpfsAdd` returns them. A `bafk…` CIDv1 and any `ipfs/<cid>` form are
  refused with `bad id`.
- References are counted per MiMei: two `MMAddRef` calls read back through
  `MMGetRef(sid, mid, "cur")` as count 2, and each `MMDelRef` removes one.
- A file imported with `IpfsAdd` and referenced this way stayed readable after
  its source file on disk was deleted.

#### Reading content by CID

- A household user session may **not** open `ipfs/<cid>` with `MMOpenUrl`
  (`1101 Permission denied`), nor a node Files-tree path; `/ipfs/<cid>` is
  `file does not exist`. `MFOpenMacFile` reads Leither's local `mac/` store,
  not IPFS content.
- `FilesStat` and `FilesCopy` do accept `ipfs/<cid>`.
- Leither's HTTP server serves `/ipfs/<cid>` with byte ranges (`206`) on the
  loopback upstream, without an RPC session.
- `MiMeiPublish` is refused for a non-author identity
  (`is not author, no publish rights`).

## 6. MiMei References

MiMei references connect one MiMei to another or retain referenced resources. They are separate from relationships inside one File-type MiMei tree.

Applications must maintain references explicitly because Leither cannot infer semantic relationships hidden inside application files. References are relevant when:

- A primary container points to a separately shareable File-type MiMei.
- One application references an independently stored object or collection.
- A container must retain an attached immutable resource.
- A provider must synchronize a direct child MiMei together with its parent.

Leither synchronization follows direct references, not an unlimited recursive object graph. Applications must therefore define their own traversal rules and use cross-MiMei references at intentional data boundaries.

## 7. Publication, Provision, and Synchronization

These operations have different meanings:

| Operation | Meaning |
|---|---|
| Backup | Commit `cur` as a new immutable version and advance `last` |
| Publish | Announce current MiMei information so the network can discover its state |
| Provide | Announce that this node can serve the MiMei's data |
| Find providers | Locate nodes that claim to serve a MiMei |
| Synchronize | Pull selected MiMei versions and data to the local node |
| Unprovide | Withdraw this node's provider announcement |
| Unpublish | Withdraw publication of the MiMei information |

#### Undocumented Variadic Options, and What `delref=false` Does

The installed `MMBackup` takes more parameters than the published spec shows.
`api/MiMei.md` documents `MMBackup(sid, mid, memo string) (ver string, err error)`,
but the live V0.24.02 signature is:

```text
MMBackup   func(sid, mid, memo string, opts ...string) (string, error)
MMAddRef   func(sid, mid string, fileids ...string) (int, error)
MMDelRef   func(sid, mid string, fileids ...string) (int, error)
```

TweetBackendApp's JavaScript entries pass `"delref=false"` into that tail on **all 39**
of their `MMBackup` calls, including pure additions. The Go port does the same on every
write path.

Probed on a live node, the option has **no observable effect**. Reading the reference
list back with `MMGetRef` after each arm, these four are indistinguishable:

| Option passed | Reference map after backup |
|---|---|
| none | unchanged |
| `delref=false` | unchanged |
| `delref=true` | unchanged |
| `zzz=nonsense` | unchanged |

Tested across a plain re-backup after a content change, and a full
`MMAddRef` → backup → `MMDelRef` → backup cycle. In every arm the added reference
survived a normal backup, and `MMDelRef` removed it regardless of the option.

The nonsense arm is the load-bearing control: **an unrecognised option is accepted
without error.** So an option being tolerated proves nothing about whether it is
implemented, and code that passes one cannot infer behaviour from the absence of a
failure. Do not read `delref=false` as a description of what backup does.

This was observed through `MMGetRef` on type-`2` (database) MiMei. An effect invisible
there — garbage collection of referenced object data, or different handling on
File-type MiMei — would not have shown up.

A distributed application should not treat a write as fully replicated merely because backup succeeded. It can track separately:

1. Local commit success.
2. Publication success.
3. Required provider announcements.
4. Replica availability according to policy.

The normal distributed lifecycle is:

```text
mutate cur
   -> backup to last
   -> publish MiMei state
   -> provide from approved nodes
   -> verify providers and replica health
```

Synchronization supports selecting versions and a source peer. Minimal synchronization without IPFS payloads is available for metadata-oriented recovery, while full synchronization retrieves the required content.

**Do not synchronize a MiMei this node already owns.** On V0.24.02, `SyncMiMei` panics
with a bounds error when the object exists but no *other* node provides it — the
provider list resolves to empty and is then indexed at `[0]`. The node reports it as a
recovered runtime error rather than a normal failure:

```text
SyncMiMei:(runtime.boundsError{x:0, y:0, signed:false, code:0x0})
  Leither/pnet.(*MiMei).getSyncInfo   dhtMiMei.go:495
```

Isolating cases: a MiMei with a genuine remote provider synchronizes cleanly; a
nonexistent MiMei is rejected earlier with a clean `30000:No MiMei mid[...]`; a MiMei
whose announced providers are all this node's own addresses panics. The node's own
background routines reach the same fault, so an application cannot avoid it entirely —
but it can avoid causing it by resolving the owning host first and skipping the
synchronize call when the owner is this node. That check is also correct on its own
terms: there is nothing to pull from yourself.

## 8. Identity, Sessions, and Rights

Leither identities are rooted in asymmetric keys. A node creates and retains a host key; its public-key digest identifies the node. The intended architecture uses certificates and passports for identity and delegated access, sessions for API use, and MiMei rights for authorization.

The upstream [key-authentication guide](https://github.com/3and4/Leither/blob/90362aef401c4d7ff53eb3573443f031f0be880f/doc/KEY_AUTH_GUIDE.md), dated 2026-09-06 and reviewed here on 2026-10-16, publishes a narrower supported contract: self-signed `CertFor=Self` passport login, the `IAuth` methods, key/certificate/passport CLI operations, and identity lookup with `GetVar(sid, "userid")`. It specifies Ed25519 with `sodiumv2` as the default for new keys. The same key derives the same user ID on different nodes, while authentication records remain node-local; synchronization of those records is not promised. `CertFor=Bind` and username/password login are outside that published contract.

The guide specifies a one-hour in-memory session lifetime and loss of sessions on node restart. Passport expiry is checked at login and does not shorten an already established session. These are upstream documented semantics, not a new runtime verification of every authentication path. In particular, they do not resolve the V0.24.62 Files-namespace and CID disclosure observations in section 5.2.

The installed CLI exposes `--key-file` on privileged operations such as `mimei publish` and `mimei setright`. Each user should have a distinct key; a shared node or application key would collapse user authority.

MiMei rights are stored per MiMei and recipient UID. The API shape is `MMSetRight(authorSID, mid, recipientUID, right)`, with `MMGetRight(sid, mid, uid)` available to retrieve the current value. The author SID identifies the authenticated authority making the change; the recipient UID is the stable identity receiving the permission. Consequently, the grant survives the expiry of the particular browser SID used by the recipient: a later authenticated SID is mapped back to the UID when File-MiMei data operations evaluate access. On MiniPC (2026-09-08), anonymous `MMOpen` returned a non-empty handle even with rights `0`, while `MFGetData` and `MFReaddir` both returned `1101:Permission denied`, through both loopback and public RPC. A successful open is not evidence of permission to read or list contents.

The `right` argument is a packed `uint64` made from Linux-style permission groups. Inside a group, read is `4`, write is `2`, and execute/use is `1`; combinations use bitwise OR, as with Unix file modes. For example, `6` is read plus write and `7` is read plus write plus execute. A larger value such as LifeDrive's `0x07276704` contains several such groups and must be interpreted using Leither's group layout rather than as one `rwx` triplet. Software changing only part of an existing grant should read it with `MMGetRight`, update the intended bits, and then call `MMSetRight` with the resulting complete mask.

Where the installed node's authorization behavior does not meet an application's requirements, applications may place a fail-closed compatibility authorization layer in front of Leither. Such a layer can enforce application rights and then execute an allowlisted operation with the owner's configured key, but this is application-level enforcement: a node administrator or anyone holding the owner key may bypass it. Applications must disclose that limitation and retain an adapter boundary for later native enforcement.

The private key file must remain outside distributed content, application repositories, logs, diagnostics, and client requests. It is an authority/signing credential, not a content-encryption key. Loss or compromise of the only owner key may cause loss or compromise of owner authority, so encrypted offline recovery and a tested future rotation procedure are required.

The CLI accepts common connection and identity options:

- `--node`: Target node URL, ID, or configured name.
- `--key-file`: Key used to authenticate or sign a request.
- `--ppt`: Passport file.
- `--dht-ids`: Selected DHT networks where supported.
- `--api-version`: Explicit API contract version.
- `--json`: Stable machine-readable command envelope.

Applications should use the corresponding APIs directly. The CLI remains valuable for administration, diagnostics, recovery, and reproducing an API operation manually.

## 9. API Boundaries

There are three distinct integration surfaces:

1. **CLI:** Complete administrative surface with structured JSON for supported commands.
2. **External API/RPC:** Programmatic counterpart of CLI operations and the preferred integration when an application needs the full node capability set.
3. **In-container MApp API:** The runtime interface supplied to application code running inside Leither. JavaScript entries receive `request` and `args` and use built-in `lapi` operations.

The public repository adds two useful forms of evidence inside those boundaries:

- `lapi/stub.go` describes the currently exported Go RPC client surface and handle/result types.
- `example/rpc/` demonstrates intended call sequences, including authentication, opening a MiMei, operating on the returned session, backup, temporary files, the node Files filesystem, and read-back.

Neither is automatically the contract of a newer installed binary. Repository Markdown, generated/client stubs, runnable examples, JavaScript MApp bindings, CLI help, and the binary can evolve at different times. A missing function in the current Go stub does not prove that the installed CLI or JavaScript runtime lacks the capability; conversely, a function described in an older Markdown file is not callable evidence.

CLI automation should request `--json` where supported and parse the documented envelope rather than terminal prose:

```json
{"exit_code": 0, "data": {}, "error": null}
```

An application-level MApp response is a separate contract and may use its own versioned success/error envelope.

TweetBackendApp demonstrates the JavaScript MApp model: plain handler files use built-in APIs for author login, MiMei creation/opening, temporary binary files, file objects, backup, publication, provider routing, and calls to other application entries. Its database-oriented persistence is an application choice and should not be generalized to file-oriented products.

The currently exposed in-container Go interface is not evidence of the complete JavaScript `lapi` surface. The installed CLI operations have corresponding APIs, but an application must confirm the exact JavaScript function names and argument contracts against its installed Leither version before freezing its adapter.

A file-oriented server application can therefore run inside Leither as dependency-free JavaScript and call File-type MiMei APIs directly. External API/RPC and CLI JSON remain important for administration, diagnostics, recovery, and capabilities not intentionally exposed to foreground clients; they do not require a separate Go or Node.js product server.

### 9.1 MApp Runtime and Entry Routing

Leither supports Go and JavaScript MApps, but their entry structure differs:

- A JavaScript MApp uses one self-contained `.js` file per entry. The filename without `.js` is the entry name. The file evaluates an expression such as `((request, args) => { ... })(request, args)` and uses the injected `lapi` global.
- A Go MApp is one flat interpreted package with a `RunMApp` function and an internal dispatch table. In container mode, the routing layer may place the effective entry name in `Request["entry"]` even when the positional `Entry` argument is empty.

MApps run inside the Leither interpreter and are not independent Node.js or native server processes. Go MApps must stay within the exact standard-library, third-party-package, and `lapi` allowlist registered by the installed interpreter.

A JavaScript MApp has **no module loading of any kind**. `require` exists as a
function, but on V0.24.02 every call fails with `GoError: Invalid module` — sibling
files in the same application, Node built-ins such as `fs` and `path`, and third-party
packages alike. Verified by probe.

Two consequences for application structure:

- Every entry file must be wholly self-contained. Shared helpers cannot be imported and
  must be repeated in each file, or generated into it by a build step before upload.
  TweetBackendApp's JavaScript entries repeat their response-wrapping helpers in 65 of
  68 files for exactly this reason; it is a platform constraint, not an oversight.
- An entry that calls `require` is dead on arrival — it throws on that line every time.
  A JavaScript entry needing filesystem, archive, or subprocess work cannot be written;
  the capability has to come from `lapi` or move out of the MApp entirely.

A JavaScript entry file must also **invoke** itself. The file is evaluated as an
expression and must end with the call, or the entry returns nothing and the caller sees
HTTP 200 with an empty body:

```javascript
((request, args) => {
    // ...
    return result
})(request, args)     // the trailing call is required
```

Deployment automatically registers the HTTP execution route:

```text
GET /entry?aid=<application-id>&ver=<version>&entry=<entry-name>
```

The documented `/entry` parameters are:

| Parameter | Meaning |
|---|---|
| `aid` | Backed-up application ID |
| `ver` | Application version, normally `last` for clients and `cur` only during controlled development |
| `entry` | JavaScript filename or Go dispatch entry |
| `mid` | Optional MiMei data source selected by the application contract |
| `nid` | Optional remote node identifier for cross-node execution |

`/entry` is useful for simple GET-style calls, operational checks, and debugging. It is not the preferred transport for chunked binary application traffic.

### 9.2 Go MApp Interpreter Contract

The TweetBackendApp Go port supplies a tested model for Go applications hosted by Leither:

- The uploaded source is compiled by Leither's `ixgo` interpreter, not executed as a prebuilt native binary.
- The application is one flat package named `lapp`; subpackages are not part of the promised MApp format.
- The package exposes one `RunMApp` function and performs its own entry dispatch.
- `go.mod` is useful for local syntax and type checking but is not read when Leither builds the uploaded application.
- Runtime imports must use the exact paths registered by Leither, including `Leither/lapi`.
- In local CLI mode the entry may arrive in the positional `Entry` argument. In deployed/container mode, verified versions may place it in `Request["entry"]` while leaving `Entry` empty. A portable dispatcher checks both.
- Do not leave legacy JavaScript MApp files whose entry names overlap the Go
  dispatcher's entries. On V0.24.05, `darwin/arm64`, a Go LifeDrive package compiled
  and dispatched correctly after those previous JavaScript server entries were
  removed. Browser objects such as `lifedrive.js` have a distinct name and are
  loaded through the web bootstrap rather than invoked as server entries. Re-probe
  mixed Go/browser packages after a Leither upgrade.
- Clients should address a backed-up version such as `last`. A verified runtime served a stale compiled `cur` after source upload until application backup created a new version.
- `--local` is suitable for read-only and dispatch probes, but a verified runtime did not provide author identity there, causing `BELoginAsAuthor` writes to fail. Container-mode testing is required for authorized mutations.

#### Interpreter-Provided Go Packages

Leither exposes an allowlisted package environment rather than unrestricted module
download, **and the allowlist changes between versions — it has shrunk.** Several
packages listed by earlier revisions of this document no longer resolve.

Verified on **V0.24.02, linux/amd64**, by compiling a one-import probe MApp on the
installed node. An unavailable import fails at compile time with
`could not import <path> (not found package)`.

| Import path | V0.24.02 | Note |
|---|---|---|
| `errors` | available | proven in production |
| `fmt` | available | proven in production |
| `io` | available | proven in production |
| `os` | available | probe returned `os.PathSeparator` |
| `sort` | available | proven in production |
| `strconv` | available | proven in production |
| `strings` | available | proven in production |
| `time` | available | probe returned `time.Second` |
| `encoding/json` | available | full `Marshal`/`Unmarshal` round trip verified, map keys sorted |
| `Leither/lapi` | available | in-container Leither API |
| `bytes` | **not found** | listed as available by earlier revisions |
| `encoding/gob` | **not found** | listed as available by earlier revisions |
| `net/http` | **not found** | listed as available by earlier revisions |
| `crypto/sha256`, `crypto/ed25519` | **not found** | no in-app hashing or signature verification |
| `unicode/utf16` | **not found** | surrogate-pair handling must be written by hand |
| `github.com/shirou/gopsutil/*` | **not found** | `mem` tested |
| `github.com/hprose/hprose-golang/v3/*` | **not found** | `rpc/core` tested |

Two consequences follow from the absences rather than the presences:

- Missing direct imports do not imply missing Leither capabilities.
  `InitStubByUrl(stub, url)` is an exported Go stub helper, not the public HTTP
  API. It provides outbound Hprose RPC without importing `net/http` or Hprose
  into the MApp. The current helper
  converts HTTP URLs to WebSocket URLs and appends `/ws3/`; supply an origin,
  not an endpoint path. Cryptographic operations likewise require a supported
  wrapper when their direct imports are unavailable. On 2026-09-08, MiniPC
  successfully ran the interpreted LifeDrive Hprose stub against the loopback
  identity service and returned an encrypted authorization response. This
  verifies the wrapper, not direct Hprose package imports.
- `encoding/json` is available and sorts map keys, so a hand-written JSON encoder is
  not required for deterministic output.

Do not infer the package set from a successful local `go build`, and do not carry a
package list forward across a version upgrade. Re-probe against the installed binary.

Import availability is not sufficient proof of correct behavior. On the verified Go-MApp runtime, `init()` functions and package-level variable initializers did not run: variables silently retained their zero values. Consequently:

- Construct dispatch tables, maps, sentinels, and configuration inside functions or an explicit per-call initialization path.
- Do not depend on imported libraries whose correctness requires package initialization until their required operations pass an installed-runtime probe.
- Treat a successful local `go build` only as a syntax/type check, not as a runtime-compatibility test.

#### Reaching lapi Capabilities Not on the Published Interface

`lapi.GetLApi()` returns a value typed as the `lapi.LApi` interface, but the object
behind it is a `*frame.LApi` whose method set is **larger**. Several operations exist
on the concrete value, are absent from `lapi.LApi`, and are absent from
`github.com/3and4/Leither/lapi` at every published version, so an application written
against the published interface cannot see them and will conclude they do not exist.

Verified present on V0.24.02, with signatures read from the live node using
`fmt.Sprintf("%T", ...)`:

| Method | Signature |
|---|---|
| `MiMeiSync` | `func(sid, dhts, mid string, param map[string]string) error` |
| `MiMeiPublish` | `func(sid, dhts, mid string) ([]lapi.DhtReply, error)` |
| `MiMeiProvide` | `func(sid, dhts, mid string) ([]lapi.DhtReply, error)` |
| `MiMeiUnprovide` | `func(sid, dhts, mid string) ([]lapi.DhtReply, error)` |
| `MiMeiUnpublish` | `func(sid, dhts, mid string) ([]lapi.DhtReply, error)` |
| `MiMeiIsProvider` | `func(sid, mid string) (bool, error)` |
| `RunMApp` | `func(entry string, req map[string]string, args []any, opts ...string) (any, error)` |

The `dhts` argument selects which DHTs to act on; `""` means all, which is what the
JavaScript runtime passes.

Reach them by asserting the handle to a locally declared interface, one per operation:

```go
type mimeiSyncer interface {
	MiMeiSync(sid, dhts, mid string, param map[string]string) error
}

if s, ok := api.(mimeiSyncer); ok {
	err := s.MiMeiSync(sid, "", mid, param)
}
```

This keeps the application importing only `Leither/lapi`. The internal `Leither/api`
package also exposes them, but its `GetLApi()` hands back `*frame.LApi` directly, which
couples the application to a node-internal type for no benefit. An assertion also
degrades safely: on a build without the method the assertion fails and the caller can
report an unavailable capability instead of failing to compile.

`Ed25519Verify` is **not** reachable by any route tested. Applications that verify
signatures must treat that as unavailable rather than degrading to a well-formedness
check that reads like verification.

#### Interpreter Memory Cost Per Request

Each Go MApp request retains interpreter compilation state that is never released.
Measured on two independent nodes, same trivial entry, 30 sequential requests, with a
JavaScript MApp on the same node as control:

| Node | Go MApp | JavaScript MApp |
|---|---|---|
| Node A | +609 MB (20.3 MB/request) | +22.7 MB (0.76 MB/request) |
| Node B | +242 MB (8.1 MB/request) | +22.7 MB (0.76 MB/request) |

A heap profile of a node holding 610 MB is dominated by compilation, not application
data: `ixgo.(*function).regInstr`, `go/types.(*Checker).recordTypeAndValue`,
`go/ssa.createFunction`. Goroutine count stays low, so this is retention, not leaked
concurrency.

The operational consequence is severe and easy to misdiagnose. Past roughly 3 GB a node
still answers a plain HTTP probe in milliseconds while every application call takes
minutes, so port-level monitoring reports the node healthy while clients time out. One
node was measured answering `HEAD /` in 8 ms and a trivial `health` entry in 2 m 3.8 s.

V0.24.05 is materially better but does not yet prove the retention is gone. A
LifeDrive probe on `darwin/arm64` uploaded 128 MiB in eight 16 MiB `upload_write`
calls, then trashed and purged the test object. The complete 12-call sequence took
about 2.44 seconds and process RSS rose from 257,984 KiB to 300,800 KiB. That
42 MiB upper bound includes upload/IPFS buffers and was not separated with a heap
profile, so it cannot be labelled entirely as interpreter retention; it also did not
return immediately after the probe.

Use large bounded chunks to reduce invocation count, monitor the RSS of the process
that owns the listening socket, and remeasure after every Leither upgrade. Scheduled
restarts remain a prudent guard until a long-run V0.24.05 heap profile demonstrates
that compilation state is released.

### 9.3 Browser Web-Application Bootstrap

Leither has two observed web-entry paths. The direct local form is:

```text
http://<node-ip>:<leither-port>/entry?aid=<application-id>&ver=last
```

A domain registered with `mimei setdomain` returns a provider-selecting bootstrap.
By contrast, Tahoe V0.24.05 returned the application's `index.html` directly for the
LAN URL, without defining `window.getParam()` and without rebasing external assets.
Neither form requires a node-root path route or an application-named URL directory,
but a web package that supports both forms must account for this difference.

The domain bootstrap performs this sequence:

```text
resolve the application MiMei and its providers
  -> expose aid, mid, version, provider address groups, and selected node
     through window.getParam()
  -> race the candidate provider addresses
  -> fetch the published application index from the first reachable provider
  -> inject that index into the bootstrap document
  -> rebase application link and script objects below
     http://<provider>/mm/<application-mid>:<version>/
  -> preload and execute the application scripts in document order
```

The observed bootstrap parameter block includes `aid`, `mid`, `ver`, `addrs`,
`ips`, `CurNode`, and `remote`. `addrs` preserves addresses grouped by provider;
`ips` is the flattened candidate list used by the bootstrap race; and `CurNode`
identifies the address that supplied the application page. Fields should be treated
as versioned bootstrap data: clients must validate their types and retain a safe
fallback for development outside Leither.

TweetWeb is the working packaging precedent. Its application MiMei contains
top-level objects such as:

```text
index.html
lifedrive.js (or another non-conflicting browser-script name)
inline CSS (or an external stylesheet with a distinct normalized short name)
hprose.js and other explicitly referenced browser dependencies
```

For a domain bootstrap, the built `index.html` can refer to these objects by bare
names rather than a node-root path or application-named subdirectory. Leither's
injector resolves the bare name against the `/mm/` base.
Vite 8 emits leading `./` paths for a relative base, so TweetWeb removes that prefix
in a post-build HTML transform. Other Leither web applications should do the same
unless the exact installed Leither version has been verified to accept `./` object
names.

A direct Tahoe `/entry` response does not perform that rewrite. A bare
`lifedrive.js` reference resolves to node-root `/lifedrive.js`, which returns 404;
the published object is available at
`/mm/<application-id>:<version>/lifedrive.js`. LifeDrive supports both entry paths
with a small inline loader in `index.html`: it prefers `window.getParam()` when a
domain bootstrap supplied it, otherwise reads `aid` and `ver` from the direct URL,
then constructs the `/mm/` module URL. This loader contains no deployment-specific
application ID or node address.

Asset naming is also version-sensitive. TweetWeb publishes `index_entry.js`, but
Tahoe V0.24.05 rejected packages containing `index.html` together with either
`index_entry.js` or `index.css` with `entry shortName[index] ... exist`: that runtime
derives the same short name from all three objects. LifeDrive therefore uses the
distinct top-level name `lifedrive.js` and embeds its CSS in `index.html`. Preserve
the bootstrap layout, but select browser object names that do not collide under the
target runtime's application-entry normalization.

After domain injection, application initialization reads `window.getParam()` rather
than hardcoding an application ID or node address. The application ID comes from
`aid` (with `mid` as a compatibility fallback), while the current RPC origin comes
from `ips[CurNode]`. In direct mode the same values come from the `/entry` URL and its
origin. Applications that support provider failover should retain and order the
remaining provider-grouped addresses instead of discarding them.

This browser bootstrap is separate from a MApp entry invocation. Supplying
`entry=<entry-name>` addresses server code; omitting `entry` loads the application's
web entry page.

### 9.4 Browser RPC and Binary Arguments

Leither exposes Hprose RPC over HTTP and WebSocket endpoints:

```text
https://<node-or-gateway>/webapi/
wss://<node-or-gateway>/ws/
```

These public endpoints are distinct from the Go `lapi.InitStubByUrl` helper,
which currently forces WebSocket `/ws3/` and sets a 24-hour default timeout.
LifeDrive uses explicit 60-second method timeouts for its local household calls.
Its iOS and browser clients, and the identity service's calls back into Leither,
already use Hprose HTTP through `/webapi/`.

The network API documentation (`Leither/api/Net.md`, section 4) also documents
outbound HTTP operations:

```go
Proxyget(sid, url string) (string, error)
Proxypost(sid, url string, header map[string]string, postContent interface{}) (string, error)
```

Both names appear in MiniPC's public `/webapi/` method list. On 2026-09-08,
`Proxyget` with an empty SID targeting the loopback identity-service health
endpoint returned `5006:Insufficient permissions Run as administrator`.
Thus public method discovery does not imply anonymous access. The checked Go
`lapi.LApi` / `INetStub` declarations do not expose these methods directly;
their availability over public RPC does not establish availability through the
injected Go MApp interface. Do not replace the fixed local household connection
with an anonymous proxy call or give pairing clients administrator credentials.

A browser client creates an Hprose client with an explicit allowlist of node methods and invokes a MApp through:

```javascript
client.RunMApp(entryName, request, args)
```

The two call channels have distinct responsibilities:

- `request` is a JSON-compatible parameter map. It carries `aid`, `ver`, API version, IDs, offsets, declared sizes, media types, and authorization material defined by the application.
- `args` carries positional values that do not belong in query parameters, especially binary chunks. A JavaScript upload entry should normalize a received chunk with `new Uint8Array(args[0])` before calling `MFSetData`.

Every foreground call should explicitly send the backed-up application `aid` and normally use `ver: "last"`. Binary uploads should use bounded chunks, authoritative offsets returned by the server, an increased but finite RPC timeout, and a separate completion call. The completion entry—not the final chunk write—performs validation and the MiMei backup.

TweetBackendApp's `upload_ipfs` entry and TweetWeb upload client demonstrate the proven transport sequence:

```text
open or reuse upload handle
  -> RunMApp(upload entry, metadata including offset, [Uint8Array chunk])
  -> repeat with returned handle
  -> RunMApp(upload entry, finished=true)
  -> convert or attach completed data
```

A File-tree application may keep bytes in a temporary file, attach the completed file to its final path under a writable node Files directory, flush that directory, and assign the resulting root CID to the File MiMei. In the installed workflow, `MFSetCid` returns the assigned File-MiMei version exposed through `last`; no catalog is required for this boundary. This is the node-Files assembly route described in Section 5.2, not evidence that the node Files root and the File-MiMei filesystem are the same namespace.

### 9.5 Proven MApp Conventions from TweetBackendApp

TweetBackendApp is a useful behavioral reference for:

- `BELoginAsAuthor()` before owner-authorized creation or mutation.
- `MMCreate` with application ID, namespaced extension, deliberate mark, data type, and rights.
- `mm://<mid>:cur/...` paths for File-MiMei tree writes and `mm://<mid>:last/...` for committed reads.
- One backup after a successful logical mutation.
- Explicit application-version response envelopes so transport failure and application rejection remain distinguishable.
- Passing `aid` and `ver: "last"` when one entry calls another.
- Binary upload chunks in `args`, with offsets and session handles in `request`.
- Logging contextual entry and object identifiers without logging private keys or raw credentials.
- Closing handles on every exit path. The Go port improves on older JavaScript entries by making this invariant explicit.

Its Database calls—`Get`, `Set`, hashes, lists, and sorted sets—belong to its type-`2` MiMei model. They are not required by the MApp runtime and must not be copied into an application whose authoritative model is File type. A basic cloud drive can use ordinary directories and files directly; JSON is needed only for optional information the filesystem cannot represent.

### 9.6 Public Web Gateway Boundary

A public HTTPS application should not expose a private Leither node's entire `/webapi/` or `/ws/` surface merely to make one MApp reachable. It should use a restricted same-origin gateway that:

- Terminates public TLS.
- Proxies only the required RPC transport.
- Allows only `RunMApp` calls for an explicit application ID and entry allowlist.
- Rejects client-selected application IDs, node administration methods, arbitrary entry names, and raw key paths.
- Applies request-size, chunk-size, rate, timeout, and origin limits.
- Preserves application envelopes while separating gateway failures from MApp failures.
- Reaches the owner-controlled Leither node through a private or authenticated upstream route.

This pattern also avoids browser mixed-content and Private Network Access restrictions that arise when a public page attempts to call LAN, RFC6598, localhost, or private tailnet addresses directly. Private or Leither-hosted applications may still connect directly to a node when their browser origin and network policy permit it.

### 9.7 Leither Application Deployment Lifecycle

A Leither application has a source name and a backed-up application identity. Its basic deployment lifecycle is:

```text
application source directory
  -> lapp uploadapp <application-name>
  -> lapp backup -a <application-name>
  -> query appId with lapp showapp --json under the same identity
  -> mimei publish --mid <appId> under the same identity
```

The application directory name identifies the source supplied to `uploadapp`. The `appId` is determined by the publisher identity and application name and becomes publishable after backup; a different publisher using the same name addresses a different application. Because the backed-up application is published through the MiMei publication command, the authoritative `appId` queried after backup is supplied as the `--mid` value.

An operator script should preserve this ordering, stop on a failed upload or backup, query the application with the same publisher identity, validate its name, and publish only the returned `appId`. It must not embed a deployment-specific publication ID. Deployment credentials remain outside the application source and source-control repository.

Application backup and application publication are distinct from an application's own data lifecycle. For a file-oriented application, `lapp backup` packages the application, while `mimei backup --mid <data-mid>` commits a File-type data tree.

## 10. Operational Model for File-Oriented Applications

A file-oriented cloud-drive application can preserve the following invariants:

- Every authoritative container is File type.
- The directory tree is the authoritative parent-child relationship.
- Canonical paths identify files and folders in the fundamental API.
- File bytes and directory mutations share one MiMei backup boundary.
- Normal reads use `last`, never partially written `cur`.
- Publish and provide status are observable separately from commit status.
- Cross-MiMei references are maintained at every partition boundary.
- All caches and indexes are reconstructable from files and JSON objects.

## 11. Failure and Recovery Principles

- A failed upload must remain outside the visible directory tree until complete.
- A failed backup leaves the prior `last` version authoritative.
- A failed publication does not invalidate the local commit, but it leaves replication pending.
- A missing provider triggers repair according to replica policy.
- Conflicting writes are preserved as explicit conflict objects rather than silently overwritten.
- Restore creates a new working state based on history; destructive history deletion is a separate policy decision.
- Removal from approved providers cannot erase a copy already taken by an uncontrolled peer. Sensitive content must be encrypted before distribution.

## 12. Common Command Families

The command families below combine the original V0.24.02 review with the V0.24.63 inspection in section 12.1:

- `mimei`: Creation, file-tree mutation, versions, references, rights, publication, provision, synchronization, and inspection.
- `files`: Mutable IPFS filesystem operations.
- `ipfs`: Immutable content add/get, pinning, provider discovery, and announcements.
- `dag`: DAG-node retrieval and inspection; confirmed in V0.24.63 (section 12.1).
- `lpki`: Keys, certificates, passports, sessions, services, and application execution.
- `lapp`: Application upload, versioning, release, restoration, and domain binding.
- `getvar` and `stat`: Node identity, health, capacity, and resource information.
- `repo`: Repository size, local references, and garbage collection.
- `dht` and `swarm`: Routing, peer discovery, and connectivity.
- `route`: Path-prefix routing to MiMei applications.
- `user`: User information associated with the node.

A storage-oriented application can initially restrict its runtime dependency to `mimei`, file operations, per-user key-file authority, security-capability detection, node health, and provider discovery. Select identity and rights integration using the supported contract and observed enforcement boundaries in sections 5.2 and 8. Application hosting, generalized DHT operations, consensus, and credit can remain separate concerns.

### 12.1 Inspecting an IPFS file from the CLI (2026-10-16)

On gen8, recursive traversal of every `Available Commands` section from
`./Leither -h` inspected **159 help entries**, including parent commands, on
V0.24.63 (amd64, API 1). The installed help exposes:

| Operation | Command | Output or boundary |
|---|---|---|
| List a File MiMei | `./Leither mimei ls <mid-or-path>` | Child names, CIDs, types and sizes; not file contents |
| Retrieve an IPFS DAG node | `./Leither dag get /ipfs/<cid>` | DAG JSON representation, including encoded `Data` and `Links` for DAG-PB nodes |
| Retrieve a file to a local path | `./Leither ipfs get /ipfs/<cid> -o <local-file>` | `get` and `--outfile` are exposed by help; the requested tweet CID was unavailable on the inspected node, so its download was not verified |
| Retrieve through a MiMei path | `./Leither mimei get <path> -o <local-file>` | `get` and `--outfile` are exposed by help; use the target runtime's supported path form |

For the `core.json` CID from the tweet listing discussed on this date:

```bash
./Leither dag get /ipfs/QmXQ9LjYkvQYyASnAzRkA5kvv5Hw9kL5t6wEWEDAPAUZvv
```

This is a DAG inspection command, not a plain-text JSON-file reader. A successful
read of an existing directory DAG on the inspected node returned:

```json
{"Data":{"/":{"bytes":"CAE"}},"Links":[]}
```

The `Data` field is encoded DAG data. For a UnixFS file, decoding its base64 bytes
alone can still leave a UnixFS structure; a multi-block file also requires following
its links. Do not describe this output as the original `core.json` text.

No dedicated `cat` or `read` command appeared anywhere in that installed command
tree. `ipfs get -o /dev/stdout` is an output-file workaround, not a separately
exposed print command; successful content retrieval through that workaround was
not established. A bare CID was rejected by both `dag get` and `ipfs get` during
inspection, so retain the `/ipfs/` prefix for these commands. This observation is
specific to these read commands and does not replace the different `files cp`
source-path rules recorded in section 5.6.

To check another installation, follow each command family recursively rather than
assuming that `ipfs -h` lists every content-inspection operation:

```bash
./Leither version
./Leither -h
./Leither dag -h
./Leither dag get -h
./Leither ipfs get -h
./Leither mimei get -h
```

## 13. Verification Sources

### 13.1 Upstream documentation repository

The most complete relevant public repository found in the 2026-10-16 review is
[3and4/Leither](https://github.com/3and4/Leither). Use its documentation, RPC stubs
and examples together. This independently maintained reference lives in
`cfa532/Leither/cfa532-docs/`. At consolidation on 2026-10-16, the fork baseline
was `3381dad3c296fd54e9a19f1be78904adf040f2e9`, containing upstream
`90362aef401c4d7ff53eb3573443f031f0be880f` and its key-authentication guide.
The upstream source files remain at their existing paths; use the
[complete source index](../UPSTREAM_INDEX.md) to navigate them and the
[maintenance guide](../MAINTENANCE.md) when synchronizing the fork.

The following links pin the upstream snapshot to
`90362aef401c4d7ff53eb3573443f031f0be880f`:

| Source | Use |
|---|---|
| [README](https://github.com/3and4/Leither/blob/90362aef401c4d7ff53eb3573443f031f0be880f/README.md) | System overview, setup and introductory CLI/HTTP examples |
| [API index](https://github.com/3and4/Leither/blob/90362aef401c4d7ff53eb3573443f031f0be880f/api/Api.md) and [MiMei API](https://github.com/3and4/Leither/blob/90362aef401c4d7ff53eb3573443f031f0be880f/api/MiMei.md) | Interface families and storage operations; verify signatures against the deployed binding |
| [Architecture documents](https://github.com/3and4/Leither/tree/90362aef401c4d7ff53eb3573443f031f0be880f/doc) | MiMei, directories, PKI and system concepts |
| [Key-authentication guide](https://github.com/3and4/Leither/blob/90362aef401c4d7ff53eb3573443f031f0be880f/doc/KEY_AUTH_GUIDE.md) | Supported self-key login contract, wire shapes, session semantics and exclusions |
| [RPC stubs](https://github.com/3and4/Leither/tree/90362aef401c4d7ff53eb3573443f031f0be880f/lapi) and [examples](https://github.com/3and4/Leither/tree/90362aef401c4d7ff53eb3573443f031f0be880f/example) | Typed interfaces and intended call sequences |
| [Install.md](https://github.com/3and4/Leither/blob/90362aef401c4d7ff53eb3573443f031f0be880f/Install.md) | Historical installation notes and changelog; its listed V0.21.03 is not the current installed version or a complete CLI reference |

Documentation-index review does not mean every API or example was rerun. Preserve
the version/date attached to each earlier observation, and use installed recursive
help for the current CLI surface. The upstream tree reviewed here does not contain
`docs/skill-schema.md`, although the installed binary advertises a
`gen-skill-doc` command; do not link to that file as an existing upstream manual.

### 13.2 Runtime and application evidence

This reference was reconciled against:

- The current public Leither documentation at `http://vzhan.cn/#doc.html`, including CLI commands, MApp structure, `/entry`, and the public `lapi` categories.
- The `3and4/Leither` repository, especially `api/MiMei.md`, `lapi/stub.go`, and the `example/rpc/files`, `example/rpc/mfile`, and backup examples.
- The installed Leither `MiMei`, application, and API documentation for File-type constants and function signatures.
- TweetBackendApp's JavaScript entries for real `request`/`args`, `RunMApp`, author-session, version, backup, publication, and binary-upload behavior.
- TweetBackendApp's Go MApp, `go/README.md`, and live-node issue record for handle lifecycle, `cur`/`last` invariants, entry dispatch, registered imports, and interpreter constraints.
- Direct probes against an installed **V0.24.02, linux/amd64** node (2026-08-25) for the
  Go package allowlist, the `lapi` methods reachable only by type assertion and their
  signatures, per-request interpreter memory growth, the `SyncMiMei` empty-provider
  panic, and the absence of JavaScript module loading. Where this document and an
  earlier revision disagree, the probe result is the later evidence: the Go package
  allowlist in particular has **shrunk** between versions, so treat any package list —
  including this one — as version-pinned.
- TweetWeb's generated Leither bootstrap page and published top-level web objects for
  the `/entry` provider race, `window.getParam()` contract, `/mm/<mid>:<ver>/` asset
  rebasing, and ordered script injection.
- TweetWeb's browser client for the deployed `/webapi/` and `/ws/` Hprose transports, explicit API method lists, timeout handling, chunked binary calls, and public-gateway boundary.

When sources differ, use this evidence order:

1. Observed behavior and help from the installed Leither binary being deployed.
2. A successful round-trip through the installed API binding used by the application.
3. Working applications against the same Leither version.
4. Current repository stubs and runnable examples.
5. Repository and website Markdown documentation.

The lower-ranked sources remain valuable for architecture, intent, and discovering capabilities to verify. Existing applications and examples are behavioral evidence, not substitutes for version pinning and capability detection.

## 14. Reference Position

> Leither is a versioned data-and-application container and distributed serving platform. MiMei provides stable identity above immutable content, supports both File and database storage models, and combines versions, references, rights, publication, provision, and synchronization. Applications that use File-type MiMei can expose an ordinary directory tree directly and add JSON sidecars only when their product needs metadata beyond filesystem semantics.

This model is more complete than describing Leither as only an IPFS serving node.
