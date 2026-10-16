# Leither/MiMei documentation consolidation

**Date:** 2026-10-16

Each document has one maintained body in this directory. The application repositories
retain short relocation notices at old paths. Their Git history preserves the original
bodies; there is no parallel archived copy to maintain. Upstream-owned files remain
at their original paths so normal fork merges remain independent of this collection.

## Canonical document map

| Previous location | Canonical document |
|---|---|
| `LifeAlbum/docs/LEITHER_AND_MIMEI.md` | [reference/LEITHER_AND_MIMEI.md](reference/LEITHER_AND_MIMEI.md) |
| `LifeAlbum/docs/zh-CN/LEITHER_AND_MIMEI.md` | [reference/zh-CN/LEITHER_AND_MIMEI.md](reference/zh-CN/LEITHER_AND_MIMEI.md) |
| `TweetBackendApp/LEITHER_ISSUES.md` | [reference/RUNTIME_ISSUES.md](reference/RUNTIME_ISSUES.md) |
| `TweetBackendApp/docs/LEITHER_DATA_AND_SYNC_CONTRACT.md` | [integrations/tweet/DATA_AND_SYNC_CONTRACT.md](integrations/tweet/DATA_AND_SYNC_CONTRACT.md) |
| `TweetWeb/docs/LEITHER_CONNECTION.md` | [integrations/tweet/WEB_CONNECTION.md](integrations/tweet/WEB_CONNECTION.md) |
| `LifeAlbum/docs/zh-CN/LIFEDRIVE_LEITHER_AUTH_REVIEW.md` | [integrations/lifedrive/KEY_AUTH_REVIEW.zh-CN.md](integrations/lifedrive/KEY_AUTH_REVIEW.zh-CN.md) |

## Source revisions before consolidation

The foundations were first imported from LifeAlbum in commit `c150a55` of this
fork. This pass removes the redundant LifeAlbum archive bodies and imports the
remaining runtime evidence and integration documents from these source revisions:

- `LifeAlbum`: `ade508c37bec3cd2d772e00f28d004eef5cdfde4`.
- `TweetBackendApp`: `77071bcbf61357e8c685ee5c45c5742194a0dde3`.
- `TweetWeb`: `6246c715de6de873630d2902a96862972d2cd26a`.

## Consolidation boundaries

- The English foundation is the general narrative; the Chinese edition is a partial
  translation. Versioned runtime reports retain their original findings and later
  corrections, with links to the narrative instead of merging incompatible dates.
- The Tweet synchronization contract remains normative for all four Tweet projects.
  Its rules are unchanged. Their agent instructions now read the central file.
- The TweetWeb connection guide and LifeDrive key-auth review retain their original
  examples and validation claims, explicitly marked as historical integration notes.
- The former document locations are migration notices, not alternate contracts.
  Incoming Markdown links are retargeted with section fragments preserved.
- Product designs, source code, user-facing setup content and release runbooks stay
  in their applications. The inventory of upstream sources is unchanged.

Publish the central Leither changes before publishing application commits that link
to the new paths. Until then, use the sibling Leither checkout for local access.
