using demokrat_back.Db;
using demokrat_back.Migrations;
using demokrat_back.Other;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;

namespace demokrat_back.Users.Auth
{
    [ApiController]
    [Route("api/auth")]
    public class AuthController(AppDbContext _database, TokenService _tokenService) : ControllerBase
    {
        [AllowAnonymous]
        [HttpPost("registration")]
        public async Task<IActionResult> Registration([FromBody] RegisterDto request)
        {
            // создать пользователя
            var user = new User
            {
                Id = Guid.NewGuid(),
                Name = request.Name,
                Email = request.Email,
                PasswordHash = BCrypt.Net.BCrypt.HashPassword(request.Password),
                IsTrainer = request.IsTrainer
            };
            // сохранить в бд
            await _database.Users.AddAsync(user);
            await _database.SaveChangesAsync();
            // токен + вернуть все вместе
            var accessToken = _tokenService.GenerateAccessToken(user);
            var refreshToken = _tokenService.GenerateRefreshToken(user);

            var response = new
            {
                user = new
                {
                    id = user.Id,
                    name = user.Name,
                    email = user.Email,
                    isTrainer = user.IsTrainer
                },
                tokens = new
                {
                    accessToken = accessToken,
                    refreshToken = refreshToken
                }
            };
            return Ok(response);
        }

        [AllowAnonymous]
        [HttpPost("login")]
        public async Task<IActionResult> Login([FromBody] LoginDto request)
        {
            var user = await _database.Users.SingleAsync(u => u.Email == request.Email);
            if (!BCrypt.Net.BCrypt.Verify(request.Password, user.PasswordHash)) return Unauthorized();

            var accessToken = _tokenService.GenerateAccessToken(user);
            var refreshToken = _tokenService.GenerateRefreshToken(user);

            var response = new
            {
                user = new
                {
                    id = user.Id,
                    name = user.Name,
                    email = user.Email,
                    isTrainer = user.IsTrainer
                },
                tokens = new
                {
                    accessToken = accessToken,
                    refreshToken = refreshToken
                }
            };
            return Ok(response); 
        }
        [HttpGet("refresh")]
        public async Task<IActionResult> RefreshTokens(string oldRefresh)
        {
            var userIdClaim = User.FindFirst("userId")?.Value;
            if (userIdClaim == null)
            {
                return Unauthorized("User not authenticated");
            }
            var userId = Guid.Parse(userIdClaim);
            var user = await _database.Users.FirstOrDefaultAsync(u => u.Id == userId);

            var dbrefresh = await _database.RefreshTokens.FirstOrDefaultAsync(r => r.UserId == userId);

            if(!_tokenService.VerifyRefreshToken(dbrefresh.HashedToken, oldRefresh)) { return Unauthorized(); }

            var newRefresh = _tokenService.GenerateRefreshToken(user);
            var newAccess = _tokenService.GenerateAccessToken(user);

            var response = new
            {
                tokens = new
                {
                    accessToken = newAccess,
                    refreshToken = newRefresh
                }
            };

            return Ok(response);
        } 
    }
}
