```
┌─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┐
│    ___o .--.               ____                                          ____                                      __                      ___o .--.    │
│   /___| |--|              /\  _`\                                       /\  _`\                                 __/\ \__                  /___| |OO|    │
│  /'   |_|  |_             \ \,\L\_\    ___     __     _ __    __        \ \,\L\_\     __    ___   __  __  _ __ /\_\ \ ,_\  __  __        /'   |_|  |_   │
│       (_    _)             \/_\__ \   /'___\ /'__`\  /\`'__\/'__`\  __o__\/_\__ \   /'__`\ /'___\/\ \/\ \/\`'__\/\ \ \ \/ /\ \/\ \           (_     _)  │
│       | |   \                /\ \L\ \/\ \__//\ \L\.\_\ \ \//\  __/    |    /\ \L\ \/\  __//\ \__/\ \ \_\ \ \ \/ \ \ \ \ \_\ \ \_\ \           | |   \   │
│       | |___/                \ `\____\ \____\ \__/.\_\\ \_\\ \____\  / \   \ `\____\ \____\ \____\\ \____/\ \_\  \ \_\ \__\\/`____ \          | |___/   │
│                               \/_____/\/____/\/__/\/_/ \/_/ \/____/  _______\/_____/\/____/\/____/ \/___/  \/_/   \/_/\/__/ `/___/> \                   │
│                                                                     /\______\                                                  /\___/                   │
│                                                                     \/______/                                                  \/__/                    │
│                                                                                                                                                         │ 
│           Professional Digital forensics, Network hacking, Stegonography, Recon, OSINT, Bluetooth, CAN and Web Exploitation Expert Secruity Team        │
└─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
```

# What is the red rabbit project 

The red rabbit project is an offensive security framework that is like no other, a framework that aims to touch on all fields of cyber security known to the developers and security team, Scare Security has finally brought a whole new relm of possibilities and easy setups for hackers across the world.

# Why choose us

This project is one of the more rare cyber security tools out there, most frameworks you see today like wifite, or even smaller frameworks like brute forcing or even basic tool frameworks just arent enough, sometimes as a developer i wondered why these programs got so much stars and downloads in a small amount of time which led me to make something diffent. While most frameworks again like wifite are frameworks that just simply automate other peoples tools like the aircrack suite since noone really can recreate such legendary tool kits or find ways to advance them. This framework is not like your average framework where it automates someone elses tool/toolkit and calls it a day. Here are some reasons why this project should be used as a framework 

> Lightweight: This program is built using the go programming language with limited third party libraries, while using perl and ruby everything is either pre installed in any average linux system or is kept in a file rather than installed in the root directory so it can be deleted later 

> User Control: Most frameworks have 1-5 tools or maybe are automation scripts that will automate certian processes and toolkits ( As wifite does which makes it a good framework ), however red rabbit allows the user to have control over the output such as using flags like --reso which you can specify if your screen is veritcle or landscape which will control how the output is formatted, and also gives the user complete control over the options in the menu in commands like parse which is a parser to dump certian keys or information in files like Pcap, GIF, PNG etc files 

> Open Source: Alot of the projects the team and i scout online are mostly compiled, non open source, or encoded which can often give the illusion that the tool is faked and does not work or the owner has something to hide, red rabbit leaves its code completely open source for security researchers and future developers to look in to 

> Large selection of utilities: the Red Rabbit project has many utilities that like to give the user multiple options on how the script should run the utility

# What does the red rabbit project have to offer 

This project has alot to offer especially when it comes to offensive cyber security, here are some examples 

> Stegonography: In terms of stegonography you can inject the following file formats with red rabbit 

| Command name  | Description                                                   | 
| ------------- | ------------------------------------------------------------- |  
|  inject gif   | Allows the user to inject GIF images with custom payloads     |
|  inject bmp   | Allows the user to inject BMP images with custom payloads     |
|  inject jpg   | Allows the user to inject JPG images with custom payloads     |
|  inject png   | Allows the user to inject PNG images at chunks with payloads  | 
|  inject webp  | Allows the user to inject WEBP images with custom payloads    |
|  inject zip   | Allows the user to inject JPG images with an archive file     |
|---------------|-------------------------------------------------------------- |

> Web and Graphical user interfaces: Red rabbit uses a mix of a few languages while its main is go it is backed by Perl and Ruby so it was essential to make up for such a noobie mistake to add web and graphical user interfaces. Red rabbit has two graphical user interfaces one for host scanning and the other which is a mix of OUI searching, stenography, and host pingers. While also looking at graphical user interfaces it also has a documentation server which is thrown onto the local host with a flashy lookin web UI ;)

> Tool sets and variety: as said about two times now this project features a wide variety of utilities in topics or fields of hacking like Stegonography, Wifi attacks, Brute forcing, Hash brute forcing, Service fuzzing, Network fuzzing, metadata and file forensics, Offline password attacks, Offline OSINT, Online OSINT, web recon, host discovery, port scanning, system automation, system handeling, and well built commands to view certian things on your system such as ranges of IP's and Interfaces as well as connected addresses.

# Getting more iformation on red rabbits functionality, utilities, and commands

Before i start this section i would like to say something that ALOT of people missed in the pre alpha version, in red rabbit version 1.0 there is a command called `start document server` which means it will start the offical red rabbit web server, once it starts the server you will be brought to this window on `localhost:8080` 

![alt text](git/cipherlogger.png)

you will be asked for the following a username and a password, as a place holder they show `Cipher Username` and `Cipher Password`

which requires the following data 

```
username -> 15 18 19 22 19 7 13 4 12 23 23 22 16 15 26 4 13 26 14 23 26 8 22 19 7
password -> 13 18 14 23 26 9 9
```

YOU CAN NOT HAVE ANY EXTRA SPACES OR THE PROGRAM WILL FAIL AND KICK YOU OUT OF THE SERVER, BE VERY CREFUL WHEN COPYING THIS. This cipher text login is not a standard algorithm and is something that is experimental right now so i figured i should add it into the web UI just to test what people think about it and how the login system should work with it, but its also for security measures. **NUMBERS MUST HAVE SPACES THEY WORK LIKE ARRAYS**

once you are done you will be brought to a web UI which shows all information on the red rabbit project such as mentions, inspirations, reasons to the project, history of the project, code bugs, developer notes and mainly the depths of the commands, how the utilities work and how you would use them since there is so much utilities i had to put it in some nicely formatted wiki and not in something like a readme `given how MASSIVE this framework goes`




