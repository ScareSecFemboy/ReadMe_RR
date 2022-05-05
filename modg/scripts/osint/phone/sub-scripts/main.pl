use warnings;
use strict;

my $phone_num = shift;

system("./modg/scripts/osint/phone/sub-scripts/sub.sh --num $phone_num");