<%@ Page Language="C#" AutoEventWireup="true" CodeBehind="default.aspx.cs" Inherits="pageLoad._default" %>

<!DOCTYPE html>

<html xmlns="http://www.w3.org/1999/xhtml">
<head runat="server">
    <title></title>
</head>
<body>
    <form id="form1" runat="server">
        <div>
            <asp:TextBox ID="myTextBox" runat="server"></asp:TextBox>
            <br />
            <br />
            <asp:Calendar ID="myCalendar" runat="server"></asp:Calendar>
            <br />
            <asp:Button ID="okButton" runat="server" Text="OK" OnClick="okButton_Click" />
            <br />
            <br />
            <asp:Label ID="resultLabel" runat="server" ></asp:Label>
            <br />
        </div>
    </form>
</body>
</html>
