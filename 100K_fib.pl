# A perl program which works under both perl5 (with perl -Mbigint)
# and perl6

my ($prev, $current) = (1,0);

for ( 0..100_000) {
    ($prev, $current) = ($current, $prev + $current);
}
print $current;

