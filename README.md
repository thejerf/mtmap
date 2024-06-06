# mtmap - type-safe multi-type maps for Go

This package provides a map-like object for Go that backs to a
`map[any]any`, but unlike direct usage of a `map[any]any`, when
retrieving a value out of the map, the type is statically-known at
compile time and thus does not need to be checked. See the example in
the godoc.

# Pull Requests

I will not be accepting pull requests for this package.

The reason for that is, there are a lot of ways of extending this, but
many of them are mutually exclusive. For instance, in the real code I
have that uses this idea, a stack trace of where all keys are set is
taken, because this is worthwhile and useful for my use, but would
ruin other people's usage.

I also have some code that can penetrate the abstraction for testing
purposes. Adding:

    func NewWith(in map[any]any) *Map {
        if testing.Testing() {
            panic("can't use NewWith in non-testing code")
        }
        return &Map{in}
    }

is useful in my testing, but some users may not appreciate the
abstraction being penetrated, even for testing code. Others may
require it.

This package as specified also only requires Go 1.18. There are many
possible ways to integrate this with range support, with different
performance tradeoffs that one user who only ever stores a few
different keys in this may be fine with, but could be a stopper for
someone storing a lot of keys.

Thus, rather than try to be all things to all people, my suggestion is
that if you like the looks of this package, but need some other
particular feature, just directly put mtmap.go into your codebase
somewhere and modify it into what you want.

If you can use this this directly, by all means feel free to do so. I
will keep it stable. But I would consider people copying & modifying
it once they see the base type tricks used to pull this off the real
primary use of this code.

As such, this is released under the Unlicense license, the closest
thing I can find to just putting this in the public domain. There are
no restrictions on this use, not even an attribution restriction.

