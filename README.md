Package errors 

Errors is a simple package the provide a custom error type that satisfies the 
error interface but provides more informations like the stardard log library, 
will print out the line number and file name where the error happened.
The error message string which has been used at creation time it's accessible by
simply calling Msg.

-It provides only one function which is errors.New("error message");
-It also provides errors.New("message",calldepth) calldepth is a non negative interger
 which specifies the calldepth, in the other New function is defaults to 1.
-The errors type provides the errros.Error() which adds line number and file name/path;
-The original error message it stored inside err.Msg

This package has been developed with simplicity in mind and as a drop it replacement
for the standard error type.
