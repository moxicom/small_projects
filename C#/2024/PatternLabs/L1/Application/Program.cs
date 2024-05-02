using Application.PluginsFactory;

Console.WriteLine("context begin\n");

var fA = new FactoryPlugA("/home/ejs-mint/GitHub/small_projects/C#/2024/PatternLabs/L1/PluginA/bin/Debug/net8.0/PluginA.dll");
var fB = new FactoryPlugB("/home/ejs-mint/GitHub/small_projects/C#/2024/PatternLabs/L1/PluginB/bin/Debug/net8.0/PluginB.dll");

var pluginA = fA.Create();
var pluginB = fB.Create();

Console.WriteLine();

pluginA.Initialize();
pluginA.Execute();
pluginA.Terminate();

Console.WriteLine();

pluginB.Initialize();
pluginB.Execute();
pluginB.Terminate();


Console.WriteLine("\ncontext end");
