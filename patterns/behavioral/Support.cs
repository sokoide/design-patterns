namespace patterns
{
    public abstract class Support
    {
        protected string name;
        protected Support next = null;
        public Support Next { get; set; }
        public string Name => name;
        public Support(string name)
        {
            this.name = name;
        }

        public abstract string Resolve(int complexity);

        protected string Resolved()
        {
            return $"Resolved by {Name}";
        }
    }

    public class L1Support : Support
    {
        public L1Support(string name) : base(name) { }
        public override string Resolve(int complexity)
        {
            if (complexity <= 10)
                return Resolved();
            else if (Next != null)
                return Next.Resolve(complexity);
            else
                throw new System.Exception("Not handled");
        }
    }
    public class L2Support : Support
    {
        public L2Support(string name) : base(name) { }
        public override string Resolve(int complexity)
        {
            if (complexity <= 20)
                return Resolved();
            else if (Next != null)
                return Next.Resolve(complexity);
            else
                throw new System.Exception("Not handled");
        }
    }
    public class L3Support : Support
    {
        public L3Support(string name) : base(name) { }
        public override string Resolve(int complexity)
        {
            if (complexity <= 30)
                return Resolved();
            else if (Next != null)
                return Next.Resolve(complexity);
            else
                throw new System.Exception("Not handled");
        }
    }
}