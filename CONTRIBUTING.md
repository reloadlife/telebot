# Code Style

## Golang

- files names should be `snake_case.go` and `snake_case_test.go` (files without test files will get **AUTO REJECTED**)
- telegram types should be `PascalCase` (e.g. `Message`, `User`, `Chat` ...) and have their own seperated file, e.g. `message.go`,
  and with `message.go` there should be `message_test.go`, to ensure JSON.Unmarshal works correctly, with any input from telegram.
- telegram types should have `json:"..."` tags, to ensure JSON.Unmarshal works correctly, with any input from telegram.
- every telegram type should have a `func (t *Type) String() string` method, to ensure `fmt.Println` works correctly.
- Anything that is exposed to user should be a private *pointer, and be exposed using a public `Interface` (
  e.g. `type Message interface { ... }`).

### file names logic

- `${snake_case}.go`
- `${snake_case}_test.go`
- if it's a telegram type: `type_${snake_case}.go`
- if it's a type related telegram function: `func_${snake_case}.go`

_**ANYTHING THAT DOES NOT FOLLOW THIS LOGIC WILL BE AUTO REJECTED**_

## Issues and Project Tasks / Commit messages

- if your PR/commit is related to an issue, please mention it in the commit message, e.g. `fix: #123`
- please, please, please, use [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) for your commit messages,
  don't just commit 'fix' or 'added X', write a proper, detailed commit message.
- if your code does not contain documents (godoc, comments, etc), in your PR/commit message, write `documentation_required` or your
  commit will be rejected by the bot.

### tests

- write tests for everything, use Telegram's official API to test the code, in the test environment (on github actions), you have access to
  `${secrets.TELEBOT_SECRET} (a bot token related to [@TeleBotUnitTestBot](https://t.me/TeleBotUnitTestBot))`, use it to test your code.
- if your code does not contain tests, in your PR/commit message, write `tests_required`, so someone else can write the tests for you, or
  your commit will be rejected by the bot. (be sure to create the `...test.go` file, or your commit will be rejected, even if you
  provide `tests_required` in your commit message)
- WATCH OUT FOR RACE CONDITIONS, use `go test -race` to test your code, if it fails, fix it, or your commit will be rejected by the bot.

