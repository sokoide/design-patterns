using System.Collections.Generic;
using System.Text;

namespace patterns
{
    public class Command : ICommand
    {
        private StringBuilder sb = new StringBuilder();
        private string commandName = "";
        public string Output => sb.ToString();

        public Command(string commandName)
        {
            this.commandName = commandName;
        }

        public void Do()
        {
            sb.Clear();
            sb.Append(commandName);
        }
    }


    public class CommandList
    {
        Stack<Command> commands = new Stack<Command>();

        public void Push(Command cmd)
        {
            commands.Push(cmd);
        }

        public void Pop()
        {
            commands.Pop();
        }
        public string Execute()
        {
            StringBuilder sb = new StringBuilder();
            while (commands.Count > 0)
            {
                Command cmd = commands.Pop();
                cmd.Do();
                sb.Append(cmd.Output);
            }
            return sb.ToString();
        }
    }
}