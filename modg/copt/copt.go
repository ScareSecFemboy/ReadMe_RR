/*

Module:  Copt.go
Package: Copt

What does this do:
	This module is the initalizer of all flags

Note:
	Structures must be capitalized

	Go doesn’t have any public,  private or protected keyword. The only mechanism to control the visibility outside the package is using the capitalized and non-capitalized formats

Capitalized Identifiers are exported. The capital letter indicates that
this is an exported identifier and is available outside the package.
Non-capitalized identifiers are not exported. The lowercase indicates
that the identifier is not exported and will only be accessed from
within the same package.go any struct which starts with a capital
letter is exported to other packages. Similarly, any struct field
which starts with capital is exported otherwise not. And also
similarly any struct method which starts with a capital letter is
exported. Let’s see an example that shows exporting and non-exporting
of structs, struct fields, and methods.

*/

package copt

import (
	"github.com/spf13/pflag"
)

var (
	flags = pflag.FlagSet{SortFlags: false}
)

type RR6_options struct {
	S_M              bool
	Screen_rotation  string
	Input            string
	Output           string
	Suppress         bool
	Image_offset     string
	Inject           bool
	Payload          string
	Type             string
	Payload_Encode   bool
	Payload_Decode   bool
	Key              string
	Extract_ZIP      bool
	INJECT_ZIP       bool
	Filepath_general string
	Hexdump          bool
	Geo              bool
	Walk             bool
	Walkerfp         bool
	Discover         bool
	Rr6_help_screen  bool
	Rr6_help_flags   bool
	Jpgchunk         string
	Pheight          string
	Pwidth           string
	Sniffc           string
	Iprange          string
	Target_mac       string
	Target_spoof     string
	Gateway_mac      string
	Brute_list       string
	Hashlist         string
	Userlist         string
	PayloadList      string
	Packet_t         string
	Per_mode         string
	Url              string
	Workers          int
	Pass_length      int
}
