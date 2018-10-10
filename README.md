errors
======

This package provides helpers to deal with error handling in go programs

## Features

### Error Slice

```go
errs := errors.NewSlice()
errs = errs.Add(errors.New("err1"))
errs = errs.Add(errors.New("err2"))
errs = errors.MergeSlice(errs, errors.New("err3"), errors.Slice{errors.New("err4)})

```

### Error Map

```go
errs := errors.NewMap()
errs = errs.Add("field1", errors.New("err1"))
errs = errs.Add("field2", errors.New("err2"))
errs = errors.MergeMap(errs, errors.Map{"field3": errors.New("err3)})
```

### Wrap, Wrapf, WithStack

There are wrappers for the most used functions from the github.com/pkg/errors package

### Marshaling according to PROD-03

`MarshalToJSON`, `MarshalToYAML` and `MarshalToObject` are supported

```go
errSlice := errors.Slice{errors.New("err1"),errors.New("err2")}
errMap := errors.Map{"f1": errors.New("err3")}
buf := &bytes.Buffer{}
errors.MarshalToJSON(buf, errSlice, errMap, errors.New("err4"))
fmt.Print(buf.String())
// prints: { "errors": ["err1", "err2", "err4"], "fieldErrors": {"f1": "err3"} }
```
