<%@ Page Language="C#" AutoEventWireup="true" CodeBehind="default.aspx.cs" Inherits="cs_asp_013._default" %>

<!DOCTYPE html>

<html xmlns="http://www.w3.org/1999/xhtml">
<head runat="server">
    <title></title>
</head>
<body>
    <form id="form1" runat="server">
        <div>
            Is
            <asp:TextBox ID="firstTextBox" runat="server"></asp:TextBox>
&nbsp;<asp:Label ID="comparisionTypeLabel" runat="server"></asp:Label>
&nbsp;<asp:TextBox ID="secondTextBox" runat="server"></asp:TextBox>
&nbsp;?
            <br />
            <br />
            <asp:CheckBox ID="myCheckBox" runat="server" />
            <br />
            <br />
            <asp:Button ID="okButton" runat="server" Text="OK" OnClick="okButton_Click" />
            <br />
            <br />
            <asp:Label ID="resultLabel" runat="server"></asp:Label>
        </div>
    </form>
</body>
</html>
