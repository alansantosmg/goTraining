<%@ Page Language="C#" AutoEventWireup="true" CodeBehind="Default.aspx.cs" Inherits="changingArrayLength.Default" %>

<!DOCTYPE html>

<html xmlns="http://www.w3.org/1999/xhtml">
<head runat="server">
    <title></title>
</head>
<body>
    <form id="form1" runat="server">
        <div>
            Hours worked on Project:
            <asp:TextBox ID="hoursWorkedProjectTextBox" runat="server"></asp:TextBox>
&nbsp;<asp:Button ID="addHoursButton" runat="server" OnClick="addHoursButton_Click" Text="Add Hours" />
            <br />
            <br />
            <asp:Label ID="resultLabel" runat="server"></asp:Label>
        </div>
    </form>
</body>
</html>
