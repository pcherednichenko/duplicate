## Duplicate files

This fast utility allows you to find and delete the same files (but with different names) in the folder

Files will be checked by hash sum

### Download

[Mac OS](download_mac/duplicate)

[Linux](download_linux/duplicate)

[Windows](download_windows/duplicate.exe)


## How to use it?

Put `duplicate` file in the folder with the files that you want to check and open it

Or set folder path using argument `-path=folder`

### Example of run on Mac OS

In current folder:
```bash
./duplicate
```

In `/folder/folder`
```bash
./duplicate -path=/folder/folder
```

Then application ask for remove files:
```
Found 10 duplicate files, delete copies? (originals will not be deleted) [y/N]:
```
Write
```bash
y [enter]
```

and files will be deleted

Delete without confirmation with flag: `noconfirm`
```bash
./duplicate -noconfirm -path=/folder/folder
```
