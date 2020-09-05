namespace patterns
{
    public class Character
    {
        protected IDrawApi Api;
        public Character(IDrawApi api) { Api = api; }

        public string Draw()
        {
            return Api.Draw();
        }
    }
}