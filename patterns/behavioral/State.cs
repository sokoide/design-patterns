namespace patterns
{

    public class StateContext
    {
        private IState state;

        public void ChangeState(IState state)
        {
            this.state = state;
        }

        public string GetState()
        {
            if (state == null)
                return "";
            else
                return state.ToString();
        }

        public void Do(string action)
        {
            if (state != null)
                state.Do(this, action);
        }
    }

    public interface IState
    {
        void Do(StateContext ctx, string str);
    }

    public class TitleState : IState
    {
        public static TitleState Instance = new TitleState();

        private TitleState() { }

        public override string ToString()
        {
            return "Title";
        }

        public void Do(StateContext ctx, string str)
        {
            if (str == "startbutton")
            {
                ctx.ChangeState(GameState.Instance);
            }
        }
    }

    public class GameState : IState
    {
        public static GameState Instance = new GameState();

        private GameState() { }

        public override string ToString()
        {
            return "Game";
        }
        public void Do(StateContext ctx, string str)
        {
            if (str == "monster")
            {
                ctx.ChangeState(GameOverState.Instance);
            }
        }
    }

    public class GameOverState : IState
    {
        public static GameOverState Instance = new GameOverState();

        private GameOverState() { }

        public override string ToString()
        {
            return "GameOver";
        }
        public void Do(StateContext ctx, string str)
        {
            if (str == "tick")
            {
                ctx.ChangeState(TitleState.Instance);
            }
        }
    }
}