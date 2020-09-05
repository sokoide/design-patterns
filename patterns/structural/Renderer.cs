namespace patterns
{
    public class RedRenderar : IDrawApi
    {
        public string Draw()
        {
            return "Drawing in Red";
        }
    }

    public class GreenRenderar : IDrawApi
    {
        public string Draw()
        {
            return "Drawing in Green";
        }
    }

}