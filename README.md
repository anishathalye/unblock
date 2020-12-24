# Unblock [![Build Status](https://github.com/anishathalye/unblock/workflows/CI/badge.svg)](https://github.com/anishathalye/unblock/actions?query=workflow%3ACI)


Unblock is a tiny utility to make shell pipes behave as if they have unlimited
buffering.

```
a | unblock | b
```

---

Generally, Unblock's behavior is **not** what you want: you want the standard
behavior where pipes have a fixed size buffer and slow readers exert
backpressure on writers. However, it can sometimes be useful to have an
unlimited buffer in user memory.

Here is one scenario where Unblock might be useful. Suppose you have a program
`slow` that produces 5000 lines of output, but does it slowly, perhaps because
it's doing lots of computation. Because the output doesn't fit on a screen, you
want to use `less` to view the output. But if you do `slow | less`, once the
pipe buffer fills up, `slow` gets blocked. If you scroll down in `less`, you'll
need too wait for `slow` to catch up and produce output. One way you might work
around this is to decouple the two processes and make slow fully materialize
its output into a file, running `slow > out.txt` and viewing the results with a
`less +F out.txt`. Unblock makes this kind of workflow easier: `slow | unblock
| less`.

Unblock buffers its input in memory, buffering as much as necessary, and writes
it out as fast as the reader can accept it. Note that Unblock's buffer is
_unlimited in size_, so if a writer produces a huge amount of output and the
reader is slow, Unblock is going to consume a lot of memory.

## Installation

**Download a binary release:**
[Unblock releases](https://github.com/anishathalye/unblock/releases).

**Install from source with `go get`:**

```bash
go get github.com/anishathalye/unblock
```

## License

Copyright (c) 2020 Anish Athalye (me@anishathalye.com). Released under the MIT
license. See [LICENSE.md](LICENSE.md) for details.
