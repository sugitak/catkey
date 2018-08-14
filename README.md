# catkey

The simplest certificate reader ever.

Can you remember what command you should run for your certificate?
`openssl rsa`, `openssl req`, `openssl ec`? And the option, `-noout`, `-text`, and what else?

`catkey` offers you the simplest solution for you... just call it!

# Usage

Just call it.

```
$ catkey whatever.pem
```

or, pipe to it.

```
$ cat whatever.pem | catkey
```

You'll be getting the appropriate `openssl` command and the result.

# Requirements

Catkey uses `openssl` to run the command.
