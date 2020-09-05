namespace patterns
{
    public class ComplexShapeMaker
    {
        public string Make()
        {
            Circle c = new Circle();
            Rectangle r = new Rectangle();
            return $"{c.Name()}-{r.Name()}";
        }
    }
}