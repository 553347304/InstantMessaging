package method

import (
	"fim_server/utils/stores/conv"
	"fim_server/utils/stores/logs"
	"cmp"
	"fmt"
	"regexp"
	"sort"
)

type listServerInterface[T cmp.Ordered] interface {
	In(T) int           // 是否在列表中    返回 index    -1 = false
	InRegex(T) bool     // 是否在列表中  支持正则
	Sort(bool) []T      // 排序 升序 true 降序 false
	Unique() []T        // 去重
	Intersect([]T) []T  // 交集
	Difference([]T) []T // 差集
	Delete(int) []T     // 删除下标	index 索引
}
type listServer[T cmp.Ordered] struct {
	Slice []T
}

//goland:noinspection GoExportedFuncWithUnexportedType	忽略警告
func List[T cmp.Ordered](slice []T) listServerInterface[T] {
	return &listServer[T]{Slice: slice}
}

func (l *listServer[T]) In(key T) int {
	for index, s := range l.Slice {
		if s == key {
			return index
		}
	}
	return -1
}
func (l *listServer[T]) InRegex(key T) bool {
	for _, s := range l.Slice {
		regex, err := regexp.Compile(fmt.Sprint(s))
		if err != nil {
			logs.Info("正则表达式编译失败:", err)
		}
		if regex.MatchString(fmt.Sprint(key)) {
			return true
		}
	}
	return false
}
func (l *listServer[T]) Sort(asc bool) []T {
	sort.Slice(l.Slice, func(i, j int) bool {
		if asc {
			return l.Slice[i] < l.Slice[j] // 升序  asc
		}
		return l.Slice[i] > l.Slice[j] // 降序 desc
	})
	return l.Slice
}
func (l *listServer[T]) Unique() []T {
	s := make([]T, 0)
	m := make(map[T]bool)
	for _, v := range l.Slice {
		if !m[v] {
			m[v] = true
			s = append(s, v)
		}
	}
	return s
}
func (l *listServer[T]) Intersect(slice []T) []T {
	s := make([]T, 0)
	m := conv.SliceMap(l.Slice)
	for _, v := range slice {
		if m[v] {
			s = append(s, v)
		}
	}
	return s
}
func (l *listServer[T]) Difference(slice []T) []T {
	s := make([]T, 0)
	m := conv.SliceMap(slice)
	for _, v := range l.Slice {
		if !m[v] {
			s = append(s, v)
		}
	}
	return s
}
func (l *listServer[T]) Delete(index int) []T {
	l.Slice = append(l.Slice[:index], l.Slice[index+1:]...)
	return l.Slice
}
