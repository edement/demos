using Microsoft.AspNetCore.Http;
using Microsoft.IdentityModel.Tokens;
using System.IdentityModel.Tokens.Jwt;
using System.Text;
using System.Threading.Tasks;

namespace demokrat_back.Middleware
{
    public class AuthMiddleware
    {
        private readonly RequestDelegate _next;
        private readonly string _jwtSecret;

        public AuthMiddleware(RequestDelegate next, IConfiguration _config)
        {
            _next = next;
            _jwtSecret = _config.GetValue<string>("AppSetings:JwtSecretKey");
        }

        public async Task InvokeAsync(HttpContext context)
        {
            var token = context.Request.Headers["Authorization"].FirstOrDefault()?.Split(" ").Last();

            if (token != null)
            {
                try
                {
                    var tokenHandler = new JwtSecurityTokenHandler();
                    var key = Encoding.UTF8.GetBytes(_jwtSecret);

                    var principal = tokenHandler.ValidateToken(token, new TokenValidationParameters
                    {
                        ValidateIssuerSigningKey = true,
                        IssuerSigningKey = new SymmetricSecurityKey(key),
                        ValidateIssuer = false,
                        ValidateAudience = false,
                        ClockSkew = TimeSpan.Zero // Убираем задержку на время
                    }, out _);

                    // Устанавливаем информацию о пользователе в HttpContext.User
                    context.User = principal;
                }
                catch
                {
                    // Если токен невалидный
                    context.Response.StatusCode = 401; // Unauthorized
                    await context.Response.WriteAsync("Invalid token");
                    return;
                }
            }

            // Передаем управление следующему компоненту в конвейере
            await _next(context);
        }
    }
}