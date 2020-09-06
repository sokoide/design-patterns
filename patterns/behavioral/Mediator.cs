namespace patterns
{
    public class Questionaire : IMediator
    {
        public Checkbox c1, c2;
        private Question q1, q2;

        public Questionaire()
        {
            CreateColleagues();
            ColleagueChanged();
        }
        public void ColleagueChanged()
        {
            q1.ColleagueEnabled(c1.Enabled);
            q2.ColleagueEnabled(c2.Enabled);
        }

        public void CreateColleagues()
        {
            c1 = new Checkbox("check1");
            c2 = new Checkbox("check2");
            q1 = new Question("question1");
            q2 = new Question("question2");
            c1.Mediator(this);
            c2.Mediator(this);
            q1.Mediator(this);
            q2.Mediator(this);
        }

        public string Test()
        {
            return c1.Enabled.ToString() + c2.Enabled.ToString() + q1.Enabled.ToString() + q2.Enabled.ToString();
        }
    }


    public class Checkbox : IColleague
    {
        private IMediator mediator;
        private bool enabled;
        public bool Enabled => enabled;
        private string title;

        public Checkbox(string title, bool enabled = true)
        {
            this.title = title;
            this.enabled = enabled;
        }
        public void ColleagueEnabled(bool enabled)
        {
            if (this.enabled != enabled)
            {
                this.enabled = enabled;
                mediator.ColleagueChanged();
            }
        }

        public void Mediator(IMediator mediator)
        {
            this.mediator = mediator;
        }
    }
    public class Question : IColleague
    {
        private IMediator mediator;
        private bool enabled;
        public bool Enabled => enabled;
        private string question;

        public Question(string question, bool enabled = true)
        {
            this.question = question;
            this.enabled = enabled;
        }
        public void ColleagueEnabled(bool enabled)
        {
            if (this.enabled != enabled)
            {
                this.enabled = enabled;
                mediator.ColleagueChanged();
            }
        }

        public void Mediator(IMediator mediator)
        {
            this.mediator = mediator;
        }
    }
}