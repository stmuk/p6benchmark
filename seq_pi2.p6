my int $n = 100_000; 
my $delta = 1.0 / $n;
my $sum = 0.0;
my int $i = 1;
while ( $i <= $n) {
  my $x = ( $i - 0.5 ) * $delta;
  $sum += 1.0 / ( 1.0 + $x * $x );
  $i += 1;
}
my $pi = 4.0 * $delta * $sum;
say $pi;
