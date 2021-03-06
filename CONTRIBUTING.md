# Contributing to cloud-sdk-go

Contributions are very welcome, these can include documentation, bug reports, issues, feature requests, feature implementations or tutorials.

We have [good first issues](https://github.com/elastic/cloud-sdk-go/labels/good%20first%20issue) for new contributors. These are issues that we've identified as good starting points for getting acquainted with our codebase, [styleguide](https://github.com/elastic/ecctl/blob/master/developer_docs/STYLEGUIDE.md) or working with Go.

- [Reporting Issues](#reporting-issues)
- [Code Contribution Guidelines](#code-contribution-guidelines)
  - [Workflow](#workflow)
  - [Commit Messages](#commit-messages)
  - [Pull Requests](#pull-requests)
- [Setting up a dev environment](#setting-up-a-dev-environment)
  - [Environment prerequisites](#environment-prerequisites)
  - [Dependency handling](#dependency-handling)
  - [Setting up your fork](#setting-up-your-fork)
- [Development](#development)
  - [Generating the client and models from a swagger definition](#generating-the-client-and-models-from-a-swagger-definition)
  - [Running unit tests](#running-unit-tests)

## Reporting Issues

If you have found an issue or defect in `cloud-sdk-go` or the latest documentation, use the GitHub [issue tracker](https://github.com/elastic/cloud-sdk-go/issues) to report the problem. Make sure to follow the template provided for you to provide all the useful details possible.

## Code Contribution Guidelines

For the benefit of all and to maintain consistency, we have come up with some simple guidelines. These will help us collaborate in a more efficient manner.

- Unless the PR is very small (e.g. fixing a typo), please make sure there is an issue for your PR. If not, make an issue with the provided template first, and explain the context for the change you are proposing.

  Your PR will go smoother if the solution is agreed upon before you've spent a lot of time implementing it. We encourage PRs to allow for review and discussion of code changes.

- We encourage PRs to be kept to a single change. Please don't work on several tasks in a single PR if possible.

- When you're ready to create a pull request, please be sure to:
  - Adhere to the project's [style guide](https://github.com/elastic/ecctl/blob/master/developer_docs/STYLEGUIDE.md).

  - Make sure you signed the [Contributor License Agreement](https://www.elastic.co/contributor-agreement/). You only need to sign the CLA once. By signing this agreement, you give us the right to distribute your code without restriction.

  - Have test cases for the new code. If you have questions about how to do this, please ask in your pull request.
  
  - Run `make format`, `make lint` and `make meta-lint`.
  
  - Ensure that [unit tests](#unit) succeed with `make unit`.
  
  - Use the provided PR template, and assign any labels which may fit your PR.
  
  - There is no need to add reviewers, the code owners will be automatically added to your PR.

### Workflow

The codebase is maintained using the "contributor workflow" where everyone without exception contributes patch proposals using "pull requests". This facilitates social contribution, easy testing and peer review.

To contribute a patch, make sure you follow this workflow:

1. Fork repository
2. Enable GitHub actions to run in your fork
3. Create topic branch
4. Commit patches

### Commit Messages

In general commits should be atomic and diffs should be easy to read.

Commit messages should be verbose by default consisting of a short subject line (50 chars max). A blank line and detailed explanatory text as separate paragraph(s), unless the title alone is self-explanatory ("trivial: Fix comment typo in main.go"). Commit messages should be helpful to people reading your code in the future, so explain the reasoning for your decisions. Further explanation here.

If a particular commit references another issue, please add the reference. For example: refs #123 or fixes/closes #1234. Using the fixes or closes keywords will cause the corresponding issue to be closed when the pull request is merged.

Example:

```console
elasticsearch: Remove unnecessary cluster list sorting

Removes the unnecessary sorting on the obtained API response since the
cluster list is already sorted by the backend.

Closes #1234
```

### Pull Requests

*Pull requests that contain changes on the code base **and** related documentation, e.g. for a new feature, shall remain a single, atomic one.*

The title of the pull request should be prefixed by the package or area that the pull request affects. Any change that affects a package in general must be prefixed by its name. Valid areas as:

- _docs_: Changes for documentation related to the project
- _ci_: Changes for the CI build system.
- _scripts_: Changes that affect ancillary scripts or the Makefile
- _trivial_: should only be used for PRs that do not change generated executable code. Notably, refactors (change of function arguments and code reorganization) and changes in behaviour should not be marked as trivial. Examples of trivial PRs are changes to comments or variable names.

Example:

```console
elasticsearch: Remove unnecessary cluster list sorting
docs: Improve the contributor documentation
trivial: Fix a couple of typos on elasticsearch.List
```

If a pull request is not to be considered for merging (yet), please prefix the title with [WIP] or use Tasks Lists in the body of the pull request to indicate tasks are pending.

The body of the pull request should contain a detailed description about what the patch does, along with any justification/reasoning. You should include references to any related discussions (for example other tickets or mailing list discussions).

At this stage one should expect comments and review from other contributors. You can add more commits to your pull request by committing them locally and pushing to your fork until you have satisfied all feedback.

## Setting up a dev environment

### Environment prerequisites

To start off you will need to have **Go 1.11** or higher installed. Make sure you also have the environment variables `$GOPATH` and `$GOBIN` (`$GOPATH/bin`) defined. If you need assistance with this please follow [golangbootcamp guide](http://www.golangbootcamp.com/book/get_setup#cha-get_setup).

### Dependency handling

`cloud-sdk-go` uses go modules to version dependencies, make sure you've got the `GO111MODULE=on` environment variable set to leverage go modules.  Running `go get` in the root folder will download all the required dependencies.

### Setting up your fork

Due to the way Go handles package imports, the best approach for working on a `cloud-sdk-go` fork is to use Git Remotes.  Here's a simple walk-through for getting started:

1. Fetch the sources with `go get -d github.com/elastic/cloud-sdk-go`

2. Change to the source directory: `cd ${GOPATH}/src/github.com/elastic/cloud-sdk-go`

3. Rename the remote to `upstream`: `git remote rename origin upstream`

4. Fork `cloud-sdk-go` in GitHub.

5. Add your fork as a new remote and name it `origin`:

    ```console
    # For SSH based credentials
    git remote add origin git@github.com:USERNAME/cloud-sdk-go.git

    # For https based credentials
    git remote add origin https://github.com/USERNAME/cloud-sdk-go
    ```

6. Verify your Git remotes with `$ git remote -v`. The result should be something like this:

    ```console
    origin    git@github.com:USERNAME/cloud-sdk-go.git (fetch)
    origin    git@github.com:USERNAME/cloud-sdk-go.git (push)
    upstream    git@github.com:elastic/cloud-sdk-go (fetch)
    upstream    git@github.com:elastic/cloud-sdk-go (push)
    ```

7. Create a new branch for your changes (the branch name is arbitrary):

    ```console
    git checkout -b issue1234
    ```

8. After making your changes, commit them to your new branch:

    ```console
    git commit -a -v -s
    ```

9. Push the changes to your new remote:

    ```console
    git push --set-upstream origin issue1234
    ```

10. You're now ready to submit a PR based upon the new branch in your forked repository.

## Development

### Generating the client and models from a swagger definition

To generate / update the client code you will need to follow these steps:

1. Update the `api/apidocs.json` file with the latest swagger specification available from the public API.
2. Run either `make swagger` or `make docker-swagger` to generate the client and models from the swagger definition.
3. Commit the latest changes to a new branch and create a pull request to be reviewed.

### Running unit tests

There's two variables that can be passed to the Makefile target:

- `TEST_UNIT_FLAGS` args controls the flags that are sent to `go test`.
- `TEST_UNIT_PACKAGE` controls which package names to test (defaults to `./...` which means all packages).

The current `TEST_UNIT_FLAGS` default to: `-timeout 10s -p 4 -race -cover`.

#### Go test flags

_For a complete list of all the `go test` flags see [Testing flags](https://golang.org/cmd/go/#hdr-Testing_flags)._

_For a complete list of all the `go build` flags see [Compile packages and dependencies](https://golang.org/cmd/go/#hdr-Compile_packages_and_dependencies)._

| Flag       | Description                                                    |
| ---------- | -------------------------------------------------------------- |
| `-timeout` | controls the max runtime for a set of unit tests in a package. |
| `-p`       | controls the number of package tests to run concurrently.      |
| `-count`   | number of times that the unit tests for a package will be run. |
| `-race`    | enables race condition detection during test execution.        |
| `-cover`   | displays the test coverage for each package.                   |

Examples:

```console
$ make unit
-> Running unit tests for cloud-sdk-go...
...
$ make unit TEST_UNIT_PACKAGE=github.com/elastic/cloud-sdk-go/pkg/api
-> Running unit tests for cloud-sdk-go...
ok      github.com/elastic/cloud-sdk-go/pkg/api 1.277s  coverage: 89.0% of statements
```
