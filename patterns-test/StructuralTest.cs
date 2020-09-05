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

    #region Composite
    public class CompositeTest
    {
        [Fact]
        public void VisitTest()
        {
            Node n1 = new Node("n1");
            Node n2 = new Node("n2");
            Leaf l1 = new Leaf("l1");
            Leaf l2 = new Leaf("l2");
            Leaf l3 = new Leaf("l3");
            n1.left = l1;
            n1.right = n2;
            n2.left = l2;
            n2.right = l3;
            Assert.Equal("l1,n1,l2,n2,l3", n1.Visit());
        }
    }
    #endregion
}
