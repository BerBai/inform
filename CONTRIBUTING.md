# Contributing to Notify

We want to make contributing to this project as easy and transparent as possible.

# Semantic Commit Messages

See how a minor change to your commit message style can make you a better programmer.

Format: `<type>(<scope>): <subject>`

`<scope>` is optional

## Example

```shell
feat: add hat wobble
^--^  ^------------^
|     |
|     +-> Summary in present tense.
|
+-------> Type: chore, docs, feat, fix, refactor, style, or test.
```

**More Examples:**

- feat: (new feature for the user, not a new feature for build script)
- fix: (bug fix for the user, not a fix to a build script)
- docs: (changes to the documentation)
- style: (formatting, missing semicolons, etc.; no production code change)
- refactor: (refactoring production code, e.g. renaming a variable)
- test: (adding missing tests, refactoring tests; no production code change)
- chore: (updating grunt tasks etc.; no production code change)


# Tests

Ideally, unit tests should accompany every newly introduced exported function. We're always striving to increase the project's test coverage.

If you need to mock an interface in your tests, follow the next steps:

1. Comment the interface you'd like to mock following the example below

```go
//go:generate mockery --name=nameOfClient --output=. --case=underscore --inpackage
type nameOfClient interface {
    ...
}

```
> Remember to set the --name argument accordingly. For real-life implementation examples, check out existing services, for example fcm.

2. Run make mock 运行 make mock

> The first time you'll also need to run make setup to download the packages required to generate your mocks

3. Use the mocked interface in your tests

```go
mockClient := newMockNameOfClient(t)
```

# Commits

Commit messages should be well formatted, and to make that "standardized", we are using Conventional Commits.

You can follow the documentation on [their website](https://www.conventionalcommits.org/).

# Pull Requests

We actively welcome your pull requests.

1. Fork the repo and create your branch from main.
2. If you've added code that should be tested, add tests.
3. If you've changed or added exported functions or types, document them.
4. We use gofumpt to format our code. Don't forget to always run make fmt before opening a new PR.
5. Ensure the test suite passes and the linter doesn't complain (make ci).

# Issues

We use GitHub issues to track public bugs. Please ensure your description is clear and has sufficient instructions to be able to reproduce the issue.

# License
By contributing to notify, you agree that your contributions will be licensed under the LICENSE file in the root directory of this source tree.
