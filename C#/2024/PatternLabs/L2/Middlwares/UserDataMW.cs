using Patterns.ChainOfResp;

namespace Patterns.Middlewares;

public static class UserDataMW
{
    public static void ProcessData(object data)
    {
        var inputDataHandler = new InputDataHandler();
        var mandatoryHandler = new MandatoryHandler();

        inputDataHandler.SetNext(mandatoryHandler);

        inputDataHandler.Handle(data);
    }

}