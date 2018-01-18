# cgotest

Sample code to demonstrate an issue I'm seeing with CGO when developing on windows:

Running this should show the following message in the standard output:

"The process cannot access the file because another process has locked a portion of the file."

The origination point for this err is in my tulipindicators lib, found here (debug branch linked to for convenience): 

https://github.com/technicalviking/tulipindicators/tree/windows_error_confusion

in the bridge.go file, on line 78

(https://github.com/technicalviking/tulipindicators/blob/windows_error_confusion/bridge.go#L78)

The package is a binding for the static C lib called "tulipindicators", found here:

https://github.com/TulipCharts/tulipindicators

Full disclosure, I'm using the lib compiled with a fork of the tulip indicators lib, since I modified the makefile to work on windows (I removed the ability to override the flags, I think)  (again, linking to a debug branch for convenience)

https://github.com/technicalviking/tiLibWinFork/tree/windows_error_confusion

printf debugging in the lib shows that the code in the C library is running and returning.  If I comment out the call to the function pointer in the preamble, the error does not occur.
Something happening in C land is making windows do funky things.  I only noticed this because I made the lib super paranoid about whether an exception would be thrown in the C code itself somehow (yes, I know *now* that C doesn't actually have 'exceptions' like higher level languages do).  I'm also wondering if the bug is in what the library is trying to do with the input floats I'm providing, since ignoring the error in the lib results in me being able to see that several output values are being set to NaN when I'm casting them back from C.double to float64.

This is as far as my knowledge takes me. I'm not asking help for debugging my code (that's my problem and I'm already working on other solutions).  What I'm asking is why this issue in C land results in windows not being able to return back to Go without (useful) error messaging.
