<%@ Page Language="C#" AutoEventWireup="true" CodeBehind="default.aspx.cs" Inherits="calendarControl._default" %>

<!DOCTYPE html>

<html xmlns="http://www.w3.org/1999/xhtml">
<head runat="server">
    <title></title>
</head>
<body>
    <form id="form1" runat="server">
        <div>
            <strong>Working with the Calendar Server Control</strong><br />
            <br />
            <asp:Calendar ID="myCalendar" runat="server" OnSelectionChanged="myCalendar_SelectionChanged"></asp:Calendar>
            <br />
            <br />
            <br />
            <asp:TextBox ID="setDateTextBox" runat="server"></asp:TextBox>
&nbsp;<asp:Button ID="setDateButton" runat="server" Text="Set Date" OnClick="setDateButton_Click" />
            <br />
            <br />
            <asp:Button ID="getDateButton" runat="server" Text="Get Date" OnClick="getDateButton_Click" />
&nbsp;&nbsp;<asp:Button ID="showDateButton" runat="server" Text="Show Date" OnClick="showDateButton_Click" />
&nbsp;<asp:Button ID="SelectedWeekButton" runat="server" Text="Selected Week" />
            <br />
            <br />
            <asp:Label ID="resultLabel" runat="server"></asp:Label>
        </div>
    </form>
</body>
</html>
