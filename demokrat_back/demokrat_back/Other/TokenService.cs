using demokrat_back.Users;
using System.Security.Claims;
using System.Text;
using Microsoft.IdentityModel.Tokens;
using System.IdentityModel.Tokens.Jwt;
using System.Security.Cryptography;
using demokrat_back.Db;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Http.HttpResults;

namespace demokrat_back.Other
{
    public class TokenService
    {
        private readonly IConfiguration _config;
        private readonly AppDbContext _database;
        public TokenService(IConfiguration config, AppDbContext database)
        {
            _config = config;
            _database = database;
        }
        //создание JWT токенов
        public string GenerateAccessToken(User user)
        {
            var claims = new List<Claim>
            {
                new Claim("Name", user.Name),
                new Claim("userId", user.Id.ToString())
            };

            var key = new SymmetricSecurityKey(
                Encoding.UTF8.GetBytes(_config.GetValue<string>("AppSettings:JwtSecretKey")!));

            var creds = new SigningCredentials(key, SecurityAlgorithms.HmacSha256);

            var tokenDescriptor = new JwtSecurityToken(
                issuer: _config.GetValue<string>("AppSettings:Issuer"),
                audience: _config.GetValue<string>("AppSettings:Audience"),
                claims: claims,
                expires: DateTime.UtcNow.AddMinutes(60),
                signingCredentials: creds
                );

            return new JwtSecurityTokenHandler().WriteToken(tokenDescriptor);
        }

        public async Task<string> GenerateRefreshToken(User user)
        {
            var randomNumber = new byte[32];
            using var rng = RandomNumberGenerator.Create();
            rng.GetBytes(randomNumber);
            var refreshToken = Convert.ToBase64String(randomNumber);

            var hashedToken = Hash(refreshToken);
            var token = new RefreshToken
            {
                HashedToken = hashedToken,
                UserId = user.Id,
                ExpiresAt = DateTime.UtcNow.AddDays(14)
            };

            _database.RefreshTokens.Add(token);
            await _database.SaveChangesAsync();

            return hashedToken;
        }
        //парсинг user.Id из токенов
        private string? GetIDClaimFromToken(string accessToken)
        {
            var handler = new JwtSecurityTokenHandler();
            var jwtToken = handler.ReadJwtToken(accessToken);

            return jwtToken.Claims.FirstOrDefault(c => c.Type == "userId")?.Value;
        }
        public bool VerifyRefreshToken(string dbRefresh, string oldRefresh)
        {
            return false;
        }
        //хэшер токена
        public static string Hash(string token)
        {
            using var sha256 = SHA256.Create();
            var bytes = Encoding.UTF8.GetBytes(token);
            var hashBytes = sha256.ComputeHash(bytes);
            return Convert.ToBase64String(hashBytes);
        }
    }
}
