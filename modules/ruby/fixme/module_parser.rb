#...........................................................................................
#
# This module is used to run ruby like commands, why is this in the fixme file path? 
# this module needs to be worked on and is most likely in this path for bugs or things to 
# do at the end of the RR6 development as ruby will be the last language to work with for this 
# project, note alot of this might be changed to perl just ruby for now since alot of the 
# wifi like tools ex DEAUTh, FAKE AP, etc are all ruby based, since perl does not have any 
# libraries to construct your own custom fake access point, for not these use the
# fake ap, ftp, net, and color modules
#
# ~ ArkAngeL43
#
#
#
#
#
#
#
#..........................................................................................

require_relative 'modules/ruby/fake-ap'
require_relative 'modules/ruby/ftp'
require_relative 'modules/ruby/net'
require_relative 'modules/ruby/color'
#require_relative 'modules/ruby/shod'

command  = ARGV[0] 
command2 = ARGV[1] 


class Commands
    # call fake ap module
    def Commands.Fakeap(ssidname, inter)
        FakeAP.reciever(inter, ssidname, true)
    end

    def self.honorable_mentions
        puts "BUSHIDO LOL"
    end

    def self.history
        puts "There is some history I guess"
    end

    def self.art
        puts ":) - wow look at this art"
    end

    def self.install
        puts "Installing shit..."
    end

    def self.script_help
        puts "You need help with this?"
    end

    def self.CMDS
        {
            :ini => {
                :exec => {
                    'opt1': lambda { Commands.script_help}
                }
            },
        }
    end

    def self.show_invalid_command(cmd, type=false)
        if type
            puts "[ - ] Invalid command for #{cmd.split(' ')[0].downcase}: '#{cmd.split(' ')[1..].join ' '}'"
        else
            puts "[ - ] Invalid command: '#{cmd}'"
        end
    end

    def self.parse_command(cmd)
        to = cmd.downcase.strip.split ' '
        @ct = to[0].to_sym
        @command = to[1..].join(' ').to_sym

        if not Commands.CMDS.keys.include? @ct
            Commands.show_invalid_command cmd
        elsif not Commands.CMDS[@ct][:exec].include? @command
            Commands.show_invalid_command(cmd, type=true)
        else
            invoke = Commands.CMDS[@ct][:exec][@command]
            invoke.call
        end
    end
end


parseda = "#{command} #{command2}"
Commands.parse_command(parseda)



