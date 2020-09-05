namespace patterns
{
    public abstract class TreeElement
    {
        public string Name;
        public TreeElement(string name) { Name = name; }
        public abstract string Visit();
    }

    public class Node : TreeElement
    {
        public TreeElement left = null;
        public TreeElement right = null;

        public Node(string name) : base(name) { }
        public override string Visit()
        {
            if (left == null && right == null)
                return Name;
            else if (right == null)
                return $"{left.Visit()},{Name}";
            else
                return $"{left.Visit()},{Name},{right.Visit()}";
        }
    }

    public class Leaf : TreeElement
    {
        public Leaf(string name) : base(name) { }
        public override string Visit() { return Name; }
    }
}