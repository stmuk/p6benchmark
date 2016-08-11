# ported from mandel.pl (from Rosetta code)
# this in turn was based on mandel.rb
# see 3rdparty
#
use v6;

sub mandelbrot($c) {
    my $z = $c;
    for (1 .. 20) {
        $z = $z * $z + $c;
        return $_ if $z.re.abs > 2;
    }
}

loop (my $y = 1; $y >= -1; $y -= 0.05) {
    loop (my $x = -2; $x <= 0.5; $x += 0.0315)
    {print mandelbrot($x + $y * i) ?? ' ' !! '#'}
    print "\n"
}
