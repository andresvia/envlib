envlib
======

A go package to manipulate the environment with a simple approach.

This is a package that helps to edit environment files like those present on many Linux OSes (like `/etc/environment`) or the environment passed to a Docker container with the flag `-env-file <file>`. Environment files look like this:

```
VARIABLE1=value1
VARIABLE2=value2
```

This package simplify editing such files.

This package comes from the necesity of quickly editing environment files when changing environments (for example, setting http_proxy on and off, and is modeled arround this particular use case)

The `Init` function can take any standard `Environ` as such for example `os.Environ()` or you can choose to build your own `Environ`.

Then you can choose to `Select` which variables to keep, or to `MergeWith` a standard environment file, and finally maybe `Persist` to store the resulting environment file.

Check the directory `examples` for some use cases. You can run all the examples with:

```
cd examples
go generate -v ./...
```

Also check the tests for some other use cases.
