# Leither Go MApp — issues

**Canonical home:** `cfa532/Leither/cfa532-docs/`
**Consolidated:** 2026-10-16 from [TweetBackendApp/LEITHER_ISSUES.md](https://github.com/cfa532/TweetBackendApp/blob/77071bcbf61357e8c685ee5c45c5742194a0dde3/LEITHER_ISSUES.md).

Historical runtime evidence, with each original version and correction retained. This relocation does not establish that an older issue persists in V0.24.63. Unqualified application source paths refer to the source TweetBackendApp revision.

---

Found porting a 68-entry backend from JS to Go. Reproduced on a live node at
**V0.23.95** and **re-verified unchanged on V0.24.02**, linux/amd64. Probes are
single files run with `./Leither lpki runapp --local <dir> x -a gen8.key`.

**Correction (V0.24.02).** P2/P3 were originally reported as missing APIs. They
are not missing — they are unpublished. `lapi.GetLApi()` returns a `*frame.LApi`
whose method set is larger than the `lapi.LApi` interface it is typed as, and it
carries `MiMeiSync`, `MiMeiPublish`, `MiMeiProvide`, `MiMeiUnprovide`,
`MiMeiUnpublish`, `MiMeiIsProvider` and `RunMApp`. Asserting the handle to a
locally declared interface reaches all seven, and this app now does. Only
`Ed25519Verify` (P7) is genuinely absent. P2 and P3 are rewritten below as what
they actually are: an interface-surface problem, not an absence.

**Not a problem:** the Go language implementation. A conformance probe covering
defer ordering, named-return-defer, panic/recover, per-iteration loop capture,
interfaces, type switches, embedding with pointer receivers, `errors.Is/As`+`%w`,
slice 3-index aliasing, 50 goroutines with WaitGroup/Mutex, channels+`select`,
rune ranging and method values passed completely. Everything below is runtime
integration and library surface.

---

## P1 — Package-level `var` initialisers and `init()` never run

`const` works; every `var` is its zero value.

```go
const c = 7                    // -> 7   ok
var i = 42                     // -> 0
var s = "literalOK"            // -> ""
var sl = []int{1, 2, 3}        // -> []
var f = makeVal()              // -> ""
func init() { x = "initOK" }   // -> ""   never runs
```

Compiles clean, fails silently, and package-level `var` is ordinary Go — most
real code and most libraries use it. Caused two faults here: a dispatch table
built in `init()` was empty so *every* entry was unknown; and two structs holding
database key names were zero, which would have written user data under **empty
key names** instead of erroring.

Fix on our side: no package-level vars at all. Tables/structs come from
functions, and an `errors.New` sentinel became an error *type* (a package-level
sentinel is `nil`, so `errors.Is` against it silently misclassifies).

## P2 — Replication APIs exist but are not on the published interface

`MiMeiPublish`, `MiMeiProvide`, `MiMeiUnprovide`, `MiMeiUnpublish`,
`MiMeiIsProvider` and `MiMeiSync` are all present on the concrete `*frame.LApi`
returned by `lapi.GetLApi()`, and all work. None is declared on the `lapi.LApi`
interface that the value is typed as, and none appears in
`github.com/3and4/Leither/lapi` at any published version — checked on `main`
(last commit 2026-08-09) and on tags `lapi/v0.1.0` and `lapi/v0.1.1`. They are
also absent from `api/MiMei.md`, which documents `MMCreate`, `MMOpen`,
`MMBackup`, `MMAddRef` and friends but no replication call.

The consequence is not that a MApp cannot replicate — it is that a MApp cannot
discover that it can. Written straightforwardly, the call does not compile, and
the natural conclusion is that the platform cannot do it. The original report
here reached exactly that conclusion.

Verified live, signatures read off the node with `fmt.Sprintf("%T", ...)`:

```
MiMeiSync        func(sid, dhts, mid string, param map[string]string) error
MiMeiPublish     func(sid, dhts, mid string) ([]lapi.DhtReply, error)
MiMeiProvide     func(sid, dhts, mid string) ([]lapi.DhtReply, error)
MiMeiUnprovide   func(sid, dhts, mid string) ([]lapi.DhtReply, error)
MiMeiUnpublish   func(sid, dhts, mid string) ([]lapi.DhtReply, error)
MiMeiIsProvider  func(sid, mid string) (bool, error)
```

These match the JavaScript call sites exactly, including the `""` second
argument for "all DHTs" — so the Go and JS surfaces are the same API, and only
the Go type declaration hides it.

**Ask:** declare these on `lapi.LApi` (or on the embedded `IMiMei`) and publish
the module. `Leither/api` does expose them, but by handing back the internal
`*frame.LApi`, which is not something an app should be asked to depend on.

**Workaround applied:** `caps.go` asserts the handle to one small interface
per method, so a build lacking a method degrades to a typed error instead of a
compile failure.

## P3 — `RunMApp` is declared on `ILApp`, which `LApi` does not embed

`RunMApp` is present on the concrete handle and works, but `LApi` embeds only
`IBackEnd`, `IAuth`, `IVarAct`, `IMiMei` and `INet` — not `ILApp` — so it is
invisible to a MApp written against the published interface. Five operations
here inherently span two owners on two nodes and cannot be split by the caller;
a follow, for example, writes to both users' accounts, which live on different
nodes.

Live signature:

```
RunMApp  func(entry string, req map[string]string, args []any, opts ...string) (any, error)
```

**Ask:** add `ILApp` to the interfaces `LApi` embeds. This is a one-line change
and is the smallest of the fixes listed here.

**Workaround applied:** same type assertion as P2.

## P4 — `Entry` is empty in container mode

Under `--local`, `RunMApp`'s `Entry` carries the entry name. Under `--id`/HTTP it
is empty; the name is only in `Request["entry"]`.

This contradicts the documentation — `examples/echo` states *"`Entry` 是入口名,
对应 CLI 的位置参数或 URL 的 `?entry=xxx`"* — and **both official examples
(`hello`, `echo`) dispatch on `switch Entry`**, so both serve nothing once
deployed. `mapp-development.md` lists it under troubleshooting with the advice
"use `--local` for development", but for a single-entry-point Go MApp this is the
difference between serving requests and not.

Filling `Entry` from the same source in both modes removes the trap.

## P5 — `.js` files anywhere in the tree disable a Go app's routing

Leither builds the entry list from `.js` basenames found **recursively**,
including `mmroot/`. With any present, the router rejects every request with
`unknown entry` before `RunMApp` runs — while correctly identifying the app as Go:

```
Dry run: app=twbe entry=health type=go (valid)
  Available entries: [bootstrap.min gtag hprose index_entry popper.min ...]
  Has Go sources: true
```

`Has Go sources: true` with a JS-derived entry list is contradictory. Practical
effect: a Go MApp cannot ship static web assets. Removing them gives
`Available entries: []` and the app works.

## P6 — Interpreter package coverage is versioned and allowlisted

The table below records the original Leither V0.23.95 probe and should be read
as historical evidence, not as the current package inventory:

| Resolves | Missing |
|---|---|
| `fmt` `io` `errors` `strings` `strconv` `sort` `time` | `bytes` `unicode/utf8` `unicode/utf16` |
| `encoding/json` `math/big` `sync` `os` | `encoding/base64` `crypto/*` `regexp` `net/http` `encoding/gob` |

**The published allowlist does not match the runtime.** [technical foundation](LEITHER_AND_MIMEI.md)
§9.2 lists `bytes`, `net/http` and `encoding/gob` as "confirmed available"; all
three are rejected on V0.24.02. It omits `encoding/json`, which resolves. A table
presented as confirmed should be generated from the interpreter's registry, or
carry the exact build it was measured on — an author trusting it would write code
that fails only at upload.

`bytes` and `unicode/utf8` are close to universal in Go. Note `encoding/json`
resolves while `unicode/utf8` does not, though json depends on it internally — so
the boundary isn't "what json needs". Cost here: hand-written JSON and base64
(~470 lines), and `json.go` spells out UTF-16 surrogate pairing.

**Updated package information:** the current Leither program also registers
`bytes`, `encoding/gob`, `fmt`, `io`, `net/http`, `os`, `strings`, `time`,
`Leither/lapi`, Hprose `rpc/core` and `rpc/websocket`, and gopsutil `disk` and
`mem`. The original claims that `bytes` and `net/http` are missing are therefore
superseded for the newer runtime. See [the source backend README](https://github.com/cfa532/TweetBackendApp/blob/77071bcbf61357e8c685ee5c45c5742194a0dde3/README.md) for exact import paths.

Import resolution alone does not guarantee native-Go behavior. The interpreter
still controls package initialization and runtime integration, so required
operations must be exercised on the exact deployed Leither version.

## P7 — No signature verification

No `Ed25519Verify` or equivalent, and with `crypto/*` missing it cannot be done
in-app. Agent authentication in this app is consequently unverified — it checks
only that a signature is 64 well-formed bytes.

## P8 — `Act` resolves nothing

Fails even for a name `GetVar` answers:

```go
l.Act("", "ver")     // ERR 5020:Variable names unavailable:ver
l.GetVar("", "ver")  // V0.23.95
```

`api/VarAct.md` marks it 草 (draft) holding 待定 APIs, so an empty registry is
presumably intended — worth saying plainly in the MApp docs, since `Act` reads
like the extension point for exactly P2/P3/P7. Unchanged on V0.24.02.

## P9 — Unwired backend stubs return `("", nil)`

`BEOpenAppDataNode` and `BELoginAsAuthor` return an empty session **and nil
error** when unwired (`lapi/backend.go`). The empty handle flows onward and dies
far away as `Zaddwithseq(): invalid key size` — reading like a key-length problem
rather than a missing identity. Also `BELoginAsAuthor` fails under `--local` with
`5:To be continued`, so no write entry is testable there.

## P10 — CLI papercuts

- `-r` takes one semicolon-separated string. Repeated `-r a=1 -r b=2` silently
  keeps only the last — indistinguishable from an entry ignoring its parameters.
  `&`, `,` and JSON forms are accepted and ignored.
- `showapp -a <name>` dies on an unrelated object:
  `30000:No MiMei mid[FPrNJIrJUfIbdWtNVirjVfv4pS-]`. `mimei show <appid>` works.
- `runapp --dry-run` requires an entry, so it can't validate an app as a whole.
- V0.24.02 removed the `-v` shorthand from `runapp` without an alias; scripts
  and docs using `-v last` now fail with `unknown shorthand flag: 'v'`. The
  replacement is `--app-ver`. (`mimei publish` kept its positional form, so that
  change was backward compatible — this one was not.)

## P12 — Node memory growth degrades a node into unusability

Observed on minipc (V0.24.02) after ~4 days uptime: the Leither process held
**3.37 GB RSS (42% of RAM)** while using only 5.2% CPU. The node stayed "healthy"
to a `HEAD /` probe but every application call stalled — a `HEAD /` took 10.6 s
and `/entry` calls, including `health` (which does no I/O at all), timed out
beyond 30 s. Static `/mm/` reads still worked intermittently, so the node looked
alive while being useless for applications.

A `systemctl restart leither-tweet` fixed it completely:

| | before | after |
|---|---|---|
| RSS | 3.37 GB | 521 MB |
| `HEAD /` | 10.6 s | 0.25 s |
| `health` | timeout >30 s | 0.34 s |
| `get_user_core_data` | timeout | 0.66 s |
| `get_tweet_feed` (10 tweets) | timeout | 0.29 s |

**Cause identified — see P17.** The growth is per-request leakage of interpreter
compilation state on Go MApp calls (~8-20 MB per request). The "gen8 sits at
87 MB" comparison originally made here was read off the wrong process; gen8's
serving process was at 12.8 GB. A restart clears it because it discards the
accumulated compilation artifacts, which is why the fix below works and why it
does not last.

Two things make this expensive to diagnose from the application side:

1. A degraded node is indistinguishable from a slow application. The client's
   only symptom is `PoolTimeoutError ... timed out after 15000ms`, which reads
   as "the app is slow".
2. The node keeps answering health probes while failing every real call, so
   client-side failover never triggers.

An RSS or request-latency figure in `getvar health` would let both operators and
clients act on this.

## P13 — Provider lists hand browsers unreachable private addresses

The domain template publishes every announced address, including
Tailscale CGNAT (`100.64.0.0/10`), VPN (`10.8.0.0/24`), and LAN
(`192.168.5.0/24`). A public page cannot use any of them:

```
Access to XMLHttpRequest at 'http://100.79.13.15:8002/...' from origin
'http://twbe.fireshare.us' has been blocked by CORS policy: The request client
is not a secure context and the resource is in more-private address space `local`.
```

Each blocked or unroutable address still consumes a client connection slot and a
15 s timeout, so a handful of them exhausts the RPC pool and *reachable* nodes
start failing with `PoolTimeoutError`. [technical foundation](LEITHER_AND_MIMEI.md) §9.5 describes the
browser restriction; the fix belongs upstream of it — announcements served to a
public origin should be filtered to publicly routable addresses.

## P14 — The node logs request credentials in plaintext

Calling an entry with a password writes it to the node log twice, before the
application runs and regardless of what the application does:

```
[frame][I] actionEntry /entry?...&username=u&password=wrongpw&version=v2
[app][D]   RunMApp PreLogin param=map[... password:wrongpw ...]
```

The first is at **Info**, so it is present in ordinary logs, not just debug. Any
entry taking a credential as a request parameter is affected — here `login`, and
`register`/`set_author_core_data`, which carry the password inside a `user` JSON
blob.

This application now redacts credentials from its own log lines
(`"password":"[redacted]"`), but that cannot help: the node has already logged
the plaintext by the time the entry is called. It affects the JavaScript
implementation identically and always has.

Worth either omitting known-secret parameter names from `actionEntry` and
`RunMApp PreLogin`, or moving both lines behind a level that production does not
enable.

## P15 — `SyncMiMei` panics when the only providers are the calling node

Both entry points — `BEMMSync` on the published interface and `MiMeiSync` on the
concrete handle — funnel into the same `SyncMiMei`, and both produce the
identical panic, so this is one bug in the node, not a property of either API:

```
MiMeiSync -> SyncMiMei:(runtime.boundsError{x:0, y:0, signed:false, code:0x0}) stack:
Leither/pnet.(*MiMei).getSyncInfo(...)   D:/workspace/src/Leither/pnet/dhtMiMei.go:495
Leither/pnet.(*MiMei).getSyncInfo2(...)  D:/workspace/src/Leither/pnet/dhtMiMei.go:404
Leither/pnet.(*MiMei).SyncMiMei(...)     D:/workspace/src/Leither/pnet/dhtMiMei.go:1061
```

`boundsError{x:0, y:0}` is index 0 of an empty slice. Three cases isolate the
trigger:

| mid | provider list | result |
|---|---|---|
| `d4lRy…` (twbe app) | has a remote provider | `<nil>` — succeeds |
| `zzzz…` (nonexistent) | none | clean `30000:No MiMei mid[...]` |
| `iq1w…` (a user) | **all six are this node's own addresses** | panic |

So the failing case is precisely: the object exists, and every announced
provider is the calling node itself. `getSyncInfo` evidently drops self from the
provider list and then indexes `[0]` without checking the result is non-empty.
A nonexistent mid is rejected earlier and never reaches that code, which is why
that case is clean.

This also explains the latency previously recorded here (5.6s / 33.5s / 16.9s
through the `sync_user` entry): the panicking calls were spending that time in
provider lookup before crashing, long enough to blow past TweetWeb's 15s client
timeout and make a `toggle_following` that succeeded server-side report failure.

**Fix:** a length check at `dhtMiMei.go:495`. The condition is already known to
the platform's own users — the JavaScript entries carry the comment "if original
tweet is on the same node, MimeiSync will throw an error" — so returning a typed
"nothing to sync from" error rather than panicking would match expectations.

### Confirmed in production, and the node does it to itself

Observed on a live node (minipc, `~/tweet/logs/Leither.log`, window
09:32-10:37). The node logs the argument that panics, one line above the error:

```
[p2p][D] SyncMiMei mids=[] err=SyncMiMei:(runtime.boundsError{x:0, y:0, ...}) stack:
[p2p][D] sync qg6sX_XmwJFFp-pqwNziMUbksOn fail SyncMiMei:(runtime.boundsError{...})
```

Across that window:

| `mids` argument | occurrences | outcome |
|---|---|---|
| `mids=[]` | 10 | panic, every time |
| `mids=[d4lRy...]` | 2 | fine |
| `mids=[SN61... 2Tcd... heWg...]` | 1 | fine |

So the trigger is exact: **`SyncMiMei` is entered with an empty mid list and
indexes `[0]` anyway.** This is the same fault as the provider experiment above,
one step further in — when nothing remains to pull from, the list `getSyncInfo`
builds is empty rather than the call being rejected.

It is not only reachable from an application. Two node-internal callers hit it
in the same window:

```
[cai][D] CheckDBProvide MiMeiProvide err=SyncMiMei:(runtime.boundsError{...})
Leither/pnet.(*MiMei).routineMiMeiSync(...)
```

`routineMiMeiSync` frames appear 23 times in the current log and 40 times in the
previous one, so the background sync routine is hitting this continuously,
independent of any MApp. An app-side workaround cannot help those.

Seven distinct panic events occurred in a 65-minute window, triggered from
`toggle_following`, `sync_user`, and `CheckDBProvide`.

**Workaround applied:** `syncIfRemote` in `caps.go` compares the owning host
against this node and skips the call entirely when they match. This covers the
application's own calls only.

## P17 — Every Go MApp request leaks megabytes of interpreter compilation state

**This supersedes P16, whose conclusion was wrong** (see the correction at the
end of this section). It is the most serious issue in this document: it makes a
Go MApp unusable in production regardless of what the app does.

### Measurement

Same node, same trivial `health` entry, 30 sequential requests, nothing else
running. RSS read from `/proc/<pid>/status` of the process that owns the
listening socket:

| node | app | RSS delta over 30 calls | per call |
|---|---|---|---|
| minipc | **Go** (`twbe`) | **+609 MB** | **20.3 MB** |
| minipc | JS (`tweet1`) | +22.7 MB | 0.76 MB |
| gen8 | **Go** (`twbe`) | **+242 MB** | **8.1 MB** |
| gen8 | JS (`tweet1`) | +22.7 MB | 0.76 MB |

The JS control is identical on both nodes to three significant figures, so the
harness is sound and the difference is the application runtime, not the node,
the entry, or the measurement.

The memory is never returned. minipc climbed 116 MB -> 740 MB during the 30
calls above and stayed there; gen8's serving process has reached **12.8 GB**
over four days of ordinary use.

### Cause

A heap profile from a live node (`/debug/pprof/heap`, 610 MB in use) is
dominated by the interpreter compiling Go source, not by application data:

```
77.44MB  12.70%  github.com/goplus/ixgo.(*function).regInstr
72.34MB  11.86%  go/types.(*Checker).recordTypeAndValue
31.00MB   5.08%  go/types.newVar
29.08MB   4.77%  github.com/goplus/ixgo.(*visitor).function   (cum 150.82MB, 24.73%)
26.51MB   4.35%  golang.org/x/tools/go/ssa.createFunction
16.42MB   2.69%  go/types.(*Checker).recordUse
15.00MB   2.46%  golang.org/x/tools/go/ssa.(*Function).newBasicBlock
```

Type-checker state, SSA function bodies and registered instructions — the
artifacts of compiling the app — are retained per request rather than compiled
once and reused, or compiled and freed. Goroutines are not the problem: the same
node showed only 69.

This is also the likely explanation for the throughput gap measured separately
here: the Go MApp serves `health` in ~200 ms against the JS app's ~9 ms on the
same node, a ~21x difference that no application logic accounts for.

### Consequence

A node dies after a few hundred Go MApp requests. At 20 MB per call, 150 calls
is 3 GB. Past roughly 3 GB the node still answers a plain HTTP probe in 8 ms
while every application call takes minutes — a `health` entry measured at
**2 m 3.8 s** — so clients time out and the node looks healthy to any monitor
that probes the port rather than the app.

Observed end to end: a browser login failed three times with "Could not fetch
user data" because the account's home node had degraded this way, while the
identical call served in **0.08 s** against the same node freshly restarted.

**No application-side workaround exists.** The allocation happens in the node
before app code runs.

### Correction to P16

P16 reported "gen8 1d 17h / 0.08 GB / 1560 calls" against "minipc 58 min /
3.28 GB / 369 calls" and concluded the growth did not track application traffic.
The gen8 figure was taken from the wrong process: that host runs two `Leither`
processes, and `ps -C Leither | head -1` returned a stray `Leither ipfs add`
helper (12 MB, idle) rather than the node. gen8's serving process was at
**12.8 GB**, not 0.08 GB.

The corrected conclusion is the opposite of P16's: growth tracks application
traffic closely — but only **Go** MApp traffic, which is why a node serving
mostly JS appears stable. Always resolve the pid from the listening socket:

```
PID=$(ss -lntp | grep ":<port>" | grep -oE "pid=[0-9]+" | head -1 | cut -d= -f2)
awk "/VmRSS/{print \$2}" /proc/$PID/status
```

## P18 — A published version is not served under `last` until the node restarts

`twbe.sh` uploads the sources, backs them up as a new version and publishes the
app mimei. `showapp` then reports the new version as `Last`, and that version is
fetchable by number — but the node keeps serving previously compiled code to
anyone who asks for `ver=last`.

Observed on the entry `get_tweet_id_list`, whose reply changed shape between
versions: a `[]lapi.ScorePair` became a list of maps, so JSON key order changed
from struct field order to sorted order, making the running version visible from
outside.

| request | ksbox | gen8 |
|---|---|---|
| `ver=last` | old | old |
| `ver=cur` | new | old |
| `ver=1603`, explicit | new | new |

Two further backups (1601, then 1603) did not move what `last` resolved to. A
node restart did. So the sources upload, back up, publish and synchronise
correctly — only the label resolution is stale, and it is stale per node: after
restarting one node, the other still served the old code under `last`.

Practical effect: everything asks for `last`. The clients send `ver: "last"`,
and this app's own cross-node calls pass it too, so a fix can be published
repeatedly and still run nowhere while every step of the deployment reports
success.

**Workaround:** restart the node after publishing, and confirm with an entry
whose output differs between versions rather than trusting `showapp`.

## P19 — `FilesRm` is refused to the MApp author session

Observed on gen8, **V0.24.24**, 2026-09-18. A Go MApp that gets its session from
`BELoginAsAuthor()` can `FilesMkdir`, `FilesCopy`, `MFTemp2Files`,
`FilesFlush` and `MFSetCid` under its own scratch root, but every `FilesRm` on
that same root fails:

```
[mapp][W] Tweed add_tweet: remove staging tree: 5006:Insufficient permissions Run as administrator
[mapp][W] Tweed update_tweet: remove staging tree: 5006:Insufficient permissions Run as administrator
[mapp][E] Tweed Error node_update_score: replace retained tree for B_xYF...: 5006:Insufficient permissions Run as administrator
```

The same session created the paths it is not allowed to remove, and the refusal
happens for both recursive (`FilesRm(sid, dir, true, false)`) and single-file
(`FilesRm(sid, file, false, false)`) calls.

Why this blocks File MiMei writes: a File MiMei is updated by copying its
committed root into a scratch directory, replacing the changed entries, then
`FilesFlush` + `MFSetCid`. `MFTemp2Files` will not overwrite an existing path,
so replacing an entry (for example `core.json` when a tweet is edited) requires
`FilesRm` first. With `FilesRm` refused:

- every edit of an existing File object fails — the client sees
  `5006:Insufficient permissions Run as administrator` from `update_tweet`;
- removing a collection member (un-like, un-bookmark, comment removal) fails;
- scratch trees are never removed, so they accumulate under the app's work root
  on every write, including writes that otherwise succeed.

Creating a new object is unaffected, because its scratch tree starts empty.

**Fix:** allow `FilesRm` for the author session on paths that session created
(or on the app's own Files root), consistent with the Files calls it is
already allowed to make there. **No application-side workaround is applied.**

## P11 — Documentation

- Broken links on vzhan.cn (HTTP 500): `doc.html`, `capabilities.html`,
  `start.html`, `docs/MAPP_CONTAINER_ARCHITECTURE.md`, `app/example/main.go`,
  `app/example/README.md`. The last two are cited by `mapp-development.md` as the
  container-architecture and 8-entry references — exactly what P4/P5 need.
- `api/VarAct.md` gives `Act(...) error`; `lapi/stub.go` declares `(any, error)`.
- `troubleshooting.md` is required reading per `agent-skill-spec.md` but is not
  linked from `mapp-development.md`.

---

## Status on V0.24.02

P4 and P5 are now documented in [technical foundation](LEITHER_AND_MIMEI.md) §9.1/§9.2, which is a real
improvement. P4 remains a trap rather than a footnote, though: both official
examples (`hello`, `echo`) still dispatch on `switch Entry`, so both compile,
pass `--local`, and serve nothing once deployed.

P2 and P3 are reclassified: the APIs exist and this app now uses them, so
neither blocks production any more. What remains is that both are unreachable
through the published interface, which is a discoverability and stability
problem rather than a functional one.

P1, P6, P7, P8 are unchanged. P15 now has a precise root cause and repro, and
P17 — the per-request interpreter memory leak — is new and is the most serious
item here. P18 is new: a published version is not served under `last` until the
node restarts, which makes a successful deployment indistinguishable from one
that changed nothing.

## Priority

0. **P19** — `FilesRm` refused to the author session: no File MiMei can be
   edited, and scratch trees leak on every write.
1. **P17** — every Go MApp request leaks 8-20 MB of interpreter state. A node
   dies after a few hundred calls, and no application-side fix is possible.
   This blocks the Go port from production on its own.
2. **P1** — silent, hits ordinary Go code, breaks libraries.
3. **P15** — a node-side panic on a case that occurs in normal use; reached by
   both sync APIs and by the node's own background routine.
4. **P18** — a deployment that reports success everywhere and changes nothing
   until the node is restarted. Trivially worked around once known; expensive
   until then, because it looks exactly like a bug that was not fixed.
5. **P2/P3** — not blocking any more, but every MApp author will hit the same
   dead end until the methods are declared on `lapi.LApi`. P3 is a one-line fix.
6. **P4/P5** — two ways a correct app serves nothing.
7. **P6/P7** — stdlib and crypto. P7 (`Ed25519Verify`) is the only capability
   still genuinely absent, and agent signatures go unverified without it.

P1, P2, P3, P4, P5, P6 and P15 have workarounds applied here. **P7 and P17 have
none.** P17 in particular is invisible to the application: the allocation
happens in the node before app code runs.
