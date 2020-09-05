namespace patterns
{
    public interface IAbstractItemFactory
    {
        IName GetItem(string name);
    }
}