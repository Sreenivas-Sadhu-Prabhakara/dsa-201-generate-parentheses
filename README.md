# Generate Parentheses

> NeetCode 150 · Stack — backtracking on two counters: open while any remain, close while it wouldn't overtake the opens; invalid strings are never even started. Go, stdlib only, tested.

## Problem

Produce all well-formed arrangements of n bracket pairs — 5 strings for n=3, and in general the Catalan number C(n), which the n=4 test (14) pins.

## Approach

Generate-and-filter wastes effort on the exponentially many invalid strings. The two legality rules — never more than n opens; never more closes than opens — can be enforced *during* construction, so the recursion only ever extends prefixes that can still become valid: branch on '(' while `open < n`, on ')' while `close < open`. The shared byte buffer grows and shrinks with the recursion (append, recurse, truncate), with `string(buf)` copying only at completed leaves. The output arrives in lexicographic-by-construction order because '(' is always tried first.

**Complexity:** O(Catalan(n) · n) — the output's own size; O(n) depth.

## Run

    go test ./...   # verify
    go run .        # demo

---

Part of a 275-repo DSA series in Go — Striver's SDE Sheet (dsa-001–191) plus the NeetCode 150 / Blind 75 additions (dsa-192–275), stdlib only, every solution tested. Browse: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
