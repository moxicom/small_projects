namespace Patterns.ChainOfResp;

public abstract class AbstractHandler : IHandler
{
    private IHandler? _nextHandler;

    public IHandler SetNext(IHandler handler)
    {
        this._nextHandler = handler;
        return handler;
    }

    public virtual object? Handle(object request)
    {
        if (this._nextHandler != null)
        {
            return this._nextHandler.Handle(request);
        }
        else
        {
            return null;
        }
    }
}