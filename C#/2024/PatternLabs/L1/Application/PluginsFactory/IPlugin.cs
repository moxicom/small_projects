namespace Patterns.PluginsFactory;

public interface IPlugin {
    void Initialize();
    void Execute();
    void Terminate();
}
