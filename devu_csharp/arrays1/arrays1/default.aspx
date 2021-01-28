<%@ Page Language="C#" AutoEventWireup="true" CodeBehind="default.aspx.cs" Inherits="arrays1._default" %>

<!DOCTYPE html>

<html xmlns="http://www.w3.org/1999/xhtml">
<head runat="server">
    <title></title>
</head>
<body>
    <form id="form1" runat="server">
        <div>
            <asp:TextBox ID="TextBox1" runat="server"></asp:TextBox>
&nbsp;<asp:TextBox ID="TextBox2" runat="server"></asp:TextBox>
&nbsp;<asp:TextBox ID="TextBox3" runat="server"></asp:TextBox>
&nbsp;<asp:TextBox ID="TextBox4" runat="server"></asp:TextBox>
&nbsp;<asp:TextBox ID="TextBox5" runat="server"></asp:TextBox>
            <br />
            <br />
            <asp:Button ID="submitButton1" runat="server" Text="Submit" OnClick="submitButton1_Click" />
            &nbsp;<asp:Button ID="retrieveButton" runat="server" Text="retrieve" OnClick="retrieveButton_Click" />
            <br />
            <br />
            <asp:Label ID="resultLabel" runat="server" ></asp:Label>
        </div>
    </form>
</body>
</html>
