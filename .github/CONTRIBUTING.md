# Contributing

When contributing to this repository, please first discuss the change you wish to make via our developer, by creating an issue or any other method with the owners of this repository before making a change.

Please note: we have a [code of conduct](https://github.com/thnkrn/comet/blob/master/.github/CODE_OF_CONDUCT.md), please follow it in all your interactions with the `Comet` project.

## Commit rules

### Conventional Commits

* We are using Conventional Commits rule to add readable meaning to commit messages
* We are following the [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) rule

#### Format for commit

[type][optional scope]: [optional REFERENCE-1234] [description]

* ex. build(husky): [BO-000] add husky and commitlint

* List of commit type
  [
  'build',
  'ci',
  'chore',
  'docs',
  'feat',
  'fix',
  'perf',
  'refactor',
  'revert',
  'style',
  'test'
  ]

## Commit types

| Commit Type | Title                    | Description                                                                                                 |
| ----------- | ------------------------ | ----------------------------------------------------------------------------------------------------------- |
| `feat`      | Features                 | A new feature                                                                                               |
| `fix`       | Bug Fixes                | A bug Fix                                                                                                   |
| `docs`      | Documentation            | Documentation only changes                                                                                  |
| `style`     | Styles                   | Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)      |
| `refactor`  | Code Refactoring         | A code change that neither fixes a bug nor adds a feature                                                   |
| `perf`      | Performance Improvements | A code change that improves performance                                                                     |
| `test`      | Tests                    | Adding missing tests or correcting existing tests                                                           |
| `build`     | Builds                   | Changes that affect the build system or external dependencies (example scopes: gulp, broccoli, npm)         |
| `ci`        | Continuous Integrations  | Changes to our CI configuration files and scripts (example scopes: Travis, Circle, BrowserStack, SauceLabs) |
| `chore`     | Chores                   | Other changes that don't modify src or test files                                                           |
| `revert`    | Reverts                  | Reverts a previous commit                                                                                   |

## Contributing to `comet`

To contribute to `comet`, follow these steps:'

1. Clone this repository.
2. Create a feature branch: `git checkout -b <branch_name>`.
3. Make your changes and commit them: `git commit -m '<commit_message>'`
4. Push to the original branch: `git push origin <branch_name>`
5. Create the pull request against `master`.
