using Patterns.PluginsFactory;

namespace PluginA;

public class PluginA : IPlugin
{
    public void Initialize()
    {
        Console.WriteLine("Plugin A is initialized");
    }

    public void Execute()
    {
        Console.WriteLine("Plugin A is executing");
    }

    public void Terminate()
    {
        Console.WriteLine("Plugin A is terminated");
    }
}
