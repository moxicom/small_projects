using Patterns.PluginsFactory;

namespace Application.PluginsFactory;

public class FactoryPlugA : Factory
{
    public FactoryPlugA(string path) : base(path) { }

    public override IPlugin Create()
    {
        IPlugin plugin = LoadPlugin(path);
        // ...
        Console.WriteLine("Do some business logic with plugin A");
        // ...
        return plugin;
    }
}
