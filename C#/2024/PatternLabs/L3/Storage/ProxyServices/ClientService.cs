using Microsoft.Extensions.Caching.Memory;

namespace L3.Storage.ProxyServices;

public interface IClientService
{
    Task<Client?> GetClient(int id);
}

public class ClientServiceProxy : IClientService
{
    private ClientService _service;
    private IMemoryCache _cache;
    private ILogger _logger;
    public ClientServiceProxy(ApplicationContext db, IMemoryCache cache, ILogger<ClientServiceProxy> logger)
    {

        this._service = new ClientService(db);
        this._cache = cache;
        this._logger = logger;
    }

    public async Task<Client?> GetClient(int id)
    {
        _cache.TryGetValue(id, out Client? client);

        if (client == null)
        {
            client = await _service.GetClient(id);
            if (client != null)
            {
                _logger.LogInformation($"cached user with id = {id}");
                _cache.Set<Client>(client.ID, client, new MemoryCacheEntryOptions().SetAbsoluteExpiration(TimeSpan.FromMinutes(5)));
            }
            else
            {
                _logger.LogInformation($"user with id = {id} was not found");
            }
        }
        else
        {
            _logger.LogInformation($"user with id = {id} found in cache");
        }
        return client;
    }
}

public class ClientService : IClientService
{
    ApplicationContext db;

    public ClientService(ApplicationContext db)
    {
        this.db = db;
    }

    public async Task<Client?> GetClient(int id)
    {
        return await db.Clients.FindAsync(id);
    }
}