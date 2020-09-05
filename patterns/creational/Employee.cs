using System;

namespace patterns
{
    public class Employee : ICloneable
    {
        public string Name;

        public Employee(string name) { Name = name; }
        public string Say() { return $"Hello, I'm {Name}"; }

        public object Clone()
        {
            return (Employee)MemberwiseClone();
        }
    }
}