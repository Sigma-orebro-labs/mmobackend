using System;
using System.Collections.Generic;
using System.Net;
using System.Net.Sockets;
using System.Threading;
using System.Threading.Tasks;

namespace MmoBackend.Server
{
    internal class MmoClientSimulator
    {
        public int RequestCount;

        public void Start(string hostname, int port, int commandsPerConnection, int simultaneousConnections)
        {
            Console.WriteLine($"Sending commands to {hostname}:{port}...");

            try
            {
                IPHostEntry ipHostInfo = Dns.GetHostEntry(hostname);
                IPAddress ipAddress = ipHostInfo.AddressList[0];
                IPEndPoint remoteEndPoint = new IPEndPoint(ipAddress, port);

                var tasks = new List<Task>();

                for (int i = 0; i < simultaneousConnections; i++)
                {
                    var task = Task.Run(() => SendCommands(remoteEndPoint, commandsPerConnection));
                    tasks.Add(task);
                }

                Task.WaitAll(tasks.ToArray());
            }
            catch (Exception e)
            {
                Console.WriteLine(e.ToString());
            }
        }

        private void SendCommands(IPEndPoint remoteEndPoint, int commandsPerConnection)
        {
            var socket = new Socket(remoteEndPoint.AddressFamily, SocketType.Stream, ProtocolType.Tcp);

            // Declare the buffers once for the connection, so that
            // they can be reused for all coming messages
            var header = new byte[MmoServer.MessageHeaderLength];
            var footer = new byte[MmoServer.MessageFooterLength];
            var body = new byte[MmoServer.MaxMessageBodyLength];
            var request = new byte[MmoServer.MaxMessageBodyLength + MmoServer.MessageHeaderLength + MmoServer.MessageFooterLength];

            request[0] = MmoServer.MessageStartMarker;

            byte bodyLength;
            byte commandCode;

            try
            {
                socket.Connect(remoteEndPoint);

                for (int i = 0; i < commandsPerConnection; i++)
                {
                    request[1] = MmoServer.GetCurrentUserPositionCommand;
                    request[2] = 1; // content length
                    request[3] = 189; // Single body byte with some value (add something with actual meaning here later...)
                    request[4] = MmoServer.MessageEndMarker;

                    socket.Send(request, 0, 1 + MmoServer.MessageHeaderLength + MmoServer.MessageFooterLength, SocketFlags.None);

                    socket.Receive(header);

                    bodyLength = header[2];

                    if (header[0] != MmoServer.MessageStartMarker)
                    {
                        // Corrupt message. Do something smart.
                    }

                    socket.Receive(body, 0, bodyLength, SocketFlags.None);
                    socket.Receive(footer);

                    if (footer[0] != MmoServer.MessageEndMarker)
                    {
                        // Corrupt message. Do something smart.
                    }

                    Interlocked.Increment(ref RequestCount);
                }
            }
            catch (Exception e)
            {
                Console.WriteLine("Unexpected exception: {0}", e);
            }
            finally
            {
                socket.Shutdown(SocketShutdown.Both);
                socket.Close();
            }
        }
    }
}