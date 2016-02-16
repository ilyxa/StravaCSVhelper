go build && strip StravaCSVhelper && mv StravaCSVhelper StravaCSVhelper-`uname -s`-`uname -m` && tar cf - StravaCSVhelper-SunOS-i86pc | bzip2 -c > StravaCSVhelper-`uname -s`-`uname -m`.tar.bz2
