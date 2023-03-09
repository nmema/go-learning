If you run this, the output will be:

```text
{"results":["banana","orange","apple","grapes"],"total_results":4}
```

But if you change `&fruits` for `fruits` in this [line](https://github.com/nmema/go-learning/blob/main/adhoc/custom-json/main.go#L29), the output will be:

```text
{"results":["banana","orange","apple","grapes"]}
```
The reason is that `json.Marshal` cannot obtain a pointer to a field when the struct is not passed as a pointer. Since we've only implemented `MarshalJSON` for a `*Fruits` and not a `Fruits`, `json.Marshal` has to fall back on its generic marshaling for the `Fruits` fields.

Resources:
- [MarshalJSON works for pointer struct but not for non-pointer struct](https://groups.google.com/g/golang-nuts/c/gW53LC_ZC98)
- [encoding/json: bad encoding of field with MarshalJSON method](https://github.com/golang/go/issues/22967)