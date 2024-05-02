using Patterns.Middlewares;

Console.WriteLine("Context begin\n");

var userdata = "A9Hk29JKhasb2SmkA82";

UserDataMW.ProcessData(userdata);

Console.WriteLine("\nContext end");

