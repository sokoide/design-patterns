namespace patterns
{
    public interface IBuilder
    {
        void AddTitle(string title);
        void AddListItem(int level, string item);

        void Clear();
    }
}