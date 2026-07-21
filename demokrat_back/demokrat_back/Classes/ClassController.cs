using demokrat_back.Classes.demokrat_back.Classes;
using demokrat_back.Db;
using demokrat_back.Users;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;
using System.Text.Json;

namespace demokrat_back.Classes
{
    [ApiController]
    [Route("api/classes")]
    public class ClassController(AppDbContext _database) : ControllerBase
    {
        [HttpPost]
        public async Task<IActionResult> CreateClass([FromBody] ClassCreateRequest request)
        {
            if (!DateTime.TryParse($"{request.Date} {request.Time}", out var dateTime))
                return BadRequest("Invalid date or time format");
            
            var userIdClaim = User.FindFirst("userId")?.Value;
            if (userIdClaim == null)
            {
                return Unauthorized("User not authenticated");
            }
            var userId = Guid.Parse(userIdClaim);

            var newClass = new Class
            {
                Id = Guid.NewGuid(),
                DateTime = DateTime.SpecifyKind(dateTime, DateTimeKind.Utc),
                Location = request.Location,
                Price = request.Price,
                TrainerId = userId,
                CreatedAt = DateTime.UtcNow
            };

            _database.Classes.Add(newClass);
            await _database.SaveChangesAsync();

            return Created();
        }

        [AllowAnonymous]
        [HttpGet]
        public List<ClassDto> GetClasses()
        {
            var classes = _database.Classes
                .Include(c => c.Trainer)
                .Include(c => c.Enrollments)
                .Select(c => new ClassDto
                {
                    Id = c.Id.ToString(),
                    Date = c.DateTime.ToString("dd-MM-yyyy"), // Преобразуем дату в строку в нужном формате
                    Time = c.DateTime.ToString("HH:mm"), // Только время
                    Location = c.Location,
                    Price = c.Price,
                    Trainer = new TrainerDto
                    {
                        Id = c.TrainerId.ToString(),
                        Name = c.Trainer.Name
                    },
                    EnrolledStudents = c.Enrollments
                        .Select(e => e.UserId.ToString()) // Список студентов, которые записались
                        .ToList()
                })
                .ToList();

            return classes;
        }

        [HttpPost("enrollments")]
        public async Task<IActionResult> CreateEnrollment([FromBody] EnrollmentDto request)
        {
            // достать данные с запроса
            var classId = request.ClassId;

            var userIdClaim = User.FindFirst("userId")?.Value;
            if (userIdClaim == null)
            {
                return Unauthorized("User not authenticated");
            }
            var userId = Guid.Parse(userIdClaim);
            // создать запись
            var enrollment = new Enrollment
            {
                UserId = userId,
                ClassId = classId
            };
            // сохранить запись
            _database.Enrollments.Add(enrollment);
            await _database.SaveChangesAsync();
            return Created();
        }

        [HttpGet("enrollments")]
        public List<ClassDto> GetEnrollments()
        {
            var userIdClaim = User.FindFirst("userId")?.Value;
                var userId = Guid.Parse(userIdClaim);
            // Получаем список классов, на которые записан пользователь
            var enrollments = _database.Enrollments
                .Where(e => e.UserId == userId)
                .Include(e => e.Class)
                    .ThenInclude(c => c.Trainer)
                .Select(e => new ClassDto
                {
                    Id = e.Class.Id.ToString(),
                    Date = e.Class.DateTime.ToString("dd.MM.yyyy"),
                    Time = e.Class.DateTime.ToString("hh.mm"),
                    Location = e.Class.Location,
                    Price = e.Class.Price,
                    Trainer = new TrainerDto
                    {
                        Id = e.Class.TrainerId.ToString(),
                        Name = e.Class.Trainer.Name
                    }
                })
                .ToList();

            return enrollments;
        }

        [HttpDelete("enrollments/{classId}")]
        public async Task<IActionResult> CancelEnrollment(Guid classId)
        {
            var userIdClaim = User.FindFirst("userId")?.Value;
            if (userIdClaim == null)
            {
                return Unauthorized("User not authenticated");
            }
            var userId = Guid.Parse(userIdClaim);

            var enrollment = await _database.Enrollments
                .FirstOrDefaultAsync(e => e.UserId == userId && e.ClassId == classId);

            if(enrollment == null) { return NotFound(); }

            _database.Enrollments.Remove(enrollment);
            await _database.SaveChangesAsync();

            return NoContent();
        }
    }
}
