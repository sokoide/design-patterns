using System;

namespace patterns
{
    public class ShapeFactory : IAbstractItemFactory
    {
        public IName GetItem(string name)
        {
            switch (name.ToLower())
            {
                case "rectangle":
                    return new Rectangle();
                case "circle":
                    return new Circle();
                default:
                    throw new ArgumentException($"{name} not supported");
            }
        }
    }
}