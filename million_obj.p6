#!/usr/bin/env perl6;
use v6;

class Foo { has $.i is rw};

for (1..1_000_000) -> $i {
    my $obj = Foo.new;
    $obj.i = $i;
}
