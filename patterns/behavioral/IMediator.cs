namespace patterns
{
    public interface IMediator
    {
        void CreateColleagues();
        void ColleagueChanged();
    }

    public interface IColleague
    {
        void Mediator(IMediator mediator);
        void ColleagueEnabled(bool enabled);
    }
}