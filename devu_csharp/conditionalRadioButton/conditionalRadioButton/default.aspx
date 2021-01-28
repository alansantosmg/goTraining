<%@ Page Language="C#" AutoEventWireup="true" CodeBehind="default.aspx.cs" Inherits="conditionalRadioButton._default" %>

<!DOCTYPE html>

<html xmlns="http://www.w3.org/1999/xhtml">
<head runat="server">
    <title>Conditional Radio Buttons</title>
    
</head>
<body>
    <form id="form1" runat="server">
        <div>
            Your note taking preferences<br />
            <br />
            <asp:RadioButton ID="pencilRadioButton" runat="server" GroupName="note" Text="Pencil" />
            <br />
            <asp:RadioButton ID="penRadioButton"  runat="server" GroupName="note" Text="Pen" />
            <br />
            <asp:RadioButton ID="phoneRadioButton" runat="server" GroupName="note" Text="Phone" />
            <br />
            <asp:RadioButton ID="tabletRadioButton" runat="server" GroupName="note" Text="Tablet" />
            <br />
            <br />
            <asp:Button ID="okButton" runat="server" Text="ok" OnClick="okButton_Click" />
            <br />
            <br />
            <asp:Image ID="resultimage" runat="server" />
            <br />
            <br />
            <asp:Label ID="resultLabel" runat="server"></asp:Label>
        </div>
    </form>
</body>
</html>
