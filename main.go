package main

import "fmt"

// 作为泛型函数的一个示例，MapKeys 采用任何类型的映射并返回其键的切片。
// 这个函数有两个类型参数——K和V；
// K 具有可比较约束，这意味着我们可以使用 == 和 != 运算符比较这种类型的值。这是 Go 中的映射键所必需的。
// V 具有 any 约束，这意味着它不受任何限制（any 是 interface{} 的别名）。
func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

// List 是具有任何类型值的单向链表。
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

// 我们可以像在常规类型上一样在泛型类型上定义方法，但是我们必须保留类型参数。类型是 List[T]，而不是 List。
func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	var m = map[int]string{1: "2", 2: "4", 4: "8"}

	// 在调用泛型函数时，我们通常可以依赖类型推断。请注意，我们不必在调用 MapKeys 时指定 K 和 V 的类型 - 编译器会自动推断它们。
	fmt.Println("keys m:", MapKeys(m))

	// 尽管我们也可以明确指定它们。
	_ = MapKeys[int, string](m)

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.GetAll())
}
