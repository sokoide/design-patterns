using patterns;
using Xunit;

namespace patterns_test
{
    #region Chain-of-responsibility
    public class ChainOfResponsibilityTest
    {
        [Theory]
        [InlineData(5, "Resolved by L1")]
        [InlineData(15, "Resolved by L2")]
        [InlineData(25, "Resolved by L3")]
        public void SupportTest(int level, string want)
        {
            L1Support l1 = new L1Support("L1");
            L2Support l2 = new L2Support("L2");
            L3Support l3 = new L3Support("L3");
            l1.Next = l2;
            l2.Next = l3;
            l3.Next = null;

            string got = l1.Resolve(level);
            Assert.Equal(want, got);
        }
    }
    #endregion

    #region Command
    public class CommandTest
    {
        [Fact]
        public void CommandsTest()
        {
            CommandList cl = new CommandList();
            cl.Push(new Command("North"));
            cl.Push(new Command("West"));
            cl.Push(new Command("South"));
            cl.Pop();
            cl.Push(new Command("West"));
            Assert.Equal("WestWestNorth", cl.Execute());
        }
    }
    #endregion
}