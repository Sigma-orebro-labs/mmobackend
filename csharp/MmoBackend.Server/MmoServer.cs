using System;
using System.Net;
using System.Net.Sockets;
using System.Threading.Tasks;

namespace MmoBackend.Server
{
    public class MmoServer
    {
        public const ushort MaxMessageBodyLength = 128;

        // Command codes
        public const byte GetCurrentUserPositionCommand = 1;
        public const byte GetCurrentUserPositionResponse = 2;

        // Message byte markers
        public const byte MessageStartMarker = 0xF0; // 1111 0000
        public const byte MessageEndMarker = 0xCC;   // 1100 1100

        public const ushort MessageHeaderLength = 4; // Header marker [1 byte] + command code [1 byte] + content length [2 bytes]
        public const ushort MessageFooterLength = 1; // Footer marker [1 byte]

        public void Start(int port)
        {
            IPHostEntry ipHostInfo = Dns.GetHostEntry("localhost");
            IPAddress ipAddress = ipHostInfo.AddressList[1];
            IPEndPoint localEndPoint = new IPEndPoint(ipAddress, port);

            var listener = new Socket(ipAddress.AddressFamily, SocketType.Stream, ProtocolType.Tcp);

            try
            {
                listener.Bind(localEndPoint);
                listener.Listen(100);

                Console.WriteLine($"Listening at port {port}...");

                while (true)
                {
                    // Program is suspended while waiting for an incoming connection.
                    var socket = listener.Accept();

                    Task.Run(() => HandleConnection(socket));
                }
            }
            catch (Exception e)
            {
                Console.WriteLine(e.ToString());
            }
        }

        private static void HandleConnection(Socket socket)
        {
            try
            {
                // Declare the buffers once for the connection, so that
                // they can be reused for all coming messages
                var header = new byte[MessageHeaderLength];
                var footer = new byte[MessageFooterLength];
                var body = new byte[MaxMessageBodyLength];
                var response = new byte[MaxMessageBodyLength + MessageHeaderLength + MessageFooterLength];

                response[0] = MessageStartMarker;

                ushort bodyLength;
                byte commandCode;

                while (true)
                {
                    socket.Receive(header);

                    if (header[0] != MessageStartMarker)
                    {
                        // Corrupt message. Do something smart.
                    }

                    commandCode = header[1];
                    bodyLength = TypeConverter.BytesToUShort(header[2], header[3]);

                    // Should be rewritten to ReceiveAsync
                    socket.Receive(body, bodyLength, SocketFlags.None);
                    socket.Receive(footer);

                    if (footer[0] != MessageEndMarker)
                    {
                        // Corrupt message. Do something smart.
                    }

                    switch (commandCode)
                    {
                        case GetCurrentUserPositionCommand:

                            // Do something smart with the message body here...

                            // 3 byte response body (x, y, z coordinates)
                            byte b1, b2;
                            TypeConverter.UShortToBytes(3, out b1, out b2);

                            response[1] = GetCurrentUserPositionResponse;
                            response[2] = b1;
                            response[3] = b2;
                            response[4] = 101; // x
                            response[5] = 102; // y
                            response[6] = 103; // z
                            response[7] = MessageEndMarker;

                            socket.Send(response, 0, 3 + MessageHeaderLength + MessageFooterLength, SocketFlags.None);

                            break;

                        default:
                            // Invalid command. Do something smart.
                            break;
                    }
                }
            }
            finally
            {
                socket.Shutdown(SocketShutdown.Both);
                socket.Close();
            }
        }
    }
}