﻿<%@ Page Language="C#" AutoEventWireup="true" CodeBehind="default.aspx.cs" Inherits="conditionalTernaryOperator._default" %>

<!DOCTYPE html>

<html xmlns="http://www.w3.org/1999/xhtml">
<head runat="server">
    <title></title>
</head>
<body>
    <form id="form1" runat="server">
        <div>
            Is
            <asp:TextBox ID="TextBox1" runat="server"></asp:TextBox>
&nbsp;equal
            <asp:TextBox ID="TextBox2" runat="server"></asp:TextBox>
&nbsp;?
            <br />
            <br />
            <asp:CheckBox ID="myCheckBox" runat="server" Text="true" />
            <br />
            <br />
            <asp:Button ID="okButton" runat="server" OnClick="okButton_Click" Text="ok" />
            <br />
            <br />
            <asp:Label ID="resultLabel" runat="server"></asp:Label>
            <br />
        </div>
    </form>
</body>
</html>