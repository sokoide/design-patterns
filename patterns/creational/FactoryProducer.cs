using System;

namespace patterns
{
    public class FactoryProducer
    {
        public static IAbstractItemFactory GetFactory(string name)
        {
            switch (name.ToLower())
            {
                case "shape":
                    return new ShapeFactory();
                case "beam":
                    return new BeamFactory();
                default: throw new ArgumentException($"{name} not supported");
            }
        }
    }
}