# proctask

A simple task definition and runner

## What is proctask?

You can determine a task using `run`, `env`, `stdin` files and get outputs as `stdout`, `stderr`. `proctask` uses user's login shell, so you can use your own profile, like `rbenv`.

## Usage

Locate `run` in a working directory. Write `env` and `stdin` in the same directory if you need.

    $ ls /path/to/dir
    env    run    stdin

Run `proctask`. If you omit command line arg, `proctask` runs in your current directory.
    
    $ proctask /path/to/dir

Then, you can see `stdout` and `stderr` of the command.

    $ ls /path/to/dir
    env    run    stderr    stdin    stdout

## Example

    $ cat env
    FOO=foo

    $ cat stdin
    aaa

    $ cat run
    #!/usr/bin/env ruby
    p gets
    p ENV["FOO"]
    STDERR.puts "Hi!"

    $ proctask
    2013/08/18 23:01:39 INFO  load env path: /tmp/work/env
    2013/08/18 23:01:40 INFO  run cmd: [/bin/bash -l -c /tmp/work/run], env: [FOO=foo USER=riywo HOME=/Users/riywo], dir: /tmp/work
    hi!
    2013/08/18 23:01:41 INFO  Command succeeded!

    $ cat stdout
    "aaa\n"
    "foo"

    $ cat stderr
    Hi!

## License

MIT

## Author

Ryosuke IWANAGA a.k.a. riywo
