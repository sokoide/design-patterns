using System;

namespace patterns
{
    public class NintenboEmployeeAdapterInheritance : Employee, INintenboEmployee
    {

        public NintenboEmployeeAdapterInheritance(string name) : base(name)
        {
        }

        public string ShowName()
        {
            return Name;
        }
    }

    public class NintenboEmployeeAdapterAggregation : INintenboEmployee
    {
        private Employee employee;
        public NintenboEmployeeAdapterAggregation(string name)
        {
            employee = new Employee(name);
        }

        public string ShowName()
        {
            return employee.Name;
        }
    }
}