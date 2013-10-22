Granitox
================================================================================

Granitox aims to be a rock solid acceptance testing framework.  Granitox is
currently under development, but it's main focus is for testing web
applications.

Goals:

- Acceptance testing.
- Organize tests into suites and groups.
- Reporting.
- Data templates and data driven testing.
- Easy http faking.
- Enable interface testing with Selenium Web Driver (or equivalent).

Granitox is built on Tcl to help accomplish these goals...

- Tcl makes it easy to build a DSL for acceptance tests.
- Tcl supports coroutines which makes faking easy to implement.
- Tcl is simple and sane (no surprises in syntax).
- Tcl is LISP-like.  Data looks like code and code looks like data.
- Tcl runs everywhere, which means fewer limitations on deployment.
- Tcl lets you build subinterpreters with their own environment.
