# Wrapcheck

[![Go Report Card](https://goreportcard.com/badge/github.com/tomarrell/wrapcheck)](https://goreportcard.com/report/github.com/tomarrell/wrapcheck)
[![Tests](https://github.com/tomarrell/wrapcheck/actions/workflows/test.yaml/badge.svg)](https://github.com/tomarrell/wrapcheck/actions/workflows/test.yaml)

A simple Go linter to check that errors from external packages are wrapped
during return to help identify the error source during debugging.

> More detail in [this article](https://blog.tomarrell.com/post/introducing_wrapcheck_linter_for_go)

## Install

Go `>= v1.16`
```bash
$ go install github.com/tomarrell/wrapcheck/v2/cmd/wrapcheck@v2
```

Wrapcheck is also available as part of the golangci-lint meta linter. Docs and
usage instructions are available
[here](https://github.com/golangci/golangci-lint). When used with golangci-lint,
configuration is integrated with the `.golangci.yaml` file.

## Configuration

You can configure wrapcheck by using a `.wrapcheck.yaml` file in either the
local directory, or in your home directory.

```yaml
# An array of strings which specify substrings of signatures to ignore. If this
# set, it will override the default set of ignored signatures. You can find the
# default set at the top of ./wrapcheck/wrapcheck.go.
ignoreSigs:
- .Errorf(
- errors.New(
- errors.Unwrap(
- errors.Join(
- .Wrap(
- .Wrapf(
- .WithMessage(
- .WithMessagef(
- .WithStack(

# An array of strings which specify regular expressions of signatures to ignore.
# This is similar to the ignoreSigs configuration above, but gives slightly more
# flexibility.
ignoreSigRegexps:
- \.New.*Error\(

# An array of glob patterns which, if any match the package of the function
# returning the error, will skip wrapcheck analysis for this error. This is
# useful for broadly ignoring packages and/or subpackages from wrapcheck
# analysis. There are no defaults for this value.
ignorePackageGlobs:
- encoding/*
- github.com/pkg/*

# ignoreInterfaceRegexps defines a list of regular expressions which, if matched
# to a underlying interface name, will ignore unwrapped errors returned from a
# function whose call is defined on the given interface.
ignoreInterfaceRegexps:
- ^(?i)c(?-i)ach(ing|e)
```

## Usage

To lint all the packages in a program:

```bash
$ wrapcheck ./...
```

## Testing

This linter is tested using `analysistest`, you can view all the test cases
under the [testdata](./wrapcheck/testdata) directory.

## TLDR

If you've ever been debugging your Go program, and you've seen an error like
this pop up in your logs.

```log
time="2020-08-04T11:36:27+02:00" level=error error="sql: error no rows"
```

Then you know exactly how painful it can be to hunt down the cause when you have
many methods which looks just like the following:

```go
func (db *DB) getUserByID(userID string) (User, error) {
	sql := `SELECT * FROM user WHERE id = $1;`

	var u User
	if err := db.conn.Get(&u, sql, userID); err != nil {
		return User{}, err // wrapcheck error: error returned from external package is unwrapped
	}

	return u, nil
}

func (db *DB) getItemByID(itemID string) (Item, error) {
	sql := `SELECT * FROM item WHERE id = $1;`

	var i Item
	if err := db.conn.Get(&i, sql, itemID); err != nil {
		return Item{}, err // wrapcheck error: error returned from external package is unwrapped
	}

	return i, nil
}
```

The problem here is that multiple method calls into the `sql` package can return
the same error. Therefore, it helps to establish a trace point at the point
where error handing across package boundaries occurs.

To resolve this, simply wrap the error returned by the `db.Conn.Get()` call.

```go
func (db *DB) getUserByID(userID string) (User, error) {
	sql := `SELECT * FROM user WHERE id = $1;`

	var u User
	if err := db.Conn.Get(&u, sql, userID); err != nil {
		return User{}, fmt.Errorf("failed to get user by ID: %v", err) // No error!
	}

	return u, nil
}

func (db *DB) getItemByID(itemID string) (Item, error) {
	sql := `SELECT * FROM item WHERE id = $1;`

	var i Item
	if err := db.Conn.Get(&i, sql, itemID); err != nil {
		return Item{}, fmt.Errorf("failed to get item by ID: %v", err) // No error!
	}

	return i, nil
}
```

Now, your logs will be more descriptive, and allow you to easily locate the
source of your errors.

```log
time="2020-08-04T11:36:27+02:00" level=error error="failed to get user by ID: sql: error no rows"
```

A further step would be to enforce adding stack traces to your errors instead
using
[`errors.WithStack()`](https://pkg.go.dev/github.com/pkg/errors?tab=doc#WithStack)
however, enforcing this is out of scope for this linter for now.

## Why?

Errors in Go are simple values. They contain no more information about than the
minimum to satisfy the interface:

```go
type Error interface {
  Error() string
}
```

This is a fantastic feature, but can also be a limitation. Specifically when you
are attempting to identify the source of an error in your program.

As of Go 1.13, error wrapping using `fmt.Errorf(...)` is the recommend way to
compose errors in Go in order to add additional information.

Errors generated by your own code are usually predictable. However, when you
have a few frequently used libraries (think `sqlx` for example), you may run
into the dilemma of identifying exactly where in your program these errors are
caused.

In other words, you want a call stack.

This is especially apparent if you are a diligent Gopher and always hand your
errors back up the call stack, logging at the top level.

So how can we solve this?

## Solution

Wrapping errors at the call site.

When we call into external libraries which may return an error, we can wrap the
error to add additional information about the call site.

e.g.

```go
...

func (db *DB) createUser(name, email, city string) error {
  sql := `INSERT INTO customer (name, email, city) VALUES ($1, $2, $3);`

  if _, err := tx.Exec(sql, name, email, city); err != nil {
    // %v verb preferred to prevent error becoming part of external API
    return fmt.Errorf("failed to insert user: %v", err)
  }

  return nil
}

...
```

This solution allows you to add context which will be handed to the caller,
making identifying the source easier during debugging.

## Contributing

As with most static analysis tools, this linter will likely miss some obscure
cases. If you come across a case which you think should be covered and isn't,
please file an issue including a minimum reproducible example of the case.

## License

This project is licensed under the MIT license. See the [LICENSE](./LICENSE) file for more
details.

