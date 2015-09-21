+++
Description = "How to setup Rust racer with multiple or nightly Rust versions"
Tags = ["blog", "personal", "vim", "how-to", "rust"]
date = "2015-09-15T20:43:24+03:00"
title = "Setup Rust Racer with multiple Rust versions"

+++


Guide to setup Rust Racer [code completion utility] with multiple/nightly Rust versions.
This will let you navigate & jump to definition between Rust versions.

### 1) Install Rust version manager

> https://github.com/sdepold/rsvm#installation

### 2) Install desired Rust version

> rsvm install 0.12.0

### 3) Switch/use desired Rust version

> rsvm use 0.12.0

### 4) rsvm - makes symbolic link to current rust version in

> ls -la $HOME/.rsvm/current
```bash
total 56K
drwxrwxr-x  2 user user 4.0K Sep 14 21:31 test
drwxrwxr-x  4 user user 4.0K Sep 15 14:33 1.2.0
drwxrwxr-x  5 user user 4.0K Sep 15 14:48 .
drwxrwxr-x  8 user user 4.0K Sep 21 22:15 .git
drwx------ 24 user user 4.0K Sep 21 22:17 ..
-rw-rw-r--  1 user user 1.1K Sep 14 21:31 LICENSE
-rw-rw-r--  1 user user  324 Sep 14 21:31 CHANGELOG.md
-rw-rw-r--  1 user user  249 Sep 14 21:31 .travis.yml
-rw-rw-r--  1 user user 9.1K Sep 14 21:31 rsvm.sh
-rw-rw-r--  1 user user 2.1K Sep 14 21:31 rsvm.fish
-rw-rw-r--  1 user user 1.3K Sep 14 21:31 install.sh
-rw-rw-r--  1 user user 1.8K Sep 14 21:31 README.md
lrwxrwxrwx  1 user user   20 Sep 15 14:48 current -> /home/user/.rsvm/1.2.0
```

### 5) Export Rust source path to that symbolic link

```
export RUST_SRC_PATH="$HOME/.rsvm/current"
```

### 6) Install Rust racer

> git clone https://github.com/phildawes/racer.git "$HOME/rust-racer"

### 7) Build Rust racer

> cd $HOME/rust-racer; cargo build --release.

### 7) Install vim-racer vim plugin (Emacs, Kate available)

> https://github.com/phildawes/racer#vim-integration

### 8) Define g:racer_cmd in .vimrc

> let g:racer_cmd = "$HOME/rust-racer/target/release/racer"

### 9) Done

> Enjoy multi Rust environment with code completion & goto-definition
