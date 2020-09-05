using System.Text;

namespace patterns
{
    public class DecoratedText
    {
        protected string Text;
        public DecoratedText(DecoratedText t)
        {
            if (t != null)
                Text = t.ToString();
            else
                Text = "";

        }

        public override string ToString()
        {
            return Text;
        }

        public int Width()
        {
            string[] comps = Text.Split('\n');
            return comps[0].Length;
        }

        public int Height()
        {
            string[] comps = Text.Split('\n');
            return comps.Length;
        }
    }

    public class PlainText : DecoratedText
    {
        public PlainText(string text) : base(null)
        {
            Text = text;
        }
    }

    public class UnderlinedText : DecoratedText
    {
        public UnderlinedText(DecoratedText t) : base(t)
        {
            int width = t.Width();
            Text = $"{t.ToString()}\n" + new string('_', width);
        }
    }

    public class SurroundedText : DecoratedText
    {
        public SurroundedText(DecoratedText t) : base(t)
        {
            string[] comps = Text.Split('\n');
            StringBuilder sb = new StringBuilder();
            foreach (string line in comps)
            {
                sb.Append($"|{line}|\n");
            }
            Text = sb.ToString().Substring(0, sb.Length - 1);
        }
    }
}