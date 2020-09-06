using System;

namespace patterns
{
    public interface IStrategy
    {
        int NextPartition();
    }

    public class RoundRobinStrategy : IStrategy
    {
        public int NumPartitions = 3;
        private int currentPartition = 0;
        public RoundRobinStrategy(int partitions) { NumPartitions = partitions; }
        public int NextPartition()
        {
            currentPartition = (currentPartition + 1) % NumPartitions;
            return currentPartition;
        }
    }

    public class RandomStrategy : IStrategy
    {
        public int NumPartitions = 3;
        private Random r = new Random();
        public RandomStrategy(int partitions) { NumPartitions = partitions; }
        public int NextPartition()
        {
            return r.Next(NumPartitions);
        }
    }
}