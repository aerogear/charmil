# Contributing

## Set up Git Hooks
Run the following command to set up git hooks for the project. 

```
make setup/git/hooks
```

The following git hooks are currently available:
- **pre-commit**:
  - This runs checks to ensure that the staged `.go` files passes formatting and standard checks using gofmt and go vet.
  