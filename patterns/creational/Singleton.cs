namespace patterns
{
    public class Singleton
    {
        private Singleton() { }
        private static Singleton instance = new Singleton();

        public static Singleton Instance => instance;
    }
}