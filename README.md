# errors 

[![GoDoc](https://godoc.org/github.com/wind85/errors?status.svg)](https://godoc.org/github.com/wind85/errors)
[![Build Status](https://travis-ci.org/wind85/errors.svg?branch=master)](https://travis-ci.org/wind85/errors)
[![Coverage Status](https://coveralls.io/repos/github/wind85/errors/badge.svg?branch=master)](https://coveralls.io/github/wind85/errors?branch=master)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Error package

Errors is a simple package the provide a custom error type that satisfies the 
error interface but improves on the standard error type adding more informations 
like the stardard log library does, it will print out the line number and file 
name where the error happened.
The error message string which has been used at creation time it's accessible by
simply calling Msg.


-It provides only one function which is errors.New("error message"), is also accept
 a secondary integer number which is the calldepth, defaults to 1, if more the 
 two arguments are passed the error string will be "Ops to many paramerts", calldepth
 cannot be negative, if a negative number is passed it will default to 1.
 The original error message it stored inside err.Msg

-It also provides a errors.Wrap function which can be used to wrap a standard error 
 type. It adds the line and file name of where it has been called. Wrap is actually
 a fuction therefore it needs to be called in order to augment the error message.

This package has been developed with simplicity in mind and as a drop it replacement
for the standard error type/package.

## How do I use it?
```
  err := errors.New("error message") // default calldepth 
```
or

```
  err := errors.New("error message", 3) // custom call depth 2
```
Wrap can be call like so:
```
  err := fmt.Errorf("%s\n","error messsage!")
  wrapped := errors.Wrap(err)() // wrapped is augmented at the moment
  wrapped := errors.Wrap(err) // wrapped is not yet augmented
  
```


### Important
-The reported error line is the one in which error.New has been called.

#### Philosophy
This software is developed following the "mantra" keep it simple, stupid or better 
known as KISS. Something so simple like a cache with auto eviction should not required 
over engineered solutions. Though it provides most of the functionality needed by 
generic configuration files, and most important of all meaning full error messages.

#### Disclaimer
This software in alpha quality, don't use it in a production environment, it's not 
even completed.

#### Thank You Notes
None.
