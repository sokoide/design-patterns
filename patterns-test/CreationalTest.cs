using System;
using patterns;
using Xunit;

namespace patterns_test
{
    #region Factory
    public class ShapeFactoryTest
    {
        private readonly ShapeFactory t;
        public ShapeFactoryTest()
        {
            t = new ShapeFactory();
        }

        [Theory]
        [InlineData("Rectangle", "Rectangle")]
        [InlineData("Circle", "Circle")]
        public void TestShapeFactory(string shapeName, string want)
        {
            IName s = t.GetItem(shapeName);
            Assert.Equal(want, s.Name());
        }
    }

    public class BeamFactoryTest
    {
        private readonly BeamFactory t;
        public BeamFactoryTest()
        {
            t = new BeamFactory();
        }

        [Theory]
        [InlineData("Rectangle", "RectangleBeam")]
        [InlineData("Circle", "CircleBeam")]
        public void TestBeamFactory(string beamName, string want)
        {
            IName s = t.GetItem(beamName);
            Assert.Equal(want, s.Name());
        }
    }
    #endregion

    #region AbstractFactory
    public class FactoryProducerTest
    {
        [Fact]
        public void FactoryTest()
        {
            IAbstractItemFactory f = FactoryProducer.GetFactory("shape");
            IName s = f.GetItem("Rectangle");
            Assert.Equal("Rectangle", s.Name());

            f = FactoryProducer.GetFactory("beam");
            s = f.GetItem("Rectangle");
            Assert.Equal("RectangleBeam", s.Name());
        }
    }
    #endregion
}
