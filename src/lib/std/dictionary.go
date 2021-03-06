package std

import "github.com/raviqqe/tisp/src/lib/core"

// Dictionary creates a new dictionary from pairs of a key and value.
var Dictionary = core.NewLazyFunction(
	core.NewSignature(
		nil, nil, "keyValuePairs",
		nil, nil, "",
	),
	func(ts ...*core.Thunk) core.Value {
		d := core.EmptyDictionary
		l := ts[0]

		for {
			v := core.PApp(core.Equal, l, core.EmptyList).Eval()
			b, ok := v.(core.BoolType)

			if !ok {
				return core.NotBoolError(v)
			} else if b {
				break
			}

			k := core.PApp(core.First, l)
			l = core.PApp(core.Rest, l)
			val := core.PApp(core.First, l)
			l = core.PApp(core.Rest, l)

			d = core.PApp(core.Insert, d, k, val)
		}

		return d
	})
