using System;
using Microsoft.VisualStudio.TestTools.UnitTesting;
using MmoBackend.Server;

namespace MmoBackend.Test
{
    [TestClass]
    public class TypeConverterTests
    {
        [TestMethod]
        public void BytesToUShort()
        {
            var v1 = TypeConverter.BytesToUShort(255, 255);
            Assert.AreEqual(ushort.MaxValue, v1);

            var v2 = TypeConverter.BytesToUShort(0, 0);
            Assert.AreEqual(ushort.MinValue, v2);

            var v3 = TypeConverter.BytesToUShort(128, 0);
            Assert.AreEqual(128, v3);

            var v4 = TypeConverter.BytesToUShort(89, 67);
            Assert.AreEqual(17241, v4);
        }

        [TestMethod]
        public void UShortToBytes()
        {
            byte b1, b2;
            TypeConverter.UShortToBytes(ushort.MaxValue, out b1, out b2);
            Assert.AreEqual(byte.MaxValue, b1);
            Assert.AreEqual(byte.MaxValue, b2);

            TypeConverter.UShortToBytes(0, out b1, out b2);
            Assert.AreEqual(byte.MinValue, b1);
            Assert.AreEqual(byte.MinValue, b2);

            TypeConverter.UShortToBytes(255, out b1, out b2);
            Assert.AreEqual(255, b1);
            Assert.AreEqual(byte.MinValue, b2);

            TypeConverter.UShortToBytes(56567, out b1, out b2);
            Assert.AreEqual(247, b1);
            Assert.AreEqual(220, b2);
        }
    }
}
