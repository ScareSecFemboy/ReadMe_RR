

blk =  "\033[0;30m"
red =  "\033[0;31m"
grn = "\033[0;32m"
yel = "\033[0;33m"
blu = "\033[0;34m"
mag = "\033[0;35m"
cyn = "\033[0;36m"
bwht = "\033[0;37m"
bblk = "\033[1;30m"
bred = "\033[1;31m"
bgrn = "\033[1;32m"
byel = "\033[1;33m"
bblu = "\033[1;34m"
bmag = "\033[1;35m"
bcyn = "\033[1;36m"
bwht = "\033[1;37m"
ublk = "\033[4;30m"
ured = "\033[4;31m"
ugrn = "\033[4;32m"
uyel = "\033[4;33m"
ublu = "\033[4;34m"
umag = "\033[4;35m"
ucyn = "\033[4;36m"
uwth = "\033[4;37m"
blkb = "\033[40m"
redb = "\033[41m"
grnb = "\033[42m"
yelb = "\033[43m"
blub = "\033[44m"
magb = "\033[45m"
cynb = "\033[46m"
whtb = "\033[47m"
blkhb = "\033[0;100m"
redhb = "\033[0;101m"
grnhb = "\033[0;102m"
yelhb = "\033[0;103m"
bluhb = "\033[0;104m"
maghb = "\033[0;105m"
cynhb = "\033[0;106m"
whthb = "\033[0;107m"
hblk = "\033[0;90m"
hred = "\033[0;91m"
hgrn = "\033[0;92m"
hyel = "\033[0;93m"
hblu = "\033[0;94m"
hmag = "\033[0;95m"
hcyn = "\033[0;96m"
hwht= "\033[0;97m"
bhblk =  "\033[1;90m"
bhred = "\033[1;91m"
bhgrn = "\033[1;92m"
bhyel = "\033[1;93m"
bhblu = "\033[1;94m"
bhmag = "\033[1;95m"
bhcyn = "\033[1;96m"
bhwht = "\033[1;97m"

module Color_out
    def Color_out.OutC(msg, color, type)
        parsed = "ini #{type}"
        C_Parser.main(parsed, color, msg)
    end
end

class C_Parser

    def self.pr(color, msg)
        puts color, msg
    end
    def self.pp(color, msg)
        pp color, msg
    end

    def self.Parser(color, msg)
        {
            :ini => {
                :format => {
                    'pc': lambda { C_Parser.pr(color, msg)},
                    'pp': lambda { C_Parser.pp(color, msg)}
                }
            },
        }
    end

    def self.main(cmd, color, msg)
        to = cmd.downcase.strip.split ' '
        @ct = to[0].to_sym
        @command = to[1..].join(' ').to_sym
        invoke = C_Parser.Parser(color, msg)[@ct][:format][@command]
        invoke.call
    end
end

# usage
#Color_out.OutC("hello world i am red", bhred, "pc")