using System;
using System.IO;
using System.Net;
using System.Net.Sockets;
using System.Threading;

// Adapted from the following:
//http://www.robertgray.net.au/posts/2012/3/connecting-to-codemasters-telemetry-feed
//https://github.com/robgray/F1Speed

namespace F1_Capture
{
    class Main_Program
    {
        // Constants
        private const int PORTNUM = 20777;
        private const string IP = "127.0.0.1";

        // This is the IP endpoint we are connecting to (i.e. the IP Address and Port F1 2012 is sending to)      
        static private IPEndPoint remoteIP;
        // This is the IP endpoint for capturing who actually sent the data
        static private IPEndPoint senderIP = new IPEndPoint(IPAddress.Any, 0);
        // UDP Socket for the connection
        static private UdpClient udpSocket;
        // Holds the latest data captured from the game
        static TelemetryPacket latestData;

        // Mutex used to protect latestData from simultaneous access by both threads
        static Mutex syncMutex = new Mutex();
        
        static void Main(string[] args)
        {
            Console.WriteLine("Listing on port " + PORTNUM + " for connections from" +IP);

            remoteIP = new IPEndPoint(IPAddress.Parse(IP), PORTNUM);

            // Set up the socket for collecting game telemetry
            try
            {
                udpSocket = new UdpClient();
                udpSocket.Client.SetSocketOption(SocketOptionLevel.Socket, SocketOptionName.ReuseAddress, true);
                udpSocket.ExclusiveAddressUse = false;
                udpSocket.Client.Bind(remoteIP);
                Console.WriteLine("Bound to socket on " + IP + ":" + PORTNUM.ToString());
            }
            catch (Exception error)
            {
               Console.WriteLine(error.ToString());
            }

            FetchData();    
        }
        // This method runs continously in the data collection thread.  It
        // waits to receive UDP packets from the game, and then writes to a bin file. 
        static private void FetchData()
        {
            Console.WriteLine("Waiting for packets");
            int n = 0;

            while (true)
            {
                // Get the data (this will block until we get a packet)
                Byte[] receiveBytes = new byte[2];
                receiveBytes[0] = 2; receiveBytes[1] = 3;//udpSocket.Receive(ref senderIP);
                syncMutex.WaitOne();
                string filename = "Packet" + n+".bin";
                File.WriteAllBytes(filename, receiveBytes);
                n++;
                syncMutex.ReleaseMutex();

                // Convert the bytes received to the shared struct
                //latestData = PacketUtilities.ConvertToPacket(receiveBytes);
                //Console.WriteLine("Time:"+ latestData.Time+"   "+"Speed: "+latestData.Speed*2.23694);
            }
        }
    }
}
