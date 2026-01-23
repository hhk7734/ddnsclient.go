# Guideline

## Test

### Unit test

- Test only public functions (black-box test)
- Name a unit test file `xxx_test.go` for `xxx.go`
- Use `package xxx_test` for `package xxx`

### E2E test

## Commit convention

- Use [Conventional Commits v1.0.0](https://www.conventionalcommits.org/en/v1.0.0/)

### Type

- `<type>(<scope>)!`: Breaking change. (Increments major version, but minor version for `v0.x.x`)
- `feat`: Adds a new feature (Increments minor version)
- `fix`: Fixes a bug (Increments patch version)
- `refactor`: Improves internal structure without functional changes
- `style`: Changes code style only (whitespace, formatting, etc.)
- `test`: Modifies test code only
- `docs`: Updates documentation only
- `chore`: Changes external influences like CI/CD, dependencies, etc.

### Scope

- `workflow`: Changes CI/CD workflow
