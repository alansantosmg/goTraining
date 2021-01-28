<%@ Page Language="C#" AutoEventWireup="true" CodeBehind="default.aspx.cs" Inherits="formatingStrings._default" %>

<!DOCTYPE html>

<html xmlns="http://www.w3.org/1999/xhtml">
<head runat="server">
    <title>Loan Application Form</title>
</head>
<body>
    <form id="form1" runat="server">
        <div>
            Loan Application Form<br />
            <br />
            Name:
            <asp:TextBox ID="nameTextBox" runat="server"></asp:TextBox>
            <br />
            <br />
            Phone Number:
            <asp:TextBox ID="phoneNumberTextBox" runat="server"></asp:TextBox>
            <br />
            <br />
            Social Security number:
            <asp:TextBox ID="ssnTextBox" runat="server"></asp:TextBox>
            <br />
            <br />
            <asp:Calendar ID="loanCalendar" runat="server"></asp:Calendar>
            <br />
            <br />
            <br />
            Salary:
            <asp:TextBox ID="salaryTextBox" runat="server"></asp:TextBox>
            <br />
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
