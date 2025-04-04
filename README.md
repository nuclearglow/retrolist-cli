# Retro CLI

## Planned Features

- support local todo.json files
- sync with remote Retrolist

## Development Setup

- Install Node 22 LTS

### Arch Linux Setup

- Dependencies and DevTools:

```shell
sudo pacman -S go pkg-config pre-commit
```

- Debugging Setup

```shell
 echo 0 | sudo tee /proc/sys/kernel/yama/ptrace_scope
```

- Linter Setup

```shell
yay -S golangci-lint
npm install -g @commitlint/cli @commitlint/config-conventional
go install golang.org/x/tools/cmd/goimports@latest
go install -v github.com/go-delve/delve/cmd/dlv@latest
go install golang.org/x/tools/gopls@latest
pre-commit install
```

### Ubuntu Linux Setup

- [Install latest Go on Ubuntu](https://ubuntuhandbook.org/index.php/2024/02/how-to-install-go-golang-1-22-in-ubuntu-22-04/)

Dependencies and DevTools:

```shell
sudo apt install libc6-dev pkg-config pre-commit
```

- Debugging Setup

```shell
 echo 0 | sudo tee /proc/sys/kernel/yama/ptrace_scope
```

- Linter Setup

```shell
go install -v github.com/golangci/golangci-lint/cmd/golangci-lint@latest
go install golang.org/x/tools/cmd/goimports@latest
go install -v github.com/go-delve/delve/cmd/dlv@latest
go install golang.org/x/tools/gopls@latest
npm install -g @commitlint/cli @commitlint/config-conventional
pre-commit install
```

### Makefile Setup

- Install Rust

```shell
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
```

- Install crates

```shell
cargo install watchexec-cli hyperfine
```
