namespace patterns
{
    public abstract class Monster
    {
        public string Name;
        public abstract int HP();
        public abstract int MP();

        public string Show()
        {
            return $"name:{Name}, HP:{HP()}, MP:{MP()}";
        }
    }

    public class Slime : Monster
    {
        public Slime(string name)
        {
            Name = name;
        }

        public override int HP()
        {
            return 4;
        }

        public override int MP()
        {
            return 0;
        }
    }
}