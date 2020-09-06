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

    #region Interpreter
    public class InterpreterTest
    {
        [Fact]
        public void TestTokenizer1()
        {
            Tokenizer t = new Tokenizer("12345");
            Assert.Single(t.Tokens);
            Assert.Equal("12345", t.Tokens[0].Value);
            Assert.Equal(Token.TokenType.Number, t.Tokens[0].Type);
        }

        [Fact]
        public void TestTokenizer2()
        {
            Tokenizer t = new Tokenizer("10+23*4");
            Assert.Equal(5, t.Tokens.Count);
            Assert.Equal("10", t.Tokens[0].Value);
            Assert.Equal(Token.TokenType.Number, t.Tokens[0].Type);
            Assert.Equal("+", t.Tokens[1].Value);
            Assert.Equal(Token.TokenType.Plus, t.Tokens[1].Type);
            Assert.Equal("*", t.Tokens[3].Value);
            Assert.Equal(Token.TokenType.Multiply, t.Tokens[3].Type);
            Assert.Equal("4", t.Tokens[4].Value);
            Assert.Equal(Token.TokenType.Number, t.Tokens[4].Type);
        }

        [Fact]
        public void TestTokenizer3()
        {
            Tokenizer t = new Tokenizer("10+23*456");
            Assert.Equal(5, t.Tokens.Count);
            Assert.Equal("456", t.Tokens[4].Value);
            Assert.Equal(Token.TokenType.Number, t.Tokens[4].Type);
        }

        [Fact]
        public void TestTokenizer4()
        {
            Tokenizer t = new Tokenizer("10+23*4-");
            Assert.Equal(6, t.Tokens.Count);
            Assert.Equal("10", t.Tokens[0].Value);
            Assert.Equal(Token.TokenType.Number, t.Tokens[0].Type);
            Assert.Equal("+", t.Tokens[1].Value);
            Assert.Equal(Token.TokenType.Plus, t.Tokens[1].Type);
            Assert.Equal("*", t.Tokens[3].Value);
            Assert.Equal(Token.TokenType.Multiply, t.Tokens[3].Type);
            Assert.Equal("4", t.Tokens[4].Value);
            Assert.Equal(Token.TokenType.Number, t.Tokens[4].Type);
            Assert.Equal("-", t.Tokens[5].Value);
            Assert.Equal(Token.TokenType.Minus, t.Tokens[5].Type);
        }

        [Fact]
        public void TestInterpreter1()
        {
            Context context = new Context("1+2");
            ExprNode node = new ExprNode(context);
            node.Parse();
            Assert.Equal("[expr][mul][num]1+[mul][num]2", node.Value);
        }
        [Fact]
        public void TestInterpreter2()
        {
            Context context = new Context("12+23*34");
            ExprNode node = new ExprNode(context);
            node.Parse();
            Assert.Equal("[expr][mul][num]12+[mul][num]23*[num]34", node.Value);
        }

        [Fact]
        public void TestInterpreter3()
        {
            Context context = new Context("12+23*34*45");
            ExprNode node = new ExprNode(context);
            node.Parse();
            Assert.Equal("[expr][mul][num]12+[mul][num]23*[num]34*[num]45", node.Value);
        }

        [Fact]
        public void TestInterpreter4()
        {
            Context context = new Context("12+23*34*45-99");
            ExprNode node = new ExprNode(context);
            node.Parse();
            Assert.Equal("[expr][mul][num]12+[mul][num]23*[num]34*[num]45-[mul][num]99", node.Value);
        }
    }
    #endregion

    #region Itelator
    public class IteratorTest
    {
        [Fact]
        public void TestIterate()
        {
            Person p3 = new Person("P3");
            Person p2 = new Person("P2", p3);
            Person p1 = new Person("P1", p2);

            int total = 1;
            Person p = p1;
            while (p.HasNext())
            {
                p = p.Next();
                total++;
            }
            Assert.Equal(3, total);
        }
    }
    #endregion

    #region Mediator
    public class MediatorTest
    {
        [Fact]
        public void TestMediator()
        {
            Questionaire q = new Questionaire();
            Assert.Equal("TrueTrueTrueTrue", q.Test());

            q.c2.ColleagueEnabled(false);
            Assert.Equal("TrueFalseTrueFalse", q.Test());
        }
    }
    #endregion

    #region Moment
    public class Moment
    {
        [Fact]
        public void TestMoment()
        {
            Person p = new Person("foo");
            PersonMomento momento = p.CreateSnapshot();
            p.Name = "bar";
            Assert.Equal("bar", p.Name);
            p.RestoreSnapshot(momento);
            Assert.Equal("foo", p.Name);
        }
    }
    #endregion

    #region Observer
    public class ObserverTest
    {
        [Fact]
        public void TestObserver()
        {
            Q q = new Q();
            QClient qc = new QClient(q);
            q.Publish("hello");
            q.Publish("hi");
            Assert.Equal("hellohi", qc.Messages);
            q.Publish("bye");
            Assert.Equal("hellohibye", qc.Messages);

            // no more message after unsubscribe
            qc.Unsubscribe();
            Assert.Equal("hellohibye", qc.Messages);

        }
    }
    #endregion

    #region State
    public class StateTest
    {
        [Fact]
        public void TestState()
        {
        }
    }
    #endregion

    #region Strategy
    public class StrategyTest
    {
        [Fact]
        public void TestStrategy()
        {
        }
    }
    #endregion

    #region TemplateMethod
    public class TemplateMethodTest
    {
        [Fact]
        public void TestTemplateMethod()
        {
        }
    }
    #endregion

    #region Visitor
    public class VisitorTest
    {
        [Fact]
        public void TestVisitor()
        {
        }
    }
    #endregion
}