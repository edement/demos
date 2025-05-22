using demokrat_back.Users;

namespace demokrat_back.Other
{
    public class RefreshToken
    {
        public Guid Id { get; set; } = Guid.NewGuid();
        public string HashedToken { get; set; } = null!;
        public Guid UserId { get; set; }
        public DateTime ExpiresAt { get; set; }
        public DateTime CreatedAt { get; set; } = DateTime.UtcNow;
        public bool IsRevoked { get; set; } = false;

        public User User { get; set; } = null!;
    }

}
