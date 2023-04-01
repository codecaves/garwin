# garwin
Win32 address resolver based on the 'arwin.c' tool (https://www.vividmachines.com/shellcode/arwin.c) 

## Usage

```bash
PS C:\> garwin32.exe user32.dll MessageBoxA
garwin - win32 address resolver - mortal - v.01
'MessageBoxA' is located at 0x75ddc110 in user32.dll
```

## Building (Windows)

```bash
go build -o bin/garwin32.exe -gccgoflags "-s -w" main.go
```