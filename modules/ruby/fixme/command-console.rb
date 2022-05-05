#...................................................................................
# This help module is still in development, while the code is fine and is being
# worked on, the output of some files using ASCII designed will just quit with 
# expected EOF errors, like text/help/advancedcm.txt:13: syntax error, unexpected 
# constant, expecting end-of-input ";    \     _,-'    L   ''--._  ..."
# so this needs to be fixed, might port into perl but for now this should stay in the 
# path of -> modules/ruby/fixme 
#
# ~ ArkAngeL43
#
#
#
#
#
#...................................................................................

require_relative 'modules/ruby/color'
require_relative 'modules/ruby/file_parser'
 

# file number to output
$select_file = ARGV[0]  || 'text/help/advancedcm.txt'.empty?
classc = ARGV[1] || '$help_txt'.empty?

$help_json = [
    "config/verified.json", 
    "config/sets.json", 
    "config/flags.json"
]
$help_txt = [
    "text/help/advancedcm.txt",
    "text/help/help_general.txt",
    "text/help/"
]

# if file is json, then use 1
if classc == "1" 
    $filename = $help_json.select{|file| file == $select_file }.join
end
# if name is txt use 2
# ruby command-console.rb
if classc == "2"
    $filename = $help_txt.select{|file| file == $select_file }.join
end

class Help
    def self.Outputfiles(filename, color)
        output   = File.open(filename)
        contents = output.read
        puts contents
    end

    def self.FormOutputFiles(fnum)
        if fnum == 1
            Help.Outputfiles($filename, "\033[31m") 
        end

    end
        
    # output all filepaths
    def self.all_help_files
        File_parser.LoadFiles(help_txt)
    end
end


Help.Outputfiles($filename, "\033[38m")