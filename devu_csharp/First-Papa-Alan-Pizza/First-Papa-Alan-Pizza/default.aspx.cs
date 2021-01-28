using System;
using System.Collections.Generic;
using System.Linq;
using System.Web;
using System.Web.UI;
using System.Web.UI.WebControls;

namespace First_Papa_Alan_Pizza
{
    public partial class _default : System.Web.UI.Page
    {
        protected void Page_Load(object sender, EventArgs e)
        {

        }

        protected void fecharPedidoButton_Click(object sender, EventArgs e)
        {
            double price = 0;

            /*
            if (babyRadioButton.Checked)
            {
                price = 10;

            }
            else if (mammaRadioButton.Checked)
            {
                price = 13;
            }
            else if (PapaRadioButton.Checked)
            {
                price = 16;
            }
            else
            { resultLabel.Text = "Escolha o tamanho da pizza"; }

    */

            price = (babyRadioButton.Checked) ? price = 10 : price;
            price = (mammaRadioButton.Checked) ? price = 13 : price;
            price = (PapaRadioButton.Checked) ? price = 16 : price;
            price = (bordaRecheadaRadioButton.Checked) ? price += 2 : price;
            price = (pepperoniCheckBox.Checked) ? price += 1.50 : price;
            price = (onionsCheckBox.Checked) ? price += 0.75 : price;
            price = (greenPeppersCheckBox.Checked) ? price += 0.50 : price; 
            price = (redPeppersCheckBox.Checked) ? price += 0.75 : price;
            price = (anchoviesCheckBox.Checked) ? price += 2 : price;

            if ((pepperoniCheckBox.Checked
                && greenPeppersCheckBox.Checked
                && anchoviesCheckBox.Checked)
                || (pepperoniCheckBox.Checked
                && redPeppersCheckBox.Checked
                && onionsCheckBox.Checked))
            { price -= 2; }

            if (bordafinaRadioButton.Checked == false
                && bordaRecheadaRadioButton.Checked == false)
            { resultLabel.Text = "Escolha o tipo de borda da pizza"; }
            else if (babyRadioButton.Checked == false
                && mammaRadioButton.Checked == false
                && PapaRadioButton.Checked == false)
            { resultLabel.Text = "Escolha o tamanho da pizza"; }
            else
            { resultLabel.Text = price.ToString(); }



               


           

           


        }
    }
}