using Patterns.PluginsFactory;

namespace Application.PluginsFactory;

public class FactoryPlugB : Factory
{
    public FactoryPlugB(string path): base(path){}

    public override IPlugin Create() {
        IPlugin plugin = LoadPlugin(path);
        // ...
        Console.WriteLine("Do some business logic with plugin B");
        // ...
        return plugin;
    }
}
