<%@ Page Language="C#" AutoEventWireup="true" CodeBehind="default.aspx.cs" Inherits="dateTimeSpans._default" %>

<!DOCTYPE html>

<html xmlns="http://www.w3.org/1999/xhtml">
<head runat="server">
    <title></title>
</head>
<body>
    <form id="form1" runat="server">
        <div>
            Working with spans of time (TimeSpan)<br />
            <br />
            <asp:Button ID="okButton" runat="server" Text="ok" OnClick="okButton_Click" />
            <br />
            <br />
            <asp:Label ID="resultLabel" runat="server" ></asp:Label>
        </div>
    </form>
</body>
</html>
