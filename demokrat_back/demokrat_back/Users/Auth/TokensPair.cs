namespace demokrat_back.Users.Auth
{
    public class TokensPair
    {
        public string RefreshToken { get; set; } = string.Empty;
        public string AccessToken { get; set; } = string.Empty;
    }
}
