package lox

import (
	"cmp"
	"maps"
	"slices"

	"github.com/samber/lo"
)

// Deprecated: Use a slice literal instead, for example []T{v1, v2}.
func SliceOf[T any](values ...T) []T {
	return values
}

// Deprecated: Use an inline closure instead, for example func(k K) bool { _, ok := m[k]; return ok }.
func MapToPredicate[K comparable, V any](m map[K]V) func(K) bool {
	return func(v K) bool {
		_, ok := m[v]
		return ok
	}
}

// Deprecated: Build a set with lo.Associate and an inline closure instead.
//
//	set := lo.Associate(s, func(v V) (K, struct{}) { return f(v), struct{}{} })
//	pred := func(k K) bool { _, ok := set[k]; return ok }
func SliceToPredicateBy[K comparable, V any](s []V, f func(V) K) func(K) bool {
	return MapToPredicate(SliceToSetBy(s, f))
}

// Deprecated: Use SliceToPredicateBy(s, Identity[V]) or lo.Associate with an inline closure instead.
func SliceToPredicate[V comparable](s []V) func(V) bool {
	return MapToPredicate(SliceToSet(s))
}

// Deprecated: Use lo.Associate(collection, func(item V) (V, struct{}) { return item, struct{}{} }) instead.
func SliceToSet[V comparable](collection []V) map[V]struct{} {
	return SliceToSetBy(collection, Identity[V])
}

// Deprecated: Use lo.Associate(collection, func(item V) (K, struct{}) { return iteratee(item), struct{}{} }) instead.
func SliceToSetBy[K comparable, V any](collection []V, iteratee func(item V) K) map[K]struct{} {
	return lo.Associate(collection, func(item V) (K, struct{}) {
		return iteratee(item), struct{}{}
	})
}

// Deprecated: Use an inline function instead, for example func(v T) T { return v }.
func Identity[T any](v T) T {
	return v
}

// Deprecated: Use lo.Ternary(condition, result, lo.Empty[T]()) instead.
func IfOrEmpty[T any](condition bool, result T) T {
	return lo.Ternary(condition, result, lo.Empty[T]())
}

// Deprecated: Use lo.TernaryF(condition, f, lo.Empty[T]) instead.
func IfOrEmptyF[T any](condition bool, f func() T) T {
	return lo.TernaryF(condition, f, lo.Empty[T])
}

// Deprecated: Use an inline closure instead, for example func(v T1) R { return f(g(v)) }.
func Compose[T1, T2, R any](f func(T2) R, g func(T1) T2) func(T1) R {
	return func(v T1) R {
		return f(g(v))
	}
}

// Deprecated: Use an inline closure instead, for example func(v T) bool { return !f(v) }.
func Not[T any](f func(T) bool) func(T) bool {
	return func(v T) bool {
		return !f(v)
	}
}

// Deprecated: Pass the callback directly to lo.Map or lo.Filter and ignore the index with _.
func IgnoreIndex[T1, T2, R any](f func(T1) R) func(T1, T2) R {
	return IgnoreSecond[T1, T2, R](f)
}

// Deprecated: Pass the callback directly to lo.Map or lo.Filter and ignore the index with _.
func IgnoreSecond[T1, T2, R any](f func(T1) R) func(T1, T2) R {
	return func(v T1, _ T2) R {
		return f(v)
	}
}

// Deprecated: Use slices.SortFunc(s, func(a, b T) int { return cmp.Compare(f(a), f(b)) }) instead.
func SortBy[T any, R cmp.Ordered](s []T, f func(T) R) {
	slices.SortFunc(s, func(a, b T) int {
		return cmp.Compare(f(a), f(b))
	})
}

// Deprecated: Use lo.Filter(collection, func(item T, _ int) bool { return lo.IsEmpty(iteratee(item)) }) instead.
func OnlyEmptyBy[T, R comparable](collection []T, iteratee func(item T) R) []T {
	return FilterWithoutIndex(collection, Compose[T, R, bool](lo.IsEmpty[R], iteratee))
}

// Deprecated: Use lo.Filter(collection, func(item T, _ int) bool { return lo.IsNotEmpty(iteratee(item)) }) instead.
func WithoutEmptyBy[T, R comparable](collection []T, iteratee func(item T) R) []T {
	return FilterWithoutIndex(collection, Compose[T, R, bool](lo.IsNotEmpty[R], iteratee))
}

// Deprecated: Use lo.Filter(collection, func(item V, _ int) bool { return predicate(item) }) instead.
func FilterWithoutIndex[V any](collection []V, predicate func(item V) bool) []V {
	return lo.Filter(collection, IgnoreSecond[V, int, bool](predicate))
}

// Deprecated: Use lo.Map(collection, func(item T, _ int) R { return iteratee(item) }) instead.
func MapWithoutIndex[T, R any](collection []T, iteratee func(item T) R) []R {
	return lo.Map(collection, IgnoreSecond[T, int, R](iteratee))
}

// Deprecated: Use a type assertion instead, for example _, ok := v.(T); return ok.
func InstanceOf[T any](v any) bool {
	switch v.(type) {
	case T:
		return true
	default:
		return false
	}
}

// Deprecated: Use lo.FilterMap(elems, func(item T, _ int) (R, bool) { r, ok := any(item).(R); return r, ok }) instead.
func FilterByType[R, T any](elems []T) []R {
	return lo.FilterMap(elems, func(item T, index int) (R, bool) {
		switch item := any(item).(type) {
		case R:
			return item, true
		default:
			return lo.Empty[R](), false
		}
	})
}

// Deprecated: Use a type switch directly instead.
func IfInstanceOfF[T, R any](v any, ifFunc func(T) R, elseFunc func(any) R) R {
	switch v := any(v).(type) {
	case T:
		return ifFunc(v)
	default:
		return elseFunc(v)
	}
}

// Deprecated: Use lo.Map(slices.Sorted(maps.Keys(m)), func(k K, _ int) lo.Entry[K, V] { return lo.Entry[K, V]{Key: k, Value: m[k]} }) instead.
//
// For lazy iteration, use loi.Map(slices.Values(slices.Sorted(maps.Keys(m))), func(k K) lo.Entry[K, V] { return lo.Entry[K, V]{Key: k, Value: m[k]} }).
func EntriesSortedByKey[K cmp.Ordered, V any](m map[K]V) []lo.Entry[K, V] {
	return lo.Map(slices.Sorted(maps.Keys(m)), func(k K, _ int) lo.Entry[K, V] {
		return lo.Entry[K, V]{Key: k, Value: m[k]}
	})
}

// Deprecated: Use lo.Map(slices.Sorted(maps.Keys(m)), func(k K, _ int) V { return m[k] }) instead.
func ValuesSortedByKey[K cmp.Ordered, V any](m map[K]V) []V {
	sorted := EntriesSortedByKey(m)
	return lo.Map(sorted, func(item lo.Entry[K, V], index int) V {
		return item.Value
	})
}

// Deprecated: Use e.Key directly instead.
func EntryKey[K cmp.Ordered, V any](e lo.Entry[K, V]) K {
	return e.Key
}

// Deprecated: Use a < b directly instead.
func Less[T cmp.Ordered](a, b T) bool {
	return a < b
}

// Deprecated: Use slices.Sorted(maps.Keys(m)) instead.
func KeysSorted[K cmp.Ordered, V any](m map[K]V) []K {
	return slices.Sorted(maps.Keys(m))
}
