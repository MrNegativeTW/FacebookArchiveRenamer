<div align="center">
  
<img src="https://fakeimg.pl/128x128/" width="128" height="128">

<h1>Facebook Archive Photos Renamer</h1>
<h4>
Rename photos downloaded from Facebook archive to human-readable DateTime format.</h4>

![](https://img.shields.io/badge/Python-3-4b8bb8.svg?style=flat-square)

<p align="center">
  <a href="#Preview">Preview</a> •
  <a href="#features">Features</a> •
  <a href="#how-it-works">How it works?</a> •
  <a href="#how-to-use">How to use?</a> •
  <a href="#license">License</a>
</p>
</div>

## Preview
| Before rename | After Rename |
|---|---|
| ![Before](https://raw.githubusercontent.com/MrNegativeTW/FacebookArchivePhotosRenamer/main/screenshots/before_rename.png) | ![After](https://raw.githubusercontent.com/MrNegativeTW/FacebookArchivePhotosRenamer/main/screenshots/after_rename.png) 

## Features
- Auto backup original photos before rename
- Auto deal with the duplicate file name by adding one second.

## How it works?

Loop through JSON file to get every message, if the message type is a photo, find that photo and rename it with the message timestamp.


## How to use?

### Download your data from Facebook

1. Go to `Settings & privacy` -> `Settings`

![](screenshots/how_0.png)

2. Choose `Your Facebook Information` -> `Download Your Information`

![](screenshots/how_1.png)
![](screenshots/how_2.png)

3. Under `Select file options`, choose `JSON` and `High`.
Under `Select information to download`, choose `Messages`.

![](screenshots/how_3.png)

4. Request a download

![](screenshots/how_4.png)


## License
```
Not decided yet.
```