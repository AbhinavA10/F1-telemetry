using System;
using System.IO;
using System.Net;
using System.Net.Sockets;
using System.Threading;

// Adapted from below link, to record udp data for playback later. 
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
        //read datagrams sent from any source.
        static private IPEndPoint senderIP = new IPEndPoint(IPAddress.Any, 0);
        // UDP Socket for the connection
        static private UdpClient udpSocket;
        // Holds the latest data captured from the game
        static TelemetryPacket latestData;

        // Mutex used to protect latestData from simultaneous access by both threads
        static Mutex syncMutex = new Mutex();

        static void Main(string[] args)
        {
            Console.WriteLine("Listening on port " + PORTNUM + " for connections from" + IP);
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
            int mode = 1; // TODO: get mode from command line

            if (mode == 0)
            {
                RecordPackets();
            }
            else if (mode == 1)
            {
                PlaybackPackets();
            }
            else if (mode == 2)
            {
                ConvertLivePackets();
            }
        }

        /// <summary>
        /// Records UDP packets, and saves it to seperate .bin files
        /// </summary>
        static private void RecordPackets()
        {
            Console.WriteLine("Waiting for packets");
            int n = 0;
            while (true)
            {
                Byte[] receiveBytes = udpSocket.Receive(ref senderIP);
                syncMutex.WaitOne();
                string filename = "Packet" + n + ".bin";
                File.WriteAllBytes(filename, receiveBytes);
                n++;
                syncMutex.ReleaseMutex();
            }
        }

        /// <summary>
        /// Plays back Binary files in the current directory
        /// </summary>
        static private void PlaybackPackets()
        {
            string path = "D:\\Packets";

            syncMutex.WaitOne();

            int count = Directory.GetFiles(path, "*").Length; Console.WriteLine(count);
            Console.ReadLine();
            for (int i = 0; i < count; i++)
            {
                string filename = "D:\\Packets\\Packet" + i + ".bin";
                Byte[] udpBytes = File.ReadAllBytes(filename);
                //TODO: publish packets onto port

            }
            syncMutex.ReleaseMutex();
        }

        static private void ConvertLivePackets()
        {
            Console.WriteLine("Waiting for packets");
            while (true)
            {
                // Get the data (this will block until we get a packet)
                Byte[] receiveBytes = udpSocket.Receive(ref senderIP);
                // Convert the bytes received to the shared struct
                latestData = PacketUtilities.ConvertToPacket(receiveBytes);
                Console.WriteLine("Time:" + latestData.Time + "   " + "Speed: " + latestData.Speed * 2.23694);
            }
        }
    }
}
