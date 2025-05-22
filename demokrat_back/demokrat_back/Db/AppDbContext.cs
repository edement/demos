using Microsoft.EntityFrameworkCore;
using demokrat_back.Classes;
using demokrat_back.Users;
using demokrat_back.Other;

namespace demokrat_back.Db
{
    public class AppDbContext(DbContextOptions<AppDbContext> options) : DbContext(options)
    {
        public DbSet<User> Users { get; set; }
        public DbSet<Class> Classes { get; set; }
        public DbSet<Enrollment> Enrollments { get; set; }
        public DbSet<RefreshToken> RefreshTokens { get; set; }
    }
}