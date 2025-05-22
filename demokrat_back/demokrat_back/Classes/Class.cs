using demokrat_back.Users;
using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;

namespace demokrat_back.Classes
{
    public class Class
    {
        [Key]
        public Guid Id { get; set; }
        [Required]
        public DateTime DateTime { get; set; }
        [Required]
        [MaxLength(200)]
        public string Location { get; set; } = string.Empty;
        [Required]
        [Column(TypeName = "decimal(10,2)")]
        public decimal Price { get; set; }
        [Required]
        public Guid TrainerId { get; set; }
        [ForeignKey(nameof(TrainerId))]
        public User Trainer { get; set; }

        [Required]
        public DateTime CreatedAt { get; set; }
        public ICollection<Enrollment> Enrollments { get; set; } = new List<Enrollment>();
    }
}
