namespace demokrat_back.Classes
{
    namespace demokrat_back.Classes
    {
        public class ClassDto
        {
            public string Id { get; set; }
            public string Date { get; set; }
            public string Time { get; set; }
            public string Location { get; set; }
            public decimal Price { get; set; }
            public TrainerDto Trainer { get; set; }
            public List<string> EnrolledStudents { get; set; } = new List<string>();
        }

        public class TrainerDto
        {
            public string Id { get; set; }
            public string Name { get; set; }
        }
    }

}
