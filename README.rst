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

- Tcl is a rapid scripting language.
- Tcl makes it easy to build a DSL for acceptance tests.
- Tcl supports coroutines which makes http faking easy and straight-forward to implement.
- Tcl is consise and sane (no surprises in syntax).
- Tcl is LISP-like.  Data looks like code and code looks like data.
- Tcl runs everywhere, which means fewer limitations on deployment.
- Tcl lets you build subinterpreters with their own environment.
- Tcl is a string based language and ultimately web data is made of strings.  Tcl excels at this.
- For some reason, practice in Tcl fosters simple, easy to use, interfaces and APIs, in a way that other languages don't seem to be able to.
