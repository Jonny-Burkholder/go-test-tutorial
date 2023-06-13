# go-test-tutorial

Go tests are flexible, easy to write, and easy to use! The built-in testing library is all that is needed for most tests. The aim of this tutorial is to demistify go testing and provide practical, well documented examples all in one repository.

## Go Testing conventions

### Naming Test Functions

Go tests are nothing scary - they're really just functions with a special signature. They must begin with the word `Test`, and take a single argument of `t *testing.T`. Notice that the word "Test" is capitalized, exporting the function. If "test" is written in lowercase, the function will not be called when the tests are run. Traditionally, the test is named after the function you are testing. So if you have a function named "Foo" (as you'll see we do in this repo), Your test function would look something like this: `func TestFoo(t *testing.T){}`. The test function has no return value.

### Where to keep tests

Where you put your go tests is up to you, and will depend in part on whether you are engaging in whitebox or blackbox testing, which we will cover in more depth later in this tutorial. If you are doing whitebox testing (that is, testing unexported functions), your tests will need to be in the same package as the functions being tested. If you are only testing exported functions, the tests may go anywhere, although conventionally they will go in a subdirectory in the package you are testing.

### Naming and Organizing Test Files

Go tests should be kept in a file whose name ends in `_test.go`. Traditionally, this will match either the package name, or the resource you're testing. So if you are testing a package called `users`, convention would be to name your test file `users_test.go`. If you are testing a larger package, it may be a good idea to break up tests into separate files. In this case I recommend naming the files after the resource being tested. Let's imagine you have a package called `structures` where you are building out your own implementations of various data structures and their implementations. It might be too much to put every test in the same file - instead, you can make a folder within your structures package named `test`, which will also be the name of your testing package. Then within the test package, you can create a file for each resource being tested. Say you're working on linked lists and red-black trees. One file could be named `linked_list_test.go`, and the other `rb_tree_test.go`, both in `package test`. Take a look at the file structure in this repo to get a visual of what this structure looks like.

If you're concerned about package collision here, don't be! In Go, the package path matches the directory structure, so `internal/models/test` is a completely separate package from `internal/controllers/test`.

## Basic Testing

Now that we understand the basics of tests, let's look at creating test functions and running a test suite.

### Writing test functions

The conventional way to write tests for go is to provide input and an expected result, then run a given function to make sure its result matches what you expected. Say you have a simple function that take an integer input and doubles the amount:

```
func Doubler(amt int) int {
    return amt*2
}
```

To test this, we'll write a simple function that calls `Doubler`, and check if it gives us the expected result. When writing this function, we'll keep in mind to follow Go testing conventions.

```
func TestDoubler(t *testing.T){
    input := 2
    expected := 4

    actual := Doubler(input)
    if actual != expected{
        t.Errorf("expected %d, got %d", expected, actual)
    }
}
```

`Errorf` is a method on type `*testing.T` that throws an error when called. This method takes formatting arguments, so it is easy to compare the expected result with the actual result in our error message. Because we have written our test correctly, if our Doubler function works as expected, no error will be thrown, and the test output will read `PASS`. Otherwise, we will receive an error telling us what went wrong.

### Running a specific test file

Let's run the tests in the repository. Navigate to the root of this project and run the following command:

`go test`

Not the result you expected?

The go test command by default looks for test files in the current directory. Since there are no test files in the root, we will give the command the path to the file we want to run. Enter this command in the terminal:

`go test -v ./internal/example`

Much better, right?

You'll see from the output that we had one test run, and pass. Great! But if you look in the file `internal/example/example_test.go`, you'll see that there are two functions. Why did only one run? It's because only one followed the correct convention of capitalizing `Test` in the name.

### Running all tests in a project

If you have multiple test files across multiple packages, it's tedious to specifiy each test file that you want to run. Thankfully, it's easy to run all the tests in a project from the project root. Run the following command:

`go test ./...`

Congratulations! You've successfully run all the tests in the project in one command.

### Run a single test

If you have tests that take a long time to run, or a large suite of tests, you may not wish to run all of them. In this case we can run a single test function. We can do this using the `-run` flag. Enter this command:

`go test ./internal/example -v -run TestUnexportedFunction`

This tells go to run only the function `TestUnexportedFunction` from our test file.

### Test flags

You may have noticed that we added the `-v` flag to our command. This stand for "verbose" and gives us more detailed information about the tests that ran. This is useful for things like benchmarking and debugging.

Another useful flag is `-cover`. Running this will tell you how many of your functions are "covered", or tested, by your tests. Run this in the terminal:
`go test -cover ./internal/example`

You should see from the result that we have covered 10% of the functions in the package - not great coverage. Don't worry, we've covered a lot more in our `test` package.

Let's see what our coverage is with both our whitebox and blackbox tests being run. Execute the command from earlier to run all the test files: `go test -cover ./...`

Notice anything curious about the ouptut? We still have only 10% covered, and our file `blackbox_test.go` outputs the message `no statements`. What's that about? Simply put, the Go command is only measuring coverage tests achieve for functions in the same package. Since our blackbox tests are in their own package named "test" (more on that later), their coverage is not measured against package "example". We can adjust this using the `-coverpkg` flag, which tells go which packages we want to measure against. Run the following command:

`go test ./... -cover -coverpkg=./internal/example`

Much better, right? Per the go 1.10 release notes, the flag `coverpkg` takes a list of comma separated values. If we wanted to measure the coverage of two specific packages, we could list them comma-separated, like so: `-coverpkg=./internal/example,./internal/second_example` (this is for demonstration purposes only, please don't name your packages this way). Another handy way to tell go to measure coverage for *all* the packages in your project is to use `-coverpkg=./...`, which will recursively look for every package under the current directory. Finally, using `coverpkg=all` will measure coverage for all packages in the project as well as *all of their dependencies*. My advice is to use this sparingly.

Another flag, `coverprofile`, generates a file which can be used for more advanced coverage analysis. We'll look more at this later.

## Advanced and Blackbox Testing

### Whitebox vs Blackbox testing

What we've seen so far is called whitebox testing - we're testing will all of the unexported (or private) resources of the repository available to us. This can be useful for getting extra coverage, but it can also hide flaws in the design if functions rely too heavily on unexported components in the package. The solution to this is blackbox testing, where we test only the API that is available outside of the package. Look in the file `internal/example/test/blackbox_test.go` to see our blackbox tests. Here we exlusively test exported functions in the package to make sure they are working as expected.

### Table Driven Tests

If you look at the first test in the file, `TestExportedFunction`

### Fuzz Testing

### Benchmarks

### Documentation Examples
