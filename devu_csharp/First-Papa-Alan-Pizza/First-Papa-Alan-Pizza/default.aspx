<%@ Page Language="C#" AutoEventWireup="true" CodeBehind="default.aspx.cs" Inherits="First_Papa_Alan_Pizza._default" %>

<!DOCTYPE html>

<html xmlns="http://www.w3.org/1999/xhtml">
<head runat="server">
    <title></title>
  
    <style type="text/css">
        .auto-style2 {
            text-align: center;
        }
        .auto-style3 {
            font-size: xx-large;
        }
        .auto-style4 {
            color: #FF0000;
        }
        .auto-style5 {
            font-size: large;
        }
    </style>
  
</head>
<body>
    <form id="form1" runat="server">
        <div class="auto-style2">
       
            <br />
       
            <asp:Image ID="logo" runat="server" Height="122px" ImageUrl="~/pizza.jpg" style="text-align: center" Width="157px" />
            <br />
            <strong><span class="auto-style3">Pizzaria Santos</span></strong><br />
       
        </div>

        <div>


            <strong><span class="auto-style5">Qual o tamanho da sua fome?</span></strong><br />


            <asp:RadioButton ID="babyRadioButton" Text="Baby - 4 fatias ($10)" GroupName ="tamanho" runat="server" />
            <br />
            <asp:RadioButton ID="mammaRadioButton" Text="Mamma - 6 fatias ($13)" GroupName ="tamanho" runat="server" />
            <br />
            <asp:RadioButton ID="PapaRadioButton" Text="Papa - 8 fatias ($16)" GroupName ="tamanho" runat="server" />
            <br />
            <br />
            <strong><span class="auto-style5">Tradicional ou moderninha?</span><br />
            </strong>
            <asp:RadioButton ID="bordafinaRadioButton" Text="Borda fina" GroupName ="borda" runat="server" />
            <br />
            <asp:RadioButton ID="bordaRecheadaRadioButton" Text="Borda recheada - Queijo Cheddar ($2)" GroupName ="borda" runat="server" />
            <br />
            <br />
            <strong><span class="auto-style5">Customize para ficar mais deliciosa:</span></strong><br />
            <asp:CheckBox ID="pepperoniCheckBox" Text="Pepperoni (+$1.50)" runat="server" />
            <br />
            <asp:CheckBox ID="onionsCheckBox" Text="Cebola (+$0.75)" runat="server" />
            <br />
            <asp:CheckBox ID="greenPeppersCheckBox" Text="Pimentão Verde (+$0.50)" runat="server" />
            <br />
            <asp:CheckBox ID="redPeppersCheckBox" Text= "Pimentão vermelho (+$0.75)" runat="server" />
            <br />
            <asp:CheckBox ID="anchoviesCheckBox" Text="Anchovas (+$2)" runat="server" />
            <br />
            <br />
            <span class="auto-style4"><strong>Oferta especial</strong></span> do dia:
            <br />
            <br />
            Economize $2.00 ao adicionar <strong><em>Pepperoni, Pimentão verde e Achovas</em></strong> OU <em><strong>Pepperoni, Pimentão vermelho e Cebola </strong></em>à sua pizza.
            <br />
            <br />
            <asp:Button ID="fecharPedidoButton" runat="server" Text="Fechar Pedido" OnClick="fecharPedidoButton_Click" />
&nbsp;<br />
            <br />
            Total R$: 
            <asp:Label ID="resultLabel" runat="server"></asp:Label>
            <br />
            <br />


        </div>

    </form>
</body>
</html>
