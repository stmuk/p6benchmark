#!/usr/bin/env perl;
package Foo;
use Moose;

has i => ( is=> 'rw');

for my $i (1..1_000_000) {
    my $obj = Foo->new;
    $obj->i($i);
}


1;

