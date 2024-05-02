using Patterns.PluginsFactory;

namespace PluginB;

public class PluginB : IPlugin
{
    public void Initialize()
    {
        Console.WriteLine("Plugin B is initialized");
    }

    public void Execute()
    {
        Console.WriteLine("Plugin B is executing");
        Console.WriteLine("Plugin B is doing something while executing...");
    }

    public void Terminate()
    {
        Console.WriteLine("Plugin B is terminated");
    }
}
