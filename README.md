Package errors 

Errors is a simple package the provide a custom error type that satisfies the 
error interface but improves on the standard error type adding more informations 
like the stardard log library does, will print out the line number and file name 
where the error happened.
The error message string which has been used at creation time it's accessible by
simply calling Msg.


-It provides only one function which is errors.New("error message"), is also accept
 a secondary integer number which is the calldepth, defaults to 1, if more the 
 two arguments are passed the error string will be "Ops to many paramerts", calldepth
 cannot be negative, if a negative number is passed it will default to 1.

-The errors type provides the errros.Error() which adds line number and file name/path;

-The original error message it stored inside err.Msg

-It also provides a errors.Wrap function which can be used to wrap a standard error type.
 It adds the line and file name of where it has been called.

This package has been developed with simplicity in mind and as a drop it replacement
for the standard error type/package.

Caveats:
-The reported error line is the one in which error.New has been called
