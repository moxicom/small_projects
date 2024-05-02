namespace Patterns.ChainOfResp;

class MandatoryHandler : AbstractHandler
{
    public override object? Handle(object data)
    {
        if (!DataInteract.HasMandatoryFields(data))
        {
            Console.WriteLine("Missing mandatory fields");
            return null;
        }

        Console.WriteLine("All mandatory fields are present");

        return base.Handle(data);
    }
}