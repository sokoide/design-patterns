using System;
using System.Collections.Generic;
using System.Text;

namespace patterns
{
    public class Q : IObservable<Message>
    {
        List<IObserver<Message>> observers = new List<IObserver<Message>>();
        public void Publish(string message)
        {
            Message m = new Message(message);
            // Notify Observers
            foreach (var observer in observers)
                observer.OnNext(m);
        }

        public IDisposable Subscribe(IObserver<Message> observer)
        {
            if (!observers.Contains(observer))
            {
                observers.Add(observer);
            }
            return new Unsubscriber<Message>(observers, observer);

        }
    }

    internal class Unsubscriber<Message> : IDisposable
    {
        private List<IObserver<Message>> _observers;
        private IObserver<Message> _observer;

        internal Unsubscriber(List<IObserver<Message>> observers, IObserver<Message> observer)
        {
            this._observers = observers;
            this._observer = observer;
        }

        public void Dispose()
        {
            if (_observers.Contains(_observer))
                _observers.Remove(_observer);
        }
    }

    public class Message
    {
        private string text;
        public string Text => text;
        public Message(string text) { this.text = text; }
    }

    public class QClient : IObserver<Message>
    {
        private StringBuilder sb = new StringBuilder();
        private IDisposable unsubscriber;

        public string Messages => sb.ToString();

        public QClient(Q q)
        {
            unsubscriber = q.Subscribe(this);
        }

        public void Unsubscribe()
        {
            unsubscriber.Dispose();
            unsubscriber = null;
        }

        public void OnCompleted()
        {
            throw new NotImplementedException();
        }

        public void OnError(Exception error)
        {
            throw new NotImplementedException();
        }

        public void OnNext(Message value)
        {
            sb.Append(value.Text);
        }
    }
}