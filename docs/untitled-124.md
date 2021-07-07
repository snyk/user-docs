# What are docker scratch based images?

##  What are docker scratch based images?

When building Docker containers you define your base image in your dockerfile. The scratch image is the smallest possible image for docker. Actually, by itself it is empty \(in that it doesn’t contain any folders or files\) and is the starting point for building out images.

In order to run binary files on a scratch image, your executables need to be statically compiled and self-contained. This means there is no compiler in the image so you’re left with just system calls.

You also wouldn’t be able to login to the container either as there isn’t a shell unless you explicitly add one.

