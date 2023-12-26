# joorvajoor

## Introduction

The webpage in question serves as a remote interface for movie playback control.
Users have the ability to initiate play, pause, and other related commands directly from this web interface.
When a command is issued, it is immediately dispatched to a Raspberry Pi through a secure and reliable web service.
The Raspberry Pi, acting as a middleman, then translates these commands into appropriate signals which are delivered to the movie playback host via a TCP (Transmission Control Protocol) socket connection. This bidirectional communication ensures seamless control of the movie experience from the convenience of the user’s browser.
