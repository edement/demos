using System.ComponentModel.DataAnnotations;

namespace demokrat_back.Users
{
    public class UserDto
    {
        public Guid Id { get; set; }
        public string Name { get; set; } = string.Empty;
        public string Email { get; set; } = string.Empty;
        public bool IsTrainer { get; set; }
    }
}
