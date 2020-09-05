using System;
using patterns;
using Xunit;

namespace patterns_test
{
    #region Adapter
    public class AdapterTest
    {
        private readonly NintenboEmployeeAdapterInheritance t;
        private readonly NintenboEmployeeAdapterAggregation t2;
        public AdapterTest()
        {
            t = new NintenboEmployeeAdapterInheritance("Scott");
            t2 = new NintenboEmployeeAdapterAggregation("Scott");
        }

        [Fact]
        public void TestShowName()
        {
            Assert.Equal("Scott", t.ShowName());
            Assert.Equal("Scott", t2.ShowName());
        }

    }
    #endregion
}
