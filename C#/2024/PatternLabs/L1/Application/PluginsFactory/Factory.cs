using System.Reflection;

namespace Patterns.PluginsFactory;

abstract public class Factory
{
    protected string path;

    public Factory(string path)
    {
        this.path = path;
    }

    abstract public IPlugin Create();

    public IPlugin LoadPlugin(string path)
    {
        Assembly assembly = Assembly.LoadFile(path);
        foreach (Type type in assembly.GetTypes())
        {
            if (typeof(IPlugin).IsAssignableFrom(type) && !type.IsInterface)
            {
                IPlugin plugin = Activator.CreateInstance(type) as IPlugin;
                if (plugin != null)
                {
                    Console.WriteLine("New plugin is created: " + plugin.GetType().Name);
                    return plugin;
                }
            }
        }
        throw new Exception("Can't create plugin from file: " + path);
    }
}