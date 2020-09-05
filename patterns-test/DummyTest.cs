using System;
using patterns;
using Xunit;

namespace patterns_test
{
    public class DummyTest
    {
        private readonly Dummy t;
        public DummyTest()
        {
            t = new Dummy();
        }

        [Fact]
        public void TestName()
        {
            Assert.Equal("Dummy", t.Name());
        }

        [Theory]
        [InlineData(1, 2, 3)]
        [InlineData(-1, 2, 1)]
        [InlineData(-100, 2, -98)]
        public void TestAdd(int a, int b, int want)
        {
            Assert.Equal(want, t.Add(a, b));
        }
    }
}
