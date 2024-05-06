using L3.Storage;
using L3.Storage.ProxyServices;
using Microsoft.EntityFrameworkCore;

var builder = WebApplication.CreateBuilder(args);


builder.Services.AddDbContext<ApplicationContext>(options => options.UseSqlite("Data Source=usercache.db"));
builder.Services.AddTransient<ClientServiceProxy>();
builder.Services.AddMemoryCache();
builder.Services.AddLogging();

builder.Logging.AddConsole();


var app = builder.Build();

app.MapGet("/", () => "Hello World!");
app.MapGet("/user/{id}", async (int id, ClientServiceProxy clientService) =>
{
    Client? client = await clientService.GetClient(id);
    if (client != null) return $"User {client.Name}  Id={client.ID}  Email={client.Email}";
    return "User not found";
});

app.Run();
