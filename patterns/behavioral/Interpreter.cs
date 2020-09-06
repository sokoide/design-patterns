using System.Collections.Generic;
using System.Text;

namespace patterns
{
    // grammar)
    // prog = expr*
    // expr = mul ('+' mul | '-' mul)*
    // mul =  num ('*' num | '/' num)*
    // num = [0-9]*

    public class Token
    {
        public enum TokenType
        {
            Number,
            Plus,
            Minus,
            Multiply,
            Divide
        }

        public TokenType Type;
        public string Value;
    }

    public class Tokenizer
    {
        private List<Token> tokens = new List<Token>();

        public List<Token> Tokens => tokens;

        public Tokenizer(string program)
        {
            for (int i = 0; i < program.Length; i++)
            {
                char c = program[i];
                if (c == ' ' || c == '\n' || c == '\r' || c == '\t') continue;
                if (c == '+' || c == '-' || c == '*' || c == '/')
                {
                    Token t = new Token();
                    t.Value = c.ToString();
                    switch (c)
                    {
                        case '+':
                            t.Type = Token.TokenType.Plus;
                            break;
                        case '-':
                            t.Type = Token.TokenType.Minus;
                            break;
                        case '*':
                            t.Type = Token.TokenType.Multiply;
                            break;
                        case '/':
                            t.Type = Token.TokenType.Divide;
                            break;
                    }
                    tokens.Add(t);
                }
                else if (IsDigit(c))
                {
                    Token t = new Token();
                    t.Type = Token.TokenType.Number;

                    StringBuilder sb = new StringBuilder();
                    sb.Append(c);

                    i++;
                    for (; i < program.Length; i++)
                    {
                        c = program[i];
                        if (!IsDigit(c))
                        {
                            break;
                        }
                        sb.Append(c);
                    }
                    i--;

                    t.Value = sb.ToString();
                    tokens.Add(t);
                }
                else
                {
                    throw new System.Exception($"Parse error {c} at {i}");
                }
            }

        }

        private bool IsDigit(char c)
        {
            return (c >= '0' && c <= '9');
        }
    }

    public class Context
    {
        int pos = 0;
        List<Token> tokens = new List<Token>();

        public Context(string program)
        {
            Tokenizer t = new Tokenizer(program);
            tokens = t.Tokens;
        }

        public Token PeekNextToken()
        {
            if (tokens.Count > pos + 1) return tokens[pos + 1];
            else return null;
        }
        public Token GetCurrentToken()
        {
            Token t = null;
            if (tokens.Count > pos) t = tokens[pos];

            return t;
        }

        public void AdvanceToken()
        {
            pos++;
        }
    }

    public abstract class InterpreterNode
    {
        protected Context context;
        public string Value;

        public InterpreterNode(Context context)
        {
            this.context = context;
        }

        public abstract void Parse();
    }

    public class ExprNode : InterpreterNode
    {
        public ExprNode(Context context) : base(context) { }
        public override void Parse()
        {
            Token currentToken = context.GetCurrentToken();
            Token nextToken = context.PeekNextToken();
            MulNode mul1 = new MulNode(context);
            mul1.Parse();
            Value = $"[expr]{mul1.Value}";

            while (nextToken != null && (nextToken.Type == Token.TokenType.Plus || nextToken.Type == Token.TokenType.Minus))
            {
                // skip op
                string op = nextToken.Value;
                context.AdvanceToken();

                MulNode mul2 = new MulNode(context);
                mul2.Parse();
                Value += op + mul2.Value;

                nextToken = context.GetCurrentToken();
            }
        }
    }
    public class MulNode : InterpreterNode
    {
        public MulNode(Context context) : base(context) { }

        public override void Parse()
        {
            Token currentToken = context.GetCurrentToken();
            Token nextToken = context.PeekNextToken();

            NumNode num1 = new NumNode(context);
            num1.Parse();
            Value = $"[mul]{num1.Value}";

            while (nextToken != null && (nextToken.Type == Token.TokenType.Multiply || nextToken.Type == Token.TokenType.Divide))
            {
                // skip op
                string op = nextToken.Value;
                context.AdvanceToken();

                NumNode num2 = new NumNode(context);
                num2.Parse();
                Value += op + num2.Value;

                nextToken = context.GetCurrentToken();
            }
        }
    }

    public class NumNode : InterpreterNode
    {
        public NumNode(Context context) : base(context) { }
        public override void Parse()
        {
            Value = $"[num]{context.GetCurrentToken().Value}";
            context.AdvanceToken();
        }
    }
}