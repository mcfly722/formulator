using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace Formulator
{
    class Node
    {
        public Node Left { get; private set; }
        public Node Right { get; private set; }
        public Node(Node left, Node right)
        {
            this.Left = left;
            this.Right = right;
        }
    }

    class Program
    {

        public static string BinaryTreeString(Node node)
        {
            var sb = new StringBuilder();
            Action<Node> f = null;
            f = n =>
            {
                if (n == null)
                    sb.Append("x");
                else
                {
                    sb.Append("(");
                    f(n.Left);
                    f(n.Right);
                    sb.Append(")");
                }
            };
            f(node);
            return sb.ToString();
        }

        static IEnumerable<Node> AllBinaryTrees(int size)
        {
            if (size == 0)
                return new Node[] { null };
            return from i in Enumerable.Range(0, size)
                   from left in AllBinaryTrees(i)
                   from right in AllBinaryTrees(size - 1 - i)
                   select new Node(left, right);
        }

        static void Main(string[] args)
        {
            foreach (var t in AllBinaryTrees(6))
                Console.WriteLine(BinaryTreeString(t));
        }
    }
}
