# convert-to-webm
## CLI video converter

Converts .mov and .mp4 to .webm
Uses ffmpeg for file convertions

To install ffmpeg in MacOS
```
brew install ffmpeg
```

Usage:
```
convert-to-webm [-h | --help]
```
```
convert-to-webm [-c] <filename1>[.mov|.mp4] <filename2>[.mov|.mp4] ...
```
```
convert-to-webm [-c] <path/to/>*
```
```
  -h | --help : prints this help
   *          : converts all [.mov | .mp4] files to .webm
```

# Installation
* Download repo
* Install Go >= 1.21
* In the repo directory run:
```
go build .
```
this will create a convert-to-webm binary (excecutable) file
* Create bin/ directory at $HOME:
```
cd $HOME
mkdir bin
```
* cd to repo directory
* within your convert-to-webm folder you'll find convert-to-webm binary
* copy or move convert-to-webm to $HOME/bin/
```
mv convert-to-webm $HOME/bin
```
* Add $HOME/bin to your PATH in .zshrc or .bashrc
  * NOTE: make a backup of your .zshrc or .bashrc file BEFORE running the next command
```
cd $HOME
cp .zshrc .zshrc.orig # making a backup .zshrc file
echo 'export PATH="$HOME/bin:$PATH"' >> .zshrc
```

Change to your Videos directory and run