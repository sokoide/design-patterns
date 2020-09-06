namespace patterns
{
    public class Character
    {
        protected IDrawApi api;
        public Character(IDrawApi api) { this.api = api; }

        public string Draw()
        {
            return api.Draw();
        }
    }
}