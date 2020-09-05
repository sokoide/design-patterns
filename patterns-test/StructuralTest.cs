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

    #region Decorator
    public class DecoratorTest
    {
        [Fact]
        public void DecorateTest()
        {
            DecoratedText t = new UnderlinedText(
                    new PlainText("foo")
            );
            Assert.Equal("foo\n___", t.ToString());

            t = new SurroundedText(
                new UnderlinedText(
                    new PlainText("foo")
            ));
            Assert.Equal("|foo|\n|___|", t.ToString());
        }
    }
    #endregion

    #region Facade
    public class FacadeTest
    {
        [Fact]
        public void MakeTest()
        {
            ComplexShapeMaker t = new ComplexShapeMaker();
            Assert.Equal("Circle-Rectangle", t.Make());
        }
    }
    #endregion

    #region Flyweight
    public class FlyweightTest
    {
        [Fact]
        public void SpreadSheetTest()
        {
            SpreadSheet s = new SpreadSheet(3, 3);
            Format f1 = new Format { FontName = "Osaka", FontSize = 12 };
            Format f2 = new Format { FontName = "Monaco", FontSize = 10 };

            for (int x = 0; x < 3; x++)
            {
                for (int y = 0; y < 3; y++)
                {
                    s.SetData(x, y, "foo");
                    if (x % 2 == 0)
                        s.SetFormat(x, y, f1);
                    else
                        s.SetFormat(x, y, f2);
                }
            }
            Assert.Equal("Monaco", s.Cell(1, 1).Format.FontName);
            Assert.Equal("Osaka", s.Cell(2, 1).Format.FontName);
        }
    }
    #endregion
}
