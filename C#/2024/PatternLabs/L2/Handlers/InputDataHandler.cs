namespace Patterns.ChainOfResp;

class InputDataHandler : AbstractHandler
{
    public override object? Handle(object data)
    {
        if (!DataInteract.FormatIsValid(data))
        {
            Console.WriteLine("Invalid format");
        }

        Console.WriteLine($"Input data is valid");

        return base.Handle(data);
    }
}