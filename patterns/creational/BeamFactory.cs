using System;

namespace patterns
{
    public class BeamFactory : IAbstractItemFactory
    {
        public IName GetItem(string name)
        {
            switch (name.ToLower())
            {
                case "rectangle":
                    return new RectangleBeam();
                case "circle":
                    return new CircleBeam();
                default:
                    throw new ArgumentException($"{name} not supported");
            }
        }
    }
}