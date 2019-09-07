PAM Talk
========

Material for the PAM talk held at the [Nairobi LUG](https://groups.google.com/forum/#!forum/nairobi-gnu) meetup on 7th September 2019.

You can also view the talk online by going to https://talks.godoc.org/github.com/kmwenja/pam-talk/talk.slide.

Run the presentation locally:
-----------------------------

1. Setup a Go environment: https://golang.org/doc/install
2. `go get -u golang.org/x/tools/cmd/present`
3. In this directory, run `present -orighost localhost -notes`
4. Visit http://localhost:3999 in your browser.

[Present Docs](https://godoc.org/golang.org/x/tools/cmd/present)

Run the demo programs:
----------------------

1. Cd into `module` in this repo.
2. Run `make`. This will build a PAM module (shared library) at `/tmp/go-pam.so`.
3. Add `/tmp/go-pam.so` to any PAM application config in `/etc/pam.d/` that you can safely test with (good examples can be `su` or `sshd`). Add the module in the `auth` section with a `sufficient` control level. An example configuration looks as follows:

```
# add this line to the top of the auth section
auth    sufficient  /tmp/go-pam.so
# other auth lines will go here
....
```

4. Make a user called `test` but don't assign them any working password: `useradd -m test`.
5. Use the application to authenticate. For example, if you are using `su` to demo, run `su test` and notice how no password is required to login to the `test` user.
6. Remove the line added to the PAM config for the app you are using to demo since having this line remain there will be a security risk to your system.


References:
-----------

- man pam (conf, functions, etc)
- https://fedetask.com/write-linux-pam-module/
- https://www.linux.com/news/understanding-pam/
- https://github.com/AmandaCameron/golang-pam-auth
