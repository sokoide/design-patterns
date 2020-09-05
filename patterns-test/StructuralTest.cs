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

    #region Bridge
    public class BridgeTest
    {
        [Fact]
        public void DrawTest()
        {
            Character maaario = new Character(new RedRenderar());
            Character luuuigi = new Character(new GreenRenderar());
            Assert.Equal("Drawing in Red", maaario.Draw());
            Assert.Equal("Drawing in Green", luuuigi.Draw());
        }
    }
    #endregion
}
