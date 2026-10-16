# Leither documentation maintained by cfa532

This is the canonical home for our Leither technical reference and the entry point
for documentation throughout this fork. The collection is maintained independently
of `3and4/Leither`, in the fork-owned `cfa532-docs/` namespace. Upstream files stay
at their original paths so their links, examples and future merges continue to work.

## Read first

- [Leither and MiMei Technical Foundation](reference/LEITHER_AND_MIMEI.md): storage,
  versions, commit and pin cost, references, identity, MApps, browser bootstrap,
  versioned runtime evidence, and the V0.24.63 CLI/IPFS inspection findings.
- [中文技术参考](reference/zh-CN/LEITHER_AND_MIMEI.md): the existing Chinese edition;
  it is a partial translation and may lag the English reference.
- [Complete source index](UPSTREAM_INDEX.md): all existing non-hidden Markdown,
  HTML and text documentation, examples' READMEs, website drafts and historical
  pages, grouped by purpose. It includes fork additions as well as inherited files.
- [Maintenance and upstream sync](MAINTENANCE.md): ownership, evidence, safe merges
  and the limits of protection from a destructive reset.

## Frequently used source documents

| Topic | Source |
|---|---|
| Overview and setup | [中文 README](../README.md), [English README](../README_EN.md) |
| API map | [API index](../api/Api.md), [MiMei operations](../api/MiMei.md) |
| Key authentication | [Supported contract](../doc/KEY_AUTH_GUIDE.md), [login example](../example/key-auth-login/README.md) |
| Concepts | [MiMei](../doc/MiMei-en.md), [directory model](../doc/Directory-en.md), [PKI](../doc/Pki-en.md) |
| Application hosting | [Application APIs](../api/App.md), [deployment](../api/LApp.md) |
| Typed RPC surface | [Go binding](../lapi/README.md), [RPC](../rpc/README.md), [examples](../example/rpc/) |
| Node operations | [macOS autostart](../leither-autostart-macos-guide.md), [macOS TUS setup](../doc/tus-server-macos-setup.md) |
| Historical context | [Install notes and changelog](../Install.md); its listed version is historical |

## Consolidated evidence and integration contracts

| Document | Scope and status |
|---|---|
| [Runtime issues](reference/RUNTIME_ISSUES.md) | Versioned historical Leither interpreter, API and operational evidence from TweetBackendApp |
| [Tweet data and synchronization contract](integrations/tweet/DATA_AND_SYNC_CONTRACT.md) | Canonical current cross-client ownership, references, routing and recovery invariants |
| [TweetWeb connection management](integrations/tweet/WEB_CONNECTION.md) | Historical server-side connection integration notes; examples are not newly revalidated |
| [LifeDrive key-auth review (中文)](integrations/lifedrive/KEY_AUTH_REVIEW.zh-CN.md) | Historical native-key design and V0.24.12 observations; not the current household auth design |
| [Migration record](MIGRATION.md) | Source revisions and the single maintained destination for every moved document |

The English and Chinese technical foundations above are maintained here as well.
Old document paths in LifeAlbum, TweetBackendApp and TweetWeb contain only
relocation notices. Repository indexes, section links and agent instructions point
to the central documents; there are no full archived copies to update separately.

## Ownership and scope

Maintain new explanations, corrections and versioned observations here. Source
Markdown and runnable examples outside this directory remain inputs to review;
they are not automatically the contract of the installed binary. Prefer installed
help and verified API behavior when a source document disagrees with the runtime.

Leither-specific integration contracts and evidence are consolidated here, grouped
by application. General product architecture, application implementation and release
runbooks stay with their applications. Links back to implementation context preserve
provenance and require the reader's existing access to private repositories.

Normal merge-based fork syncs preserve committed fork-only additions. This directory
is not immune to a force reset or a future conflict at the same path. Follow the
maintenance guide rather than assuming that its name or an ignore rule protects it.

The first consolidation is dated 2026-10-16. The source baseline and full document
inventory are recorded in the source index. New document links become available on GitHub
when the consolidation changes are committed and pushed to this fork.
