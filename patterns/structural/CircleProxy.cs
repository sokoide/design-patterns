namespace patterns
{
    public class CircleProxy : IName
    {
        private Circle circle;
        private static object innerlock = new object();

        public string Name()
        {
            if (circle == null)
            {
                lock (innerlock)
                {
                    if (circle == null)
                    {
                        circle = new Circle();
                    }
                }
            }
            return circle.Name();
        }
    }
}