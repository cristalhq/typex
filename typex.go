package typex

// Ptr of a value.
func Ptr[T any](value T) *T {
	return &value
}

// ZeroOf a type.
func ZeroOf[T any]() T {
	var zero T
	return zero
}

// Ident return the given value.
func Ident[T any](v T) T {
	return v
}

// OnlyNamed fields.
type OnlyNamed struct {
	_ struct{}
}

// Incomparable type. Embed in other structs.
type Incomparable struct {
	_ [0]func()
}

// Do returns value or panics on error.
func Do[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}

// Do2 returns values or panics on error.
func Do2[T, U any](t T, u U, err error) (T, U) {
	if err != nil {
		panic(err)
	}
	return t, u
}

// Do3 returns values or panics on error.
func Do3[T, U, V any](t T, u U, v V, err error) (T, U, V) {
	if err != nil {
		panic(err)
	}
	return t, u, v
}

// Do4 returns values or panics on error.
func Do4[T, U, V, W any](t T, u U, v V, w W, err error) (T, U, V, W) {
	if err != nil {
		panic(err)
	}
	return t, u, v, w
}

// If is just a ternary operator. Don't use it.
func If[T any](c bool, a, b T) T {
	if c {
		return a
	}
	return b
}
