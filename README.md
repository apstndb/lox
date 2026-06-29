# lox

This module is **retired**.

It was a thin wrapper around [`github.com/samber/lo`](https://github.com/samber/lo) and Go 1.23 stdlib helpers (`slices`, `maps`, `cmp`). Downstream repositories now use those libraries directly, or keep domain-specific helpers in their own `internal` packages.

All exported APIs are marked deprecated. Do not add new dependents.

## Migration

| lox API | Replacement |
| --- | --- |
| `IfOrEmpty` / `IfOrEmptyF` | `lo.Ternary` / `lo.TernaryF` + `lo.Empty` |
| `MapWithoutIndex` / `FilterWithoutIndex` | `lo.Map` / `lo.Filter` with `_` index |
| `EntriesSortedByKey` / `KeysSorted` | `slices.Sorted(maps.Keys(m))` and map lookups |
| `InstanceOf` | type switch or `_, ok := v.(T)` |
| `Identity` | inline `func(v T) T { return v }` or direct comparison |
| Predicate helpers (`Compose`, `Not`, `SliceToPredicateBy`, …) | move to caller `internal` package if still useful |

## Former dependents

- `spanner-mycli`, `spanner-mcp`, `spannerplanviz` — migrated to `lo` / stdlib
- `spanneropttools` — helpers live in `internal/lox`
- `spanvalue-backup` — unmaintained snapshot; not updated
