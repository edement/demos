using demokrat_back.Classes;
using System.ComponentModel.DataAnnotations;

namespace demokrat_back.Users
{
    public class User
    {
        [Key]
        public Guid Id { get; set; }
        [Required]
        [MaxLength(100)]
        public string Name { get; set; } = string.Empty;
        [Required]
        [MaxLength(255)]
        public string Email { get; set; } = string.Empty;
        [Required]
        public string PasswordHash { get; set; } = string.Empty;
        [Required]
        public bool IsTrainer { get; set; }
        [Required]
        public DateTime CreatedAt { get; set; } = DateTime.UtcNow;

        public ICollection<Class> ClassesCreated { get; set; } = new List<Class>();
        public ICollection<Enrollment> Registrations { get; set; } = new List<Enrollment>();
    }
}
