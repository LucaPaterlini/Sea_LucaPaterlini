# SEA

This service goal is to get the records of an indefinite size
csv, parse and uniform each record and send each to a server.


In this test I have used:
- proto3 as syntax for the protocol of communication via gRPC
- I have handled the kill and interrupt signal opening a channel
  and using a select with a case handling the receive of a new signal
  in both the for loops of the client and the server
- I have used Mongodb as dataase and mgo as ODM (Object Data Mapper)
  
   
