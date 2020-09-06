namespace patterns
{
    public class Person : Iterator<Person>
    {
        public string Name;
        private Person next;
        public Person(string name, Person next = null)
        {
            Name = name;
            this.next = next;
        }

        public void SetNext(Person person)
        {
            next = person;
        }

        public override bool HasNext()
        {
            return next != null;
        }

        public override Person Next()
        {
            return next;
        }

        public PersonMomento CreateSnapshot()
        {
            return new PersonMomento(Name);
        }

        public void RestoreSnapshot(PersonMomento momento)
        {
            Name = momento.Name;
        }
    }
}