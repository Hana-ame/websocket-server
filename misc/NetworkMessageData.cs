using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class NetworkMessageData
{
    public enum MessageID
    {
        C2S_Regist,
        S2C_Regist,
        C2S_Login,
        S2C_Login,
    }

    struct C2S_Regist
    {
        public string name;
        public string id;
        public string password;
    }

    struct S2C_Regist
    {
        public int result;
    }

    struct C2S_Login
    {
        public string id;
        public string password;
    }

    struct S2C_Login
    {
        public int result;
        public string name;
    }
}


