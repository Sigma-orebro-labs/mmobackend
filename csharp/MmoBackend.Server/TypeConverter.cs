namespace MmoBackend.Server
{
    public static class TypeConverter
    {
        public static ushort BytesToUShort(byte b1, byte b2)
        {
            return (ushort)((b2 << 8) | b1 << 0);
        }

        public static void UShortToBytes(ushort number, out byte b1, out byte b2)
        {
            b2 = (byte)(number >> 8);
            b1 = (byte)number;
        }
    }
}
