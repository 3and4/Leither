# Leither Data and Synchronization Contract

**Canonical home:** `cfa532/Leither/cfa532-docs/`
**Consolidated:** 2026-10-16 from [TweetBackendApp/docs/LEITHER_DATA_AND_SYNC_CONTRACT.md](https://github.com/cfa532/TweetBackendApp/blob/77071bcbf61357e8c685ee5c45c5742194a0dde3/docs/LEITHER_DATA_AND_SYNC_CONTRACT.md).

Canonical current contract for the Tweet backend and all three clients. This move changes its location only; all ownership, reference, publication, recovery and synchronization rules are preserved.

---

This document is the canonical cross-project contract for synchronization and object ownership in:

- `TweetBackendApp` — shared backend
- `Tweet-iOS` — iOS client
- `Tweet` — Android client
- `TweetWeb` — Web client

Read this document before changing object creation, references, node routing, synchronization, profile loading, detail loading, comments, replies, or the `get_*`, `resync_user`, and `refresh_tweet` APIs.

## Root and Access Nodes

An object's root node is its authoritative write location. An access node may hold and serve a synchronized copy.

New tweets are published on the author's root node and then propagate to nodes
that provide the author. A P2P network gives no deterministic upper bound for
that propagation, so an access node can already serve the current User while
still lacking one of that User's newly-published Tweets.

Once an access node is announced as a provider of the Tweet itself, Leither is
responsible for bringing that provider copy up to date and keeping it updated
from the root. This is an eventual guarantee rather than an immediate read
barrier: clients must not assume the first read after `MiMeiProvide` is already
current.

## Publication and Provider Roles

Anyone with the required publishing rights can publish or republish a MiMei,
provided the publishing node holds the latest version of its data. Publishing
is not restricted to the original author or the author's root node. This also
applies to File MiMeis.

TweetWeb can therefore be published from either gen8 or minipc. Before publishing,
verify that the chosen node's MiMei data is up to date; a local `last` version
label alone does not establish that it matches the latest version elsewhere.
Bring a stale publishing copy up to date before republishing it.

Publishing increases the MiMei version number. Verify the new version after
publication; do not assume the pre-publication version remains current.

Providers serve copies of MiMei data, like CDN nodes. `MiMeiProvide` announces
that serving role; it does not itself grant publishing rights. A provider may
also republish when it has the required rights and the latest data.

Provider status has stronger synchronization meaning than merely finding the
object through another node. After a node successfully takes the provider role,
Leither owns propagation of subsequent root updates to that copy. Clients may
use explicit recovery APIs when they require an immediate refresh, but must not
continually re-sync a provider as a substitute for Leither's replication.

Do not prescribe `publish -> provide` as a mandatory publication sequence, or
treat an explicit `provide` call as the missing publication step merely because
remote discovery fails. If one node can find the MiMei but another reports
`no provider found`, investigate publication and cross-node discovery; that
observation alone does not establish that the publisher omitted a required
`MiMeiProvide` call.

## Object Ownership and Storage

### User

A user object is authoritative on the user's root node.

### Tweet

A tweet is stored on its own author's root node. When a tweet is created, its author user object must contain a Mimei reference to the tweet:

```text
User --reference--> Tweet
```

The tweet author's identity determines the tweet's storage node.

### Comment and Reply

A comment is also a Tweet object, but its storage rule is different from a top-level tweet.

A comment is stored on the root node of its parent tweet/comment's author, regardless of who wrote the comment. The comment writer remains represented by `comment.authorId` inside the comment data.

When a comment is created, its parent object must contain a Mimei reference to the comment:

```text
Parent Tweet/Comment --reference--> Comment/Reply
```

For `add_comment`, these identities must not be confused:

```text
hostid         = parent author's root host ID
tweetauthorid  = parent author's user ID (legacy name)
comment.authorId = user who wrote the comment
tweetid        = parent tweet/comment ID
```

`tweetauthorid` is a legacy and misleading field name. It means the parent object's author ID, not the comment writer's ID.

Replies follow the same rule: a reply is stored on its parent comment author's root node and is directly referenced by that parent comment.

## One-Level Reference Synchronization

Leither follows Mimei references one level below the synchronized parent. It does not recursively synchronize the full descendant graph.

```text
User
└── Tweet                  included when User is synchronized
    └── Comment            not included by the User synchronization
        └── Reply          not included by the User synchronization
```

Likewise:

- Synchronizing a User includes that User's directly referenced Tweets.
- Synchronizing a Tweet includes that Tweet's directly referenced Comments.
- Synchronizing a Comment includes that Comment's directly referenced Replies.
- Synchronization stops after that one child layer.

The sorted lists used for pagination, such as a tweet's comment list, do not replace Mimei references. Both the list entry and the parent-to-child reference must be maintained.

## Backend Reference Invariants

Successful creation must establish these references before the updated parent is backed up and published:

```text
create tweet:   MMAddRef(user, userId, tweetId)
create comment: MMAddRef(parent, parentId, commentId)
delete tweet:   MMDelRef(user, userId, tweetId)
delete comment: MMDelRef(parent, parentId, commentId)
```

After adding or deleting a reference, back up and publish the parent object so other nodes can observe its current child graph.

If comment creation is delegated to the parent author's remote root node, that node owns creation of the comment and the parent-to-comment reference. Synchronizing the updated parent back to the initiating node should then carry its direct comment children.

## Dual storage formats

New users, tweets and comments/replies use File MiMeis. New message stores
use Database MiMeis.
Existing Database and File records retain their format and MID; existing File
accounts and message histories remain addressable. Existing File node indexes
remain usable, while nodes without one use their node application database.
New tweets use the File adapter with native author-to-tweet and parent-to-comment
references. This policy does not repair or migrate previously published records.

Format dispatch remains per object. Existing File MiMeis use a core JSON file
and per-member collection directories. Either format can reference tweets or
comments of the other format. Directory entries never substitute for native
MiMei references. Password-derived identities and binary package containers keep
their existing File type; they are not new social-record storage.

In the reference invariants above, committing the parent means `MMBackup` for a
database parent or `FilesFlush` followed by `MFSetCid` for a File parent. Never
call `MMBackup` after `MFSetCid`. Publication and one-level synchronization remain
required separately. The target Leither runtime must be checked for visibility
of native references committed with the File root before deployment.

The v2/v3 response envelopes are unchanged. File User/Tweet payloads optionally
report `storageFormat: "tweet-file-v1"`; `health` advertises the formats the server
supports. Old backends remain usable for database data but cannot serve File
objects. Keep File-capable roots and serving/discovery nodes for existing File
records. `health` reports `creationFormat: "mixed"` and a `creationFormats` map:
`user`, `tweet` and `comment` are `tweet-file-v1`; `messages` is `database`.
It continues advertising support for both formats. This policy does not migrate
existing records or add serialization or cross-node transactions.

## Read and Recovery APIs

Normal reads trust the access node:

- `get_user` reads a user copy.
- `get_tweet` reads a tweet/comment copy.
- `get_comments` reads the direct children listed by a tweet/comment.

Explicit synchronization APIs are temporary client-side remedies while Leither synchronization is unreliable:

- `resync_user` synchronizes a User from its root and therefore also its directly referenced Tweets. It does not synchronize the Tweets' Comments.
- `refresh_tweet` synchronizes a Tweet or Comment from its root and therefore also its directly referenced Comments or Replies. It does not recursively synchronize deeper descendants.

After `refresh_tweet`, clients still call `get_comments` to load and render the refreshed direct children.

## Current Client Policy

Routine screen opening should use normal read APIs rather than forcing synchronization.

Deep-link Tweet Detail loading uses a two-stage ordinary-read policy:

1. Resolve the author first with actual `get_user` requests, because a User
   normally has more providers and its provider lookup completes faster.
2. Call `get_tweet` on that proven author access node with
   `fromdetailview="true"`. When the node can serve the Tweet and is not already
   its provider, this asks the node to take the Tweet provider role so Leither
   keeps its copy updated.
3. If that node cannot yet serve the newly-published Tweet, discover providers
   by `tweetId` and retry `get_tweet` there. Tweet provider discovery is the
   slower fallback because the Tweet normally has fewer providers. If those
   reads also fail, stop; do not cycle through more User providers.

Cached author or Tweet data may be displayed immediately, but it never suppresses
this server read sequence.

On iOS and Android, explicit user recovery is attached to pull-to-refresh:

- Profile pull: `resync_user` for the User and direct Tweets.
- Tweet Detail pull: `refresh_tweet`, then `get_comments` for direct Comments.
- Comment Detail pull: `refresh_tweet`, then `get_comments` for direct Replies.

iOS and Android main-feed pull-to-refresh first call `sync_user` for appUser on its access
node when that node differs from appUser's root. This pulls appUser's existing
root state without scanning or explicitly syncing followed Users. The pull
waits for that request, then reloads page zero from the local cache and calls
`get_tweet_feed` as before. It does not require a new-tweet count before syncing.
If the sync fails, the error is logged and the existing reload still proceeds.
The separate session-opening check still runs once per session.

Periodic detail refreshes use ordinary reads and do not force `refresh_tweet`.

TweetWeb currently has no pull-to-refresh interaction. Any Web recovery synchronization must therefore have an explicit, intentional trigger; it must not be accidentally hidden inside a general-purpose cached read helper.

Do not remove the explicit recovery APIs merely because Leither is expected to perform synchronization. They remain necessary until Leither's behavior is verified to be reliable in production. Conversely, do not run these heavier recovery operations on every routine screen opening without a documented need.

## Backend Synchronization Policy

### Feed refactor decision — 2026-10-09

An appUser can follow hundreds of users. Do not turn feed opening or a main-feed
pull into a network check or forced synchronization against every followed root.
Leither provider replication is responsible for keeping followed Users and their
direct Tweet references current. The root still scans local following tweet lists
and checks local provider status; the refactor removes per-following forced sync,
not that local scan. Main-feed pull-to-refresh recovers appUser alone on the access
node, then reads the feed already assembled on its root. It does not discover
posts the root has not yet collected. This deliberately accepts replication delay
in exchange for avoiding network work proportional to the number of followings.

The backend mirrors the client policy above. A node that already provides a
Mimei is kept current by Leither's own replication. The backend keeps no
version or score bookkeeping of its own to detect staleness; the former
`node_update_score`, `node_get_score` and `node_update_mid_by_score` entries
were a workaround for unreliable early replication and have been removed.
`MiMeiIsProvider` — a local table lookup, not a network call — decides whether
provider setup or an explicit recovery pull is needed:

- **Following-feed collection trusts provider replication.** On appUser's root
  node, `update_following_tweets` checks `MiMeiIsProvider` for each followed User
  before reading it. If false, it calls `MiMeiProvide` for that User directly,
  without `MiMeiSync`. It then reads the local tweet list; Leither supplies
  updates to the User and its directly referenced Tweets. This is the behavior
  of `update_following_tweets`. Providing is not an immediate
  freshness barrier, so newly available data may appear on a later feed check.
- **Explicit profile and detail recovery force the sync.** These actions reach
  `refresh_tweet`, `resync_user` and `sync_user`.
  Being a provider promises the copy will catch up, not that it already has.
  On an access node (one that is not the object's root) each of these calls
  `MiMeiSync` directly; on the root node there is nothing to pull. No score
  comparison gates the sync.
- **Access-node feed catch-up still pulls appUser.** The non-root branch of
  `update_following_tweets` calls `MiMeiSync` for appUser to retrieve the feed
  assembled on its root. It does not walk or explicitly sync followed Users.
- **Taking or keeping a copy checks `MiMeiIsProvider` first.** Following an
  account, saving a tweet, quoting a tweet and `mimei_provide` want possession,
  not freshness; replication supplies the rest. A node holding no copy at all is
  not a provider, so the check costs nothing there.

Pulling an object back after a write is not a case the backend has. Clients send
every mutation to the account's root node themselves, so no node forwards a
write and then synchronizes the result back. The delegated-write branches still
present in the legacy `.js` entries are dead paths, not a pattern to copy.

Routine reads still force nothing. `get_tweet` with `fromdetailview="true"`
announces the Tweet only when this node is not already a provider. After that
announcement, Leither—not a client polling loop—is responsible for keeping the
provider copy synchronized; the completion time remains unbounded in the P2P
system.

## Review Checklist

Before changing creation, loading, or synchronization code, verify:

1. Is the object being written to the correct root node?
2. For comments/replies, is routing based on the parent author's root node rather than the comment writer's root node?
3. Does the parent receive the correct child Mimei reference?
4. Is the parent backed up and published after its reference changes?
5. Does deletion remove both the pagination-list entry and Mimei reference?
6. Is the code relying on only one level of synchronization?
7. Does the recovery call synchronize the correct parent level for the missing data?
8. Are normal reads and explicit recovery operations kept distinct?
9. Have impacts been checked across iOS, Android, Web, and backend callers?
10. Does a forced synchronization serve explicit user recovery, rather than work
    `MiMeiIsProvider` shows is unnecessary?
