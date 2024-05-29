namespace Patterns.ChainOfResp;

class InputDataHandler : AbstractHandler
{
    public override object? Handle(object data)
    {
        if (!DataInteract.FormatIsValid(data))
        {
            Console.WriteLine("Invalid format");
            // or throw an exception
            return null;
        }

        Console.WriteLine($"Input data is valid");

        return base.Handle(data);
    }
}