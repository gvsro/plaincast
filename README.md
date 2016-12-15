# Plaincast

This is a small [DIAL](http://www.dial-multiscreen.org) server that emulates Chromecast-like devices, and implements the YouTube app. It only renders the audio, not the video, so it is very lightweight and can run headless.

It can be used as media server, for example on the [Raspberry Pi](http://www.raspberrypi.org/).

## Installation

I'm going to assume you're running Linux for this installation guide, preferably Debian Jessie (or newer when their time comes). Debian before Jessie contains too old versions of certain packages.

First, make sure you have the needed dependencies installed:

 *  golang 1.3 (1.1+ might also work, but 1.0 certainly doesn't)
 *  libmpv-dev
 *  youtube-dl (see also 'notes on youtube-dl' below)

These can be installed in one go under Debian Jessie:

    $ sudo apt-get install golang libmpv-dev youtube-dl

If you haven't already set up a Go workspace, create one now. Some people like to set it to their home directory, but you can also set it to a separate directory. In any case, set the environment variable `$GOROOT` to this path:

    $ mkdir golang
    $ cd golang
    $ export GOPATH="`pwd`"

Then get the required packages and compile:

    $ go get -u github.com/aykevl/plaincast

To run the server, run the executable `bin/plaincast` relative to your Go workspace. Any Android phone with YouTube app (or possibly iPhone, but I haven't tested) on the same network should recognize the server and it should be possible to play the audio of videos on it. The Chrome extension doesn't yet work.

    $ bin/plaincast

## Notes on youtube-dl

youtube-dl is often too old to be used for downloading YouTube streams. You can try to run `youtube-dl -U`, but it may say that it won't update because it has been installed via a package manager. To fix this, uninstall youtube-dl, and install it via pip. On Debian, this is as easy as running:

    $ sudo apt-get remove youtube-dl
    $ sudo apt-get install python-pip
    $ sudo pip install youtube-dl

Afterwards, you can update youtube-dl using:

    $ sudo pip install --upgrade youtube-dl

It is advisable to run this regularly as it has to keep up with YouTube updates. Certainly first try updating youtube-dl when plaincast stops working.

## Thanks

I would like to thank the creators of [leapcast](https://github.com/dz0ny/leapcast). Leapcast is a Chromecast emulator, which was essential in the process of reverse-engineering the YouTube protocol and better understanding the DIAL protocol.
