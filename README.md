# README

## manga-desktop

manga-desktop is desktop application for managing and downloading manga from server for read at later time
offline without connection.

This application is named as mangav4 and build with wails and vue and for now it only support for windows.

## Live Development

To develop this app there are some requirement and setup some development environment to do.

### Requirement

- Windows 10/11
- Golang latest
- pnpm latest
- git
- mingw-w64

### Setup

- Install wails cli
  ```
  go install github.com/wailsapp/wails/v2/cmd/wails@latest
  ```

- Clone the repo
  ```
  git clone https://github.com/twincats/manga-desktop.git
  ```

- Setting Path Environment variable of dependency dll libvips
  ```
  extract libvips.7z in folder and setting path Environment variable to that folder
  libvips.7z is located in REPO\build\windows\libvips.7z
  ```

- run dev in terminal
  ```
  wails dev
  ```

## Building

To build a redistributable, production mode package, use `wails build`.
