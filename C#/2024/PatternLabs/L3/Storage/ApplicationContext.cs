using Microsoft.EntityFrameworkCore;

namespace L3.Storage;

public class ApplicationContext : DbContext
{
    public DbSet<Client> Clients { get; set; }

    public ApplicationContext(DbContextOptions<ApplicationContext> options) : base(options) =>
        Database.EnsureCreated();

    protected override void OnModelCreating(ModelBuilder modelBuilder)
    {
        // инициализация БД начальными данными
        modelBuilder.Entity<Client>().HasData(
                new Client { ID = 1, Name = "Tom", Email = "example_tom@gmail.com" },
                new Client { ID = 2, Name = "Ben", Email = "example_ben@gmail.com" },
                new Client { ID = 3, Name = "Max", Email = "example_max@gmail.com" }
        );
    }
}