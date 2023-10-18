using Microsoft.EntityFrameworkCore;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using WpfPostgres.Models;

namespace WpfPostgres.Data
{
    public class ApplicationContext : DbContext
    {
        public DbSet<User> Users { get; set; } = null!;

        // Comment method to make a migration:
        // Use Package Manager Console
        // `Add-Migration Initial`
        // `Update-Database`
        //public ApplicationContext()
        //{
        //    Database.EnsureCreated();
        //}
        protected override void OnConfiguring(DbContextOptionsBuilder optionsBuilder)
        {
            optionsBuilder.UseNpgsql("Host=localhost;Port=5432;Database=entity_test;Username=postgres;Password=314159");
        }
    }
}
