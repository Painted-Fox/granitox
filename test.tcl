package require tcltest
package require http

tcltest::test mrwerewolf {
    This tests that we can get to the site.
} -body {
    set t [http::geturl http://mrwerewolf.com]
    set body [http::data $t]
    http::cleanup $t

    # See if we contain the word "mrwerewolf"
    string match -nocase *mrwerewolf* $body
} -result 1
