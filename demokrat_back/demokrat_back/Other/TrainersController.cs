using demokrat_back.Classes.demokrat_back.Classes;
using demokrat_back.Db;
using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;

namespace demokrat_back.Other
{
    [ApiController]
    [Route("api/trainers")]
    public class TrainersController(AppDbContext _database) : ControllerBase
    {
        [HttpGet("classes")]
        public List<ClassDto> GetClasses()
        {
            var userIdClaim = User.FindFirst("userId")?.Value;
            if(userIdClaim == null) { return null; }
            var userId = Guid.Parse(userIdClaim);

            var classes = _database.Classes
                .Where(c => c.TrainerId == userId)
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
    }
}
