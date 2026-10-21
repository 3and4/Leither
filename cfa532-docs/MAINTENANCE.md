# Documentation ownership and upstream sync

## Ownership

`cfa532-docs/` belongs to this fork. It contains our canonical technical foundation,
translations, navigation and dated observations. Upstream `README*`, `Install.md`,
`api/`, `doc/`, `lapi/`, `rpc/`, `example/` and `www/` remain at their original paths.
Do not move, rename or patch those files merely to organize this collection.

This layout uses ordinary Git history. A merge retains additions that exist only
on the fork. It does not automatically replace this directory from upstream.
If upstream ever adds the same path, resolve the conflict explicitly and preserve
this collection's ownership and evidence. No `merge=ours` attribute, ignored
tracked file, nested repository or generated copy is used as a protection mechanism.

## Safe upstream synchronization

Finish and commit intentional documentation changes before syncing. Keep a clean
working tree and stay on the intended fork branch. Review these commands before
running them; they are the maintenance procedure, not an automated hook:

```bash
git status --short
git branch --show-current
git remote -v
git fetch upstream
git log --oneline HEAD..upstream/main
git diff HEAD...upstream/main -- cfa532-docs/
git merge upstream/main
```

After the merge, inspect `git status`, the changes to upstream documentation and
all conflicts. Verify that this collection and its links still exist. Update the
source index for added or removed source documents, record the upstream revision
reviewed, then commit any reconciliation before pushing to `origin`.

A clean merge is not a technical review: new upstream statements do not invalidate
an older measured observation without a newer versioned check. Keep both dates
and explain a changed result instead of silently rewriting history.

GitHub's ordinary Sync fork/Update branch workflow merges upstream changes into
the fork. Do not select an option that discards fork changes. A force sync,
`git reset --hard upstream/main`, replacing the branch with upstream, or deleting
this directory can remove fork-only work. A path alone cannot prevent those
operations. Preserve normal repository backups; a separate repository or separate
branch would be needed to isolate documentation from a deliberate replacement of
this branch, and is outside this layout.

## Updating the documents

1. Update the English technical reference here; label the version, architecture,
   API contract, observation date and evidence for behavior-specific claims.
2. Distinguish installed help, successful runtime reads, source contracts and
   historical notes. Do not mark an unavailable file download as verified.
3. Keep source links pinned when a statement depends on a particular revision.
   Use local relative links for navigation to current source files.
4. Consolidate Leither/MiMei reference material, runtime evidence and Leither-specific
   integration contracts here. Keep general product designs and release runbooks in
   their application repositories. Repair relative links during cross-repo moves.
5. Label translation coverage explicitly. Old application-repository paths contain
   relocation notices only; preserve earlier bodies through Git history. Maintain
   each document once here, and preserve section fragments in incoming links.
   Follow MIGRATION.md for canonical destinations and source revisions.
6. Refresh the source index when the tracked documentation inventory changes.
7. Review Markdown links, headings, code fences and `git diff --check`. Do not run
   tests unless the user asks. Do not commit or push without user authorization.

The index is a navigation inventory, not a second copy of upstream documents or a
claim that every source file has been verified against the current runtime.
