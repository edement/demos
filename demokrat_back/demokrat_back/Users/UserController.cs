using demokrat_back.Db;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;

namespace demokrat_back.Users
{
    [ApiController]
    [Route("api/users")]
    public class UserController(AppDbContext _database) : ControllerBase
    {
        [Authorize]
        [HttpGet("me")]
        public async Task<IActionResult> GetMe()
        {
            // доставать id пользователя из jwt
            var userIdClaim = User.FindFirst("userId")?.Value;
            if (userIdClaim == null)
            {
                return Unauthorized("User not authenticated");
            }
            var userId = Guid.Parse(userIdClaim);

            var user = await _database.Users.SingleAsync(u => u.Id == userId);
            UserDto response = new UserDto
            {
                Id = user.Id,
                Name = user.Name,
                Email = user.Email,
                IsTrainer = user.IsTrainer
            };
            return Ok(response);
        }
    }
}
