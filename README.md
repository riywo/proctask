# proctask

A simple task definition and runner

## Usage:

Write `proctask.json` in a working directory. Write `stdin` in the same directory if you need.

    $ ls /path/to/dir
    proctask.json    stdin

Run `proctask`. If you omit command line arg, `proctask` runs in your current directory.
    
    $ proctask /path/to/dir

Then, you can see `stdout` and `stderr` of the command.

    $ ls /path/to/dir
    proctask.json    stderr    stdin    stdout

### proctask.json

    {
      "command": ["ruby", "-e", "p ENV"]
    , "env": {
        "FOO": "foo"
      }
    }

## License

MIT

## Author

Ryosuke IWANAGA a.k.a. riywo
