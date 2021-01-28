<%@ Page Language="C#" AutoEventWireup="true" CodeBehind="default.aspx.cs" Inherits="conditionalLogic._default" %>

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
&nbsp;equal to
            <asp:TextBox ID="TextBox2" runat="server"></asp:TextBox>
&nbsp;?
            <br />
            <br />
            <asp:CheckBox ID="coolCheckBox1" runat="server" Text="Are you Cool?" />
            <br />
            <br />
            If you could only eat one food for the rest of your life, what would you eat?
            <br />
            <br />
            <asp:RadioButton ID="pizzaRadioButton" runat="server" GroupName="food" Text="Pizza" />
            <br />
            <asp:RadioButton ID="saladRadioButton" runat="server" GroupName="food" Text="Salad" />
            <br />
            <asp:RadioButton ID="peanutRadioButton" runat="server" GroupName="food" Text="Peanut" />
            <br />
            <br />
            <asp:Button ID="Button1" runat="server" OnClick="Button1_Click" Text="ok" />
            <br />
            <br />
            <asp:Label ID="resultLabel" runat="server"></asp:Label>
        </div>
    </form>
</body>
</html>
