# convert-to-webm
## CLI video converter

Converts .mov and .mp4 to .webm
Uses ffmpeg for convertions

To install in MacOS
```
brew install ffmpeg
```

Usage:
```
convert-to-webm <filename>[.mov | .mp4]
```
```
convert-to-webm [-h | --help]
```
```
convert-to-webm <path/to/><filename>[.mov | .mp4]
convert-to-webm <filename>[.mov | .mp4]
convert-to-webm <path/to/>*
convert-to-webm *

-h | --help : prints this help
*           : converts all [.mov | .mp4] files to .webm
```