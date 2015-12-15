using System;

namespace MmoBackend.Server
{
    class Program
    {
        static void Main(string[] args)
        {
            if (args.Length > 0 && args[0] == "--client")
            {
                if (args.Length < 3)
                {
                    Console.WriteLine("Usage: --client <simultaneousConnections> <commandsPerConnection>");
                    return;
                }

                var startTime = DateTime.Now;

                var client = new MmoClientSimulator();

                var simultaneousConnections = int.Parse(args[1]);
                var commandsPerConnection = int.Parse(args[2]);

                client.Start("127.0.0.1", port: 11000, commandsPerConnection: commandsPerConnection, simultaneousConnections: simultaneousConnections);

                var endTime = DateTime.Now;
                var ellapsedSeconds = (endTime - startTime).TotalSeconds;

                Console.WriteLine("Finished");
                Console.WriteLine($"Requests sent: {client.RequestCount} in {ellapsedSeconds} seconds");
                Console.WriteLine($" => {client.RequestCount / ellapsedSeconds} reqs/s, using {simultaneousConnections} simultaneous connections");
            }
            else
            {
                var server = new MmoServer();

                server.Start(11000);
            }

            Console.WriteLine("\nPress ENTER to continue...");
            Console.Read();
        }
    }
}
