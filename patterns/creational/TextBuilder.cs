using System;
using System.Text;

namespace patterns
{
    public class TextBuilder : IBuilder
    {
        StringBuilder sb = new StringBuilder();

        public void AddListItem(int level, string item)
        {
            sb.Append(new String('*', level));
            sb.Append(" ");
            sb.Append(item);
            sb.Append("\n");
        }

        public void AddTitle(string title)
        {
            sb.Append("=== ");
            sb.Append(title);
            sb.Append(" ===\n");
        }

        public void Clear()
        {
            sb.Clear();
        }

        public override string ToString()
        {
            return sb.ToString();
        }
    }
}