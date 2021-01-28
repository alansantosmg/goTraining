<%@ Page Language="C#" AutoEventWireup="true" CodeBehind="default.aspx.cs" Inherits="my2WebApp._default" %>

<!DOCTYPE html>

<html xmlns="http://www.w3.org/1999/xhtml">
<head runat="server">
    <title>My web app2</title>
</head>
<body>
    <form id="form1" runat="server">
        <div>
            How old are you?
            <asp:TextBox ID="myAgeTextBox" runat="server"></asp:TextBox>
            <br />
            <br />
            How much money do you have in your pocket?
            <asp:TextBox ID="myMoneyTextBox" runat="server"></asp:TextBox>
            <br />
            <br />
            <asp:Button ID="resultButton" runat="server" OnClick="resultButton_Click" Text="Click me to see the fortune" />
            <br />
            <br />
            <asp:Label ID="resultLabel" runat="server"></asp:Label>
        </div>
    </form>
</body>
</html>
