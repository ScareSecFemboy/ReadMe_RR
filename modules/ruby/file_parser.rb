require_relative 'color'

class File_parser
    def File_parser.RunOutput(filename, color) 
        output   = File.open(filename)
        contents = output.read
        puts color, contents
    end
    def File_parser.SelectOutputFiles(filenum)
    end
    def File_parser.LoadFiles(array) 
        for i in array do
            puts i
            File_parser.OutputFiles
        end
    end
end
