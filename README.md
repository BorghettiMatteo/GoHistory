# GoHistory

## What is this?

This is a simple go program that uses `golang.design/x/clipboard` to handle xserver clipboard section and, via `clipGui.sh` and [zenity](https://help.gnome.org/users/zenity/stable/) bash utils to visualize it

## How it works?

Once executed, go will keep watching the clipboard section of xserver and writes it to a dump file, encoded in such a way that the bash script can handle whenever is called.

## requirements

This project requires 
* `zenity` (as stated up above)
* `xorg/dev`, installable with your favourite package manager!    

## How to use it

1. Clone this repo in a place were you want your files to be dumped
2. Use your system keyboard-shortcut creator to map your favourite shortcut (for me is *super + v*) and map `bash path/to/clipGui.sh`

## Configuration File

*goHistory* uses a config.xml file that handle all the needed parameter to run smoothily.

### Why XML?

You may as, *why the hell using xml as config file format?*, well I lost a bet to a friend of mine, so here we are.

This is the template of the config file (as of today)

```xml
<Configuration>
    <clipGui></clipGui>
    <DumpFilePath></DumpFilePath>
    <BufferLenght></BufferLenght>
    <BackUpFrequency></BackUpFrequency>
    <BackUpStrategy></BackUpStrategy>
</Configuration>
```

1. **ClipGui** is the path to `clipGui.sh` script
2. **DumpFIlePath** is the path to the file where the clipboard history will be saved
3. **BufferLenght** is the lenght of the *in-memory* clipboard buffer
4. **BackUpFrequency** is a cron expression for scheduling the backups
5. **BackUpStrategy** as of today, `aws` and `cron` are the only one supported


## Running goClipboard

to run just spawn a shell and run `./goClipboardExecutable path/to/the/conifig.xml/file`


## Future

This is still a projecet in its early stages, so more to come, such as:
* implementing backup on S3 bucket
* resting history from backup
* improving gui, deprecating the bash script
* better handling of concurrent access to the dump files 

Open a discussion and let's talk!