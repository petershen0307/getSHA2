# Code level
## bottom up
1. structure filter
    1. isSkip()
    1. genFilter()
1. calculateSHA2()
1. genWalkCallback() will return filepath.Walk() callback function
1. Start() go through all hard disk

# Library
* use hex to encode binary data to string format

# Output
## Is a map[string][]string, key is hex(hash); value is path list

# Further
The memory usage may not be the limitation. There are about 146,968 files in OS win10 directory "C:\Windows".

If I use 1GB memory, I can store almost 3 million[1GB / (260+64)Byte] records. ps: 260 is the maximum path, 64 is sha256 with hex string format. 260 + 64 Byte is one record maximum memory usage if no duplicate files.