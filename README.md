# GomusicX
A terminal music player written in go language
---
### iOS
Oto requires these frameworks:
AVFoundation.framework
AudioToolbox.framework
Add them to "Linked Frameworks and Libraries" on your Xcode project.
### Linux
ALSA is required. On Ubuntu or Debian, run this command:

`apt install libasound2-dev`
### On RedHat-based linux distributions, run:
`dnf install alsa-lib-devel`

In most cases this command must be run by root user or through sudo command.
### FreeBSD, OpenBSD
BSD systems are not tested well. If ALSA works, Oto should work.
# usage method
    Press the K key to play music, and press the Q key to exit the program