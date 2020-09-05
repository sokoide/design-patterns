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

    #region Singleton
    public class SingletonTest
    {
        [Fact]
        public void Sameinstance()
        {
            Singleton s = Singleton.Instance;
            Assert.Equal(s.GetHashCode(), Singleton.Instance.GetHashCode());
        }
    }
    #endregion

    #region Builder
    public class TextBuilderTest
    {
        private readonly TextBuilder t;
        public TextBuilderTest()
        {
            t = new TextBuilder();
        }
        [Fact]
        public void BuildTest()
        {
            t.AddTitle("A");
            t.AddListItem(1, "item1");
            t.AddTitle("B");
            t.AddListItem(1, "item2");
            t.AddListItem(2, "item2a");
            t.AddListItem(2, "item2b");
            Assert.Equal("=== A ===\n* item1\n=== B ===\n* item2\n** item2a\n** item2b\n", t.ToString());
        }
    }

    public class MdBuilderTest
    {
        private readonly MdBuilder t;
        public MdBuilderTest()
        {
            t = new MdBuilder();
        }
        [Fact]
        public void BuildTest()
        {
            t.AddTitle("A");
            t.AddListItem(1, "item1");
            t.AddTitle("B");
            t.AddListItem(1, "item2");
            t.AddListItem(2, "item2a");
            t.AddListItem(2, "item2b");
            Assert.Equal("# A\n* item1\n# B\n* item2\n  * item2a\n  * item2b\n", t.ToString());
        }
    }
    #endregion
}
